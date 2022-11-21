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
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.PingSwag"
                        }
                    }
                }
            }
        },
        "/todos": {
            "get": {
                "description": "get all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TODOS"
                ],
                "summary": "Get All TODOS",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.GetTodosSuccessSwag"
                        }
                    }
                }
            },
            "post": {
                "description": "create todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TODOS"
                ],
                "summary": "Create TODO",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateTodo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.CreateTodoSuccessSwag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.CreateTodoFailureSwag"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "get all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TODOS"
                ],
                "summary": "Get All TODOS",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.GetTodosSuccessSwag"
                        }
                    }
                }
            },
            "put": {
                "description": "update todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TODOS"
                ],
                "summary": "Update TODO",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateTodo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.UpdateTodoSuccessSwag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.UpdateTodoFailureSwag"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TODOS"
                ],
                "summary": "Delete TODO",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.UpdateTodoSuccessSwag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.UpdateTodoFailureSwag"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.PingSwag": {
            "type": "object",
            "properties": {
                "health": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "request.CreateTodo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Desc TODO"
                },
                "title": {
                    "type": "string",
                    "example": "Title TODO"
                }
            }
        },
        "views.CreateTodoFailureSwag": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "error": "Title is required"
                    }
                },
                "error": {
                    "type": "string",
                    "example": "BAD_REQUEST"
                },
                "message": {
                    "type": "string",
                    "example": "CREATE TODO FAIL"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "views.CreateTodoPayload": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "example": "Desc TODO"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Title TODO"
                }
            }
        },
        "views.CreateTodoSuccessSwag": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "CREATE TODO SUCCESS"
                },
                "payload": {
                    "$ref": "#/definitions/views.CreateTodoPayload"
                },
                "status": {
                    "type": "integer",
                    "example": 201
                }
            }
        },
        "views.GetTodoPayload": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "example": "Desc TODO"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Title TODO"
                }
            }
        },
        "views.GetTodosSuccessSwag": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "CREATE TODO SUCCESS"
                },
                "payload": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/views.GetTodoPayload"
                    }
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "views.UpdateTodoFailureSwag": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "error": "Title is required"
                    }
                },
                "error": {
                    "type": "string",
                    "example": "BAD_REQUEST"
                },
                "message": {
                    "type": "string",
                    "example": "UPDATE TODO FAIL"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "views.UpdateTodoSuccessSwag": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "UPDATE TODO SUCCESS"
                },
                "payload": {
                    "$ref": "#/definitions/views.CreateTodoPayload"
                },
                "status": {
                    "type": "integer",
                    "example": 201
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
