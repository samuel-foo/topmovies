package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client is the db client for MongoDB
type Client interface {
	Disconnect(ctx context.Context) error
	SaveMovie(ctx context.Context, m *Movie) error
	GetMovies(ctx context.Context) ([]*Movie, error)
	GetMovie(ctx context.Context, id string) (*Movie, error)
	AddMovieReview(ctx context.Context, movieID string, review *Review) error
}

// NewClient returns a new MongoDB client
func NewClient(u *url.URL) (Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(u.String()))
	if err != nil {
		log.Printf("Error creating new mongodb client: %s\n", err)
		return nil, err
	}

	// ping to check if connection ok
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("Error pinging mongodb: %s\n", err)
		return nil, err
	}

	return &client{
		mongo:               mongoClient,
		defaultDatabaseName: u.Path,
	}, nil
}

// client is a concrete implementation of Client interface
type client struct {
	mongo               *mongo.Client
	defaultDatabaseName string
}

// Disconnect the mongodb client
func (c *client) Disconnect(ctx context.Context) error {
	if err := c.mongo.Disconnect(ctx); err != nil {
		log.Printf("Error disconnectiong mongodb client: %s\n", err)
		return err
	}

	return nil
}

// SaveMovie saves a movie
func (c *client) SaveMovie(ctx context.Context, m *Movie) error {

	col := c.defaultDatabase().Collection("movie")
	r, err := col.InsertOne(ctx, m)
	if err != nil {
		log.Printf("Error inserting movie: %s\n", err)
		return err
	}
	log.Printf("Movie insert success, inserted id=%v", r.InsertedID)

	return nil
}

// GetMovies returns movies
func (c *client) GetMovies(ctx context.Context) ([]*Movie, error) {
	curs, err := c.defaultDatabase().Collection("movie").Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Error retrieving movies: %s\n", err)
		return nil, err
	}
	defer curs.Close(ctx)

	movies := make([]*Movie, 0)
	for curs.Next(ctx) {
		mov := new(Movie)
		if err := curs.Decode(mov); err != nil {
			log.Printf("Error decoding movie: %s\n", err)
			return nil, err
		}
		movies = append(movies, mov)
	}

	if err := curs.Err(); err != nil {
		log.Printf("Error occurs while looping through movies: %s\n", err)
		return nil, err
	}

	return movies, nil
}

// GetMovie returns a single movie by movieID
func (c *client) GetMovie(ctx context.Context, movieID string) (*Movie, error) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	m := new(Movie)
	err := c.defaultDatabase().Collection("movie").FindOne(ctx,
		bson.M{"_id": id}).Decode(m)
	if err != nil {
		log.Printf("Error retrieving movie: %s\n", err)
		return nil, err
	}

	return m, nil
}

// AddMovieReview func
func (c *client) AddMovieReview(ctx context.Context, movieID string, review *Review) error {

	id, _ := primitive.ObjectIDFromHex(movieID)
	u := bson.M{
		"$push": bson.M{
			"reviews": review,
		},
	}
	result, err := c.defaultDatabase().Collection("movie").UpdateOne(ctx, bson.M{"_id": id}, u)
	if err != nil {
		log.Printf("Error adding movie review: %s\n", err)
		return err
	}
	if result.MatchedCount == 0 {
		msg := fmt.Sprintf("Error movie %s does not exist, review not added\n", movieID)
		log.Println(msg)
		return errors.New(msg)
	}

	return nil
}

// Returns default database
func (c *client) defaultDatabase() *mongo.Database {
	return c.mongo.Database(c.defaultDatabaseName)
}
