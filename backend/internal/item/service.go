package item

import (
	"github.com/akekapong78/workflow/internal/constant"
	"github.com/akekapong78/workflow/internal/model"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (s Service) CreateItem(req model.RequestItem) (model.ResponseItem, error) {
	item := model.Item{
		Title:    req.Title,
		Price:    req.Price,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,
		OwnerID:  1,
	}

	id, err := s.Repository.CreateItem(&item)
	if err != nil {
		return model.ResponseItem{}, err
	}
	
	return model.ResponseItem{
		ID:       id,
		Title:    item.Title,
		Price:    item.Price,
		Quantity: item.Quantity,
		Status:   item.Status,
	}, nil
}

func (s Service) GetItem(id string, userId int) (model.ResponseItem, error) {
	item, err := s.Repository.GetItem(id, userId)
	if err != nil {
		return model.ResponseItem{}, err
	}

	return model.ResponseItem{
		ID:       item.ID,
		Title:    item.Title,
		Price:    item.Price,
		Quantity: item.Quantity,
		Status:   item.Status,
	}, nil
}

func (s Service) GetItems(userId int) ([]model.ResponseItem, error) {
	items, err := s.Repository.GetItems(userId)
	if err != nil {
		return nil, err
	}

	var responseItems []model.ResponseItem
	for _, item := range items {
		responseItems = append(responseItems, model.ResponseItem{
			ID:       item.ID,
			Title:    item.Title,
			Price:    item.Price,
			Quantity: item.Quantity,
			Status:   item.Status,
		})
	}
	return responseItems, nil
}

func (s Service) UpdateItem(id string, req model.RequestItem, userId int) (model.ResponseItem, error) {
	// check if item exists
	item, err := s.Repository.GetItem(id, userId)
	if err != nil {
		return model.ResponseItem{}, err
	}

	item.Title = req.Title
	item.Price = req.Price
	item.Quantity = req.Quantity

	err = s.Repository.UpdateItem(&item)
	if err != nil {
		return model.ResponseItem{}, err
	}

	return model.ResponseItem{
		ID:       item.ID,
		Title:    item.Title,
		Price:    item.Price,
		Quantity: item.Quantity,
		Status:   item.Status,
	}, nil
}

func (s Service) UpdateItemStatus(id string, req model.RequestUpdateItemStatus, userId int) (model.ResponseItem, error) {
	// check if item exists
	item, err := s.Repository.GetItem(id, userId)
	if err != nil {
		return model.ResponseItem{}, err
	}

	item.Status = req.Status

	err = s.Repository.UpdateItem(&item)
	if err != nil {
		return model.ResponseItem{}, err
	}

	return model.ResponseItem{
		ID:       item.ID,
		Title:    item.Title,
		Price:    item.Price,
		Quantity: item.Quantity,
		Status:   item.Status,
	}, nil
}