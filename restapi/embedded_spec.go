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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "restapi for sigenu",
    "title": "sigapi",
    "version": "1.0.0"
  },
  "paths": {
    "/students": {
      "get": {
        "tags": [
          "student"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the students",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/student"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
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
    "student": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "last-name": {
          "type": "string"
        },
        "middle-name": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "restapi for sigenu",
    "title": "sigapi",
    "version": "1.0.0"
  },
  "paths": {
    "/students": {
      "get": {
        "tags": [
          "student"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the students",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/student"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
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
    "student": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "last-name": {
          "type": "string"
        },
        "middle-name": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
}
