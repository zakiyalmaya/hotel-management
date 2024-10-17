package room

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=RoomRepository --output=../mocks --outpkg=mocks
type RoomRepository interface {
	Create(room *model.RoomEntity) error
	GetByName(name string) (*model.RoomEntity, error)
	GetAll(request *model.GetAllRoomRequest) ([]*model.RoomEntity, error)
	Update(room *model.RoomEntity) error
}