package main

import "github.com/gorilla/mux"

type CardImageBytes struct {
	EntityID   string `json:"entityid"`
	ImageBytes []byte `json:"imagebytes"`
}

type CardImageRequest struct {
	EntityID string `json:"entityid"`
}

type CardImageBatchRequest struct {
	Cards []CardImageRequest `json:"cards"`
}

type CardBytesBatch struct {
	Images []CardImageBytes `json:"images"`
}

type PostAdvertisement struct {
	UserID            string `json:"userid"`
	IsSelling         string `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type PostAdvertisementResult struct {
	AdvertisementPosted bool   `json:"advertisementposted"`
	ID                  string `json:"id"`
	Message             string `json:"message"`
}

type UpdateAdvertisement struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         string `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UpdateAdvertisementResult struct {
	AdvertisementUpdated bool   `json:"advertisementupdated"`
	Message              string `json:"message"`
}

type DeleteAdvertisementResult struct {
	AdvertisementDeleted bool   `json:"advertisementdeleted"`
	AdvertisementID      string `json:"id"`
	Message              string `json:"message"`
}

type GetAdvertisementResult struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	Message           string `json:"message"`
	//ImageBytes        []byte `json:"imagebytes"`
}

type GetAdvertisementsResult struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	ImageBytes        []byte `json:"imagebytes"`
}

type AdvertisementList struct {
	Advertisements []GetAdvertisementsResult `json:"advertisements"`
}

type GetUserAdvertisementResult struct {
	AdvertisementID   string `json:"advertisementid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	ImageBytes        []byte `json:"imagebytes"`
}

type UserAdvertisementList struct {
	UserAdvertisements []GetUserAdvertisementResult `json:"useradvertisements"`
}

type TypeAdvertisementList struct {
	TypeAdvertisements []GetAdvertisementsResult `json:"typeadvertisements"`
}

type DeleteUserAdvertisementResult struct {
	AdvertisementsDeleted bool `json:"advertisementsdeleted"`
	// ---- Maybe an array of AD ID's that were deleted? ----
	Message string `json:"message"`
}

type GetTextbookAdvertisementsResult struct {
	AdvertisementID   string `json:"advertisementid"`
	UserID            string `json:"userid"`
	Isselling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	TextbookID        string `json:"textbookid"`
	TextbookName      string `json:"textbookname"`
	Edition           string `json:"edition"`
	Quality           string `json:"quality"`
	Author            string `json:"author"`
	ModuleCode        string `json:"modulecode"`
	ImageBytes        []byte `json:"imagebytes"`
}

type TextbookAdvertisementList struct {
	Textbooks []GetTextbookAdvertisementsResult `json:"textbooks"`
}

type GetTutorAdvertisementsResult struct {
	AdvertisementID   string `json:"advertisementid"`
	UserID            string `json:"userid"`
	Isselling         bool   `json:"isselling"`
	Advertisementtype string `json:"advertisementtype"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	TutorID           string `json:"tutorid"`
	Subject           string `json:"subject"`
	Yearcompleted     string `json:"yearcompleted"`
	Venue             string `json:"venue"`
	Notesincluded     string `json:"notesincluded"`
	Terms             string `json:"terms"`
	Modulecode        string `json:"modulecode"`
	ImageBytes        []byte `json:"imagebytes"`
}

type TutorAdvertisementList struct {
	Tutors []GetTutorAdvertisementsResult `json:"tutors"`
}

type GetAccomodationAdvertisementsResult struct {
	AdvertisementID      string `json:"advertisementid"`
	UserID               string `json:"userid"`
	Isselling            bool   `json:"isselling"`
	Advertisementtype    string `json:"advertisementtype"`
	Price                string `json:"price"`
	Description          string `json:"description"`
	AccomodationID       string `json:"accomodationid"`
	Accomodationtypecode string `json:"accomodationtypecode"`
	Location             string `json:"location"`
	Distancetocampus     string `json:"distancetocampus"`
	InsitutionName       string `json:"institutionname"`
	ImageBytes           []byte `json:"imagebytes"`
}

