## Building
## Local and Run
```
go build -o service && ./service
```
## Docker Image
```
GOOS=linux CGO_ENABLED=0 go build -installsuffix cgo -ldflags '-extldflags "-static"' -o service
```

## Curl Quick Tests
### Blocked
```
curl --request POST \
  --header "Content-Type:application/json" \
  --data '{"ip":"1.0.0.1", "countries":["US"]}' \
  http://localhost:8080
```
### Allowed
curl --request POST \
  --header "Content-Type:application/json" \
  --data '{"ip":"1.0.0.1", "countries":["US", "AU"]}' \
  http://localhost:8080
```