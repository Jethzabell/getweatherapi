# get weather API

A Golang app using [https://api.mapbox.com](https://api.mapbox.com) and [https://api.darksky.net](https://api.darksky.net) API

## Running Locally

Download, clone or fork this repository

```sh
git clone git@github.com:Jethzabell/getweatherapi.git # or clone your own fork
cd getweatherapi
go run main.go
```

Your app should now be running on [localhost:3000](http://localhost:3000/).

## Documentation

Know the ZipCode you want to use.
- You will use this as a query parameter; [localhost:3000?zipCode=inputZipCode](http://localhost:3000?zipCode=inputZipCode).

### Response Payload Example
Get Request  [localhost:3000?zipCode=27703](http://localhost:3000?zipCode=27703)

```json
{
		"placeName":"Durham, NC",
		"temperature":"74F",
		"summary":"Today is going to be sunny",
}
```
