# Forecast

Forecast processes the location & returns temperature for that location

## Usage 

For Windows/Linux without docker

```bash
go run *go
```
For Windows/Linux with docker

```bash
docker-compose up
```
## Examples
Post Forecast to recieve total price 
```bash
GET /forecast HTTP/1.1
User-Agent: PostmanRuntime/7.29.0
Host: localhost:8080
Accept-Language: en-us
Accept-Encoding: gzip, deflate
Connection: Keep-Alive
Co-Ordinates: 38.2527° N, 84.7585° W
```
## Tests
Tests have been added for several validations & errors, to check type
```bash
go test ../forecast/pkg/http
```