type AccomodationAdvertisementList struct {
	Accomodations []GetAccomodationAdvertisementsResult `json:"accomodations"`
}

type GetNoteAdvertisementsResult struct {
	AdvertisementID   string `json:"advertisementid"`
	UserID            string `json:"userid"`
	Isselling         bool   `json:"isselling"`
	Advertisementtype string `json:"advertisementtype"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	NoteID            string `json:"noteid"`
	ModuleCode        string `json:"modulecode"`
	ImageBytes        []byte `json:"imagebytes"`
}

type NoteAdvertisementList struct {
	Notes []GetNoteAdvertisementsResult `json:"notes"`
}

type Textbook struct {
	ModuleCode string `json:"modulecode"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type TextbookResult struct {
	TextbookAdded bool   `json:"textbookadded"`
	TextbookID    string `json:"id"`
	Message       string `json:"message"`
}

type UpdateTextbook struct {
	TextbookID string `json:"id"`
	ModuleCode string `json:"modulecode"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type UpdateTextbookResult struct {
	TextbookUpdated bool   `json:"textbookupdated"`
	Message         string `json:"message"`
}

type TextbookFilter struct {
	ModuleCode string `json:"modulecode"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type TextbookFilterResult struct {
	ModuleCode string `json:"modulecode"`
	ID         string `'json:"id"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type TextbookList struct {
	Textbooks []TextbookFilterResult `json:"textbooks"`
}

type DeleteTextbookResult struct {
	TextbookDeleted bool   `json:"Textbookdeleted"`
	TextbookID      string `json:"id"`
	Message         string `json:"message"`
}

type Note struct {
	ModuleCode string `json:"modulecode"`
}

type NoteResult struct {
	NoteAdded bool   `json:"noteadded"`
	NoteID    string `json:"id"`
	Message   string `json:"message"`
}

type UpdateNote struct {
	NoteID     string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type UpdateNoteResult struct {
	NoteUpdated bool   `json:"noteupdated"`
	Message     string `json:"message"`
}

type NoteFilter struct {
	ModuleCode string `json:"modulecode"`
}

type NoteFilterResult struct {
	ID         string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type NoteList struct {
	Notes []NoteFilterResult `json:"notes"`
}

type DeleteNoteResult struct {
	NoteDeleted bool   `json:"Notedeleted"`
	NoteID      string `json:"id"`
	Message     string `json:"message"`
}

type Tutor struct {
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type TutorResult struct {
	TutorAdded bool   `json:"tutoradded"`
	TutorID    string `json:"id"`
	Message    string `json:"message"`
}

type UpdateTutor struct {
	TutorID       string `json:"id"`
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type UpdateTutorResult struct {
	TutorUpdated bool   `json:"tutorupdated"`
	Message      string `json:"message"`
}

type TutorFilter struct {
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type TutorFilterResult struct {
	ID            string `json:"id"`
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type TutorList struct {
	Tutors []TutorFilterResult `json:"tutors"`
}

type DeleteTutorResult struct {
	TutorDeleted bool   `json:"Tutordeleted"`
	TutorID      string `json:"id"`
	Message      string `json:"message"`
}

type Accomodation struct {
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type AccomodationResult struct {
	AccomodationAdded bool   `json:"accomodationadded"`
	AccomodationID    string `json:"id"`
	Message           string `json:"message"`
}

type UpdateAccomodation struct {
	AccomodationID       string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type UpdateAccomodationResult struct {
	AccomodationUpdated bool   `json:"accomodationupdated"`
	Message             string `json:"message"`
}

type AccomodationFilter struct {
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type AccomodationFilterResult struct {
	ID                   string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type AccomodationList struct {
	Accomodations []AccomodationFilterResult `json:"accomodations"`
}

type DeleteAccomodationResult struct {
	AccomodationDeleted bool   `json:"Accomodationdeleted"`
	AccomodationID      string `json:"id"`
	Message             string `json:"message"`
}

type Server struct {
	router *mux.Router
}
type Config struct {
	CRUDHost        string
	CRUDPort        string
	FILEMANAGERHost string
	FILEMANAGERPort string
	ListenServePort string
}
