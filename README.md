# GetWeatherAPI

A Golang app using [https://api.mapbox.com](https://api.mapbox.com) and [https://api.darksky.net](https://api.darksky.net) API

## Running Locally

Download, clone or fork this repository

```sh
git clone git@github.com:Jethzabell/getweatherapi.git # or clone your own fork
cd getweatherapi
go run main.go
```

Your app should now be running on [localhost:1323](http://localhost:1323/).

## Documentation

Know the ZipCode you want to use.
- You will use this as a query parameter; [localhost:1323?zipcode=inputZipCode](http://localhost:1323?zipcode=inputZipCode).

### Response Payload Example
Get Request  [localhost:1323?zipcode=27703](http://localhost:1323?zipcode=27703)

```json
{
		"placeName":"Durham, NC",
		"temperature":"74F",
		"summary":"Today is going to be sunny",
}
```

### User and Technical Documentation
[https://drive.google.com/open?id=1E_zEND0ZqzUZdj5EQ3VDwwdlbb01HMKwBWwhGT_I5mI](https://drive.google.com/open?id=1E_zEND0ZqzUZdj5EQ3VDwwdlbb01HMKwBWwhGT_I5mI)
