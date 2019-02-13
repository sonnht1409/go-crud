# go-crud

## Install

"github.com/gin-gonic/gin"

"gopkg.in/mgo.v2/bson"

"gopkg.in/mgo.v2"

## Run

go run main.go

## APIs

### Find by username
```
curl -X GET \
  http://localhost:8080/v1/users/abc \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"username":"ab",
	"password":"123"
}'
```

### Sign up
```
curl -X POST \
  http://localhost:8080/v1/users/sign-up \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"username":"ab",
	"password":"123"
}'
```

### Sign in
```
curl -X POST \
  http://localhost:8080/v1/users/sign-in \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"username":"abc",
	"password":"123"
}'
```
