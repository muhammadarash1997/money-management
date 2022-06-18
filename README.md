# Money Management
This app is for help user to make records of their income or expenses
## API Spec

### Register User
	Request:
	- Method: POST
	- Endpoint: "/api/user/register"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"name" "string",
			"email" "string",
			"password" "string"
		}
	Response:
	{
		"code": "integer",
		"status": "string",
		"data": "null"
	}

### Login User
	Request:
	- Method: POST
	- Endpoint: "/api/user/login"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"name" "string",
			"password" "string"
		}
	Response:
	{
		"code": "integer",
		"status": "string",
		"data": {
			"id": "string",
			"name": "string",
			"email": "string",
			"token": "string"
		}
	}

### Create Wallet
	Request:
	- Method: POST
	- Endpoint: "/api/wallet"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"user_id": "string",
			"name": "string",
			"currency": "string",
			"amount": "integer",
		}
	Response:
		{
			"code": "integer",
			"status": "string",
			"data": {
				"id": "string",
				"name": "string",
				"currency": "string",
				"amount": "integer",
			}
		}

### Get Transactions of Wallet
	Request:
	- Method: GET
	- Endpoint: "/api/transactions/:limit/:wallet_id"
	- Header:
		- Content-Type: application/json
	Response:
	{
		"code": "integer",
		"status": "string",
		"data": [
			{
				"id": "string",
				"amount": "integer",
				"category": "string",
				"note": "string",
				"time": "date"
			},
			{
				"id": "string",
				"amount": "integer",
				"category": "string",
				"note": "string",
				"time": "date"
			}
		]
	}

### Create Transaction
	Request:
	- Method: POST
	- Endpoint: "/api/transaction"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"wallet_id": "string",
			"amount": "integer",
			"category": "string",
			"note": "string",
			"time": "date"
		}
	Response:
		{
			"code": "integer",
			"status": "string",
			"data": "null"
		}

### Edit Transaction
	Request:
	- Method: PUT
	- Endpoint: "/api/transaction"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"id": "string",
			"wallet_id": "string",
			"amount": "integer",
			"category": "string",
			"note": "string",
			"time": "date"
		}
	Response:
	{
		"code": "integer",
		"status": "string",
		"data": "null"
	}

### Delete Transaction
	Request:
	- Method: PUT
	- Endpoint: "/api/transaction/:transaction_id"
	- Header:
		- Content-Type: application/json
	Response:
	{
		"code": "integer",
		"status": "string",
		"data": "null"
	}
