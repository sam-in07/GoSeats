
package booking

import "sync"

type MemoryStore struct {
	mu       sync.Mutex
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (s *MemoryStore) Book(b Booking) (Booking, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check whether this seat is already booked for this movie.
	key := b.MovieID + ":" + b.SeatID

	if _, exists := s.bookings[key]; exists {
		return Booking{}, ErrSeatAlreadybooked
	}

	s.bookings[key] = b
	return b, nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Booking

	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}

	return result
}


// func (s *MemoryStore) ListBookings(movieID string) []Booking {
// 	var result []Booking

// 	for _, b := range s.bookings {
// 		if b.MovieID == movieID {
// 			result = append(result, b)
// 		}
// 	}

// 	return result
// }
