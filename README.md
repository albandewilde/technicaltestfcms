# technicaltestfcms

Technical test from [here](https://gist.github.com/rbretecher-fcms/9e10fbc8418cf7b52c0bc22bd4c2af83)

## Server

The server listen on the adress `0.0.0.0:8254`.

### Start the server

There is two way to start the server, natively with Go or in a container with Docker.

To run the server natively use the command `make run`. You'll need to have at least Go 1.17 installed.

To run the server in a container use the command `make ctn-run`. You'll need to have docker installed.

### API

#### Route `/`

Tell if the server is up, always return a string.

#### Route `/alert/`

To manage an alert

##### Method `POST`

Use the `POST` http method to post a new alert. The body is a json that represent an alert.

Ex:  
Http `POST` request with url `/alert/` and body:
```json
{
		"user_id": "bob",
		"name": "my_alerte",
		"min_price": 100,
		"max_price": 1000,
		"city": "Paris"
}
```
Willn't return any thing except an http code.

#### Route `/alerts/`

To manage alerts.

##### Method `GET`

Use the `GET` http method to get the list of all alerts of an user. The body with a json string is required to be the username.

Ex:  
Http `GET` request with url `/alerts/` and body:
```json
"bob"
```
Will return a json like follow:
```json
[
	{
		"id": "a",
		"user_id": "bob",
		"date": "0001-01-01",
		"name": "alrt1",
		"max_price": 10,
		"city": "Nantes"
	},
	{
		"id": "b",
		"user_id": "bob",
		"date": "0001-01-01",
		"name": "alrt2",
		"min_price": 1,
		"max_price": 50,
		"city": "Orléans"
	}
]
```
