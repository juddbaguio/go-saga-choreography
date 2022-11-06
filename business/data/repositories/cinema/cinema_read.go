package cinema

import (
	"strconv"

	"github.com/juddbaguio/go-saga-choreography/business/data/entities"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"gorm.io/gorm"
)

func (c *Repository) GetCinemaList() (*[]domain.Cinema, error) {
	var cinemaList *[]domain.Cinema = &[]domain.Cinema{}
	var queryCinemaList *[]entities.Cinema = &[]entities.Cinema{}
	if err := c.dbConn.Find(queryCinemaList).Error; err != nil {
		return nil, err
	}

	for _, cinema := range *queryCinemaList {
		*cinemaList = append(*cinemaList, cinema.ToDomain())
	}

	return cinemaList, nil
}

func (c *Repository) GetCinemaById(cinemaId int) (*domain.Cinema, error) {
	var queryCinema *entities.Cinema = &entities.Cinema{}

	if err := c.dbConn.Where("id = ?", cinemaId).First(queryCinema).Error; err != nil {
		return nil, err
	}

	return &domain.Cinema{
		ID:       queryCinema.ID,
		FirstRow: queryCinema.FirstRow,
		LastRow:  queryCinema.LastRow,
		Columns:  queryCinema.Columns,
	}, nil
}

func (c *Repository) GetCinemaSeats(filter domain.CinemaSeatsFilter) (domain.SeatMap, error) {
	var querySeatList *[]entities.CinemaSeat = &[]entities.CinemaSeat{}

	if err := c.dbConn.Where(&entities.CinemaSeat{
		CinemaID: filter.CinemaID,
		MovieID:  filter.MovieID,
		Date:     filter.Schedule.Date,
		TimeSlot: filter.Schedule.TimeSlot,
	}).Find(querySeatList).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	seatMap := make(domain.SeatMap)
	for _, cinemaSeat := range *querySeatList {
		parsedInt, _ := strconv.Atoi(string(cinemaSeat.SeatNo[1]))
		seatMap[cinemaSeat.SeatNo] = domain.Seat{
			Row:     string(cinemaSeat.SeatNo[0]),
			Column:  parsedInt,
			IsTaken: cinemaSeat.IsTaken,
		}
	}

	return seatMap, nil
}
