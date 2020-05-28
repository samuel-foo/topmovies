package main

import (
	"log"
	"net/http"
	"regexp"
	"topmovies/service/mongodb"

	"github.com/gin-gonic/gin"
)

// NewMovieHandler instantiates a movie handler
func NewMovieHandler(mongoClient mongodb.Client) *MovieHandler {
	return &MovieHandler{
		mongoClient: mongoClient,
	}
}

// MovieHandler is a http handler for Movie object
type MovieHandler struct {
	mongoClient mongodb.Client
}

var (
	regexMovieID        = regexp.MustCompile(`/top-movies/movies/([[:xdigit:]]{24})`)
	regexMovieReviewPut = regexp.MustCompile(`/top-movies/movies/([[:xdigit:]]{24})/reviews`)
)

// GetMovies get a list of movies
func (h *MovieHandler) GetMovies(ctx *gin.Context) {
	movies, err := h.mongoClient.GetMovies(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, movies)
}

// GetMovie gets a single movie
func (h *MovieHandler) GetMovie(ctx *gin.Context) {
	movieID := ctx.Param("movieID")
	movie, err := h.mongoClient.GetMovie(ctx, movieID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

// AddMovie adds a movie
func (h *MovieHandler) AddMovie(ctx *gin.Context) {
	m := new(mongodb.Movie)
	if err := ctx.ShouldBindJSON(m); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Movie received: %+v\n", m)

	if err := h.mongoClient.SaveMovie(ctx, m); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// AddMovieReview adds a movie review
func (h *MovieHandler) AddMovieReview(ctx *gin.Context) {
	movieID := ctx.Param("movieID")

	review := new(mongodb.Review)
	if err := ctx.ShouldBindJSON(review); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.mongoClient.AddMovieReview(ctx, movieID, review); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
