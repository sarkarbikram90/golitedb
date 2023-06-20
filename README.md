# golitedb
GoliteDB is a simple in-memory database
It exposes /record for CRUD operations and /record/query for querying.
The database is represented by a map called database, where the key is the record ID, and the value is an instance of the Record struct.
To run the code, save it to a file (e.g., golitedb.go) and execute go run golitedb.go in your terminal. 
The server will start on port 8080.
You can interact with the database using HTTP requests. 
For example, you can create a new record using a POST request to http://localhost:8080/record with a JSON payload containing the record data. 
To retrieve a record, send a GET request to http://localhost:8080/record?id=<record_id>. 
Similarly, you can update or delete a record using PUT and DELETE requests, respectively. 
The /record/query endpoint allows you to perform a simple text-based search for records that contain a specific query string.
