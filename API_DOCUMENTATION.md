# API Documentation

###### Author: Patrick Archer (@pjarcher913)

###### Last Updated: 21 February 2020

This document contains...

## Getting Started

...

## Authentication

...

## Errors

...

## Requests

...

## Responses

...

## API Collection

### #01 HOME_POST

##### Handler: `/src/main.go :: prh_Home_POST(http.ResponseWriter, *http.Request)`

##### API URL: ```/:message```

HOME_POST returns a JSON response containing various messages.
`{"param"}` echoes `:message`.

Sample POST Request URL:
```
localhost:3000/giveMeEasterEgg
```

Response:
```
{
    "msg": "Hey, you found an Easter Egg!",
    "param": "giveMeEasterEgg",
    "time": "2020-02-20 22:48:48.0836616 +0000 UTC"
}
```

If multiple parameters are passed, the API will return a 404 error.

Sample POST Request URL:
```
localhost:3000/giveMeEasterEgg/secondParameter
```

Response:
```
404 page not found
```
