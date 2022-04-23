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
        "contact": {
            "name": "Sonr Inc.",
            "url": "https://sonr.io",
            "email": "team@sonr.io"
        },
        "license": {
            "name": "OpenGLv3"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/blob/download/:cid": {
            "get": {
                "description": "DownloadBlob downloads a file from IPFS given its CID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Download File",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/blob/remove/:cid": {
            "get": {
                "description": "RemoveBlob deletes a file from IPFS given its CID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Remove Blob",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/blob/upload": {
            "post": {
                "description": "UploadBlob uploads a file to IPFS and returns its CID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Upload File",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bucket/deactivate/:cid": {
            "get": {
                "description": "DeactivateBucket disables a bucket for a registered application via HTTP.",
                "produces": [
                    "application/json"
                ],
                "summary": "Deactivate Bucket",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.23.0",
	Host:             "localhost:8081",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Highway API",
	Description:      "Manage your Sonr Powered services and blockchain registered types with the Highway API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}