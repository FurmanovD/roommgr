# RoomManager service

## Business problem
Two companies, COKE and PEPSI, are sharing an office building but as they are competitors, they don’t trust each other. Tomorrow is COLA day (for one day), that the two companies are celebrating. They are gathering a number of business partners in the building. In order to optimize space utilization, they have decided to set-up a joint booking system where any user can book one of the 20 meeting rooms available, 10 from each company **(C01, C02, ..., C10 and P01, P02, ...., P10)**.

The booking system has the following functionalities:

- Users can see meeting rooms availability
- Users can book meeting rooms by the hour (first come first served)
- Users can cancel their own reservations

## Config
Service's configuration example is in the file: **./config/example-config.ini**

**NOTE:** For now it contains configuration **secrets** that **must not be used** in production and should be changed to loading secrets from external secure storage like Hashicorp Vault.

To run the service locally, create a copy of **example-config.ini** and rename it to **config.ini**

## Run
To run the system you need to run the following command: `make run`
It will spin up the docker containers of a Database (MySQL) and a RoomManager service.

To access the database, you can use address: `localhost:3306`
To access the service API, send HTTP requests to the address: `localhost:8081`

To debug from IDE, you can run DB container alone by run command: `make run-deps`

Stop all the fun: `make stop`

## Initial DB data
After the very first run of the Database, a manual operation of "migration applying" is required.
Using any database/sql client, apply scripts from files: **db/migration/0/1_schema.sql** and **db/migration/0/1_schema.sql**
MySQL container will bind the local **./test/db/** folder to save all data, so you can stop the system at any time - the data will persist.

## API
The base path for the API v1 supported is: **http://localhost:8081/api/v1**
The next listed endpoints supported atm.:
- **GET /rooms** - returns a list of all rooms registered in the system
- **GET /rooms/<roomID>** - returns a room by ID
- **GET /rooms/:roomID/reservations** - returns all reservations of the **roomID** specified
- **GET /rooms/:roomId/reservations/:rsrvID** - returns a room reservation by ID
- **POST /rooms/-/reservations** - creates a new reservation getting all the information required from the POST body of JSON view:
    ```
    {
        "user_token": 3,
        "room_id": 1,
        "start_time": "2022-02-17T08:00:00Z"
    }
    ```
- **DELETE /rooms/-/reservations/:rsrvID** - marks reservation as deleted using its ID and userID that created the reservation. The body is JSON of view:
    ```
    {
        "user_token": 3
    }
    ```

**NOTE:** Examples of tha API usage are in Postman collection file: **test/postman/RoomService.postman_collection.json**