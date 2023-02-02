package dto

import (
	"eh-digital-shift/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Booking struct {
	ClientID string `json:"client_id" bson:"client_id"`
	Number   int    `json:"number" bson:"number"`
}

func (b *Booking) ToBookingModel() *models.Booking {
	return &models.Booking{
		ClientID: b.ClientID,
		Number:   b.Number,
	}
}

func (b *Booking) ToMapFromMongoResult(bookings []bson.M) ([]Booking, error) {
	cRow := make(chan Booking)
	for _, booking := range bookings {
		go b.unmarshallMap(cRow, booking)
	}

	result := b.mappingResult(cRow, len(bookings))

	return result, nil
}

func (b *Booking) unmarshallMap(c chan Booking, row bson.M) {
	bsonBytes, err := bson.Marshal(row)
	if err != nil {
		return
	}
	var result Booking
	err = bson.Unmarshal(bsonBytes, &result)
	if err != nil {
		return
	}
	c <- result
}

func (b *Booking) mappingResult(c chan Booking, n int) []Booking {
	var results []Booking
	count := 0
	for i, open := <-c; open; {
		count++

		results = append(results, i)

		if count == n {
			break
		}
	}

	return results
}
