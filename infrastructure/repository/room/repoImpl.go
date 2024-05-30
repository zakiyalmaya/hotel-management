package room

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/zakiyalmaya/hotel-management/model"
)

type roomRepoImpl struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) RoomRepository {
	return &roomRepoImpl{
		db: db,
	}
}

func (r *roomRepoImpl) Create(room *model.RoomEntity) error {
	query := "INSERT INTO rooms (name, floor, type, price, status, description) VALUES (:name, :floor, :type, :price, :status, :description)"
	_, err := r.db.NamedExec(query, room)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (r *roomRepoImpl) GetByName(name string) (*model.RoomEntity, error) {
	room := &model.RoomEntity{}
	query := "SELECT id, name, floor, type, price, status, description, created_at, updated_at FROM rooms WHERE name = ?"
	
	err := r.db.Get(room, query, name)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return room, nil
}

func (r *roomRepoImpl) GetAll(request *model.GetAllRoomRequest) ([]*model.RoomEntity, error) {
	var params []interface{}
	rooms := make([]*model.RoomEntity, 0)
	
	query := "SELECT id, name, floor, type, price, status, description, created_at, updated_at FROM rooms WHERE TRUE"
	if request.Floor != nil {
		query += " AND floor = ?"
		params = append(params, request.Floor)
	}

	if request.Status != nil {
		query += " AND status = ?"
		params = append(params, request.Status)
	}

	res, err := r.db.Query(query, params...)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for res.Next() {
		room := &model.RoomEntity{}
		if err := res.Scan(
			&room.ID,
			&room.Name,
			&room.Floor,
			&room.Type,
			&room.Price,
			&room.Status,
			&room.Description,
			&room.CreatedAt,
			&room.UpdatedAt,
		); err != nil {
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepoImpl) Update(room *model.RoomEntity) error {
	query := "UPDATE rooms SET"
	if room.Floor != 0 {
		query += " floor = :floor,"
	}

	if room.Price != 0 {
		query += " price = :price,"
	}

	if room.Status != 0 {
		query += " status = :status,"
	}

	if room.Description != nil {
		query += " description = :description,"
	}

	query = query[:len(query)-1]
	query += " WHERE name = :name"

	_, err := r.db.NamedExec(query, room)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}
