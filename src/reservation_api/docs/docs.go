// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/dateTimeSlots": {
            "get": {
                "description": "Get all date time slots",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.DateTimeSlot"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "API Healthcheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reservations": {
            "get": {
                "description": "Get all reservations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.Reservation"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.Reservation"
                        }
                    }
                }
            }
        },
        "/reservation/{id}": {
            "put": {
                "description": "Cancel reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.Reservation"
                        }
                    }
                }
            }
        },
        "/timeslots": {
            "get": {
                "description": "Get all timeslots",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.TimeSlot"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DateTimeSlot": {
            "type": "object",
            "required": [
                "date",
                "day",
                "time_slot"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "day": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "time_slot": {
                    "$ref": "#/definitions/model.TimeSlot"
                }
            }
        },
        "model.GuestPersona": {
            "type": "object",
            "properties": {
                "dietary_requirements": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "food_preferences": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "guest_name": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                }
            }
        },
        "model.Reservation": {
            "type": "object",
            "required": [
                "dts_id",
                "email",
                "first_name",
                "guest_count",
                "last_name",
                "phone_number"
            ],
            "properties": {
                "dts_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "guest_count": {
                    "type": "integer"
                },
                "guest_persona": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.GuestPersona"
                    }
                },
                "id": {
                    "type": "string"
                },
                "is_cancelled": {
                    "type": "boolean"
                },
                "is_rescheduled": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                }
            }
        },
        "model.TimeSlot": {
            "type": "object",
            "properties": {
                "from_hour": {
                    "type": "integer"
                },
                "from_miunutes": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "to_hour": {
                    "type": "integer"
                },
                "to_miunutes": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}