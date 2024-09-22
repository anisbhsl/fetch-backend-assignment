Fetch Backend Assignment
----

This assignment consists of two endpoints to process a receipt and calculate its point. The original repo is [here](https://github.com/fetch-rewards/receipt-processor-challenge).

### To Run?

```
make run-app
```

The app will run at `localhost:3000`

### To test

```
make test
```

## Endpoints Available

| Endpoint | Method | Desc |
|----------|--------|--------|
| `/api/v1/receipts` | `Post` | stores a given receipt payload |
| `/api/v1/receipts/{id}/points` | `Get` | calculates the points for given receipt id|
