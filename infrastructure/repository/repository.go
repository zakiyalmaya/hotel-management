package repository

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/booking"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/guest"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/room"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/user"
)

type Repositories struct {
	db          *sqlx.DB
	RedCl       *redis.Client
	RoomRepo    room.RoomRepository
	GuestRepo   guest.GuestRepository
	BookingRepo booking.BookingRepository
	UserRepo    user.UserRepository
}

func NewRespository(db *sqlx.DB, redcl *redis.Client) *Repositories {
	return &Repositories{
		db:          db,
		RedCl:       redcl,
		RoomRepo:    room.NewRoomRepository(db),
		GuestRepo:   guest.NewGuestRepository(db),
		BookingRepo: booking.NewBookingRepository(db),
		UserRepo:    user.NewUserRepository(db),
	}
}

func DBConnection(dbfile string) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", dbfile)
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		return nil
	}

	createRoomTable(db)
	createBookingTable(db)
	createGuestTable(db)
	createUserTable(db)
	return db
}

func RedisClient(redisHost, redisPort string) *redis.Client {
	option := &redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "",
		DB:       0,
	}

	redcl := redis.NewClient(option)
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		_, err := redcl.Ping(ctx).Result()
		if err == nil {
			log.Println("Connected to Redis")
			break
		}

		log.Println("Failed to connect to Redis. Retrying...")
		time.Sleep(2 * time.Second)
		if i == 9 {
			log.Panicln("Could not connect to Redis:", err)
		}
	}

	return redcl
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

func createBookingTable(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		register_number VARCHAR(255) UNIQUE NOT NULL,
		guest_id INTEGER NOT NULL,
		room_name VARCHAR(255) NOT NULL,
		check_in DATE NOT NULL,
		check_out DATE NOT NULL,
		paid_amount DECIMAL(10, 2) NOT NULL,
		payment_method INTEGER NOT NULL,
		payment_status INTEGER NOT NULL,
		additional_request TEXT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating bookings table: ", err.Error())
	}
}

func createGuestTable(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS guests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		identity_number VARCHAR(255) UNIQUE NOT NULL,
		date_of_birth DATE NOT NULL,
		phone_number VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating guests table: ", err.Error())
	}
}

func createUserTable(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		username VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating users table: ", err.Error())
	}
}
