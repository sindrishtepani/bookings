package dbrepo

import (
	"errors"
	"log"
	"time"

	"github.com/sindrishtepani/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("roomID == failure case")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 99 {
		return errors.New("roomID == failure case")
	}
	return nil
}

// HasAvailabilityByDatesByRoomID returns true if availability exists and false if it doesn't
func (m *testDBRepo) HasAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	if start.After(end) {
		return false, errors.New("failing bc start date after end date")
	}
	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	// if the start date is after 2049-12-31, then return empty slice,
	// indicating no rooms are available;
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return rooms, errors.New("some error")
	}

	if start.After(t) {
		return rooms, nil
	}

	// otherwise, put an entry into the slice, indicating that some room is
	// available for search dates
	room := models.Room{
		ID: 1,
	}
	rooms = append(rooms, room)

	return rooms, nil
}

// GetRoomByID gets a room type by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "me@here.ca" {
		return 1, "", nil
	}

	return 1, "", errors.New("didn't pass me@here.ca")
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reseravtions []models.Reservation

	return reseravtions, nil
}

func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reseravtions []models.Reservation

	return reseravtions, nil
}

func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var reseravtion models.Reservation

	return reseravtion, nil
}

func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

func (m *testDBRepo) GetRoomRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction

	return restrictions, nil
}

func (m *testDBRepo) InsertBlockForRoomRestriction(id int, startDate time.Time) error {
	return nil
}

func (m *testDBRepo) DeleteBlockByID(room_restriction_id int) error {
	return nil
}
