package validators

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"log"
	"strings"
)

const MARTA_DIRECTIONS = "Directions"
const MARTA_LINES = "Lines"
const MARTA_STATIONS = "Stations"

type AliasStore struct {
	Aliases []string `json:"aliases"`
	Name    string   `json:"name"`
}

type MartaEntities struct {
	Directions []AliasStore `json:"directions"`
	Lines      []AliasStore `json:"lines"`
	Stations   []AliasStore `json:"stations"`
}

type MartaEntitiesValidator struct {
	Entities MartaEntities
}

func NewMartaEntitiesValidator() (mev MartaEntitiesValidator) {
	mev = MartaEntitiesValidator{}
	file, _ := ioutil.ReadFile("data/schemas/entities.json")
	err := json.Unmarshal([]byte(file), &mev.Entities)

	if err != nil {
		panic("Holy shit we can't find the entities file panic panic panic")
	}

	return
}

func (mev MartaEntitiesValidator) Validate() {

}

func (mev MartaEntitiesValidator) Coerce(entityType string, value string) (string, error) {
	var result interface{}

	if entityType == MARTA_DIRECTIONS {
		result = findAlias(mev.Entities.Directions, value)
	} else if entityType == MARTA_LINES {
		result = findAlias(mev.Entities.Lines, value)
	} else if entityType == MARTA_STATIONS {
		result = findAlias(mev.Entities.Stations, value)
	} else {
		return "", errors.New("Unrecognized entity type: " + entityType)
	}

	if result == nil {
		return "", errors.New(fmt.Sprintf("No %s found for %s", entityType, value))
	}

	return result.(AliasStore).Name, nil
}

func (mev MartaEntitiesValidator) GetEntities(entityType string) ([]string, error) {
	var result interface{}

	if entityType == MARTA_DIRECTIONS {
		result = funk.Get(mev.Entities.Directions, "Name")
	} else if entityType == MARTA_LINES {
		result = funk.Get(mev.Entities.Lines, "Name")
	} else if entityType == MARTA_STATIONS {
		result = funk.Get(mev.Entities.Stations, "Name")
	} else {
		return nil, errors.New("Unrecognized entity type: " + entityType)
	}

	return result.([]string), nil
}

func findAlias(entities []AliasStore, value string) interface{} {
	log.Printf("Looking for aliases of %s", value)
	return funk.Find(entities, func(store AliasStore) bool {
		return strings.ToUpper(store.Name) == strings.ToUpper(value) || funk.Contains(store.Aliases, value)
	})
}
