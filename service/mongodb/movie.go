package mongodb

// Movie object
type Movie struct {
	ID                  string `bson:"_id,omitempty"`
	Name                string
	StoryLine           string
	Genre               string
	Rating              string
	DirectedBy          string
	WrittenBy           string
	DateInTheaters      string
	DateOnDiscStreaming string
	RunTime             int16
	Studio              string
	Reviews             []*Review
}

// Review object
type Review struct {
	User    string
	Rating  uint8
	Comment string
}
