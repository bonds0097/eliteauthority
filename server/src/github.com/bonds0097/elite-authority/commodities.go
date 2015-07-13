package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	categories = [...]string{"Chemicals", "Consumer Items", "Foods", "Industrial Materials",
		"Legal Drugs", "Machinery", "Medicines", "Metals", "Minerals", "Salvage", "Slavery",
		"Technology", "Textiles", "Waste", "Weapons"}
)

type Commodity struct {
	Name            string `bson:"_id" json:"name"`
	Category        string `bson:"category" json:"category"`
	IsRare          bool   `bson:"is_rare" json:"isRare"`
	GalacticAverage uint   `bson:"galactic_average,omitempty" json:"galacticAverage,omitempty"`
}

func getCommodities(w http.ResponseWriter, r *http.Request) (apiErr *ApiError) {
	db := dbMaster.spawnDb()
	defer db.stop()

	commodities, err := db.getCommodities()
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
	if err != nil {
		return &ApiError{JSON_PARSE_ERROR, http.StatusBadRequest}
	}

	err = commodity.validateAndNormalize()
	if err != nil {
		return &ApiError{"Failed to validate commodity.\n" + err.Error(), http.StatusBadRequest}
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
func (c *Commodity) validateAndNormalize() (err error) {
	// Name
	if c.Name == "" || len(c.Name) < FIElD_LENGTH_MIN || len(c.Name) > FIELD_LENGTH_MAX {
		return errors.New("Invalid or missing commodity name.")
	}

	c.Name = normalizeString(c.Name)

	// Category
	if c.Category == "" {
		return errors.New("Missing category.")
	}

	categoryFound := false
	for _, category := range categories {
		if c.Category == category {
			categoryFound = true
			break
		}
	}

	if !categoryFound {
		return errors.New("Category does not exist.")
	}

	return
}
