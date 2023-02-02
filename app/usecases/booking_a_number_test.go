package usecases

import (
	"context"
	"testing"

	"eh-digital-shift/app/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBookingANumber_Invoke(t *testing.T) {
	type fields struct {
		mongo *mongoServiceMock
	}
	type args struct {
		ctx     context.Context
		booking *dto.Booking
	}
	tests := []struct {
		name    string
		fields  fields
		mocks   func(a args, f fields)
		args    args
		wantErr bool
	}{
		{
			name:   "success booking a number",
			fields: fields{mongo: &mongoServiceMock{}},
			args: args{
				ctx: context.Background(),
				booking: &dto.Booking{
					ClientID: "1",
					Number:   10,
				},
			},
			mocks: func(a args, f fields) {
				filter := bson.D{{"number", a.booking.Number}}
				f.mongo.On("FindAll", a.ctx, filter).Return([]bson.M{}, nil)

				filter = bson.D{{"client_id", a.booking.ClientID}}
				f.mongo.On("FindAll", a.ctx, filter).Return([]bson.M{}, nil)

				f.mongo.On("Insert", a.ctx, a.booking).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Fail booking when number exist in bd",
			fields: fields{mongo: &mongoServiceMock{}},
			args: args{
				ctx: context.Background(),
				booking: &dto.Booking{
					ClientID: "1",
					Number:   10,
				},
			},
			mocks: func(a args, f fields) {
				filter := bson.D{{"number", a.booking.Number}}
				f.mongo.On("FindAll", a.ctx, filter).Return([]bson.M{{"client_id": "2", "number": 10}}, nil)

				filter = bson.D{{"client_id", a.booking.ClientID}}
				f.mongo.On("FindAll", a.ctx, filter).Return([]bson.M{}, nil)

				f.mongo.On("Insert", a.ctx, a.booking).Return(nil)
			},
			wantErr: true,
		},
		{
			name:   "Fail booking when client has a number",
			fields: fields{mongo: &mongoServiceMock{}},
			args: args{
				ctx: context.Background(),
				booking: &dto.Booking{
					ClientID: "1",
					Number:   10,
				},
			},
			mocks: func(a args, f fields) {
				filter := bson.D{{"number", a.booking.Number}}
				f.mongo.On("FindAll", a.ctx, filter).Return([]bson.M{}, nil)

				filter = bson.D{{"client_id", a.booking.ClientID}}
				f.mongo.On("FindAll", a.ctx, filter).Return([]bson.M{{"client_id": "1", "number": 11}}, nil)

				f.mongo.On("Insert", a.ctx, a.booking).Return(nil)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.mocks(tt.args, tt.fields)
				ban := NewBookingANumber(tt.fields.mongo)

				if err := ban.Invoke(tt.args.ctx, tt.args.booking); (err != nil) != tt.wantErr {
					t.Errorf("Invoke() error = %v, wantErr %v", err, tt.wantErr)
				}
				tt.fields.mongo.AssertExpectations(t)
			},
		)
	}
}
