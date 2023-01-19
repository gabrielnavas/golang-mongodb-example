package models

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}
