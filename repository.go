package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateBook(ctx context.Context, book Book) (Book, error)
	FindBook(ctx context.Context, id primitive.ObjectID) (*Book, error)
}

func NewRepository(db *mongo.Database) *mongoRepository {
	return &mongoRepository{db: db}
}

type mongoRepository struct {
	db *mongo.Database
}

func (m *mongoRepository) CreateBook(ctx context.Context, book Book) (Book, error) {
	if book.ID.IsZero() {
		book.ID = primitive.NewObjectID()
	}

	_, err := m.db.Collection("books").InsertOne(ctx, book)
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (m *mongoRepository) FindBook(ctx context.Context, id primitive.ObjectID) (*Book, error) {
	var book Book
	filter := bson.M{
		"_id": id,
	}

	if err := m.db.Collection("books").FindOne(ctx, filter).Decode(&book); err != nil {
		return nil, err
	}

	return &book, nil
}
