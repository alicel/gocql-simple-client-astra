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

Install Go runtime (see https://go.dev/doc/install).

## Usage

To run this client, 
