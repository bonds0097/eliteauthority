package main

import (
	"errors"
	"regexp"
	"strings"
)

type Module struct {
	Name  string `bson:"name" json:"name"`
	Size  uint   `bson:"size" json:"size"`
	Class string `bson:"class" json:"mClass"`
	Price uint   `bson:"price" json:"price"`
}

func (m *Module) validate() (Module, error) {
	if len(m.Name) < FIElD_LENGTH_MIN || len(m.Name) > FIELD_LENGTH_MAX {
		return *m, errors.New("Invalid Module name.")
	}

	m.Name = normalizeString(m.Name)

	m.Class = strings.ToUpper(m.Class)

	validClass, _ := regexp.Compile("[A-Z]")
	if !validClass.MatchString(m.Class) {
		return *m, errors.New("Invalid module class. Should be a single letter.")
	}

	return *m, nil
}
