package item

import (
	"github.com/akekapong78/workflow/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Database: db,
	}
}

func (r Repository) CreateItem(item *model.Item) (uint, error) {
	err := r.Database.Create(&item).Error
	if err != nil {
		return 0, err
	}

	return item.ID, nil
}

func (r Repository) GetItem(id string, userId int) (model.Item, error) {
	var item model.Item
	err := r.Database.Where("owner_id = ?", userId).First(&item, id).Error
	if err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (r Repository) GetItems(userId int) ([]model.Item, error) {
	var items []model.Item
	err := r.Database.Where("owner_id = ?", userId).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r Repository) UpdateItem(item *model.Item) error {
	err := r.Database.Save(&item).Error
	if err != nil {
		return err
	}

	return nil
}

