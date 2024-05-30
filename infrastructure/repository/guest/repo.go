package guest

import "github.com/zakiyalmaya/hotel-management/model"

type GuestRepository interface {
	Create(guest *model.GuestEntity) error
	GetByID(id int) (*model.GuestEntity, error)
}