package handler

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"eh-digital-shift/app/dto"
	"github.com/labstack/echo/v4"
)

func TestBooking_BookingANumber(t *testing.T) {
	type fields struct {
		BookingANumberUseCase *bookingANumberUseCaseMock
		ListBookingUseCase    *listBookingUseCaseMock
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocks   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "return OK when create a booking",
			fields: fields{
				BookingANumberUseCase: &bookingANumberUseCaseMock{},
				ListBookingUseCase:    &listBookingUseCaseMock{},
			},
			args: args{ctx: getContext(http.MethodPost, `{"client_id":"1", "number":10}`)},
			mocks: func(a args, f fields) {
				bookingDto := &dto.Booking{}
				_ = a.ctx.Bind(bookingDto)
				f.BookingANumberUseCase.On(
					"Invoke", a.ctx.Request().Context(), bookingDto,
				).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail when create a booking",
			fields: fields{
				BookingANumberUseCase: &bookingANumberUseCaseMock{},
				ListBookingUseCase:    &listBookingUseCaseMock{},
			},
			args: args{ctx: getContext(http.MethodPost, `{"client_id":"1", "number":10}`)},
			mocks: func(a args, f fields) {
				bookingDto := &dto.Booking{}
				_ = a.ctx.Bind(bookingDto)
				f.BookingANumberUseCase.On(
					"Invoke", a.ctx.Request().Context(), bookingDto,
				).Return(errors.New("this is an error"))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.mocks(tt.args, tt.fields)
				b := NewBooking(tt.fields.BookingANumberUseCase, tt.fields.ListBookingUseCase)
				if err := b.BookingANumber(tt.args.ctx); (err != nil) != tt.wantErr {
					t.Errorf("BookingANumber() error = %v, wantErr %v", err, tt.wantErr)
				}
				tt.fields.BookingANumberUseCase.AssertExpectations(t)
			},
		)
	}
}

func TestBooking_ListBooking(t *testing.T) {
	type fields struct {
		BookingANumberUseCase *bookingANumberUseCaseMock
		ListBookingUseCase    *listBookingUseCaseMock
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocks   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "Show list of bookings",
			fields: fields{
				BookingANumberUseCase: &bookingANumberUseCaseMock{},
				ListBookingUseCase:    &listBookingUseCaseMock{},
			},
			args: args{ctx: getContext(http.MethodGet, ``)},
			mocks: func(a args, f fields) {
				f.ListBookingUseCase.On(
					"Invoke", a.ctx.Request().Context(),
				).Return([]dto.Booking{}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.mocks(tt.args, tt.fields)
				b := NewBooking(tt.fields.BookingANumberUseCase, tt.fields.ListBookingUseCase)
				if err := b.ListBooking(tt.args.ctx); (err != nil) != tt.wantErr {
					t.Errorf("ListBooking() error = %v, wantErr %v", err, tt.wantErr)
				}
				tt.fields.ListBookingUseCase.AssertExpectations(t)
			},
		)
	}
}

func getContext(method string, body string) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(method, "/", nil)
	req.Body = io.NopCloser(strings.NewReader(body))
	h := http.Header{}
	h.Set("Content-Type", "application/json")
	req.Header = h
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}
