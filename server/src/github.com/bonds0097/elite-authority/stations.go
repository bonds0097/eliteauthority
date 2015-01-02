package main

import (
	// "encoding/json"
	// "github.com/bonds0097/webUtils"
	// "net/http"
	"time"
)

type System struct {
	Name        string `bson:"name" json:"name"`
	Coordinates struct {
		x float64 `bson:"x" json:"y"`
		y float64 `bson:"y" json:"y"`
		z float64 `bson:"z" json:"z"`
	} `bson:"coordinates" json:"coordinates"`
}

type Station struct {
	Name             string             `bson:"_id" json:"name"`
	Type             string             `bson:"type" json:"type"`
	System           System             `bson:"system" json:"system"`
	DistanceFromJump float64            `bson:"distance_from_jump" json:"distanceFromJump"`
	Wealth           string             `bson:"wealth" json:"wealth"`
	Population       string             `bson:"population" json:"population"`
	Economies        []string           `bson:"economies" json:"economies"`
	Facilities       []string           `bson:"facilities" json:"facilities"`
	Allegiance       string             `bson:"allegiance" json:"allegiance"`
	Faction          string             `bson:"faction" json:"faction"`
	Government       string             `bson:"government" json:"government"`
	HasBlackMarket   bool               `bson:"has_black_market" json:"hasBlackMarket"`
	Outfitting       map[string]float64 `bson:"outfitting" json:"outfitting"`
	Shipyard         map[string]float64 `bson:"shipyard" json:"shipyard"`
	Commodities      map[string]struct {
		Category  string    `bson:"category" json:"category"`
		IsRare    bool      `bson:"is_rare" json:"isRare"`
		SellPrice int       `bson:"sell_price" json:"sellPrice"`
		BuyPrice  int       `bson:"buy_price" json:"buy_price"`
		Demand    int       `bson:"demand" json:"demand"`
		Supply    int       `bson:"supply" json:"demand"`
		Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	} `bson:"commodities,omitempty" json:"commodities,omitempty"`
	Slug string `bson:"slug" json:"sSlug"`
}

// func listStations(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {

// }

// func addStation(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
// 	// Parse JSON response and unmarshal to a Commodity.
// 	var commodity Stations
// 	err := json.NewDecoder(r.Body).Decode(&commodity)

// 	// If unable to parse or commodity is missing data, return error.
// 	if err != nil || !commodity.isValid() {
// 		return &ApiError{JSON_PARSE_ERROR, http.StatusBadRequest}
// 	}

// 	// Title Case the name and category of the commodity, to make pretty.
// 	commodity.Name = normalizeString(commodity.Name)
// 	commodity.Category = normalizeString(commodity.Category)

// 	commodity.Slug, err = webUtils.GenerateSlug(commodity.Name)
// 	if err != nil {
// 		return &ApiError{"Commodity name invalid. Valid characters are alphanumeric and ~_-.", http.StatusBadRequest}
// 	}

// 	db := dbMaster.spawnDb()
// 	defer db.stop()

// 	// Add commodity to database.
// 	err = db.addCommodity(&commodity)
// 	if err != nil {
// 		return &ApiError{err.Error(), http.StatusConflict}
// 	}

// 	// If all is well, return the commodity back to the requestor.
// 	json.NewEncoder(w).Encode(commodity)

// 	return nil
// }
