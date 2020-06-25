package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (s *Server) handlepostadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handlePostAdvertisement Has Been Called!")
		//get JSON payload
		fmt.Println("Gcloud test 1 remove this if you see it in console.")
		postAdvertisement := PostAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&postAdvertisement)
		//handle for bad JSON provided

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Could not read body of request into proper JSON format for posting an advertisement.\n Please check that your data is in the correct format.")
			return
		}

		//create byte array from JSON payload
		requestByte, _ := json.Marshal(postAdvertisement)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/advertisement", "application/json", bytes.NewBuffer(requestByte))

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to post an advertisement!")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post advertisement to the CRUD service")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var postAdvertisementResponse PostAdvertisementResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&postAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding post Advertisement response ")
			return
		}
		//convert struct back to JSON
		js, jserr := json.Marshal(postAdvertisementResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format! ")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload

		updateAdvertisement := UpdateAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&updateAdvertisement)
		//handle for bad JSON provided

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		client := &http.Client{}

		//create byte array from JSON payload
		requestByte, _ := json.Marshal(updateAdvertisement)

		//post to crud service
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/advertisement", bytes.NewBuffer(requestByte))

		//check for response error from CRUD service
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update advertisement")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request update advertisement to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		//close the request
		defer resp.Body.Close()

		//create new response struct
		var updateAdvertisementResponse UpdateAdvertisementResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateAdvertisementResponse)
		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding update Advertisement response ")
			return
		}
		js, jserr := json.Marshal(updateAdvertisementResponse)
		//convert struct back to JSON
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handledeleteuseradvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get User ID from URL
		userid := r.URL.Query().Get("id")
		if userid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "UserID not properly provided in URL")
			fmt.Println("UserID not properly provided in URL")
			return
		}
		client := &http.Client{}

		//post to crud service
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/useradvertisements?id="+userid, nil)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete an advertisement")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		//close the request
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request delete user advertisements to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		//create new response struct
		var deleteUserAdvertisementResponse DeleteUserAdvertisementResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteUserAdvertisementResponse)
		//handle for bad Response recieved from CRUD service
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Advertisement response ")
			return
		}
		//convert struct back to JSON
		js, jserr := json.Marshal(deleteUserAdvertisementResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetmodulecodes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/modulecode")

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve modulecode information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get modulecode data "+bodyString)
			fmt.Println("An internal error has occured whilst trying to get modulecode data " + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		moduleCodeList := ModuleCodeList{}
		moduleCodeList.Modulecodes = []ModuleCode{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&moduleCodeList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get modulecode response ")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(moduleCodeList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handleremoveadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get Advertisement ID from URL
		advertisementid := r.URL.Query().Get("id")

		//Check if Advertisement ID is null
		if advertisementid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementID not properly provided in URL")
			fmt.Println("AdvertisementID not properly provided in URL")
			return
		}
		client := &http.Client{}

		//post to crud service
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/advertisement?id="+advertisementid, nil)
		if respErr != nil {

			//check for response error of 500
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete an advertisement")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		//close the request
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request delete advertisement to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		//create new response struct
		var deleteAdvertisementResponse DeleteAdvertisementResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Advertisement response ")
			return
		}
		//convert struct back to JSON
		js, jserr := json.Marshal(deleteAdvertisementResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetuseradvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get User ID from URL

		userid := r.URL.Query().Get("id")
		advertisementType := r.URL.Query().Get("adverttype")
		resultlimit := r.URL.Query().Get("limit")

		//Check if no Advertisement type was provided in the URL
		if advertisementType == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementType not properly provided in URL")
			fmt.Println("AdvertisementType not properly provided in URL")
			return
		}
		//Check if User ID provided is null
		if userid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "User ID not properly provided in URL")
			fmt.Println("User ID not proplery provided in URL")
			return
		}

		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/useradvertisements?id=" + userid + "&adverttype=" + advertisementType + "&limit=" + resultlimit)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve user advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get a users advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get a users advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		switch {
		case advertisementType == "TXB":
			textbookAdvertList := TextbookAdvertisementList{}
			textbookAdvertList.Textbooks = []GetTextbookAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&textbookAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range textbookAdvertList.Textbooks {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				textbookAdvertList.Textbooks[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.

			if textbookAdvertList.Textbooks[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range textbookAdvertList.Textbooks {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range textbookAdvertList.Textbooks {
						if advertisement.AdvertisementID == image.EntityID {
							textbookAdvertList.Textbooks[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(textbookAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		case advertisementType == "TUT":
			tutorAdvertList := TutorAdvertisementList{}
			tutorAdvertList.Tutors = []GetTutorAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&tutorAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range tutorAdvertList.Tutors {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				tutorAdvertList.Tutors[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if tutorAdvertList.Tutors[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range tutorAdvertList.Tutors {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range tutorAdvertList.Tutors {
						if advertisement.AdvertisementID == image.EntityID {
							tutorAdvertList.Tutors[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(tutorAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		case advertisementType == "ACD":
			accomodationAdvertList := AccomodationAdvertisementList{}
			accomodationAdvertList.Accomodations = []GetAccomodationAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&accomodationAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range accomodationAdvertList.Accomodations {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				accomodationAdvertList.Accomodations[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if accomodationAdvertList.Accomodations[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range accomodationAdvertList.Accomodations {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range accomodationAdvertList.Accomodations {
						if advertisement.AdvertisementID == image.EntityID {
							accomodationAdvertList.Accomodations[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(accomodationAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		case advertisementType == "NTS":
			noteAdvertList := NoteAdvertisementList{}
			noteAdvertList.Notes = []GetNoteAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&noteAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range noteAdvertList.Notes {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				noteAdvertList.Notes[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if noteAdvertList.Notes[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range noteAdvertList.Notes {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range noteAdvertList.Notes {
						if advertisement.AdvertisementID == image.EntityID {
							noteAdvertList.Notes[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(noteAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		default:
			fmt.Println("Default Hit!")
		}
	}
}

func (s *Server) handlegetadvertisementbytype() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get Advertisement type from URL

		advertisementType := r.URL.Query().Get("adverttype")
		resultlimit := r.URL.Query().Get("limit")
		isSelling := r.URL.Query().Get("selling")
		priceFilter := r.URL.Query().Get("price")
		//Check if no Advertisement type was provided in the URL
		if advertisementType == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementType not properly provided in URL")
			fmt.Println("AdvertisementType not properly provided in URL")
			return
		}

		ModuleCodeFilter := r.URL.Query().Get("modulecode")
		NameFilter := r.URL.Query().Get("name")
		l := strings.NewReplacer(" ", "+")
		Newname := l.Replace(NameFilter)
		EditionFilter := r.URL.Query().Get("edition")
		QualityFilter := r.URL.Query().Get("quality")
		AuthorFilter := r.URL.Query().Get("author")
		SubjectFilter := r.URL.Query().Get("subject")
		YearcompletedFilter := r.URL.Query().Get("yearcompleted")
		VenueFilter := r.URL.Query().Get("venue")
		NotesincludedFilter := r.URL.Query().Get("notes")
		TermsFilter := r.URL.Query().Get("terms")
		AccomodationtypecodeFilter := r.URL.Query().Get("acdType")
		LocationFilter := r.URL.Query().Get("location")
		DistancetocampusFilter := r.URL.Query().Get("distance")
		InsitutionNameFilter := r.URL.Query().Get("institution")
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort +
			"/advertisementtype?adverttype=" + advertisementType + "&price=" + priceFilter + "&limit=" + resultlimit + "&selling=" + isSelling +
			"&modulecode=" + ModuleCodeFilter + "&name=" + Newname + "&edition=" + EditionFilter + "&quality=" + QualityFilter + "&author=" + AuthorFilter +
			"&subject=" + SubjectFilter + "&yearcompleted=" + YearcompletedFilter + "&venue=" + VenueFilter + "&notes=" + NotesincludedFilter + "&terms=" + TermsFilter +
			"&acdType=" + AccomodationtypecodeFilter + "&location=" + LocationFilter + "&distance=" + DistancetocampusFilter + "&institution=" + InsitutionNameFilter)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data: \n"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data: \n" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		switch {
		case advertisementType == "TXB":
			textbookAdvertList := TextbookAdvertisementList{}
			textbookAdvertList.Textbooks = []GetTextbookAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&textbookAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}

			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range textbookAdvertList.Textbooks {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				textbookAdvertList.Textbooks[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if textbookAdvertList.Textbooks[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range textbookAdvertList.Textbooks {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range textbookAdvertList.Textbooks {
						if advertisement.AdvertisementID == image.EntityID {
							textbookAdvertList.Textbooks[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}
			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(textbookAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		case advertisementType == "TUT":
			tutorAdvertList := TutorAdvertisementList{}
			tutorAdvertList.Tutors = []GetTutorAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&tutorAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range tutorAdvertList.Tutors {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				tutorAdvertList.Tutors[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if tutorAdvertList.Tutors[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range tutorAdvertList.Tutors {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range tutorAdvertList.Tutors {
						if advertisement.AdvertisementID == image.EntityID {
							tutorAdvertList.Tutors[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(tutorAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		case advertisementType == "ACD":
			accomodationAdvertList := AccomodationAdvertisementList{}
			accomodationAdvertList.Accomodations = []GetAccomodationAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&accomodationAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range accomodationAdvertList.Accomodations {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				accomodationAdvertList.Accomodations[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if accomodationAdvertList.Accomodations[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range accomodationAdvertList.Accomodations {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range accomodationAdvertList.Accomodations {
						if advertisement.AdvertisementID == image.EntityID {
							accomodationAdvertList.Accomodations[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(accomodationAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		case advertisementType == "NTS":
			noteAdvertList := NoteAdvertisementList{}
			noteAdvertList.Notes = []GetNoteAdvertisementsResult{}
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&noteAdvertList)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding get Advertisement response ")
				return
			}
			/* // ------- This segment is used to get images for every advertisement of the type requested as seperate calls to the filemanager service. -----------------

			for i, advertisement := range noteAdvertList.Notes {

				req, respErr := http.Get("http://" + config.FILEMANAGERHost + ":" + config.FILEMANAGERPort + "/cardimage?entityid=" + advertisement.AdvertisementID)

				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to retrieve image for advertisement -->" + advertisement.AdvertisementID)
					return
				}
				if req.StatusCode != 200 {
					w.WriteHeader(req.StatusCode)
					fmt.Fprint(w, "Request to DB can't be completed...")
					fmt.Println("Request to DB can't be completed...")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement image: "+bodyString)
					fmt.Println("An internal error has occured whilst trying to get advertisement image: " + bodyString)
					return
				}

				cardImageBytes := CardImageBytes{}
				decoder := json.NewDecoder(req.Body)
				err := decoder.Decode(&cardImageBytes)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get Advertisement Image response ")
					return
				}

				noteAdvertList.Notes[i].ImageBytes = cardImageBytes.ImageBytes

			}
			//	----------- End of code segment for single calls to the filemanager service ------------ */

			// ---------- This segment is used to get images for every advertisement of the type requested but as 1 call to the filemanager service as a batch request.
			// Instansiate object to capture advertisement ID.
			if noteAdvertList.Notes[0].AdvertisementID != "" {
				cardImageRequest := CardImageRequest{}

				cardImageBatchRequest := CardImageBatchRequest{}
				cardImageBatchRequest.Cards = []CardImageRequest{}

				for _, advertisement := range noteAdvertList.Notes {
					cardImageRequest.EntityID = advertisement.AdvertisementID
					cardImageBatchRequest.Cards = append(cardImageBatchRequest.Cards, cardImageRequest)
				}

				requestByte, _ := json.Marshal(cardImageBatchRequest)
				req, respErr = http.Post("http://"+config.FILEMANAGERHost+":"+config.FILEMANAGERPort+"/cardimagebatch", "application/json", bytes.NewBuffer(requestByte))

				//check for response error of 500
				if respErr != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, respErr.Error())
					fmt.Println("Error in communication with CRUD service endpoint for request to get file details")
					return
				}
				if req.StatusCode != 200 {
					fmt.Println("Request to DB can't be completed to get file details")
				}
				if req.StatusCode == 500 {
					w.WriteHeader(500)
					bodyBytes, err := ioutil.ReadAll(req.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					fmt.Fprintf(w, "Database error occured upon retrieval: "+bodyString)
					fmt.Println("Database error occured upon retrieval: " + bodyString)
					return
				}

				defer req.Body.Close()

				cardimages := CardBytesBatch{}
				cardimages.Images = []CardImageBytes{}

				decoder = json.NewDecoder(req.Body)
				err = decoder.Decode(&cardimages)
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprint(w, err.Error())
					fmt.Println("Error occured in decoding get batch images response ")
					return
				}

				// This double for loop might be a bit slow could look at a better implementation.
				for _, image := range cardimages.Images {
					for i, advertisement := range noteAdvertList.Notes {
						if advertisement.AdvertisementID == image.EntityID {
							noteAdvertList.Notes[i].ImageBytes = image.ImageBytes
							break
						}
					}
				}
			}

			//	----------- End of code segment for batch calls to the filemanager service ------------ */

			//convert struct back to JSON
			js, jserr := json.Marshal(noteAdvertList)
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, jserr.Error())
				fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
				return
			}

			//return success back to Front-End user
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		default:
			fmt.Println("Default Hit!")
		}

	}
}

func (s *Server) handlegetadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get Advertisement ID from URL

		advertisementid := r.URL.Query().Get("id")

		//Check if no Advertisement ID was provided in the URL
		if advertisementid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementID not properly provided in URL")
			fmt.Println("AdvertisementID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/advertisement?id=" + advertisementid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var getAdvertisementResponse GetAdvertisementResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(getAdvertisementResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetalladvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get request from URL
		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/advertisements")
		//handle for bad Response recieved from CRUD service
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, respErr.Error())
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()
		getAdvertisementList := AdvertisementList{}

		//create new response struct for JSON list
		getAdvertisementList.Advertisements = []GetAdvertisementsResult{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getAdvertisementList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(getAdvertisementList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from Pizza List Result...")
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetadvertisementbyposttype() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get Ad post type from URL
		advertisementposttype := r.URL.Query().Get("advertposttype")

		//Check if Advertisement Post Type is not provided in URL
		if advertisementposttype == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Post type not properly provided in URL")
			fmt.Println("Post type not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/advertisementposttype?advertposttype=" + advertisementposttype)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		getTypeAdvertisementResponse := TypeAdvertisementList{}
		getTypeAdvertisementResponse.TypeAdvertisements = []GetAdvertisementsResult{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getTypeAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}
		//convert struct back to JSON
		js, jserr := json.Marshal(getTypeAdvertisementResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleaddtextbook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textbook := Textbook{}
		err := json.NewDecoder(r.Body).Decode(&textbook)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Could not read body of request into proper JSON format for posting a textbook.\n Please check that your data is in the correct format.")
			return
		}
		requestByte, _ := json.Marshal(textbook)
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/textbook", "application/json", bytes.NewBuffer(requestByte))
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to post a textbook!")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post Textbook to the CRUD service")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		defer req.Body.Close()
		var addTextbookResponse TextbookResult

		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&addTextbookResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding post Textbook response ")
			return
		}
		js, jserr := json.Marshal(addTextbookResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format! ")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdatetextbook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateTextbook := UpdateTextbook{}
		err := json.NewDecoder(r.Body).Decode(&updateTextbook)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		client := &http.Client{}
		// Create request
		requestByte, _ := json.Marshal(updateTextbook)
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/textbook", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update textbook")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request update textbook to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		defer resp.Body.Close()

		var updateTextbookResponse UpdateTextbookResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateTextbookResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding update Textbook response ")
			return
		}
		js, jserr := json.Marshal(updateTextbookResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegettextbooksbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textbookfilter := TextbookFilter{}
		textbookfilter.ModuleCode = r.URL.Query().Get("modulecode")
		textbookfilter.Name = r.URL.Query().Get("name")
		textbookfilter.Edition = r.URL.Query().Get("edition")
		textbookfilter.Quality = r.URL.Query().Get("quality")
		textbookfilter.Author = r.URL.Query().Get("author")

		// create a new http GET request to the crud and send it the filter headers that was sent to this service.
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/textbooks?modulecode=" + textbookfilter.ModuleCode + "&name=" + textbookfilter.Name + "&edition=" + textbookfilter.Edition + "&quality=" + textbookfilter.Quality + "&author=" + textbookfilter.Author)

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, respErr.Error())
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		// define textbook List that is going to be sent back from the crud.s
		textbookList := TextbookList{}
		textbookList.Textbooks = []TextbookFilterResult{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&textbookList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get textbooklist response ")
			return
		}
		js, jserr := json.Marshal(textbookList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremovetextbook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textbookid := r.URL.Query().Get("id")
		if textbookid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "TextbookID not properly provided in URL")
			fmt.Println("TextbookID not properly provided in URL")
			return
		}
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/textbook?id="+textbookid, nil)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete an textbook")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request delete textbook to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		var deleteTextbookResponse DeleteTextbookResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteTextbookResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Textbook response ")
			return
		}
		js, jserr := json.Marshal(deleteTextbookResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

// ---- NOTES ----

func (s *Server) handleaddnote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		note := Note{}
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Could not read body of request into proper JSON format for posting a note.\n Please check that your data is in the correct format.")
			return
		}
		requestByte, _ := json.Marshal(note)
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/note", "application/json", bytes.NewBuffer(requestByte))
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to post a note!")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post Note to the CRUD service")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		defer req.Body.Close()
		var addNoteResponse NoteResult

		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&addNoteResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding post Note response ")
			return
		}
		js, jserr := json.Marshal(addNoteResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format! ")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdatenote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateNote := UpdateNote{}
		err := json.NewDecoder(r.Body).Decode(&updateNote)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		client := &http.Client{}
		// Create request
		requestByte, _ := json.Marshal(updateNote)
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/note", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update note")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request update note to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		defer resp.Body.Close()

		var updateNoteResponse UpdateNoteResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateNoteResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding update Note response ")
			return
		}
		js, jserr := json.Marshal(updateNoteResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetnotesbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notefilter := NoteFilter{}

		notefilter.ModuleCode = r.URL.Query().Get("modulecode")

		// create a new http GET request to the crud and send it the filters in the hedaer that was sent to this service.
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/notes?modulecode=" + notefilter.ModuleCode)

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, respErr.Error())
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()
		// define notes List that is going to be sent back from the crud.s
		noteList := NoteList{}
		noteList.Notes = []NoteFilterResult{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&noteList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get notelist response ")
			return
		}
		js, jserr := json.Marshal(noteList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremovenote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteid := r.URL.Query().Get("id")
		if noteid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Note ID not properly provided in URL")
			fmt.Println("Note ID not properly provided in URL")
			return
		}
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/note?id="+noteid, nil)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete a note")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request delete note to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		var deleteNoteResponse DeleteNoteResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteNoteResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Note response ")
			return
		}
		js, jserr := json.Marshal(deleteNoteResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

// ---- TUTORS ----

func (s *Server) handleaddtutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tutor := Tutor{}
		err := json.NewDecoder(r.Body).Decode(&tutor)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Could not read body of request into proper JSON format for posting a tutor.\n Please check that your data is in the correct format.")
			return
		}
		requestByte, _ := json.Marshal(tutor)
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/tutor", "application/json", bytes.NewBuffer(requestByte))
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to post a tutor!")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post tutor to the CRUD service")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		defer req.Body.Close()
		var addTutorResponse TutorResult

		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&addTutorResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding post Tutor response ")
			return
		}
		js, jserr := json.Marshal(addTutorResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format! ")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdatetutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateTutor := UpdateTutor{}
		err := json.NewDecoder(r.Body).Decode(&updateTutor)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		client := &http.Client{}
		// Create request
		requestByte, _ := json.Marshal(updateTutor)
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/tutor", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update tutor")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request update tutor to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		defer resp.Body.Close()

		var updateTutorResponse UpdateTutorResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateTutorResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding update Tutor response ")
			return
		}
		js, jserr := json.Marshal(updateTutorResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegettutorsbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tutorfilter := TutorFilter{}

		tutorfilter.ModuleCode = r.URL.Query().Get("modulecode")
		tutorfilter.Subject = r.URL.Query().Get("subject")
		tutorfilter.YearCompleted = r.URL.Query().Get("yearcompleted")
		tutorfilter.Venue = r.URL.Query().Get("venue")
		tutorfilter.NotesIncluded = r.URL.Query().Get("notesincluded")
		tutorfilter.Terms = r.URL.Query().Get("terms")

		// create a new http GET request to the crud and send it the filters in the hedaer that was sent to this service.
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/tutors?modulecode=" + tutorfilter.ModuleCode + "&subject=" + tutorfilter.Subject + "&yearcompleted=" + tutorfilter.YearCompleted + "&venue=" + tutorfilter.Venue + "&notesincluded=" + tutorfilter.NotesIncluded + "&terms=" + tutorfilter.Terms)

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, respErr.Error())
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		// define tutor List that is going to be sent back from the crud.s
		tutorList := TutorList{}
		tutorList.Tutors = []TutorFilterResult{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&tutorList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get tutorlist response ")
			return
		}
		js, jserr := json.Marshal(tutorList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremovetutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tutorid := r.URL.Query().Get("id")
		if tutorid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "tutor ID not properly provided in URL")
			fmt.Println("tutor ID not properly provided in URL")
			return
		}
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/tutor?id="+tutorid, nil)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete a tutor")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request delete tutor to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		var deleteTutorResponse DeleteTutorResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteTutorResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Tutor response ")
			return
		}
		js, jserr := json.Marshal(deleteTutorResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

// ---- ACCOMODATION ----

func (s *Server) handleaddaccomodation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accomodation := Accomodation{}
		err := json.NewDecoder(r.Body).Decode(&accomodation)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Could not read body of request into proper JSON format for posting an accomodation.\n Please check that your data is in the correct format.")
			return
		}
		requestByte, _ := json.Marshal(accomodation)
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/accomodation", "application/json", bytes.NewBuffer(requestByte))
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to post a Accomodation!")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post Accomodation to the CRUD service")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}

		defer req.Body.Close()
		var addAccomodationResponse AccomodationResult

		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&addAccomodationResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding post Accomodation response ")
			return
		}
		js, jserr := json.Marshal(addAccomodationResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format! ")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateaccomodation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateAccomodation := UpdateAccomodation{}
		err := json.NewDecoder(r.Body).Decode(&updateAccomodation)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		client := &http.Client{}
		// Create request
		requestByte, _ := json.Marshal(updateAccomodation)
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/accomodation", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update accomodation")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request update accomodation to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		defer resp.Body.Close()

		var updateAccomodationResponse UpdateAccomodationResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateAccomodationResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding update Accomodation response ")
			return
		}
		js, jserr := json.Marshal(updateAccomodationResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetaccomodationsbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accomodationfilter := AccomodationFilter{}

		accomodationfilter.AccomodationTypeCode = r.URL.Query().Get("accomodationtypecode")
		accomodationfilter.InstitutionName = r.URL.Query().Get("institutionname")
		accomodationfilter.Location = r.URL.Query().Get("location")
		accomodationfilter.DistanceToCampus = r.URL.Query().Get("distancetocampus")

		// create a new http GET request to the crud and send it the filters in the hedaer that was sent to this service.
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/accomodations?accomodationtypecode=" + accomodationfilter.AccomodationTypeCode + "&institutionname=" + accomodationfilter.InstitutionName + "&location=" + accomodationfilter.Location + "&distancetocampus=" + accomodationfilter.DistanceToCampus)

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, respErr.Error())
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()
		// define accomodation List that is going to be sent back from the crud.s
		accomodationList := AccomodationList{}
		accomodationList.Accomodations = []AccomodationFilterResult{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&accomodationList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get accomodationlist response ")
			return
		}
		js, jserr := json.Marshal(accomodationList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremoveaccomodation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accomodationid := r.URL.Query().Get("id")
		if accomodationid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "accomodation ID not properly provided in URL")
			fmt.Println("accomodation ID not properly provided in URL")
			return
		}
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/accomodation?id="+accomodationid, nil)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete a accomodation")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request delete accomodation to the CRUD service")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed with request: "+bodyString)
			fmt.Println("Request to DB can't be completed with request: " + bodyString)
			return
		}
		var deleteAccomodationResponse DeleteAccomodationResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteAccomodationResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Accomodation response ")
			return
		}
		js, jserr := json.Marshal(deleteAccomodationResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

/*func (s *Server) handlegetimage(id) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		advertisementid := r.URL.Query().Get("id")
		if advertisementid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "accomodation ID not properly provided in URL")
			fmt.Println("accomodation ID not properly provided in URL")
			return
		}
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/image?id=" + advertisementid)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, respErr.Error())
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get image for advertisement"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get image for advertisement" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()
	}
}*/
