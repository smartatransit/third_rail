// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-22 17:40:22.317984 -0400 EDT m=+8.662391742

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "SMARTA Support",
            "email": "smartatransit@gmail.com"
        },
        "license": {
            "name": "GNU General Public License v3.0",
            "url": "https://github.com/smartatransit/third_rail/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/live/alerts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "MARTA alerts sourced from their official twitter account",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Alerts from various MARTA sources",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.alertResponse"
                        }
                    }
                }
            }
        },
        "/live/schedule/line/{line}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Given a line, return the current live schedule",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Schedule By Line",
                "parameters": [
                    {
                        "type": "string",
                        "description": "RED, GOLD, BLUE, GREEN",
                        "name": "line",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.response"
                        }
                    }
                }
            }
        },
        "/live/schedule/station/{station}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Given a station, return the current live schedule",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Schedule By Station",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TODO: Enter all stations as enum?",
                        "name": "station",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.response"
                        }
                    }
                }
            }
        },
        "/smart/parking": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get available parking information as informed by twitter",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Parking Information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.parkingResponse"
                        }
                    }
                }
            }
        },
        "/static/directions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all available directions",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Directions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.directionsResponse"
                        }
                    }
                }
            }
        },
        "/static/lines": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all available lines",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Lines",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.linesResponse"
                        }
                    }
                }
            }
        },
        "/static/location": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get nearest station given a lat and lng",
                "produces": [
                    "application/json"
                ],
                "summary": "Get nearest station given a lat and lng",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Latitude",
                        "name": "latitutde",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.stationsLocationResponse"
                        }
                    },
                    "400": {}
                }
            }
        },
        "/static/schedule/station": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get MARTA's scheduled times for arrival for all stations",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Static Schedule By Station",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.response"
                        }
                    }
                }
            }
        },
        "/static/stations": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all available stations",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Stations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.stationsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.alertResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/marta_schemas.Alerts"
                }
            }
        },
        "controllers.directionsData": {
            "type": "object",
            "properties": {
                "directions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controllers.directionsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.directionsData"
                }
            }
        },
        "controllers.linesData": {
            "type": "object",
            "properties": {
                "lines": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controllers.linesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.linesData"
                }
            }
        },
        "controllers.parkingData": {
            "type": "object",
            "properties": {
                "station": {
                    "type": "object",
                    "$ref": "#/definitions/marta_schemas.Station"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "controllers.parkingResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.parkingData"
                    }
                }
            }
        },
        "controllers.response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.responseData"
                    }
                }
            }
        },
        "controllers.responseData": {
            "type": "object",
            "properties": {
                "schedule": {
                    "type": "object",
                    "$ref": "#/definitions/marta_schemas.Schedule"
                },
                "station": {
                    "type": "object",
                    "$ref": "#/definitions/marta_schemas.Station"
                }
            }
        },
        "controllers.stationsData": {
            "type": "object",
            "properties": {
                "stations": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controllers.stationsLocationResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/marta_schemas.StationLocation"
                    }
                }
            }
        },
        "controllers.stationsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.stationsData"
                }
            }
        },
        "marta_schemas.Alerts": {
            "type": "object",
            "properties": {
                "bus": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "desc": {
                                "type": "string"
                            },
                            "expires": {
                                "type": "string"
                            },
                            "id": {
                                "type": "string"
                            },
                            "text": {
                                "type": "string"
                            },
                            "title": {
                                "type": "string"
                            }
                        }
                    }
                },
                "rail": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "desc": {
                                "type": "string"
                            },
                            "expires": {
                                "type": "string"
                            },
                            "id": {
                                "type": "string"
                            },
                            "text": {
                                "type": "string"
                            },
                            "title": {
                                "type": "string"
                            }
                        }
                    }
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "marta_schemas.Schedule": {
            "type": "object",
            "properties": {
                "destination": {
                    "type": "string"
                },
                "event_time": {
                    "type": "string"
                },
                "next_arrival": {
                    "type": "string"
                },
                "next_station": {
                    "type": "string"
                },
                "train_id": {
                    "type": "string"
                },
                "waiting_seconds": {
                    "type": "string"
                },
                "waiting_time": {
                    "type": "string"
                }
            }
        },
        "marta_schemas.Station": {
            "type": "object",
            "properties": {
                "direction": {
                    "type": "string"
                },
                "line": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "marta_schemas.StationLocation": {
            "type": "object",
            "properties": {
                "distance": {
                    "type": "number"
                },
                "location": {
                    "type": "string"
                },
                "station_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "third-rail.services.ataper.net",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "SMARTA API",
	Description: "API to serve you SMARTA data",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}