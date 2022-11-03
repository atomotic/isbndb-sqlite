# isbndb-sqlite

Load the [isbndb dump](http://pilimi.org/isbndb.html) from [Pirate Library Mirror](http://pilimi.org/) into SQLite.

1. [download](http://pilimi.org/isbndb-downloads.html) the jsonl dump (torrent via tor)
2. decompress it

    the file is ~19G

    	ls -lah isbndb_2022_09.jsonl 
    	-rw-rw-r-- 1 raffaele raffaele 19G nov  1 11:14 isbndb_2022_09.jsonl

    30 mil lines

    	wc -l isbndb_2022_09.jsonl 
    	30859865 isbndb_2022_09.jsonl

3. build

    	go build

4. run and wait a couple of hours

    	./isbndb-sqlite


5. result: you get a 25G SQLite (5.5 zstd compressed)

        -rw-r--r--  1 raffaele raffaele  25G nov  3 13:08 isbndb.db
        -rw-r--r--  1 raffaele raffaele 5,5G nov  3 13:08 isbndb.db.zst



## schema

        sqlite3 isbndb.db .schema
        CREATE TABLE isbndb (
            isbn text generated always as (json_extract(data, '$.isbn13')), 
            title text generated always as (json_extract(data, '$.title')), 
            data text
        );
        CREATE INDEX title_idx ON isbndb (title);
        CREATE INDEX isbn_idx ON isbndb (isbn);  

## queries
      
        sqlite> select count(*) from isbndb;
        30859865

        sqlite> select * from isbndb where isbn=9788899970222;
        9788899970222|La Gabbia Di Vetro|{"isbn": "889997022X", "msrp": "0.00", "pages": 265, "title": "La Gabbia Di Vetro", "isbn13": "9788899970222", "authors": ["Colin Wilson"], "language": "it", "title_long": "La Gabbia Di Vetro", "date_published": "2018"}


