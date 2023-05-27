package handlers

import "dewetour/repositories"

type handlerTrip struct {
	TripCountry repositories.TripRepository
}

func HandlerTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}