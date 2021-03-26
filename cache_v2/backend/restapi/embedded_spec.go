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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "EM Backend server APIs",
    "title": "Backend APIs",
    "version": "1.1.0"
  },
  "paths": {
    "/search": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Return the 10 results",
        "parameters": [
          {
            "type": "integer",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "emp_name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "properties": {
                "value": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "The specified offset is not valid"
          },
          "405": {
            "description": "Method not allowed"
          }
        }
      }
    },
    "/searchbyID": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Return the search result",
        "parameters": [
          {
            "type": "string",
            "name": "key",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Data found",
            "schema": {
              "properties": {
                "value": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Key not found"
          },
          "405": {
            "description": "Method not allowed"
          }
        }
      }
    },
    "/submit": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "summary": "Save data in database",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/record"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Saved"
          },
          "400": {
            "description": "The specified key, value pair is invalid (requires strings)"
          },
          "405": {
            "description": "Method not allowed"
          }
        }
      }
    }
  },
  "definitions": {
    "record": {
      "type": "object",
      "required": [
        "key",
        "value"
      ],
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "EM Backend server APIs",
    "title": "Backend APIs",
    "version": "1.1.0"
  },
  "paths": {
    "/search": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Return the 10 results",
        "parameters": [
          {
            "type": "integer",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "emp_name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "properties": {
                "value": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "The specified offset is not valid"
          },
          "405": {
            "description": "Method not allowed"
          }
        }
      }
    },
    "/searchbyID": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Return the search result",
        "parameters": [
          {
            "type": "string",
            "name": "key",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Data found",
            "schema": {
              "properties": {
                "value": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Key not found"
          },
          "405": {
            "description": "Method not allowed"
          }
        }
      }
    },
    "/submit": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "summary": "Save data in database",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/record"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Saved"
          },
          "400": {
            "description": "The specified key, value pair is invalid (requires strings)"
          },
          "405": {
            "description": "Method not allowed"
          }
        }
      }
    }
  },
  "definitions": {
    "record": {
      "type": "object",
      "required": [
        "key",
        "value"
      ],
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    }
  }
}`))
}
