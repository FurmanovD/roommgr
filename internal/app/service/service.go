package service

import (
	"github.com/FurmanovD/roommgr/internal/pkg/db/apidbconvert/v1"
	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
)

type serviceImpl struct {
	cfg       Config
	db        *repository.Repository
	converter apidbconvert.APIDBConverter
}

func NewService(
	cfg Config,
	db *repository.Repository,
	converter apidbconvert.APIDBConverter,
) RoomManagerService {
	return &serviceImpl{
		cfg:       cfg,
		db:        db,
		converter: converter,
	}
}
