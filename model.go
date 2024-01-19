package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id"`
	Author string             `bson:"author"`
	Title  string             `bson:"title"`
	Likes  int                `bson:"likes"`
}
