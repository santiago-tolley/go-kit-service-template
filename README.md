# Go-Kit Service Template

The following code is a template for setting up micorservices using go-kit

# Setup
Initialize go modules
``` 
go mod init
```

# Running And Using The App
Start the server with: 
``` 
go run ./cmd/service/main.go 
```

Post to `localhost:8080/dostuff/` with the following JSON body
```
{
	"value": 2
}
```

The result should be 
``` 
{
    "result": 4,
    "err": null
}
```