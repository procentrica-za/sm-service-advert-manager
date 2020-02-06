package main

import "github.com/gorilla/mux"

type PostAdvertisement struct {
	UserID            string `json:"userid"`
	AdvertisementType string `json:'advertisementtype"`
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
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid`
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

type DeleteUserAdvertisementResult struct {
	AdvertisementsDeleted bool `json:"advertisementsdeleted"`
	// ---- Maybe an array of AD ID's that were deleted? ----
	Message string `json:"message"`
}

type GetAdvertisementResult struct {
	UserID            string `json:"userid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type GetAdvertisementsResult struct {
	AdvertisementID   string `json:"advertisementid"`
	UserID            string `json:"userid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type AdvertisementList struct {
	Advertisements []GetAdvertisementsResult `json:"advertisements"`
}

type GetUserAdvertisementResult struct {
	AdvertisementID   string `json:"advertisementid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UserAdvertisementList struct {
	UserAdvertisements []GetUserAdvertisementResult `json:"useradvertisements"`
}

type TypeAdvertisementList struct {
	TypeAdvertisements []GetAdvertisementsResult `json:"typeadvertisements"`
}

type Server struct {
	router *mux.Router
}
type Config struct {
	CRUDHost string
	CRUDPort string
}
