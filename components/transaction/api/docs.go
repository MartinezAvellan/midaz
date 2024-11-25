// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Discord community",
            "url": "https://discord.gg/DnhqKwkGv3"
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
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/accounts/{account_id}/operations": {
            "get": {
                "description": "Get all Operations with the input ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "Get all Operations by account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/Pagination"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "items": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/Operation"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/accounts/{account_id}/operations/{operation_id}": {
            "get": {
                "description": "Get an Operation with the input ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "Get an Operation by account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Operation ID",
                        "name": "operation_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Operation"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/asset-rates": {
            "post": {
                "description": "Create an AssetRate with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AssetRates"
                ],
                "summary": "Create an AssetRate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "AssetRate Input",
                        "name": "asset-rate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/assetrate.CreateAssetRateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/assetrate.AssetRate"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/asset-rates/{asset_rate_id}": {
            "get": {
                "description": "Get an AssetRate with the input ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AssetRates"
                ],
                "summary": "Get an AssetRate by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AssetRate ID",
                        "name": "asset_rate_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "AssetRate Input",
                        "name": "asset-rate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/assetrate.CreateAssetRateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/assetrate.AssetRate"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/portfolios/{portfolio_id}/operations": {
            "get": {
                "description": "Get all Operations with the input ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "Get all Operations by portfolio",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Portfolio ID",
                        "name": "portfolio_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/Pagination"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "items": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/Operation"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/portfolios/{portfolio_id}/operations/{operation_id}": {
            "get": {
                "description": "Get an Operation with the input ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "Get an Operation by portfolio",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Portfolio ID",
                        "name": "portfolio_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Operation ID",
                        "name": "operation_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Operation"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/transactions": {
            "get": {
                "description": "Get all Transactions with the input metadata or without metadata",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get all Transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/Pagination"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "items": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/transaction.Transaction"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/transactions/dsl": {
            "post": {
                "description": "Create a Transaction with the input DSL file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create a Transaction using DSL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Transaction DSL file",
                        "name": "transaction",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/transaction.Transaction"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/transactions/json": {
            "post": {
                "description": "Create a Transaction with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create a Transaction using JSON",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction Input",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/transaction.CreateTransactionInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/transaction.Transaction"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/transactions/{transaction_id}": {
            "get": {
                "description": "Get a Transaction with the input ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get a Transaction by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "transaction_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/transaction.Transaction"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a Transaction with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Update a Transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "transaction_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction Input",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/transaction.UpdateTransactionInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/transaction.Transaction"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/transactions/{transaction_id}/operations/{operation_id}": {
            "patch": {
                "description": "Update an Operation with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "Update an Operation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "transaction_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Operation ID",
                        "name": "operation_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Operation Input",
                        "name": "operation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateOperationInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Operation"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Amount": {
            "description": "Amount structure for marshaling/unmarshalling JSON.",
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 1500
                },
                "scale": {
                    "type": "number",
                    "example": 2
                }
            }
        },
        "Balance": {
            "description": "Balance structure for marshaling/unmarshalling JSON.",
            "type": "object",
            "properties": {
                "available": {
                    "type": "number",
                    "example": 1500
                },
                "onHold": {
                    "type": "number",
                    "example": 500
                },
                "scale": {
                    "type": "number",
                    "example": 2
                }
            }
        },
        "Operation": {
            "description": "Operation is a struct designed to encapsulate response payload data.",
            "type": "object",
            "properties": {
                "accountAlias": {
                    "type": "string",
                    "example": "@person1"
                },
                "accountId": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "amount": {
                    "$ref": "#/definitions/Amount"
                },
                "assetCode": {
                    "type": "string",
                    "example": "BRL"
                },
                "balance": {
                    "$ref": "#/definitions/Balance"
                },
                "balanceAfter": {
                    "$ref": "#/definitions/Balance"
                },
                "chartOfAccounts": {
                    "type": "string",
                    "example": "1000"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "deletedAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "Credit card operation"
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "ledgerId": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "organizationId": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "portfolioId": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "status": {
                    "$ref": "#/definitions/operation.Status"
                },
                "transactionId": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "type": {
                    "type": "string",
                    "example": "creditCard"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                }
            }
        },
        "Pagination": {
            "description": "Pagination is a struct designed to encapsulate pagination response payload data.",
            "type": "object",
            "properties": {
                "items": {},
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "UpdateOperationInput": {
            "description": "UpdateOperationInput is a struct design to encapsulate payload data.",
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 256,
                    "example": "Credit card operation"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                }
            }
        },
        "assetrate.AssetRate": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "baseAssetCode": {
                    "type": "string"
                },
                "counterAssetCode": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ledgerId": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "organizationId": {
                    "type": "string"
                },
                "scale": {
                    "type": "number"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "assetrate.CreateAssetRateInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "baseAssetCode": {
                    "type": "string"
                },
                "counterAssetCode": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "scale": {
                    "type": "number"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "model.Amount": {
            "type": "object",
            "required": [
                "asset",
                "value"
            ],
            "properties": {
                "asset": {
                    "type": "string"
                },
                "scale": {
                    "type": "integer",
                    "minimum": 0
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "model.Distribute": {
            "type": "object",
            "required": [
                "to"
            ],
            "properties": {
                "remaining": {
                    "type": "string"
                },
                "to": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.FromTo"
                    }
                }
            }
        },
        "model.FromTo": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "amount": {
                    "$ref": "#/definitions/model.Amount"
                },
                "chartOfAccountsG": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "isFrom": {
                    "type": "boolean"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "remaining": {
                    "type": "string"
                },
                "share": {
                    "$ref": "#/definitions/model.Share"
                }
            }
        },
        "model.Send": {
            "type": "object",
            "required": [
                "asset",
                "source",
                "value"
            ],
            "properties": {
                "asset": {
                    "type": "string"
                },
                "scale": {
                    "type": "integer",
                    "minimum": 0
                },
                "source": {
                    "$ref": "#/definitions/model.Source"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "model.Share": {
            "type": "object",
            "required": [
                "percentage"
            ],
            "properties": {
                "descWhatever": {
                    "type": "boolean"
                },
                "percentage": {
                    "type": "integer"
                },
                "percentageOfPercentage": {
                    "type": "integer"
                }
            }
        },
        "model.Source": {
            "type": "object",
            "required": [
                "from"
            ],
            "properties": {
                "from": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.FromTo"
                    }
                },
                "remaining": {
                    "type": "string"
                }
            }
        },
        "operation.Status": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "ACTIVE"
                },
                "description": {
                    "type": "string",
                    "maxLength": 256,
                    "example": "Active status"
                }
            }
        },
        "transaction.CreateTransactionInput": {
            "type": "object",
            "required": [
                "distribute",
                "send"
            ],
            "properties": {
                "chartOfAccountsGroupName": {
                    "type": "string",
                    "maxLength": 256
                },
                "code": {
                    "type": "string",
                    "maxLength": 100
                },
                "description": {
                    "type": "string",
                    "maxLength": 256
                },
                "distribute": {
                    "$ref": "#/definitions/model.Distribute"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "pending": {
                    "type": "boolean"
                },
                "send": {
                    "$ref": "#/definitions/model.Send"
                }
            }
        },
        "transaction.Status": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 100
                },
                "description": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "transaction.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "amountScale": {
                    "type": "number"
                },
                "assetCode": {
                    "type": "string"
                },
                "chartOfAccountsGroupName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "destination": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "ledgerId": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "operations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Operation"
                    }
                },
                "organizationId": {
                    "type": "string"
                },
                "parentTransactionId": {
                    "type": "string"
                },
                "source": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "$ref": "#/definitions/transaction.Status"
                },
                "template": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "transaction.UpdateTransactionInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 256
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3002",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Midaz Transaction API",
	Description:      "This is a swagger documentation for the Midaz Transaction API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
