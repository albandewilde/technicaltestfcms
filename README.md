# technicaltestfcms

Technical test from [here](https://gist.github.com/rbretecher-fcms/9e10fbc8418cf7b52c0bc22bd4c2af83)

## Server

The server listen on the adress `0.0.0.0:8254`.

### API

#### Route `/`

Tell if the server is up, always return a string.

#### Route `/alert/`

To manage an alert

##### Method `POST`

Use the `POST` http method to post a new alert.

#### Route `/alerts/`

To manage alerts.

##### Method `GET`

Use the `GET` http method to get the list of all request.
