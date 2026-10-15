package booking

import "errors"

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatID]; exists {
		return errors.New("seat is booked yooo")
	}
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {

}
