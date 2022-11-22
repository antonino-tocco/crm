# CRM API

## INSTALL DEPENDENCIES
The only dependency to install is: gorilla/mux. All other import are native go package.
You can install it with this command: `go get github.com/gorilla/mux`

## API

### GET CUSTOMERS
GET /customers

Return a list of customers

### GET CUSTOMER
GET /customers/{id}

Return the customer with the given id or 404 if not found

### ADD CUSTOMER
POST /customers

Add customer to the list

### UPDATE CUSTOMER
PATCH /customers/{id}

Update customer with the given id if it exists or return 404 if not found

### DELETE CUSTOMER
DELETE /customers/{id}

Delete customer with the given id if it exists or return 404 if not found