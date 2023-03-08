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

## Optional environment variables

- NAME - Specify a server name added to the response body
