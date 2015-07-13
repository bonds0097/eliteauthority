package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	EDSC_URL          string = "http://edstarcoordinator.com/api.asmx/GetSystems"
	EDSC_TIME         string = "2006-01-02 15:04:05"
	EDSC_UPDATE_FILE  string = "edsc_update.txt"
	EDSC_DEFAULT_TIME string = "2000-01-01 00:00:00"

	MIN_SEARCH_LEN int = 3
)

// A Star system with known 3D coordinates.
type System struct {
	Name        string    `bson:"_id" json:"name"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coord,omitempty"`
}

// A search for a star system.
type SystemSearch struct {
	Name string `json:"name"`
}

// Result of a system search.
type SearchResult struct {
	Systems []string `json:"systems"`
}

// A request to the EDSC API.
type EdscRequest struct {
	Data struct {
		Ver        int `json:"ver"`
		OutputMode int `json:"outputmode"`
		Filter     struct {
			Knownstatus int    `json:"knownstatus"`
			CR          int    `json:"cr"`
			Date        string `json:"date"`
		} `json:"filter"`
	} `json:"data"`
}

// A response from the EDSC API.
type EdscResponse struct {
	Data struct {
		Date    string   `json:"date"`
		Systems []System `json:"systems"`
	} `json:"d"`
}

// Searches for a star system and returns an array of named matches.
func findSystems(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	var search SystemSearch
	err := json.NewDecoder(r.Body).Decode(&search)

	// Validate request.
	if err != nil || search.Name == "" {
		return &ApiError{"Could not parse search request. Make sure to include 'name' field.",
			http.StatusBadRequest}
	}

	if len(search.Name) < MIN_SEARCH_LEN {
		return &ApiError{"Minimum search length is three (3) characters.", http.StatusBadRequest}
	}

	// Perform database search.
	db := dbMaster.spawnDb()
	defer db.stop()

	results, err := db.findSystems(search.Name)
	if err != nil {
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// Send back data and return.
	json.NewEncoder(w).Encode(&SearchResult{results})

	return
}

// Gets an updated list of systems and coordinates from the EDSC API.
func getEdscUpdate() (err error) {
	// Get time of last update.
	lastUpdate := lastEdscUpdateTime()

	// Create request body.
	var edscRequest EdscRequest
	edscRequest.Data.Ver = 2
	edscRequest.Data.OutputMode = 2
	edscRequest.Data.Filter.CR = 5
	edscRequest.Data.Filter.Date = lastUpdate

	requestBody, err := json.Marshal(&edscRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Make API request.
	client := &http.Client{}
	req, err := http.NewRequest("POST", EDSC_URL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println("There was an error accessing the EDSC API.")
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
		} else {
			log.Printf("Response Status: %s\n", resp.Status)
		}

		return
	}
	defer resp.Body.Close()

	// Parse returned JSON.
	body, _ := ioutil.ReadAll(resp.Body)

	var edscResponse EdscResponse
	err = json.Unmarshal(body, &edscResponse)
	if err != nil {
		log.Printf("There was an error parsing the response from EDSC.\nError: %s", err)
	}

	// Add systems to database.
	db := dbMaster.spawnDb()
	defer db.stop()

	if len(edscResponse.Data.Systems) > 0 {
		err = db.addSystems(edscResponse.Data.Systems)
		if err != nil {
			log.Fatalf("There was an error inserting systems into the database.\nError: %s", err.Error())
			return
		}
	}

	// Update time of update.
	err = updateEdscUpdateTime(edscResponse.Data.Date)
	if err != nil {
		log.Fatalf("There was an error updating the EDSC Update time.\nError: %s", err.Error())
		return
	}

	// If all went well, log a successful operation.
	log.Println("Successfully retrieved systems update from EDSC and inserted into database.")

	return
}

// Returns the last time we received an update from EDSC.
func lastEdscUpdateTime() (t string) {
	b, err := ioutil.ReadFile(EDSC_UPDATE_FILE)

	// If we can't read the file, we'll return a default time.
	if err != nil {
		t = EDSC_DEFAULT_TIME
		return
	}

	// Make sure the data is a valid EDSC time.
	t = string(b)
	if _, err = time.Parse(EDSC_TIME, t); err != nil {
		t = EDSC_DEFAULT_TIME
		return
	}

	return
}

// Updates our last EDSC update time.
func updateEdscUpdateTime(time string) (err error) {
	err = ioutil.WriteFile(EDSC_UPDATE_FILE, []byte(time), 0744)
	return
}

// Validate system by ensuring it exists in database and correct coordinates.
func (s *System) validate() (*System, error) {
	db := dbMaster.spawnDb()
	defer db.stop()

	result, err := db.getSystem(s.Name)

	return result, err
}
