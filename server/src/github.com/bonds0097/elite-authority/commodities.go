package main

import (
	"encoding/json"
	"github.com/bonds0097/webUtils"
	"net/http"
	"time"
)

type Commodity struct {
	Name            string `bson:"_id" json:"name"`
	Category        string `bson:"category" json:"category"`
	IsRare          bool   `bson:"is_rare" json:"isRare"`
	GalacticAverage int    `bson:"galactic_average" json:"galacticAverage"`
	Stations        map[string]struct {
		System    System    `bson:"system" json:"system"`
		SellPrice int       `bson:"sell_price" json:"sellPrice"`
		BuyPrice  int       `bson:"buy_price" json:"buyPrice"`
		Demand    int       `bson:"demand" json:"demand"`
		Supply    int       `bson:"supply" json:"supply"`
		Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	} `bson:"stations,omitempty" json:"stations,omitempty"`
	Slug string `bson:"slug" json:"cSlug"`
}

func listCommodities(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	db := dbMaster.spawnDb()
	defer db.stop()

	commodities, err := db.listCommodities()
	if len(commodities) == 0 {
		return &ApiError{"No commodities found.", http.StatusNotFound}
	}

	if err != nil {
		return &ApiError{DB_ERROR, http.StatusInternalServerError}
	}

	// If all is well, return the commodity list to the requestor.
	json.NewEncoder(w).Encode(commodities)

	return nil
}

// Parses POST request and adds commodity to database.
func addCommodity(w http.ResponseWriter, r *http.Request) *ApiError {
	// Parse JSON response and unmarshal to a Commodity.
	var commodity Commodity
	err := json.NewDecoder(r.Body).Decode(&commodity)

	// If unable to parse or commodity is missing data, return error.
	if err != nil || !commodity.isValid() {
		return &ApiError{JSON_PARSE_ERROR, http.StatusBadRequest}
	}

	// Title Case the name and category of the commodity, to make pretty.
	commodity.Name = normalizeString(commodity.Name)
	commodity.Category = normalizeString(commodity.Category)

	commodity.Slug, err = webUtils.GenerateSlug(commodity.Name)
	if err != nil {
		return &ApiError{"Commodity name invalid. Valid characters are alphanumeric and ~_-.", http.StatusBadRequest}
	}

	db := dbMaster.spawnDb()
	defer db.stop()

	// Add commodity to database.
	err = db.addCommodity(&commodity)
	if err != nil {
		return &ApiError{err.Error(), http.StatusConflict}
	}

	return nil
}

// Checks if commodity is valid, all data is required except for station data.
func (c *Commodity) isValid() (validity bool) {
	if c.Name == "" || c.Category == "" || c.GalacticAverage == 0 {
		return false
	}

	return true
}
