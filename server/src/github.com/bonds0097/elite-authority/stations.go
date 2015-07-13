package main

import (
	"encoding/json"
	"errors"
	"github.com/bonds0097/webUtils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

type Station struct {
	ID               bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Name             string        `bson:"name" json:"name"`
	Type             string        `bson:"type" json:"type"`
	System           *System       `bson:"system,omitempty" json:"system,omitempty"`
	DistanceFromJump float64       `bson:"distance_from_jump" json:"distanceFromJump"`
	Wealth           string        `bson:"wealth" json:"wealth"`
	Population       string        `bson:"population" json:"population"`
	Economies        []string      `bson:"economies" json:"economies"`
	Facilities       []string      `bson:"facilities" json:"facilities"`
	Allegiance       string        `bson:"allegiance" json:"allegiance"`
	Faction          string        `bson:"faction" json:"faction"`
	Government       string        `bson:"government" json:"government"`
	HasBlackMarket   bool          `bson:"has_black_market" json:"hasBlackMarket"`
	Outfitting       []Module      `bson:"outfitting,omitempty" json:"outfitting,omitempty"`
	Shipyard         []Ship        `bson:"shipyard,omitempty" json:"shipyard,omitempty"`
	Snapshots        []Snapshot    `bson:"snapshots,omitempty" json:"snapshots,omitempty"`
	Slug             string        `bson:"slug" json:"slug"`
	Timestamp        time.Time     `bson:"timestamp" json::"timestamp"`
}

// Returns all stations with basic data.
func getRecentStations(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	// Get Most Recent Stations From Database
	db := dbMaster.spawnDb()
	defer db.stop()

	stations, err := db.getRecentStations()
	if err != nil {
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// Return to requester.
	json.NewEncoder(w).Encode(stations)

	return
}

// Returns data on a specific station.
func getStationDetails(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	// Get Station Slug
	slug := mux.Vars(r)["slug"]

	// Request station data from database.
	db := dbMaster.spawnDb()
	defer db.stop()

	station, err := db.getStationDetails(slug)
	if err != nil {
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// Return to Requester.
	json.NewEncoder(w).Encode(&station)

	return
}

// Updates a station with new data.

// Adds a new station to the database.
func addStation(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	// Parse JSON response and unmarshal to a Commodity.
	var station Station
	err := json.NewDecoder(r.Body).Decode(&station)

	// If unable to parse JSON data, return error.
	if err != nil {
		return &ApiError{JSON_PARSE_ERROR, http.StatusBadRequest}
	}

	// Validate station data and normalize it for DB entry.
	if err = station.validateAndNormalize(); err != nil {
		return &ApiError{"Failed to validate station object.\n" + err.Error(), http.StatusBadRequest}
	}

	// Add station to database.
	db := dbMaster.spawnDb()
	defer db.stop()

	err = db.addStation(&station)
	if err != nil {
		return &ApiError{err.Error(), http.StatusConflict}
	}

	// If all is well, return the station back to the requestor.
	json.NewEncoder(w).Encode(&station)

	return
}

// Validates a station object. Required data is name and system.
func (s *Station) validateAndNormalize() (err error) {
	// Name
	fieldLength := len(s.Name)
	if s.Name == "" || fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX {
		return errors.New("Station name invalid length or missing.")
	}

	s.Name = normalizeString(s.Name)

	// Type
	fieldLength = len(s.Type)
	if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX) {
		return errors.New("Station type invalid length.")
	}

	s.Type = normalizeString(s.Type)

	// System
	s.System, err = s.System.validate()
	if err != nil {
		return errors.New("Error validating system: " + err.Error())
	}

	// DistanceFromJump
	if s.DistanceFromJump < 0 {
		return errors.New("Station distance from jump must be positive.")
	}

	// Wealth
	fieldLength = len(s.Wealth)
	if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX) {
		return errors.New("Station wealth invalid length.")
	}

	s.Wealth = normalizeString(s.Wealth)

	// Population
	fieldLength = len(s.Wealth)
	if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX) {
		return errors.New("Station population invalid length.")
	}

	s.Population = normalizeString(s.Population)

	// Economies
	for index, _ := range s.Economies {
		fieldLength = len(s.Economies[index])
		if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
			fieldLength > FIELD_LENGTH_MAX) {
			return errors.New("Station economies element invalid length.")
		}

		s.Economies[index] = normalizeString(s.Economies[index])
	}

	// Facilities
	for index, _ := range s.Facilities {
		fieldLength = len(s.Facilities[index])
		if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
			fieldLength > FIELD_LENGTH_MAX) {
			return errors.New("Station facilities element invalid length.")
		}

		s.Facilities[index] = normalizeString(s.Facilities[index])
	}

	// Allegiance
	fieldLength = len(s.Allegiance)
	if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX) {
		return errors.New("Station allegiance invalid length.")
	}

	s.Allegiance = normalizeString(s.Allegiance)

	// Faction
	fieldLength = len(s.Faction)
	if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX) {
		return errors.New("Station faction invalid length.")
	}

	s.Faction = normalizeString(s.Faction)

	// Government
	fieldLength = len(s.Government)
	if fieldLength > 0 && (fieldLength < FIElD_LENGTH_MIN ||
		fieldLength > FIELD_LENGTH_MAX) {
		return errors.New("Station government invalid length.")
	}

	s.Government = normalizeString(s.Government)

	// Outfitting
	for index, _ := range s.Outfitting {
		s.Outfitting[index], err = s.Outfitting[index].validate()
		if err != nil {
			return errors.New("Error validating outfitting entry: " + err.Error())
		}
	}

	// Shipyard
	for index, _ := range s.Shipyard {
		s.Shipyard[index], err = s.Shipyard[index].validate()
		if err != nil {
			return errors.New("Error validating shipyard entry: " + err.Error())
		}
	}

	// Snapshots
	for index, _ := range s.Snapshots {
		err = s.Snapshots[index].validate()
		if err != nil {
			return errors.New("Error validating snapshot entry: " + err.Error())
		}

		s.Snapshots[index].Station = s.Name
	}

	// Slug
	s.Slug, err = webUtils.GenerateSlug(s.System.Name + "-" + s.Name)
	if err != nil {
		return errors.New("Station name invalid. Failed to generate slug.")
	}

	// Timestamp
	s.Timestamp = time.Now()

	return
}
