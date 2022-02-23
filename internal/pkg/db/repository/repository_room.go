package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
	"github.com/FurmanovD/roommgr/pkg/sqldb"
)

type roomRepositoryImpl struct {
	db sqldb.SqlDB
}

// ================= interface instantiation =================================================

func NewRoomRepository(db sqldb.SqlDB) RoomRepository {
	return &roomRepositoryImpl{
		db: db,
	}
}

// ================= interface methods =================================================
func (r *roomRepositoryImpl) GetRoom(ctx context.Context, id int) (*automodel.Room, error) {
	room, err := automodel.Rooms(
		automodel.RoomWhere.ID.EQ(id),
	).One(ctx, r.db.Connection())

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return room, nil
}

func (r *roomRepositoryImpl) GetRooms(ctx context.Context, filter RoomFilter) (automodel.RoomSlice, error) {
	var qm []qm.QueryMod
	if !filter.IncludeDeleted {
		qm = append(qm, automodel.RoomWhere.DeletedAt.EQ(null.TimeFromPtr(nil)))
	}
	rooms, err := automodel.Rooms(
		qm...,
	).All(ctx, r.db.Connection())

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return rooms, nil
}

func (r *roomRepositoryImpl) GetReservation(
	ctx context.Context,
	executor Transaction,
	id int,
) (*ReservationJoinUserJoinRoom, error) {
	var exec boil.ContextExecutor
	if executor != nil {
		exec = executor.Executor()
	} else {
		exec = r.db.Connection()
	}

	joins, err := r.getReservationJoins(
		ctx,
		exec,
		automodel.ReservationWhere.ID.EQ(id),
	)
	if err != nil {
		return nil, NewDBError("reading room reservation[ID:%d] failed: %w", id, err)
	}
	if len(joins) == 0 {
		return nil, nil
	}

	return joins[0], nil
}

func (r *roomRepositoryImpl) GetRoomReservations(
	ctx context.Context,
	executor Transaction,
	roomID int,
	startFrom time.Time,
) ([]*ReservationJoinUserJoinRoom, error) {
	var exec boil.ContextExecutor
	if executor != nil {
		exec = executor.Executor()
	} else {
		exec = r.db.Connection()
	}

	joins, err := r.getReservationJoins(
		ctx,
		exec,
		automodel.ReservationWhere.RoomID.EQ(roomID),
		automodel.ReservationWhere.DeletedAt.EQ(null.TimeFromPtr(nil)),
	)
	if err != nil {
		return nil, NewDBError("reading room[ID:%d] reservations failed: %w", roomID, err)
	}
	if len(joins) == 0 {
		return nil, nil
	}

	return joins, nil
}

func (r *roomRepositoryImpl) GetDBReservations(
	ctx context.Context,
	executor Transaction,
	id int,
	userID int,
	roomID int,
	startingFrom time.Time,
) (automodel.ReservationSlice, error) {
	var exec boil.ContextExecutor
	if executor != nil {
		exec = executor.Executor()
	} else {
		exec = r.db.Connection()
	}

	var qmods []qm.QueryMod
	if id > 0 {
		qmods = append(qmods, automodel.ReservationWhere.ID.EQ(id))
	}
	if roomID > 0 {
		qmods = append(qmods, automodel.ReservationWhere.RoomID.EQ(roomID))
	}
	if userID > 0 {
		qmods = append(qmods, automodel.ReservationWhere.UserID.EQ(userID))
	}
	if !startingFrom.Equal(time.Time{}) {
		qmods = append(qmods, automodel.ReservationWhere.StartTime.GTE(startingFrom))
	}

	reservations, err := automodel.Reservations(
		qmods...,
	).All(ctx, exec)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return reservations, nil
}

