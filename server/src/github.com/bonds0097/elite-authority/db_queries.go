package main

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Adds a commodity object to the database.
func (db *MongoDbConnection) addCommodity(c *Commodity) (err error) {
	err = db.session.DB("elite_authority").C("commodities").Insert(c)

	// Check if error is complaining about duplication and return new error with cleaner message.
	if mgo.IsDup(err) {
		err = errors.New("That commodity already exists.")
	}

	return
}

// Returns single commodity.
func (db *MongoDbConnection) getCommodity(id string) (*Commodity, error) {
	var c Commodity
	err := db.session.DB("elite_authority").C("commodities").FindId(id).One(&c)

	return &c, err
}

// Returns data for commodity.
func (db *MongoDbConnection) getCommodities() (c []Commodity, err error) {
	err = db.session.DB("elite_authority").C("commodities").Find(nil).All(&c)

	return
}

// Returns an array of commodity names.
func (db *MongoDbConnection) listCommodities() (s []string, err error) {
	var c []Commodity
	err = db.session.DB("elite_authority").C("commodities").Find(nil).Select(bson.M{"_id": 1}).Sort("_id").All(&c)

	for _, commodity := range c {
		s = append(s, commodity.Name)
	}

	return
}

// Checks if a commodity exists.
func (db *MongoDbConnection) doesCommodityExist(id string) (exists bool, err error) {
	count, err := db.session.DB("elite_authority").C("commodities").FindId(id).Count()
	if count > 0 {
		exists = true
	}

	return
}

// Updates Commodity and Station to reflect latest snapshot data.
func (db *MongoDbConnection) propagateSnapshot(s Snapshot) (err error) {
	// Fill in update data.
	station_key := "stations." + s.Station
	updateData := bson.M{"$set": bson.M{station_key: bson.M{"sell_price": s.SellPrice}}}

	err = db.session.DB("elite_authority").C("commodities").Update(bson.M{"_id": s.Commodity}, updateData)

	return
}

// Adds a snapshot to the database.
func (db *MongoDbConnection) addSnapshot(s *Snapshot) (err error) {
	err = db.session.DB("elite_authority").C("snapshots").Insert(s)

	return
}

// Checks if a station exists.
func (db *MongoDbConnection) doesStationExist(id string) (exists bool, err error) {
	count, err := db.session.DB("elite_authority").C("stations").FindId(id).Count()
	if count > 0 {
		exists = true
	}

	return
}

// Returns a single system object that matches the input id.
func (db *MongoDbConnection) getSystem(id string) (*System, error) {
	var s System
	err := db.session.DB("elite_authority").C("systems").FindId(id).One(&s)

	return &s, err
}

// Add systems to the database.
func (db *MongoDbConnection) addSystems(systems []System) (err error) {
	var args []interface{}
	for index, _ := range systems {
		args = append(args, systems[index])
	}

	err = db.session.DB("elite_authority").C("systems").Insert(args...)

	return
}

// Seearches for systems that match the input string.
func (db *MongoDbConnection) findSystems(name string) (s []string, err error) {
	var systems []System

	// We want to search by phrase, so modify the search string.
	name = "^" + name

	// db.session.DB("elite_authority").C("systems").EnsureIndexKey("_id")
	err = db.session.DB("elite_authority").C("systems").
		Find(bson.M{"_id": bson.M{"$regex": bson.RegEx{name, "i"}}}).Select(bson.M{"_id": 1}).All(&systems)

	for _, system := range systems {
		s = append(s, system.Name)
	}

	return
}

// Adds a commodity object to the database.
func (db *MongoDbConnection) addStation(s *Station) (err error) {
	err = db.session.DB("elite_authority").C("stations").Insert(s)

	// Check if error is complaining about duplication and return new error with cleaner message.
	if mgo.IsDup(err) {
		err = errors.New("That station already exists.")
	}

	return
}

// Returns Recent Stations with Basic Data
func (db *MongoDbConnection) getRecentStations() (s []Station, err error) {
	err = db.session.DB("elite_authority").C("stations").Find(nil).
		Select(bson.M{"snapshots": 0, "outfitting": 0, "shipyard": 0}).Sort("timestamp").Limit(20).
		All(&s)

	return
}

// Returns All Data for Given Station
func (db *MongoDbConnection) getStationDetails(slug string) (*Station, error) {
	var s Station
	err := db.session.DB("elite_authority").C("stations").Find(bson.M{"slug": slug}).One(&s)

	return &s, err
}
