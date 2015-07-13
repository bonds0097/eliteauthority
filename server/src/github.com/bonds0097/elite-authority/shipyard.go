package main

import (
	"errors"
)

type Ship struct {
	Name  string `bson:"name" json:"name"`
	Price uint   `bson:"price" json:"price"`
}

func (s *Ship) validate() (Ship, error) {
	if len(s.Name) < FIElD_LENGTH_MIN || len(s.Name) > FIELD_LENGTH_MAX {
		return *s, errors.New("Invalid ship name.")
	}

	s.Name = normalizeString(s.Name)

	return *s, nil
}
