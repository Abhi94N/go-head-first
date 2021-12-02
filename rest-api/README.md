
# GO Mod
* `go mod init package`W
* `go run` or `go mod tidy` will update dependency for you in `go.mod` file
* `go.sum` file shows you dependencies
* `go mod why -m {dependency}` 


## Example Endpoints
```bash
/albums

GET – Get a list of all albums, returned as JSON.
POST – Add a new album from request data sent as JSON.

/albums/:id
GET – Get an album by its ID, returning the album data as JSON.
```