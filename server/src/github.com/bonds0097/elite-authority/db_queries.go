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

// Returns all commodities, basic data only.
func (db *MongoDbConnection) listCommodities() (c []Commodity, err error) {
	err = db.session.DB("elite_authority").C("commodities").Find(nil).Select(bson.M{"stations": 0}).All(&c)

	return
}

func (db *MongoDbConnection) getCommodityNames() (s []string, err error) {
	var c []Commodity
	err = db.session.DB("elite_authority").C("commodities").Find(nil).Select(bson.M{"_id": 1}).All(&c)

	for _, commodity := range c {
		s = append(s, commodity.Name)
	}

	return
}

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

func (db *MongoDbConnection) doesStationExist(id string) (exists bool, err error) {
	count, err := db.session.DB("elite_authority").C("stations").FindId(id).Count()
	if count > 0 {
		exists = true
	}

	return
}
