package main

import (
	"context"
	"log"
	"net/url"
	"topmovies/service/mongodb"

	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()

	// Initialize MongoDB client
	u := url.URL{
		Scheme: "mongodb",
		Host:   "127.0.0.1:27017",
		Path:   "movie", // database
		// User:   url.UserPassword("", ""),
	}
	mongoClient, err := mongodb.NewClient(&u)
	if err != nil {
		log.Fatalf("Error creating new mongodb client\n")
	}
	defer mongoClient.Disconnect(ctx)

	movieHandler := NewMovieHandler(mongoClient)

	router := gin.Default()
	router.GET("/top-movies/movies", movieHandler.GetMovies)
	router.GET("/top-movies/movies/:movieID", movieHandler.GetMovie)
	router.POST("/top-movies/movies", movieHandler.AddMovie)
	router.PUT("/top-movies/movies/:movieID/reviews", movieHandler.AddMovieReview)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
