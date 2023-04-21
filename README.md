# http-echo

Simple server image that list information about the request in the response body. *Note* that header names are in canonical case.

# Run

`docker run -it --rm -e NAME=server1 -p 8080:8080 ghcr.io/attiand/http-echo:latest`

## Test

`curl http://localhost:8080 | jq .`
```json
{
  "name": "server1",
  "request": {
    "url": "/",
    "method": "GET",
    "headers": {
      "Accept": [
        "*/*"
      ],
      "User-Agent": [
        "curl/7.76.1"
      ]
    }
  }
}
```

## Set Cookies
`curl -I -c cookies http://localhost:8080/setcookie?JSESSIONID=ABC123`

```
HTTP/1.1 200 OK
Set-Cookie: JSESSIONID=ABC123
Date: Fri, 21 Apr 2023 11:36:48 GMT
```

`curl -b cookies http://localhost:8080 | jq .`

```json
{
  "name": "server1",
  "request": {
    "url": "/",
    "method": "GET",
    "headers": {
      "Accept": [
        "*/*"
      ],
      "Cookie": [
        "JSESSIONID=ABC123"
      ],
      "User-Agent": [
        "curl/7.87.0"
      ]
    }
  }
}

```

## Optional environment variables

- NAME - Specify a server name added to the response body
