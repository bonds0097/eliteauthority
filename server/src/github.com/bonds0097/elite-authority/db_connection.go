package main

import (
	"gopkg.in/mgo.v2"
)

type MongoDbConnection struct {
	session *mgo.Session
}

// Spawn a new  database connection with a copy of the master session.
func (db *MongoDbConnection) spawnDb() *MongoDbConnection {
	return &MongoDbConnection{db.session.Copy()}
}

// Start database connection. Should only be used once.
func (db *MongoDbConnection) start() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	db.session = session
}

// Close database connection.
func (db *MongoDbConnection) stop() {
	db.session.Close()
}
