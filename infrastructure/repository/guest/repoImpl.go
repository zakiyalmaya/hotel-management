package guest

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/hotel-management/model"
)

type guestRepoImpl struct {
	db *sqlx.DB
}

func NewGuestRepository(db *sqlx.DB) GuestRepository {
	return &guestRepoImpl{db: db}
}

func (g *guestRepoImpl) Create(guest *model.GuestEntity) error {
	query := "INSERT INTO guests (first_name, last_name, identity_number, date_of_birth, phone_number, email) VALUES (:first_name, :last_name, :identity_number, :date_of_birth, :phone_number, :email)"
	_, err := g.db.NamedExec(query, guest)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (g *guestRepoImpl) GetByID(id int) (*model.GuestEntity, error) {
	guest := &model.GuestEntity{}
	query := "SELECT id, first_name, last_name, identity_number, date_of_birth, phone_number, email, created_at, updated_at FROM guests WHERE id = ?"
	
	err := g.db.Get(guest, query, id)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return guest, nil
}