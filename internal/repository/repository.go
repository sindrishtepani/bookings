package repository

import (
	"time"

	"github.com/sindrishtepani/bookings/internal/models"
)

type DataseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	HasAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
}
