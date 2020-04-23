package main

func (s *Server) routes() {
	s.router.HandleFunc("/advertisement", s.handlepostadvertisement()).Methods("POST")     // Unit Tested
	s.router.HandleFunc("/advertisement", s.handleupdateadvertisement()).Methods("PUT")    // Unit Tested
	s.router.HandleFunc("/advertisement", s.handleremoveadvertisement()).Methods("DELETE") // Unit Tested
	s.router.HandleFunc("/advertisement", s.handlegetadvertisement()).Methods("GET")       // Unit Tested

	s.router.HandleFunc("/useradvertisements", s.handlegetuseradvertisements()).Methods("GET")
	//s.router.HandleFunc("/useradvertisements", s.handledeleteuseradvertisements()).Methods("DELETE")

	s.router.HandleFunc("/advertisementtype", s.handlegetadvertisementbytype()).Methods("GET") // Unit Tested

	//s.router.HandleFunc("/advertisementposttype", s.handlegetadvertisementbyposttype()).Methods("GET")
	//s.router.HandleFunc("/advertisements", s.handlegetalladvertisements()).Methods("GET")

	s.router.HandleFunc("/textbooks", s.handlegettextbooksbyfilter()).Methods("GET")
	s.router.HandleFunc("/textbook", s.handleaddtextbook()).Methods("POST") // Unit Tested
	//s.router.HandleFunc("/textbook", s.handleupdatetextbook()).Methods("PUT")
	s.router.HandleFunc("/textbook", s.handleremovetextbook()).Methods("DELETE") // Unit Tested

	s.router.HandleFunc("/notes", s.handlegetnotesbyfilter()).Methods("GET")
	s.router.HandleFunc("/note", s.handleaddnote()).Methods("POST") // Unit Tested
	//s.router.HandleFunc("/note", s.handleupdatenote()).Methods("PUT")
	s.router.HandleFunc("/note", s.handleremovenote()).Methods("DELETE") // Unit Tested

	s.router.HandleFunc("/tutors", s.handlegettutorsbyfilter()).Methods("GET")
	s.router.HandleFunc("/tutor", s.handleaddtutor()).Methods("POST") // Unit Tested
	//s.router.HandleFunc("/tutor", s.handleupdatetutor()).Methods("PUT")
	s.router.HandleFunc("/tutor", s.handleremovetutor()).Methods("DELETE") // Unit Tested

	s.router.HandleFunc("/accomodations", s.handlegetaccomodationsbyfilter()).Methods("GET")
	s.router.HandleFunc("/accomodation", s.handleaddaccomodation()).Methods("POST") // Unit Tested
	//s.router.HandleFunc("/accomodation", s.handleupdateaccomodation()).Methods("PUT")
	s.router.HandleFunc("/accomodation", s.handleremoveaccomodation()).Methods("DELETE") // Unit Tested

	s.router.HandleFunc("/modulecode", s.handlegetmodulecodes()).Methods("GET")

	//s.router.HandleFunc("/image", s.handlegetimage()).Methods("GET")

}
