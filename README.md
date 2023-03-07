# http-echo

Simple server image that list information about the request in the response body. *Note* that header name case is changed to canonical format.

## Build

`podman build . --tag http-echo`

## Run

`podman run -it --rm -p 8080:8080 http-echo`

