package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/room"
)

type Repositories struct {
	db       *sqlx.DB
	RoomRepo room.RoomRepository
}

func NewRespository(db *sqlx.DB) *Repositories {
	return &Repositories{
		db:       db,
		RoomRepo: room.NewRoomRepository(db),
	}
}

func DBConnection() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "./hotel.db")
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		return nil
	}

	createRoomTable(db)
	return db
}

func createRoomTable(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS rooms (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) UNIQUE NOT NULL,
		floor INTEGER NOT NULL,
		type VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		status INTEGER NOT NULL,
		description TEXT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating rooms table: ", err.Error())
	}
}
