// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-01-19 11:59:21.185721 -0300 -03 m=+0.059121904

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
            "name": "Trust Wallet",
            "url": "https://t.me/wallecore"
        },
        "license": {
            "name": "MIT License",
            "url": "https://raw.githubusercontent.com/trustwallet/redemption/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/metrics": {
            "get": {
                "description": "Get application metrics",
                "tags": [
                    "status"
                ],
                "summary": "Get Metrics",
                "operationId": "metrics"
            }
        },
        "/status": {
            "get": {
                "description": "Get application status",
                "tags": [
                    "status"
                ],
                "summary": "Get Status",
                "operationId": "health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.Success"
                        }
                    }
                }
            }
        },
        "/v1/address/{platform}": {
            "get": {
                "description": "Get public address from a main wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get public address",
                "operationId": "public_address",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "default": "714",
                        "description": "the platform Id",
                        "name": "platform",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.Address"
                        }
                    }
                }
            }
        },
        "/v1/link/{code}": {
            "get": {
                "description": "Get a specific link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "redeem"
                ],
                "summary": "Get specific link",
                "operationId": "get_link",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "the link code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.Link"
                        }
                    }
                }
            },
            "post": {
                "description": "Update a specific link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "redeem"
                ],
                "summary": "Update specific link",
                "operationId": "update_link",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "the link code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Link",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/redemption.UpdateLink"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.Link"
                        }
                    }
                }
            }
        },
        "/v1/links": {
            "get": {
                "description": "Get all referral links",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "redeem"
                ],
                "summary": "Get all links",
                "operationId": "get_all_links",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Provider name",
                        "name": "provider",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.Links"
                        }
                    }
                }
            }
        },
        "/v1/links/create": {
            "post": {
                "description": "Create code and links for referral from a specific asset",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "redeem"
                ],
                "summary": "Create code for referral",
                "operationId": "create_links",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Links",
                        "name": "links",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/redemption.CreateLinks"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.Links"
                        }
                    }
                }
            }
        },
        "/v1/referral/redeem": {
            "post": {
                "description": "Redeem the referral code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "redeem"
                ],
                "summary": "Redeem code",
                "operationId": "create_links",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Redeem",
                        "name": "redeem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/redemption.Redeem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/redemption.RedeemResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "redemption.Address": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "coin": {
                    "type": "integer"
                },
                "coin_name": {
                    "type": "string"
                }
            }
        },
        "redemption.Asset": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "token_id": {
                    "type": "string"
                }
            }
        },
        "redemption.Assets": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/redemption.Asset"
                    }
                },
                "coin": {
                    "type": "integer"
                },
                "used": {
                    "type": "boolean"
                }
            }
        },
        "redemption.CreateLinks": {
            "type": "object",
            "properties": {
                "asset": {
                    "type": "object",
                    "$ref": "#/definitions/redemption.Assets"
                },
                "link_count": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                }
            }
        },
        "redemption.Link": {
            "type": "object",
            "properties": {
                "asset": {
                    "type": "object",
                    "$ref": "#/definitions/redemption.Assets"
                },
                "code": {
                    "type": "string"
                },
                "created_date": {
                    "type": "string"
                },
                "expiration_date": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "redemption.Links": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/redemption.Link"
            }
        },
        "redemption.Observer": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "coin": {
                    "type": "integer"
                },
                "publicKeys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "redemption.Observers": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/redemption.Observer"
            }
        },
        "redemption.Redeem": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "observers": {
                    "type": "object",
                    "$ref": "#/definitions/redemption.Observers"
                }
            }
        },
        "redemption.RedeemResult": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "object",
                    "$ref": "#/definitions/redemption.Assets"
                },
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "result_id": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "redemption.Success": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "redemption.UpdateLink": {
            "type": "object",
            "properties": {
                "expiration_date": {
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "TrustWallet Redeem API",
	Description: "Redemption API for Trust Wallet",
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
