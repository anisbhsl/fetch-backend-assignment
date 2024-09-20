Requirements:

Design API endpoints:

- `/receipts/process`
POST
Payload: JSON
response: JSON containing id for the receipt

use UUID for the ID

- The ID returned will be passed into `/receipts/{id}/points` to get the number of points the receipt was awarded
GET
response: JSON object containing the number of points awarded

Can use in-memory database

There are rules to define how many points should be awarded to a receipt
