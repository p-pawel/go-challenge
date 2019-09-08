# go-challenge


This project contains a solution to one recruitment challenge. Out of consideration of the possible parallel recruitment processes I am not quoting requirements. 

## Execution

Run:
```
docker-compose up --build
```
and API should be available at port 3000, e.g. http://localhost:3000/booking


## Test data available upon startup

### Destinations

| id  | Destination |
| --- | ------  |
| 1   | Mars
| 2   | Moon
| 3   | Pluto
| 4   | Asteroid Belt
| 5   | Europa
| 6   | Titan
| 7   | Ganymede


### Launchpads

| id  | Launchpad                   | SpaceX site code (if any) |
| --- | --------------------------- | ------------------------- |
| 1   | Le Centre Spatial Guyanais  | -                         |
| 2   | Космодро́м Байкону́р          | -                         |
| 3   | Vandenberg 3W               | `vafb_slc_3w`             |
| 4   | Vandenberg 4E               | `vafb_slc_4e`             |
| 5   | Cape Canaveral Historic 39A | `ccafs_slc_40`            |
| 6   | Cape Canaveral              | `ccafs_slc_40`            |
| 7   | Kwajalein Atoll             | `kwajalein_atoll`         |





## Examples

### Create booking

#### Happy scenario

URL:
```
POST http://localhost:3000/booking
```

Request:
```
{
	"FirstName": "John",
	"LastName": "Doe",
	"Gender": "M",
    "LaunchDate": "2019-09-08T21:40:00+02:00",
    "LaunchpadId": 5,
    "DestinationId": 2
}
```

Response (status code 200 - OK):
```
{
    "Id": 21,
    "FirstName": "John",
    "LastName": "Doe",
    "Gender": "M",
    "Birthday": "",
    "LaunchpadId": 5,
    "DestinationId": 2,
    "LaunchDate": "2019-09-08T21:40:00+02:00"
}
```


 
#### Trip destination not matching destinations for given launchpad for given day


URL:
```
POST http://localhost:3000/booking
```

Request:
```
{
	"FirstName": "John",
	"LastName": "Doe",
	"Gender": "M",
	"LaunchDate": "2019-09-08T21:40:00+02:00",
	"LaunchpadId": 5,
	"DestinationId": 6
}
```

Response (status code 422 - unprocessable entity):
```
["Destinations mismatch, requested destination #6 while for day #6 and pad #5 only destination #2 is allowed"]
```



#### Trip overlaps with another booking

URL:
```
POST http://localhost:3000/booking
```

Request:
```
{
	"FirstName": "John",
	"LastName": "Doe",
	"Gender": "M",
	"LaunchDate": "2019-09-08T21:40:00+02:00",
	"LaunchpadId": 5,
	"DestinationId": 2
}
```

Response (status code 422 - unprocessable entity):
```
["Launchpad #5 for day 2019-09-08 21:40:00 +0200 CEST is not available"]
```


#### Trip overlaps with SpaceX launches


URL:
```
POST http://localhost:3000/booking
```

Request:
```
{
	"FirstName": "John",
	"LastName": "Doe",
	"Gender": "M",
	"LaunchDate": "2019-10-17T21:10:00+02:00",
	"LaunchpadId": 5,
	"DestinationId": 6
}
```

Response (status code 422 - unprocessable entity):
```
["Launchpad #5 for day 2019-10-17 21:10:00 +0200 CEST is used by SpaceX"]
```

### Retrieve all bookings
URL:
```
GET http://localhost:3000/booking
```

Response (status code 422 - unprocessable entity):
```
[
    {
        "Id": 20,
        "FirstName": "John",
        "LastName": "Doe",
        "Gender": "M",
        "Birthday": "",
        "LaunchpadId": 5,
        "DestinationId": 5,
        "LaunchDate": "2019-10-16T19:10:00Z"
    },
    {
        "Id": 21,
        "FirstName": "John",
        "LastName": "Doe",
        "Gender": "M",
        "Birthday": "",
        "LaunchpadId": 5,
        "DestinationId": 2,
        "LaunchDate": "2019-09-08T19:40:00Z"
    }
]
```

### Delete booking
URL:
```
DELETE http://localhost:3000/booking/20
```

Response (status code 202 - accepted)


