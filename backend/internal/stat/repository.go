package stat

import (
	"fmt"
	"time"

	"github.com/vnkot/piklnk/pkg/db"
	"gorm.io/datatypes"
)

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{database}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currDate := datatypes.Date(time.Now())

	repo.Database.Find(&stat, "link_id = ? AND date = ?", linkId, currDate)

	if stat.ID == 0 {
		repo.Database.Create(&Stat{
			Clicks: 1,
			LinkID: linkId,
			Date:   currDate,
		})
	} else {
		stat.Clicks += 1
		repo.Database.Save(&stat)
	}
}

func (repo *StatRepository) GetGroupStat(linkID uint, by string, from, to time.Time) ([]GetGroupStatResponse, error) {
	var results []GetGroupStatResponse

	dateFormat := "YYYY-MM-DD"
	if by == "month" {
		dateFormat = "YYYY-MM"
	}

	query := `
        SELECT 
            to_char(date::timestamp, $1) as period,
            SUM(clicks)::text as count
        FROM stats
        WHERE link_id = $2
          AND date BETWEEN $3 AND $4
        GROUP BY to_char(date::timestamp, $1)
        ORDER BY period ASC
    `

	err := repo.Database.Raw(query, dateFormat, linkID, from, to).Scan(&results).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get group stats: %w", err)
	}

	return results, nil
}
