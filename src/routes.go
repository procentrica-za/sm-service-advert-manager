package main

func (s *Server) routes() {
	s.router.HandleFunc("/advertisement", s.handlepostadvertisement()).Methods("POST")
	s.router.HandleFunc("/advertisement", s.handleupdateadvertisement()).Methods("PUT")
	s.router.HandleFunc("/advertisement", s.handleremoveadvertisement()).Methods("DELETE")
	s.router.HandleFunc("/advertisement", s.handlegetadvertisement()).Methods("GET")
	s.router.HandleFunc("/advertisements", s.handlegetadvertisements()).Methods("GET")
	s.router.HandleFunc("/useradvertisements", s.handlegetuseradvertisements()).Methods("GET")
	s.router.HandleFunc("/typeadvertisement", s.handlegettypeadvertisements()).Methods("GET")
	s.router.HandleFunc("/useradvertisements", s.handleremoveuseradvertisements()).Methods("DELETE")
}
