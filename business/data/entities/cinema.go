package entities

import "github.com/juddbaguio/go-saga-choreography/business/domain"

type Cinema struct {
	ID       int    `gorm:"primaryKey;column:id"`
	FirstRow string `gorm:"column:first_row"`
	LastRow  string `gorm:"column:last_row"`
	Columns  int    `gorm:"column:columns"`
}

func (c *Cinema) TableName() string {
	return "Cinema"
}

func (c *Cinema) ToDomain() domain.Cinema {
	return domain.Cinema{
		ID:       c.ID,
		FirstRow: c.FirstRow,
		LastRow:  c.LastRow,
		Columns:  c.Columns,
	}
}

type CinemaSeat struct {
	CinemaID  int    `gorm:"primaryKey;column:cinema_id;autoIncrement:false"`
	MovieID   int    `gorm:"primaryKey;column:movie_id;autoIncrement:false"`
	Date      string `gorm:"primaryKey;type:date;column:date"`
	TimeSlot  string `gorm:"primaryKey;type:varchar(50);column:time_slot"`
	SeatNo    string `gorm:"primaryKey;type:varchar(50);column:seat_no"`
	BookingID int    `gorm:"type:int;column:booking_id"`
	IsTaken   bool   `gorm:"column:is_taken"`
}

func (cs *CinemaSeat) TableName() string {
	return "Cinema_Seat"
}

type CinemaMovie struct {
	CinemaID         int    `gorm:"primaryKey;column:cinema_id;autoIncrement:false"`
	MovieID          int    `gorm:"primaryKey;column:movie_id;autoIncrement:false"`
	StartAiringDate  string `gorm:"type:date;column:start_airing_date"`
	EndAiringDate    string `gorm:"type:date;column:end_airing_date"`
	StartingTimeSlot string `gorm:"column:starting_time_slot"` // 24-hour format
	Interval         int    `gorm:"column:interval"`
}

func (cm *CinemaMovie) TableName() string {
	return "Cinema_Movie"
}
