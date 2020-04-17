# sql

adapter for database/sql package.

Hides overhead of error handling and usage of vendor pkg.

Should be internal pkg.

keep it as reusable as possible.

Vision:

for now there is only Insert(), but idea was to provide at least the most common sql operations in a way 
which can be reused in any project

## main interfaces

- Client: adapter interface to hide direct dependency on vendor client.
- Result: Adapter for vendor sql result. contains basic methods like number of affected rows and last isnerted id


## current implementations

- MySqlClient
    - implements Client
    - depends on DB
    - hides overhead of vendor client's usage
    - **SHOULD** hide direct dependencies on database/sql
- MySqlResult
    - implements Result
    - simple implementation which only delegates logic to internal dependency
    - embeds database/sql.Result
- DB
    - implements no interface
    - embeds *database/sql.DB
    - creates a new DB connection
    - provides Close() to close connection when necessary
    
## dependencies

- mariadb: see docker-compose
- mysql: see docker-compose

## tests
you need one of the containers running and make sure the test is using the correct url


go 1.13: `go test -tags=integration ./...`

go < 1.13 `GO111MODULE=on go test -tags=integration`
