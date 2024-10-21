package guest

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=GuestService --output=../mocks --outpkg=mocks
type GuestService interface {
	Create(guest *model.GuestEntity) error
	GetByID(id int) (*model.GuestResponse, error)
}