func (r *roomRepositoryImpl) InsertReservation(
	ctx context.Context,
	executor Transaction,
	reservation *automodel.Reservation,
) (*automodel.Reservation, error) {
	var exec boil.ContextExecutor
	if executor != nil {
		exec = executor.Executor()
	} else {
		exec = r.db.Connection()
	}

	reservation.CreatedAt = time.Now().UTC()

	if err := reservation.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, NewDBError(
			"insert reservation [%v] to a DB failed: %w",
			reservation,
			err,
		)
	}

	return reservation, nil
}

func (r *roomRepositoryImpl) DeleteReservation(
	ctx context.Context,
	executor Transaction,
	reservationID int,
) error {
	var exec boil.ContextExecutor
	if executor != nil {
		exec = executor.Executor()
	} else {
		exec = r.db.Connection()
	}

	reservation := &automodel.Reservation{
		ID:        reservationID,
		DeletedAt: null.TimeFrom(time.Now().UTC()),
	}

	if _, err := reservation.Update(
		ctx,
		exec,
		boil.Whitelist(automodel.ReservationColumns.DeletedAt),
	); err != nil {
		return NewDBError(
			"mark reservation[ID:%d] deleted failed: %w",
			reservationID,
			err,
		)
	}
	return nil
}

// ================= utility methods =================================================
func (r *roomRepositoryImpl) getReservationJoins(
	ctx context.Context,
	executor boil.ContextExecutor,
	queryMods ...qm.QueryMod,
) ([]*ReservationJoinUserJoinRoom, error) {
	var joins []*ReservationJoinUserJoinRoom

	queryModsAll := append(r.getReservationQueryMods(), queryMods...)

	err := automodel.NewQuery(queryModsAll...).Bind(ctx, executor, &joins)
	if err != nil {
		// no error, no records found
		if err == sql.ErrNoRows || err.Error() == ErrBindNoSQLRows.Error() {
			return nil, nil
		}

		return nil, err
	}

	return joins, nil
}

func (r *roomRepositoryImpl) getReservationQueryMods() []qm.QueryMod {
	fields := r.generateColumnNames(
		map[string][]string{
			automodel.TableNames.Reservations: {
				automodel.ReservationColumns.ID,
				automodel.ReservationColumns.StartTime,
				automodel.ReservationColumns.EndTime,
				automodel.ReservationColumns.RoomID,
				automodel.ReservationColumns.UserID,
			},

			automodel.TableNames.Rooms: {
				automodel.RoomColumns.Name,
				automodel.RoomColumns.CompanyID,
			},

			automodel.TableNames.Users: {
				automodel.UserColumns.FirstName,
				automodel.UserColumns.LastName,
				automodel.UserColumns.Email,
			},

			automodel.TableNames.Companies: {
				automodel.CompanyColumns.Name,
			},
		})

	return []qm.QueryMod{
		// fields
		qm.Select(fields...),
		// From
		qm.From(automodel.TableNames.Reservations),
		// Joins
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s.%s ",
				automodel.TableNames.Rooms,
				automodel.ReservationColumns.RoomID,
				automodel.TableNames.Rooms,
				automodel.RoomColumns.ID,
			)),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s.%s ",
				automodel.TableNames.Users,
				automodel.ReservationColumns.UserID,
				automodel.TableNames.Users,
				automodel.UserColumns.ID,
			)),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s.%s = %s.%s ",
				automodel.TableNames.Companies,
				automodel.TableNames.Rooms,
				automodel.RoomColumns.CompanyID,
				automodel.TableNames.Companies,
				automodel.CompanyColumns.ID,
			)),
	}
}

func (r *roomRepositoryImpl) generateColumnNames(names map[string][]string) []string {
	res := make([]string, 0)
	for table, columns := range names {
		for _, column := range columns {
			if column == "*" {
				res = append(res, fmt.Sprintf("`%s`.*", table))
			} else {
				res = append(res,
					fmt.Sprintf(
						"`%s`.`%s` as \"%s.%s\" ",
						table, column,
						table, column,
					))
			}
		}
	}

	return res
}
