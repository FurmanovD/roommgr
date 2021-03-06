// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package automodel

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Room is an object representing the database table.
type Room struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CompanyID int         `boil:"company_id" json:"company_id" toml:"company_id" yaml:"company_id"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *roomR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L roomL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoomColumns = struct {
	ID        string
	Name      string
	CompanyID string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	CompanyID: "company_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

var RoomTableColumns = struct {
	ID        string
	Name      string
	CompanyID string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "rooms.id",
	Name:      "rooms.name",
	CompanyID: "rooms.company_id",
	CreatedAt: "rooms.created_at",
	UpdatedAt: "rooms.updated_at",
	DeletedAt: "rooms.deleted_at",
}

// Generated where

var RoomWhere = struct {
	ID        whereHelperint
	Name      whereHelpernull_String
	CompanyID whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpernull_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "`rooms`.`id`"},
	Name:      whereHelpernull_String{field: "`rooms`.`name`"},
	CompanyID: whereHelperint{field: "`rooms`.`company_id`"},
	CreatedAt: whereHelpertime_Time{field: "`rooms`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`rooms`.`updated_at`"},
	DeletedAt: whereHelpernull_Time{field: "`rooms`.`deleted_at`"},
}

// RoomRels is where relationship names are stored.
var RoomRels = struct {
	Company      string
	Reservations string
}{
	Company:      "Company",
	Reservations: "Reservations",
}

// roomR is where relationships are stored.
type roomR struct {
	Company      *Company         `boil:"Company" json:"Company" toml:"Company" yaml:"Company"`
	Reservations ReservationSlice `boil:"Reservations" json:"Reservations" toml:"Reservations" yaml:"Reservations"`
}

// NewStruct creates a new relationship struct
func (*roomR) NewStruct() *roomR {
	return &roomR{}
}

// roomL is where Load methods for each relationship are stored.
type roomL struct{}

var (
	roomAllColumns            = []string{"id", "name", "company_id", "created_at", "updated_at", "deleted_at"}
	roomColumnsWithoutDefault = []string{"name", "company_id", "updated_at", "deleted_at"}
	roomColumnsWithDefault    = []string{"id", "created_at"}
	roomPrimaryKeyColumns     = []string{"id"}
)

type (
	// RoomSlice is an alias for a slice of pointers to Room.
	// This should almost always be used instead of []Room.
	RoomSlice []*Room
	// RoomHook is the signature for custom Room hook methods
	RoomHook func(context.Context, boil.ContextExecutor, *Room) error

	roomQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	roomType                 = reflect.TypeOf(&Room{})
	roomMapping              = queries.MakeStructMapping(roomType)
	roomPrimaryKeyMapping, _ = queries.BindMapping(roomType, roomMapping, roomPrimaryKeyColumns)
	roomInsertCacheMut       sync.RWMutex
	roomInsertCache          = make(map[string]insertCache)
	roomUpdateCacheMut       sync.RWMutex
	roomUpdateCache          = make(map[string]updateCache)
	roomUpsertCacheMut       sync.RWMutex
	roomUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var roomBeforeInsertHooks []RoomHook
var roomBeforeUpdateHooks []RoomHook
var roomBeforeDeleteHooks []RoomHook
var roomBeforeUpsertHooks []RoomHook

var roomAfterInsertHooks []RoomHook
var roomAfterSelectHooks []RoomHook
var roomAfterUpdateHooks []RoomHook
var roomAfterDeleteHooks []RoomHook
var roomAfterUpsertHooks []RoomHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Room) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Room) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Room) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Room) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Room) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Room) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Room) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Room) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Room) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roomAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRoomHook registers your hook function for all future operations.
func AddRoomHook(hookPoint boil.HookPoint, roomHook RoomHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		roomBeforeInsertHooks = append(roomBeforeInsertHooks, roomHook)
	case boil.BeforeUpdateHook:
		roomBeforeUpdateHooks = append(roomBeforeUpdateHooks, roomHook)
	case boil.BeforeDeleteHook:
		roomBeforeDeleteHooks = append(roomBeforeDeleteHooks, roomHook)
	case boil.BeforeUpsertHook:
		roomBeforeUpsertHooks = append(roomBeforeUpsertHooks, roomHook)
	case boil.AfterInsertHook:
		roomAfterInsertHooks = append(roomAfterInsertHooks, roomHook)
	case boil.AfterSelectHook:
		roomAfterSelectHooks = append(roomAfterSelectHooks, roomHook)
	case boil.AfterUpdateHook:
		roomAfterUpdateHooks = append(roomAfterUpdateHooks, roomHook)
	case boil.AfterDeleteHook:
		roomAfterDeleteHooks = append(roomAfterDeleteHooks, roomHook)
	case boil.AfterUpsertHook:
		roomAfterUpsertHooks = append(roomAfterUpsertHooks, roomHook)
	}
}

// One returns a single room record from the query.
func (q roomQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Room, error) {
	o := &Room{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "automodel: failed to execute a one query for rooms")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Room records from the query.
func (q roomQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoomSlice, error) {
	var o []*Room

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "automodel: failed to assign all query results to Room slice")
	}

	if len(roomAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Room records in the query.
func (q roomQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: failed to count rooms rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q roomQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "automodel: failed to check if rooms exists")
	}

	return count > 0, nil
}

// Company pointed to by the foreign key.
func (o *Room) Company(mods ...qm.QueryMod) companyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.CompanyID),
	}

	queryMods = append(queryMods, mods...)

	query := Companies(queryMods...)
	queries.SetFrom(query.Query, "`companies`")

	return query
}

// Reservations retrieves all the reservation's Reservations with an executor.
func (o *Room) Reservations(mods ...qm.QueryMod) reservationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`reservations`.`room_id`=?", o.ID),
	)

	query := Reservations(queryMods...)
	queries.SetFrom(query.Query, "`reservations`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`reservations`.*"})
	}

	return query
}

// LoadCompany allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (roomL) LoadCompany(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRoom interface{}, mods queries.Applicator) error {
	var slice []*Room
	var object *Room

	if singular {
		object = maybeRoom.(*Room)
	} else {
		slice = *maybeRoom.(*[]*Room)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roomR{}
		}
		args = append(args, object.CompanyID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roomR{}
			}

			for _, a := range args {
				if a == obj.CompanyID {
					continue Outer
				}
			}

			args = append(args, obj.CompanyID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`companies`),
		qm.WhereIn(`companies.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Company")
	}

	var resultSlice []*Company
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Company")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for companies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for companies")
	}

	if len(roomAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Company = foreign
		if foreign.R == nil {
			foreign.R = &companyR{}
		}
		foreign.R.Rooms = append(foreign.R.Rooms, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CompanyID == foreign.ID {
				local.R.Company = foreign
				if foreign.R == nil {
					foreign.R = &companyR{}
				}
				foreign.R.Rooms = append(foreign.R.Rooms, local)
				break
			}
		}
	}

	return nil
}

