package repository

import (
	"github.com/vnkot/piklnk/pkg/db"
)

type StatRepository struct {
	database *db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{
		database: database,
	}
}

func (r *StatRepository) AddClick(linkID uint, ipAddress string, useragent string) {

}
