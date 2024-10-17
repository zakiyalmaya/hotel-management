package guest

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=GuestRepository --output=../mocks --outpkg=mocks
type GuestRepository interface {
	Create(guest *model.GuestEntity) error
	GetByID(id int) (*model.GuestEntity, error)
}