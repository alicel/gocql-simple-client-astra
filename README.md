# gocql-simple-client-astra
Minimal client application that uses the GoCQL driver and connects to Astra using mTLS. Taken from this example https://community.datastax.com/questions/3753/has-anyone-managed-to-connect-to-astra-from-gocql.html.

## Preparation
Create a keyspace called `community` and create the following table:

```
CREATE TABLE community.cities_by_rank (
    rank int PRIMARY KEY,
    city text,
    country text
);
```
Populate this table with the data in cities_by_rank.csv:

``` COPY community.cities_by_rank FROM 'cities_by_rank.csv' WITH DELIMITER=',' AND HEADER=TRUE; ```

Install Go version 1.18 (see https://go.dev/doc/install).

From the root directory of this project, download the appropriate version of GoCQL by running:

```go get github.com/gocql/gocql@ce100a15a6899a7f42fbdc588874a36afcadc921```

Note: this is an unreleased commit hash with a fix that is needed when connecting to Astra. Once this fix is merged into GoCQL, you will just need to download the latest GoCQL driver (removing from the command above everything from the `@` onwards).

Download the Secure Connect Bundle for your Astra cluster and unzip it somewhere in a dedicated directory.

Edit the `astra_gocql_connect.go` file replacing the missing values as indicated in the comments.

## Usage

To run this client, build it with `go build astra_gocql_connect.go` and then run the executable `./astra_gocql_connect`.

The expected output is:
```
According to independent.co.uk, the top 5 most liveable cities in 2019 were:
        Rank 1: Vienna, Austria
        Rank 2: Melbourne, Australia
        Rank 3: Sydney, Australia
        Rank 4: Osaka, Japan
        Rank 5: Calgary, Canada
```
