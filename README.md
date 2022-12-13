# Introduction

This API was built for my Intro to APIs elective as my final project. It is used for retrieving information on players in the NFL.

I created 3 scrapers in Python to scrape from the NFL website and populate a mySQL database, which then I used GO to interact with the DB and send requested information (only for educational purposes)

## Use Cases

There are many reasons to use this API, most common being to retrieve NFL data on teams, players or player stats for a mobile or web application


## Endpoints

### Teams

#### Example Response Object
```javascript
{
    "response": [
        {
            "TEAM_KEY": 1,
            "Name": "Los Angeles Rams",
            "City": "Los Angeles",
            "Website": "https://www.therams.com/",
            "Wins": 3,
            "Losses": 9,
            "Ties": 0,
            "Position": "4th",
            "Division": "NFC West",
            "Coach": "Sean McVay",
            "Stadium": "SoFi Stadium",
            "Owners": "Stan Kroenke",
            "Established": 1937
        },
        {
            "TEAM_KEY": 2,
            "Name": "New York Giants",
            "City": "New York",
            "Website": "https://www.giants.com/",
            "Wins": 7,
            "Losses": 4,
            "Ties": 1,
            "Position": "3rd",
            "Division": "NFC East",
            "Coach": "Brian Daboll",
            "Stadium": "MetLife Stadium",
            "Owners": "John Mara",
            "Established": 1925
        },
        {
            "TEAM_KEY": 3,
            "Name": "Philadelphia Eagles",
            "City": "Philadelphia",
            "Website": "https://www.philadelphiaeagles.com/",
            "Wins": 11,
            "Losses": 1,
            "Ties": 0,
            "Position": "1st",
            "Division": "NFC East",
            "Coach": "Nick Sirianni",
            "Stadium": "Lincoln Financial Field",
            "Owners": "Jeffrey Lurie",
            "Established": 1933
        },
        ...
    ],
    "statusDescription": "Success!",
    "status_code": 200
}
```

#### Routes

Return all teams
```http
GET /api/teams/
``` 

Return Team By Id
```http
GET /api/teams/<id>
GET /api/teams/1
```

Return Team By Name
```http
GET /api/teams/<name>
GET /api/teams/Seattle-Seahawks
```

### Players

#### Example Response Object
```javascript
{
    "response": [
        {
            "PLAYER_KEY": 1,
            "Name": "Genard Avery",
            "Jersey_Number": 59,
            "Position": "OLB",
            "Status": "ACT",
            "Height": 72,
            "Weight": 250,
            "Experience": 5,
            "College": "Memphis",
            "TEAM_KEY": 17
        },
        {
            "PLAYER_KEY": 2,
            "Name": "Kenjon Barner",
            "Jersey_Number": 38,
            "Position": "RB",
            "Status": "CUT",
            "Height": 69,
            "Weight": 195,
            "Experience": 8,
            "College": "Oregon",
            "TEAM_KEY": 17
        },
        ...
       
      ],
    "statusDescription": "Success!",
    "status_code": 200
```

#### Routes
Get All Players
```http
GET /api/players/
```

Get Player By ID
```http
GET /api/players/<id>
GET /api/players/6
```

Get Player By Team ID
```http
GET /api/players/teamid/<teamId>
GET /api/players/teamid/1
```

Get Player By Name
```http
GET /api/players/<name>
GET /api/players/Sauce-Gardner
```

Get Player By Status
```http
GET /api/players/status/<status>
GET /api/players/status/ACT
```

Get Player By Team Name
```http
GET /api/players/team/<teamName>
GET /api/players/team/Cincinnati-Bengals
```

Get Players By Position
```http
GET /api/players/position/<Position>
GET /api/players/position/QB
```


### Stats

#### Example Response Object
```javascript
{
    "response": {
        "PLAYER_KEY": 400,
        "Pass_Yds": 4160,
        "Pass_Yds_Att": 0,
        "Pass_Att": 511,
        "Pass_Cmp": 336,
        "Pass_Cmp_Percent": 0,
        "Pass_TD": 33,
        "Pass_Int": 0,
        "Pass_Rate": 103.4,
        "Pass_1st": 216,
        "Pass_1st_Percent": 0,
        "Pass_20_Plus": 0,
        "Pass_40_Plus": 0,
        "Pass_Lng": 57,
        "Pass_Sck": 21,
        "Pass_Sck_Yds": 0,
        "Rush_Yds": 280,
        "Rush_Att": 47,
        "Rush_TD": 2,
        "Rush_20_Plus": 0,
        "Rush_40_Plus": 0,
        "Rush_Lng": 20,
        "Rush_1st": 19,
        "Rush_1st_Percent": 0,
        "Rush_FUM": 3
    },
    "statusDescription": "Success!",
    "status_code": 200
}
```

#### Routes
Get All Stats
```http
GET /api/stats/
```

Get Stats By Player ID
```http
GET /api/stats/<id>
GET /api/stats/400
```

Get Stats By Player Name
```http
GET /api/stats/<name>
GET /api/stats/Patrick-Mahomes
```






The `statusDescription` attribute contains the status of the request, indicating whether it was successful or if there was some error
The `status_code` gives the HTTPS status code of the request.

The `response` attribute contains any data that was received from the request. Stored in JSON objects. It is either a single JSON object or a list of JSON objects depending on the request.


## Status Codes

The API returns the following status codes:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |

