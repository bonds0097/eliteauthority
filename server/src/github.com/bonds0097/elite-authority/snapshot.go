package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

// Market snapshot that captures the state of a specific commodity at a specific stations at a
// given time.
type Snapshot struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Commodity       string        `bson:"commodity" json:"commodity"`
	Station         string        `bson:"station" json:"station"`
	SellPrice       int           `bson:"sell_price" json:"sellPrice"`
	BuyPrice        int           `bson:"buy_price" json:"buyPrice"`
	Demand          int           `bson:"demand" json:"demand"`
	Supply          int           `bson:"supply" json:"supply"`
	GalacticAverage int           `bson:"galactic_average" json:"galacticAverage"`
	Timestamp       time.Time     `bson:"timestamp" json:"timestamp"`
}

// Adds a new market snapshot to the database.
// NOTE: Both the station and commodity have to exist before this function is called or you get
// an error.
func addSnapshot(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	// Turn the JSON request into a Snapshot object
	var snapshot Snapshot
	err := json.NewDecoder(r.Body).Decode(&snapshot)
	// Validate the snapshot object. All fields are mandatory.
	if err != nil || !snapshot.isValid() {
		apiErr.Code = http.StatusBadRequest
		apiErr.Err = "Unable to parse snapshot object. Note that all attributes are required."
		return
	}

	// Get database session.
	db := dbMaster.spawnDb()
	defer db.stop()

	// Check that the referenced commodity exists.
	exists, err := db.doesCommodityExist(snapshot.Commodity)
	if err != nil {
		log.Println(err.Error())
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	if !exists {
		return &ApiError{"Commodity must exist before creating a snapshot for it.", http.StatusBadRequest}
	}

	// // Check that the referenced station exists.
	// exists, err = db.doesStationExist(snapshot.Station)
	// if err != nil {
	// 	return &ApiError{DB_ERROR, http.StatusInternalServerError}
	// }

	if !exists {
		return &ApiError{"Station must exist before creating a snapshot for it.", http.StatusBadRequest}
	}

	// Set the timestamp on the snapshot to current server time.
	snapshot.Timestamp = time.Now()

	// Insert the snapshot into the database.
	err = db.addSnapshot(&snapshot)
	if err != nil {
		b, _ := json.Marshal(snapshot)
		log.Println(string(b))
		log.Println(err.Error())
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// Update the commodity and station with snapshot.
	err = db.propagateSnapshot(snapshot)
	if err != nil {
		log.Println(err.Error())
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// Return nil if no errours encountered until now.
	return nil
}

func (s *Snapshot) isValid() bool {
	if s.Commodity == "" || s.Station == "" || s.SellPrice == 0 || s.BuyPrice == 0 || s.Demand == 0 ||
		s.Supply == 0 || s.GalacticAverage == 0 {
		return false
	}

	return true
}
