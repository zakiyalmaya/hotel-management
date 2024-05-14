package constant

type RoomStatus int

const (
	RoomStatusAvailable  RoomStatus = iota + 1
	RoomStatusBooked
	RoomStatusMaintenance
	RoomStatusUnavailable
)

var mapRoomStatus = map[RoomStatus]string{
	RoomStatusAvailable:   "Available",
	RoomStatusBooked:      "Booked",
	RoomStatusMaintenance: "Maintenance",
	RoomStatusUnavailable: "Unavailable",
}

func (r RoomStatus) Enum() string {
	if val, ok := mapRoomStatus[r]; ok {
		return val
	}

	return "Unknown"
}
