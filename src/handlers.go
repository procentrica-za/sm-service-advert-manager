package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)



func (s *Server) handlepostadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handlePostAdvertisement Has Been Called!")
		postAdvertisement := PostAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&postAdvertisement)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Could not read body of request into proper JSON format for posting an advertisement.\n Please check that your data is in the correct format.")
			return
		}
		requestByte, _ := json.Marshal(postAdvertisement)
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/advertisement", "application/json", bytes.NewBuffer(requestByte))
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

		defer req.Body.Close()
		var postAdvertisementResponse PostAdvertisementResult

		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&postAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding post Advertisement response ")
			return
		}
		js, jserr := json.Marshal(postAdvertisementResponse)
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

func (s *Server) handleupdateadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateAdvertisement := UpdateAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&updateAdvertisement)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		client := &http.Client{}
		// Create request
		requestByte, _ := json.Marshal(updateAdvertisement)
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/advertisement", bytes.NewBuffer(requestByte))
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
		defer resp.Body.Close()

		var updateAdvertisementResponse UpdateAdvertisementResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding update Advertisement response ")
			return
		}
		js, jserr := json.Marshal(updateAdvertisementResponse)
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

func (s *Server) handledeleteuseradvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query().Get("id")
		if userid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "UserID not properly provided in URL")
			fmt.Println("UserID not properly provided in URL")
			return
		}
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/useradvertisements?id="+userid, nil)
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
		var deleteUserAdvertisementResponse DeleteUserAdvertisementResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteUserAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Advertisement response ")
			return
		}
		js, jserr := json.Marshal(deleteUserAdvertisementResponse)
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

func (s *Server) handleremoveadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		advertisementid := r.URL.Query().Get("id")
		if advertisementid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementID not properly provided in URL")
			fmt.Println("AdvertisementID not properly provided in URL")
			return
		}
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/advertisement?id="+advertisementid, nil)
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
		var deleteAdvertisementResponse DeleteAdvertisementResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding delete Advertisement response ")
			return
		}
		js, jserr := json.Marshal(deleteAdvertisementResponse)
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

func (s *Server) handlegetuseradvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query().Get("id")
		if userid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "User ID not properly provided in URL")
			fmt.Println("User ID not proplery provided in URL")
			return
		}
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/useradvertisements?id=" + userid)
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
		defer req.Body.Close()
		
		
		getUserAdvertisementResponse := UserAdvertisementList{}
		getUserAdvertisementResponse.UserAdvertisements = []GetUserAdvertisementResult{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getUserAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get User Advertisement response ")
			return
		}
		js, jserr := json.Marshal(getUserAdvertisementResponse)
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

func (s *Server) handlegetadvertisementbytype() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		advertisementtype := r.URL.Query().Get("adverttype")
		if advertisementtype == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementID not properly provided in URL")
			fmt.Println("AdvertisementID not properly provided in URL")
			return
		}
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/advertisementtype?adverttype=" + advertisementtype)
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
		defer req.Body.Close()
		getTypeAdvertisementResponse := TypeAdvertisementList{}
		getTypeAdvertisementResponse.TypeAdvertisements = []GetAdvertisementsResult{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getTypeAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}
		js, jserr := json.Marshal(getTypeAdvertisementResponse)
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

func (s *Server) handlegetadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		advertisementid := r.URL.Query().Get("id")
		if advertisementid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "AdvertisementID not properly provided in URL")
			fmt.Println("AdvertisementID not properly provided in URL")
			return
		}
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/advertisement?id=" + advertisementid)
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
		defer req.Body.Close()
		var getAdvertisementResponse GetAdvertisementResult
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getAdvertisementResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}
		js, jserr := json.Marshal(getAdvertisementResponse)
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

func (s *Server) handlegetalladvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/advertisements")
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
		defer req.Body.Close()
		getAdvertisementList := AdvertisementList{}
		getAdvertisementList.Advertisements = []GetAdvertisementsResult{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getAdvertisementList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}
		js, jserr := json.Marshal(getAdvertisementList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from Pizza List Result...")
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleaddtextbook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
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
		// define a new client to send an http request.	
		client := &http.Client{}
		// create a new http GET request to the crud and send it the JSON body that was sent to this service.
		req, respErr := http.NewRequest("GET", "http://" + config.CRUDHost + ":" + config.CRUDPort + "/textbooks", r.Body)

		// check for any CRUD errors (CRUD being down etc)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve textbook information")
			return
		}

		// Request the response back from the CRUD service 
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		// Error check all possible status codes from the response received from the crud.
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post Textbook to the CRUD service")
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
		// define textbook List that is going to be sent back from the crud.s
		textbookList := TextbookList{}
		textbookList.Textbooks = []TextbookFilterResult{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&textbookList)
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
	return func(w http.ResponseWriter, r *http.Request){
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
		// define a new client to send an http request.	
		client := &http.Client{}
		// create a new http GET request to the crud and send it the JSON body that was sent to this service.
		req, respErr := http.NewRequest("GET", "http://" + config.CRUDHost + ":" + config.CRUDPort + "/notes", r.Body)

		// check for any CRUD errors (CRUD being down etc)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve Notes information")
			return
		}

		// Request the response back from the CRUD service 
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		// Error check all possible status codes from the response received from the crud.
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post note to the CRUD service")
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
		// define notes List that is going to be sent back from the crud.s
		noteList := NoteList{}
		noteList.Notes = []NoteFilterResult{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&noteList)
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
	return func(w http.ResponseWriter, r *http.Request){
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
		// define a new client to send an http request.	
		client := &http.Client{}
		// create a new http GET request to the crud and send it the JSON body that was sent to this service.
		req, respErr := http.NewRequest("GET", "http://" + config.CRUDHost + ":" + config.CRUDPort + "/tutors", r.Body)

		// check for any CRUD errors (CRUD being down etc)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve Tutors information")
			return
		}

		// Request the response back from the CRUD service 
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		// Error check all possible status codes from the response received from the crud.
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post Tutors to the CRUD service")
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
		// define tutor List that is going to be sent back from the crud.s
		tutorList := TutorList{}
		tutorList.Tutors = []TutorFilterResult{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&tutorList)
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
	return func(w http.ResponseWriter, r *http.Request){
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
		// define a new client to send an http request.	
		client := &http.Client{}
		// create a new http GET request to the crud and send it the JSON body that was sent to this service.
		req, respErr := http.NewRequest("GET", "http://" + config.CRUDHost + ":" + config.CRUDPort + "/accomodations", r.Body)

		// check for any CRUD errors (CRUD being down etc)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve Accomodation information")
			return
		}

		// Request the response back from the CRUD service 
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		// Error check all possible status codes from the response received from the crud.
		if resp.StatusCode != 200 {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to request post Accomodation to the CRUD service")
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
		// define accomodation List that is going to be sent back from the crud.s
		accomodationList := AccomodationList{}
		accomodationList.Accomodations = []AccomodationFilterResult{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&accomodationList)
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