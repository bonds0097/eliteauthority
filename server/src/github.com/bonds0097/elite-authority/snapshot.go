package main

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

// Market snapshot that captures the state of a specific commodity at a specific stations at a
// given time.
type Snapshot struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Commodity       *Commodity    `bson:"commodity" json:"commodity"`
	Station         string        `bson:"station" json:"station"`
	SellPrice       uint          `bson:"sell_price" json:"sellPrice"`
	BuyPrice        uint          `bson:"buy_price" json:"buyPrice"`
	Demand          uint          `bson:"demand" json:"demand"`
	Supply          uint          `bson:"supply" json:"supply"`
	GalacticAverage uint          `bson:"galactic_average" json:"galacticAverage"`
	Timestamp       time.Time     `bson:"timestamp" json:"timestamp"`
}

// Adds a new market snapshot to the database.
// NOTE: Both the station and commodity have to exist before this function is called or you get
// an error.
func addSnapshot(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	// Turn the JSON request into a Snapshot object
	var snapshot Snapshot
	err := json.NewDecoder(r.Body).Decode(&snapshot)
	if err != nil {
		return &ApiError{JSON_PARSE_ERROR, http.StatusBadRequest}
	}

	// Validate the snapshot object. All fields are mandatory.
	err = snapshot.validate()
	if err != nil {
		return &ApiError{"Market snapshot failed to validate.\n" + err.Error(), http.StatusBadRequest}
	}

	// Get database session.
	db := dbMaster.spawnDb()
	defer db.stop()

	// Insert the snapshot into the database.
	err = db.addSnapshot(&snapshot)
	if err != nil {
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// Return nil if no errours encountered until now.
	return nil
}

func (s *Snapshot) validate() (err error) {
	// Create DB session.
	db := dbMaster.spawnDb()
	defer db.stop()

	// Commodity
	commodityExists, err := db.doesCommodityExist(s.Commodity.Name)
	if !commodityExists {
		return errors.New("Cannot create snapshot for non-existent commodity.")
	}

	if err != nil {
		return
	}

	s.Commodity, err = db.getCommodity(s.Commodity.Name)
	if err != nil {
		return
	}

	// Station
	if s.Station != "" {
		var stationExists bool
		stationExists, err = db.doesStationExist(s.Station)
		if !stationExists {
			return errors.New("Cannot create snapshot for non-existent station.")
		}

		if err != nil {
			return
		}
	}

	// Update timestamp.
	s.Timestamp = time.Now()

	return
}
