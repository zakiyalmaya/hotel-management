package room

import (
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/model"
)

type roomSvcImpl struct {
	repos *repository.Repositories
}

func NewRoomServiceImpl(repos *repository.Repositories) RoomService {
	return &roomSvcImpl{repos: repos}
}

func (r *roomSvcImpl) Create(room *model.RoomEntity) error {
	err := r.repos.RoomRepo.Create(room)
	if err != nil {
		return err
	}

	return nil
}

func (r *roomSvcImpl) GetByName(name string) (*model.RoomResponse, error) {
	room, err := r.repos.RoomRepo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return &model.RoomResponse{
		ID:          room.ID,
		Name:        room.Name,
		Floor:       room.Floor,
		Type:        room.Type,
		Price:       room.Price,
		Status:      room.Status.Enum(),
		Description: room.Description,
	}, nil
}

func (r *roomSvcImpl) GetAll(request *model.GetAllRoomRequest) ([]*model.RoomResponse, error) {
	rooms, err := r.repos.RoomRepo.GetAll(request)
	if err != nil {
		return nil, err
	}

	roomResponses := make([]*model.RoomResponse, len(rooms))
	for i, room := range rooms {
		roomResponses[i] = &model.RoomResponse{
			ID:          room.ID,
			Name:        room.Name,
			Floor:       room.Floor,
			Type:        room.Type,
			Price:       room.Price,
			Status:      room.Status.Enum(),
			Description: room.Description,
		}
	}

	return roomResponses, nil
}

func (r *roomSvcImpl) Update(request *model.UpdateRoomRequest) error {	
	var roomEntity model.RoomEntity
	roomEntity.Name = request.Name

	if request.Floor != nil {
		roomEntity.Floor = *request.Floor
	}

	if request.Price != nil {
		roomEntity.Price = *request.Price
	}

	if request.Status != nil {
		roomEntity.Status = constant.RoomStatus(*request.Status)
	}

	if request.Description != nil {
		roomEntity.Description = request.Description
	}

	err := r.repos.RoomRepo.Update(&roomEntity)
	if err != nil {
		return err
	}	

	return nil
}
