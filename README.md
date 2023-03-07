# http-echo

Simple server image that list information about the request in the response body. Note* that header names are in canonical case.

# Run

`docker run -it --rm -e NAME=server1 -p 8080:8080 ghcr.io/attiand/http-echo:latest`

## Test

`curl http://localhost:8080`

## Optional env variables

- NAME: Specify a server name added to the response body
