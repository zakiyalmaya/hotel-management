package guest

import "github.com/zakiyalmaya/hotel-management/model"

type Service interface {
	Create(guest *model.GuestEntity) error
	GetByID(id int) (*model.GuestResponse, error)
}