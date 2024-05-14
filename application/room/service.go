package room

import "github.com/zakiyalmaya/hotel-management/model"

type Service interface {
	Create(room *model.RoomEntity) error
	GetByName(name string) (*model.RoomResponse, error)
	GetAll(request *model.GetAllRoomRequest) ([]*model.RoomResponse, error)
	Update(request *model.UpdateRoomRequest) error
}