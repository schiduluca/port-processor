# ports-processor
The purpose of this service is to parse a large JSON file and insert the records into a storage of choice. For this example Redis was used.

### Steps to run the service
At the root of the project there's a `docker-compose.yaml` file which can be used to start the project locally.
To start the project just run the `docker-compose up` command.

Once the service will finish processing the JSON file, the container will shut down with a message `finished processing` (if no errors happened), while the redis container will still be running.

Because of time limitation the testing should be done manually to see if actually the data has been stored.
1. ``docker exec  -it {redis-container-id}  /bin/sh``
2. ``redis-cli``
3. do the redis queries for verifying: `dbsize`, `get`, etc.