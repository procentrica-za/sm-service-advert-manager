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
