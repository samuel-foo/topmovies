# Highlights
- Clean REST APIs design
- MongoDB client save nested json review list for each movie
- MongoDB client exposing data ignostic interface

# External packages:
- Official golang MongoDB driver: "go.mongodb.org/mongo-driver/mongo"
- HTTP web api framework using Gin

# TopMovies
TopMovies is a movie reviews microservice. 

It provides endpoints to:
- save a movie
- list all movies
- get a single movie including its reviews
- adds a movie review

# Setup
1. Create a "movie" collection on your local MongoDB instance
2. Change directory to "cmd/topmovies" folder and run "go run main.go":
```bash
$ cd cmd/topmovies
$ go run main.go
```

# Save a movie
POST http://localhost:8080/top-movies/movies

Request body:
```json
{
	"Name": "Pheonix Soar",
	"StoryLine": "Some heroic movie",
	"Genre": "Drama",
	"Rating": "PG",
	"DirectedBy": "Jay Zee",
	"DateInTheaters": "2015-10-23T00:00:00Z",
	"RunTime": 120,
	"Studio": "Warners Bro"
}
```

# List all movies
GET http://localhost:8080/top-movies/movies

Response body:
```json
[
    {
        "ID": "5ea297c7e2d7d3f827378bc8",
        "Name": "Pheonix Soar",
        "StoryLine": "Some heroic movie",
        "Genre": "Drama",
        "Rating": "PG",
        "DirectedBy": "Jay Zee",
        "WrittenBy": "",
        "DateInTheaters": "2015-10-23T00:00:00Z",
        "DateOnDiscStreaming": "",
        "RunTime": 120,
        "Studio": "Warners Bro",
        "Reviews": [
            {
                "User": "Sam Johnson",
                "Rating": 5,
                "Comment": "Very good movie!!"
            },
            {
                "User": "Joe Lee",
                "Rating": 4,
                "Comment": "All right, not too bad"
            }
        ]
    },
    {
        "ID": "5ea4199fc228967f27dbc09b",
        "Name": "Star Wars: Episode IV - A New Hope",
        "StoryLine": "New Jedi is here",
        "Genre": "ScienceFiction",
        "Rating": "PG",
        "DirectedBy": "Geroge Lucas",
        "WrittenBy": "",
        "DateInTheaters": "1977-09-01T00:00:00Z",
        "DateOnDiscStreaming": "",
        "RunTime": 140,
        "Studio": "Dreamworks",
        "Reviews": [
            {
                "User": "Joe Williamson",
                "Rating": 3,
                "Comment": "Too much action, not much good story line."
            },
            {
                "User": "Lindsay Hunter",
                "Rating": 4,
                "Comment": "Good movie, very encouraging, good ending, like it!"
            }
        ]
    }
]
```

# Get a movie and its reviews
GET http://localhost:8080/top-movies/movies/5ea297c7e2d7d3f827378bc8

Response body:
```json
{
    "ID": "5ea297c7e2d7d3f827378bc8",
    "Name": "Pheonix Soar",
    "StoryLine": "Some heroic movie",
    "Genre": "Drama",
    "Rating": "PG",
    "DirectedBy": "Jay Zee",
    "WrittenBy": "",
    "DateInTheaters": "2015-10-23T00:00:00Z",
    "DateOnDiscStreaming": "",
    "RunTime": 120,
    "Studio": "Warners Bro",
    "Reviews": [
        {
            "User": "Sam Johnson",
            "Rating": 5,
            "Comment": "Very good movie!!"
        },
        {
            "User": "Joe Lee",
            "Rating": 4,
            "Comment": "All right, not too bad"
        }
    ]
}
```

# Add a movie review
PUT http://localhost:8080/top-movies/movies/5ea297c7e2d7d3f827378bc8/reviews

Request body:
```json
{
	"User": "Joe Lee",
	"Rating": 4,
	"Comment": "All right, not too bad"
}
```