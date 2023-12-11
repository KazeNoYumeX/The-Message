// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/games": {
            "post": {
                "description": "Start a new game",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Start a new game",
                "parameters": [
                    {
                        "description": "Players",
                        "name": "players",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateGameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.CreateGameResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/heartbeat": {
            "get": {
                "description": "Check if the server is alive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "heartbeat"
                ],
                "summary": "Check if the server is alive",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/v1/player/{id}/player-cards/": {
            "get": {
                "description": "GetPlayerCardsByPlayerId",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "player_cards"
                ],
                "summary": "GetPlayerCards",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.PlayerCardsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateGameRequest": {
            "type": "object",
            "properties": {
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.PlayerInfo"
                    }
                }
            }
        },
        "request.CreateGameResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "request.PlayerCardsResponse": {
            "type": "object",
            "properties": {
                "player_cards": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                }
            }
        },
        "request.PlayerInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "127.0.0.1:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "The Message API",
	Description:      "This is an online version of the \"The Message\" board game backend API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}