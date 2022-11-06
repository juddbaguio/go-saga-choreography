package http_booking

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
)

func (c *Container) InitRoutes(r *chi.Mux) {
	r.Post("/book", c.HandleCreateBooking)
}

func (c *Container) HandleCreateBooking(w http.ResponseWriter, r *http.Request) {
	var payload *domain.Booking = &domain.Booking{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	booking, err := c.bookingService.CreateBooking(*payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}
