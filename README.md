Fetch Backend Assignment
----

Submitted By: Anish Bhusal

This assignment consists of two endpoints to process a receipt and calculate its points. The original repo for assignment is [here](https://github.com/fetch-rewards/receipt-processor-challenge).

### To Run?

1. Using `go run`:

```
go run main.go
```

2. Using `Makefile`:

```
make run-app
```

The app will run at `localhost:3000`

3. Using docker:

```
docker build .
```

```
docker run -p 3000:3000 <image_id>
```

### To test

```
make test
```

## Endpoints Available

| Endpoint | Method | Desc |
|----------|--------|--------|
| `/api/v1/receipts/process` | `POST` | stores the given receipt payload |
| `/api/v1/receipts/{id}/points` | `GET` | calculates the points for given receipt id|

## Design Considerations

1. This webservice has been written in `Go` using `net/http` and `mux` as router.
2. Data is persisted in-memory using a simple key-value implementation.
3. Each endpoint handler follows request validation, business logic implementation and response components.
4. Routes are grouped together in `routes/routes.go` making it easier to navigate and maintain consistency.
5. Logging is done for each request using `zap` logger.
