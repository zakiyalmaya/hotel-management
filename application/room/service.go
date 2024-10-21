package room

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=RoomService --output=../mocks --outpkg=mocks
type RoomService interface {
	Create(room *model.RoomEntity) error
	GetByName(name string) (*model.RoomResponse, error)
	GetAll(request *model.GetAllRoomRequest) ([]*model.RoomResponse, error)
	Update(request *model.UpdateRoomRequest) error
}