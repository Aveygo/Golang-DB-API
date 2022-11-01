# Golang-DB-API
Simple Golang project that opens an endpoint to store key/value pairs 
```
go run .
```

Add a pair with the endpoint:

```
localhost:8080/set?key=your_key&value=your_value
```

And retrieve them with:

```
localhost:8080/get?key=your_key
```
