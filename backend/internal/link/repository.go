package link

import (
	"github.com/vnkot/piklnk/pkg/db"

	"gorm.io/gorm"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link, userID uint) error {
	result := repo.Database.Where("id =? and user_id = ?", link.ID, userID).Updates(link)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *LinkRepository) Delete(linkID uint, userID uint) error {
	result := r.Database.Where("id = ? AND user_id = ?", linkID, userID).Delete(&Link{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) GetCount(userId uint) int64 {
	var count int64

	repo.Database.
		Table("links").
		Where("deleted_at is null and user_id = ?", userId).
		Count(&count)

	return count
}

func (repo *LinkRepository) GetAll(limit, offset int, userId uint) []Link {
	var links []Link

	repo.Database.
		Table("links").
		Where("deleted_at is null and user_id = ?", userId).
		Order("id").
		Offset(offset).
		Limit(limit).
		Scan(&links)

	return links
}
