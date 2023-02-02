package usecases

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"eh-digital-shift/app/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func TestListBookings_Invoke(t *testing.T) {
	type fields struct {
		mongo *mongoServiceMock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocks   func(a args, f fields)
		want    []dto.Booking
		wantErr bool
	}{
		{
			name:   "success show a list",
			fields: fields{mongo: &mongoServiceMock{}},
			args:   args{ctx: context.Background()},
			mocks: func(a args, f fields) {
				f.mongo.On("FindAll", a.ctx, bson.D{}).Return(
					[]bson.M{{"client_id": "1", "number": 10}}, nil,
				)
			},
			want: []dto.Booking{
				{
					ClientID: "1",
					Number:   10,
				},
			},
			wantErr: false,
		},
		{
			name:   "success show a list",
			fields: fields{mongo: &mongoServiceMock{}},
			args:   args{ctx: context.Background()},
			mocks: func(a args, f fields) {
				f.mongo.On("FindAll", a.ctx, bson.D{}).Return(
					[]bson.M{}, errors.New("error getting data"),
				)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.mocks(tt.args, tt.fields)
				l := NewListBookings(tt.fields.mongo)
				got, err := l.Invoke(tt.args.ctx)
				if (err != nil) != tt.wantErr {
					t.Errorf("Invoke() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Invoke() got = %v, want %v", got, tt.want)
				}
				tt.fields.mongo.AssertExpectations(t)
			},
		)
	}
}
