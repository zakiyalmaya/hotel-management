package guest

import (
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/model"
)

type guestSvcImpl struct {
	repos *repository.Repositories
}

func NewGuestServiceImpl(repos *repository.Repositories) Service {
	return &guestSvcImpl{repos: repos}
}

func (g *guestSvcImpl) Create(guest *model.GuestEntity) error {
	err := g.repos.GuestRepo.Create(guest)
	if err != nil {
		return err
	}

	return nil
}

func (g *guestSvcImpl) GetByID(id int) (*model.GuestResponse, error) {
	guest, err := g.repos.GuestRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &model.GuestResponse{
		ID:          guest.ID,
		Name:        guest.FirstName + " " + guest.LastName,
		Identity:    guest.Identity,
		DateOfBirth: guest.DateOfBirth.Format("02 January 2006"),
		PhoneNumber: guest.PhoneNumber,
		Email:       guest.Email,
	}, nil
}
