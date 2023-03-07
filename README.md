# http-echo

Simple server image that list information about the request in the response body. Currently only header information. *Note* that header names are in canonical case.

# Run

`docker run -it --rm -p 8080:8080 ghcr.io/attiand/http-echo:latest`

## Optional env variables

- NAME: Specify a server name added to the response body
