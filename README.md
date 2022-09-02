# gocql-simple-client-astra
Minimal client application that uses the GoCQL driver and connects to a [**DataStax Astra**](https://astra.datastax.com/) database using mTLS. 

This code was created in response to a post on the [**DataStax Community**](https://community.datastax.com/) site _"How do I connect to Astra with the gocql driver?"_.

## Prerequisites
1. To compile and run the code, download and install the Go language (version 1.18) - see https://go.dev/doc/install.

2. Install the gocql driver. From the root directory of this project, run:

```go get github.com/gocql/gocql@ce100a15a6899a7f42fbdc588874a36afcadc921```

Note: this is an unreleased commit hash with a fix that is needed when connecting to Astra. Once this fix is merged into GoCQL, you will just need to download the latest GoCQL driver (removing from the command above everything from the `@` onwards).

3. Create an [**Astra** database](https://astra.datastax.com). 

4. [Download the secure connect bundle](https://docs.datastax.com/en/astra/aws/doc/dscloud/astra/dscloudObtainingCredentials.html) for your Astra database.

   Unzip your copy of `secure-connect-your_astra_db.zip` which will contain the following files:
   ```
   ca.crt
   cert
   cert.pfx
   config.json
   cqlshrc
   identity.jks
   key
   trustStore.jks
   ```

   You will need these files to configure the SSL/TLS options.

   The `cqlshrc` file contains the connection details for `_cqlshrc_host` and `_cqlshrc_port` in the code. Here is an example:
   ```
   [connection]
   hostname = 31fecf38-2491-4d43-b6ce-22562679f1b8-us-east1.db.astra.datastax.com
   port = 34567
   ssl = true
   ```

5. Create a keyspace called `community` and create the following table:

    ```
    CREATE TABLE community.cities_by_rank (
        rank int PRIMARY KEY,
        city text,
        country text
    );
    ```
   
   Load the data in [`cities_by_rank.csv`](cities_by_rank.csv) using the cqlsh `COPY FROM` command:
   ```
   cqlsh> COPY cities_by_rank (rank,city,country) FROM './cities_by_rank.csv' WITH header = true;
   ```
6. Edit the `astra_gocql_connect.go` file replacing the missing values as indicated in the comments.


## Build and execute the code
Satisfy the required [Prerequisites](#prerequisites) in the above section.

Clone this repo, then build [`astra_gocql_connect.go`](astra_gocql_connect.go) with:
```
$ go build astra_gocql_connect.go
```

Finally, run the executable:
```
$ ./astra_gocql_connect
```

The expected output is:
```
According to independent.co.uk, the top 5 most liveable cities in 2019 were:
        Rank 1: Vienna, Austria
        Rank 2: Melbourne, Australia
        Rank 3: Sydney, Australia
        Rank 4: Osaka, Japan
        Rank 5: Calgary, Canada
```

## Sample data
The CSV file contains the top 10 most liveable cities of 2019 rated by The Independent.
```
 rank | city       | country
------+------------+-----------
    1 |     Vienna |   Austria
    2 |  Melbourne | Australia
    3 |     Sydney | Australia
    4 |      Osaka |     Japan
    5 |    Calgary |    Canada
    6 |  Vancouver |    Canada
    7 |    Toronto |    Canada
    8 |      Tokyo |     Japan
    9 | Copenhagen |   Denmark
   10 |   Adelaide | Australia
```

Source: Helen Coffey, 'This is the world's most liveable city', _The Independent_, 4 September 2019 (accessed 2 May 2020), https://www.independent.co.uk/travel/news-and-advice/vienna-city-best-quality-life-study-ranked-austria-sydney-melbourne-osaka-a9091016.html

## Credit
Huge thanks to [@dougwettlaufer](https://github.com/dougwettlaufer) and Erick Ramirez for contributing this solution.
