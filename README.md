# Receipt Processor Challenge

## Basic Usage

1. Run `go run .`
1. The server will be mounted at `localhost:8080`.

## Testing

1. Run `go test`.

Tests exist that cover the logic for calculating points of a given receipt
(`calculate_points_test.go`) and coverage of the basic API (`server_test.go`).

## cURL Tests

If you've got `node` and `python` installed, you can run `./test.sh`, which is a
quickly-builty script designed to test the API iteratively and quickly. Postman,
Thunder Client, etc. are all options available, but this was quick enough to
scrap together that it seemed unwarranted to go much further at this point.
