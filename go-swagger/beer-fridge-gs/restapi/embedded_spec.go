// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.grohm.go-workshop.beer-fridge.v1+json"
  ],
  "produces": [
    "application/io.grohm.go-workshop.beer-fridge.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Beer fridge service build with go-swagger",
    "title": "A beer fridge service",
    "version": "1.0.0"
  },
  "paths": {
    "/beers": {
      "get": {
        "tags": [
          "beers"
        ],
        "operationId": "getAllBeers",
        "parameters": [
          {
            "type": "number",
            "format": "float",
            "name": "volume-percentage",
            "in": "query"
          },
          {
            "type": "string",
            "name": "title",
            "in": "query"
          },
          {
            "type": "string",
            "name": "origin",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 10,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the beer operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/beer"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "beers"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/beer"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/beer"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/beers/{id}": {
      "delete": {
        "tags": [
          "beers"
        ],
        "operationId": "destroyOne",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/temperature": {
      "get": {
        "tags": [
          "fridge"
        ],
        "operationId": "getTemperature",
        "responses": {
          "200": {
            "description": "return the current fridge temperature",
            "schema": {
              "$ref": "#/definitions/temperature"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "beer": {
      "type": "object",
      "required": [
        "title",
        "origin",
        "volume-percentage"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "origin": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "volume-percentage": {
          "type": "number",
          "format": "float",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "temperature": {
      "type": "integer",
      "format": "int64",
      "readOnly": true
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.grohm.go-workshop.beer-fridge.v1+json"
  ],
  "produces": [
    "application/io.grohm.go-workshop.beer-fridge.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Beer fridge service build with go-swagger",
    "title": "A beer fridge service",
    "version": "1.0.0"
  },
  "paths": {
    "/beers": {
      "get": {
        "tags": [
          "beers"
        ],
        "operationId": "getAllBeers",
        "parameters": [
          {
            "type": "number",
            "format": "float",
            "name": "volume-percentage",
            "in": "query"
          },
          {
            "type": "string",
            "name": "title",
            "in": "query"
          },
          {
            "type": "string",
            "name": "origin",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 10,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the beer operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/beer"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "beers"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/beer"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/beer"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/beers/{id}": {
      "delete": {
        "tags": [
          "beers"
        ],
        "operationId": "destroyOne",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/temperature": {
      "get": {
        "tags": [
          "fridge"
        ],
        "operationId": "getTemperature",
        "responses": {
          "200": {
            "description": "return the current fridge temperature",
            "schema": {
              "$ref": "#/definitions/temperature"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "beer": {
      "type": "object",
      "required": [
        "title",
        "origin",
        "volume-percentage"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "origin": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "volume-percentage": {
          "type": "number",
          "format": "float",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "temperature": {
      "type": "integer",
      "format": "int64",
      "readOnly": true
    }
  }
}`))
}