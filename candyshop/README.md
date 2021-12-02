* cmd - allows for different driver actors 
* pkg
  * where crud service operations resides
    * adding
    * deleting
    * reading
    * updating
  * also includes router in `http` in pkg
    * seperate directory based on type of API(`rest` or `soap`)
      *  handler.go - where all endpoints are 
  * storage
    * provide sotrage units
    * [Cassandra](https://phoenixnap.com/kb/install-cassandra-on-ubuntu)
    * `docker run -p 9042:9042 --rm --name cassandra -d cassandra:latest`
        ```bash
              use candy_shop_db;
                * *********Todo: setup candy_shop_db keyspace in cassandra*********
            CREATE KEYSPACE candy_shop_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

            CREATE TABLE IF NOT EXISTS candies (
                candy_id uuid,
                category text,
                name text,
                price float,
                PRIMARY KEY (candy_id, category)
            );

            INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Chocolate', 'KitKat', 1.99);
            INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Sour', 'Skittles', 1.89);
            ***********************************************************************
          ```
    
