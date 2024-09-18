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

func (s Service) CreateItem(req model.RequestItem, userId uint) (model.ResponseItem, error) {
	item := model.Item{
		Title:    req.Title,
		Price:    req.Price,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,
		OwnerID:  userId,
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

func (s Service) GetItem(id string, userId uint, role string) (model.ResponseItem, error) {
	item, err := s.Repository.GetItem(id, userId, role)
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

func (s Service) GetItems(userId uint, role string) ([]model.ResponseItem, error) {
	items, err := s.Repository.GetItems(userId, role)
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

func (s Service) UpdateItem(id string, req model.RequestItem, userId uint, role string) (model.ResponseItem, error) {
	// check if item exists
	item, err := s.Repository.GetItem(id, userId, role)
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

func (s Service) UpdateItemStatus(id string, req model.RequestUpdateItemStatus, role string) (model.ResponseItem, error) {
	// check if item exists
	item, err := s.Repository.GetItem(id, 0, role)
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