// LoadReservations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (roomL) LoadReservations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRoom interface{}, mods queries.Applicator) error {
	var slice []*Room
	var object *Room

	if singular {
		object = maybeRoom.(*Room)
	} else {
		slice = *maybeRoom.(*[]*Room)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roomR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roomR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`reservations`),
		qm.WhereIn(`reservations.room_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load reservations")
	}

	var resultSlice []*Reservation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice reservations")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on reservations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for reservations")
	}

	if len(reservationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Reservations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &reservationR{}
			}
			foreign.R.Room = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RoomID {
				local.R.Reservations = append(local.R.Reservations, foreign)
				if foreign.R == nil {
					foreign.R = &reservationR{}
				}
				foreign.R.Room = local
				break
			}
		}
	}

	return nil
}

// SetCompany of the room to the related item.
// Sets o.R.Company to related.
// Adds o to related.R.Rooms.
func (o *Room) SetCompany(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Company) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `rooms` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"company_id"}),
		strmangle.WhereClause("`", "`", 0, roomPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CompanyID = related.ID
	if o.R == nil {
		o.R = &roomR{
			Company: related,
		}
	} else {
		o.R.Company = related
	}

	if related.R == nil {
		related.R = &companyR{
			Rooms: RoomSlice{o},
		}
	} else {
		related.R.Rooms = append(related.R.Rooms, o)
	}

	return nil
}

// AddReservations adds the given related objects to the existing relationships
// of the room, optionally inserting them as new records.
// Appends related to o.R.Reservations.
// Sets related.R.Room appropriately.
func (o *Room) AddReservations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Reservation) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RoomID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `reservations` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"room_id"}),
				strmangle.WhereClause("`", "`", 0, reservationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RoomID = o.ID
		}
	}

	if o.R == nil {
		o.R = &roomR{
			Reservations: related,
		}
	} else {
		o.R.Reservations = append(o.R.Reservations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &reservationR{
				Room: o,
			}
		} else {
			rel.R.Room = o
		}
	}
	return nil
}

// Rooms retrieves all the records using an executor.
func Rooms(mods ...qm.QueryMod) roomQuery {
	mods = append(mods, qm.From("`rooms`"))
	return roomQuery{NewQuery(mods...)}
}

// FindRoom retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRoom(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Room, error) {
	roomObj := &Room{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `rooms` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, roomObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "automodel: unable to select from rooms")
	}

	if err = roomObj.doAfterSelectHooks(ctx, exec); err != nil {
		return roomObj, err
	}

	return roomObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Room) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("automodel: no rooms provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roomColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	roomInsertCacheMut.RLock()
	cache, cached := roomInsertCache[key]
	roomInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			roomAllColumns,
			roomColumnsWithDefault,
			roomColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(roomType, roomMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(roomType, roomMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `rooms` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `rooms` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `rooms` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, roomPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "automodel: unable to insert into rooms")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == roomMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "automodel: unable to populate default values for rooms")
	}

CacheNoHooks:
	if !cached {
		roomInsertCacheMut.Lock()
		roomInsertCache[key] = cache
		roomInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Room.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Room) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	roomUpdateCacheMut.RLock()
	cache, cached := roomUpdateCache[key]
	roomUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			roomAllColumns,
			roomPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("automodel: unable to update rooms, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `rooms` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, roomPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(roomType, roomMapping, append(wl, roomPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to update rooms row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "automodel: failed to get rows affected by update for rooms")
	}

	if !cached {
		roomUpdateCacheMut.Lock()
		roomUpdateCache[key] = cache
		roomUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q roomQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to update all for rooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to retrieve rows affected for rooms")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoomSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("automodel: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `rooms` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, roomPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to update all in room slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to retrieve rows affected all in update all room")
	}
	return rowsAff, nil
}

var mySQLRoomUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Room) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("automodel: no rooms provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roomColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRoomUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	roomUpsertCacheMut.RLock()
	cache, cached := roomUpsertCache[key]
	roomUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			roomAllColumns,
			roomColumnsWithDefault,
			roomColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			roomAllColumns,
			roomPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("automodel: unable to upsert rooms, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`rooms`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `rooms` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(roomType, roomMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(roomType, roomMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "automodel: unable to upsert for rooms")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == roomMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(roomType, roomMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "automodel: unable to retrieve unique values for rooms")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "automodel: unable to populate default values for rooms")
	}

CacheNoHooks:
	if !cached {
		roomUpsertCacheMut.Lock()
		roomUpsertCache[key] = cache
		roomUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Room record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Room) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("automodel: no Room provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), roomPrimaryKeyMapping)
	sql := "DELETE FROM `rooms` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to delete from rooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "automodel: failed to get rows affected by delete for rooms")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q roomQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("automodel: no roomQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to delete all from rooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "automodel: failed to get rows affected by deleteall for rooms")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoomSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(roomBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `rooms` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, roomPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "automodel: unable to delete all from room slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "automodel: failed to get rows affected by deleteall for rooms")
	}

	if len(roomAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Room) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRoom(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoomSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoomSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `rooms`.* FROM `rooms` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, roomPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "automodel: unable to reload all in RoomSlice")
	}

	*o = slice

	return nil
}

// RoomExists checks if the Room row exists.
func RoomExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `rooms` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "automodel: unable to check if rooms exists")
	}

	return exists, nil
}
