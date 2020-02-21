# API Documentation

###### Author: Patrick Archer (@pjarcher913)

###### Last Updated: 21 February 2020

###### Copyright (c) 2020 Patrick Archer

This document contains...

## TABLE OF CONTENTS

...

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

### #01 POST_HOME

##### Handler: `/src/main.go :: prh_POST_Home(http.ResponseWriter, *http.Request)`

##### API URL: ```/:message```

POST_HOME returns a JSON response of 3 root keys with 3 corresponding string values.

Even though this API takes a POST request, nothing actually gets posted in any database.
The trigger simply requires a POST request with one parameter appended to the home route URL,
but the request gets functionally-treated like a GET.

- `{msg}` always returns "Hey, you found an Easter Egg!".
- `{"param"}` echoes `:message`.
- `{"time"}` returns the timestamp that the server completed drafting its response to the http request.

__*Sample Request URL:*__
```
localhost:3000/giveMeEasterEgg
```

__*Response:*__
```
{
    "msg": "Hey, you found an Easter Egg!",
    "param": "giveMeEasterEgg",
    "time": "2020-02-20 22:48:48.0836616 +0000 UTC"
}
```

If multiple parameters are passed, the response will be a 404 error.

__*Sample Request URL:*__
```
localhost:3000/giveMeEasterEgg/secondParameter
```

__*Response:*__
```
404 page not found
```
