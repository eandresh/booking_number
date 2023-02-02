package httpServer

func (s *Server) Routes() {

	root := s.Server.Group(s.cfg.Prefix)
	root.POST("/booking", s.app.Booking.BookingANumber)
	root.GET("/booking-list", s.app.Booking.ListBooking)
}
