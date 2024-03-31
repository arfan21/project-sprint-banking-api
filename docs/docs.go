// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.synapsis.id"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/balance": {
            "get": {
                "description": "Get balance from user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Get Balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.BalanceGetResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add balance to user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Add Balance",
                "parameters": [
                    {
                        "description": "Payload balance add request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.TransactionAddBalanceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/v1/balance/history": {
            "get": {
                "description": "Get transaction from user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Get Transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit data",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.TransactionGetResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/v1/image": {
            "post": {
                "description": "Upload image to s3",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image Uploader"
                ],
                "summary": "Upload Image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/v1/transaction": {
            "post": {
                "description": "Transfer balance to bank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Transfer Balance",
                "parameters": [
                    {
                        "description": "Payload balance transfer request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.TransactionTransferBalanceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Payload user Login Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.UserLoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Payload user Register Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.UserLoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_arfan21_project-sprint-banking-api_internal_model.BalanceGetResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.TransactionAddBalanceRequest": {
            "type": "object",
            "required": [
                "addedBalance",
                "currency",
                "senderBankAccountNumber",
                "senderBankName",
                "transferProofImg"
            ],
            "properties": {
                "addedBalance": {
                    "type": "number",
                    "minimum": 0
                },
                "currency": {
                    "type": "string"
                },
                "senderBankAccountNumber": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "senderBankName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "transferProofImg": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.TransactionGetResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string"
                },
                "source": {
                    "$ref": "#/definitions/github_com_arfan21_project-sprint-banking-api_internal_model.TransactionSourceResponse"
                },
                "transactionId": {
                    "type": "string"
                },
                "transferProofImg": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.TransactionSourceResponse": {
            "type": "object",
            "properties": {
                "bankAccountNumber": {
                    "type": "string"
                },
                "bankName": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.TransactionTransferBalanceRequest": {
            "type": "object",
            "required": [
                "balances",
                "fromCurrency",
                "recipientBankAccountNumber",
                "recipientBankName"
            ],
            "properties": {
                "balances": {
                    "type": "number",
                    "minimum": 0
                },
                "fromCurrency": {
                    "type": "string"
                },
                "recipientBankAccountNumber": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                },
                "recipientBankName": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.UserLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 5
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.UserLoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_internal_model.UserRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 5
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.ErrValidationResponse": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_project-sprint-banking-api_pkg_pkgutil.HTTPResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "meta": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "project-sprint-banking-api",
	Description:      "This is a sample server cell for project-sprint-banking-api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
