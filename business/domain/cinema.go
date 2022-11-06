package domain

type Cinema struct {
	ID       int    `json:"id"`
	FirstRow string `json:"first_row"`
	LastRow  string `json:"last_row"`
	Columns  int    `json:"columns"`
}

type CinemaSeats struct {
	CinemaID int             `json:"cinema_id"`
	MovieID  int             `json:"movie_id"`
	Schedule BookingSchedule `json:"schedule"`
	Seatlist []Seat          `json:"seats"`
}

type CinemaSeatsFilter struct {
	CinemaID int
	MovieID  int
	Schedule BookingSchedule
}
