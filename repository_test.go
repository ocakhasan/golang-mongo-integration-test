package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositorySuite struct {
	suite.Suite
	repository   Repository
	testDatabase *TestDatabase
}

func (suite *RepositorySuite) SetupSuite() {
	suite.testDatabase = SetupTestDatabase()
	suite.repository = NewRepository(suite.testDatabase.DbInstance)
}

func (suite *RepositorySuite) TearDownSuite() {
	suite.testDatabase.TearDown()
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *RepositorySuite) TestCreateBook() {
	suite.Run("when id is not provided", func() {
		book := Book{
			Author: "Irvin D. Yalom",
			Title:  "Staring at the Sun: Overcoming the Terror of Death",
			Likes:  100,
		}

		createdBook, createBookErr := suite.repository.CreateBook(context.Background(), book)

		suite.Nil(createBookErr)
		suite.Equal(createdBook.Title, "Staring at the Sun: Overcoming the Terror of Death")
		suite.Equal(createdBook.Author, "Irvin D. Yalom")
		suite.False(createdBook.ID.IsZero())
	})

	suite.Run("when id is provided", func() {
		book := Book{
			ID:     primitive.NewObjectID(),
			Author: "Dostoyevksi",
			Title:  "Notes From the Underground",
			Likes:  100,
		}

		createdBook, createBookErr := suite.repository.CreateBook(context.Background(), book)

		suite.Nil(createBookErr)
		suite.Equal(createdBook, book)
	})
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *RepositorySuite) TestFindBook() {
	suite.Run("when there is no record", func() {
		id := primitive.NewObjectID()

		foundBook, findBookErr := suite.repository.FindBook(context.Background(), id)

		suite.Equal(findBookErr, mongo.ErrNoDocuments)
		suite.Nil(foundBook)
	})

	suite.Run("when there is record for given id", func() {
		book := Book{
			Author: "Dostoyevksi",
			Title:  "Notes From the Underground",
			Likes:  100,
		}

		createdBook, createBookErr := suite.repository.CreateBook(context.Background(), book)
		suite.Nil(createBookErr)

		id := createdBook.ID

		foundBook, findBookErr := suite.repository.FindBook(context.Background(), id)

		suite.Nil(findBookErr)
		suite.Equal(*foundBook, createdBook)
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
