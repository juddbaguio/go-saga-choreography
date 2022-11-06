package domain

import "fmt"

type Seat struct {
	Row     string `json:"row"`
	Column  int    `json:"column"`
	IsTaken bool   `json:"is_taken"`
}

type SeatMap map[string]Seat

func (s Seat) GenerateSeatNumber() string {
	return fmt.Sprintf("%s%d", s.Row, s.Column)
}
