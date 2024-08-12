package tests

import (
	"context"
	"fmt"
	"goPractice/task_manager/domain"
	"goPractice/task_manager/repositories"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepoSuite struct {
	suite.Suite

	repository repositories.UserRepository
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func (suite *UserRepoSuite) SetupSuite() {
	fmt.Println("Setting up test suite")
	options := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), options)
	require.NoError(suite.T(), err, "Failed to connect to MongoDB")

	suite.client = client
	db := suite.client.Database("task-manager-test")
	suite.db = db
	suite.collection = db.Collection("users")
	suite.repository = repositories.NewUserRepository(db)
}
func (suite *UserRepoSuite) TearDownSuite() {
	//drop user collection
	defer suite.client.Disconnect(context.TODO())
	defer suite.collection.Drop(context.TODO())
}

// delete all users after each test
func (suite *UserRepoSuite) TearDownTest() {

	//delete all users after each test
	suite.collection.DeleteMany(context.TODO(), bson.M{})
}
func (suite *UserRepoSuite) TestCreateUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser1",
		Password: "testpassword",
	}
	_, err := suite.repository.CreateUser(user)
	require.NoError(suite.T(), err, "Failed to create user")

}
func (suite *UserRepoSuite) TestGetUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser2",
		Password: "testpassword",
	}
	_, err := suite.repository.CreateUser(user)
	require.NoError(suite.T(), err, "Failed to create user")

	user, err = suite.repository.FilterUser(user.Username)
	require.NoError(suite.T(), err, "Failed to get user")
	require.Equal(suite.T(), user.ID, user.ID)
}

func (suite *UserRepoSuite) TestPromoteUser() {

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser3",
		Password: "testpassword",
		Role:     "user",
	}
	_, err := suite.repository.CreateUser(user)
	require.NoError(suite.T(), err, "Failed to create user")
	//promote it
	err = suite.repository.PromoteUser(user.ID.Hex())
	require.NoError(suite.T(), err, "Failed to promote user")
	//get it
	promotedUser, err := suite.repository.FilterUser(user.Username)
	fmt.Println(promotedUser)
	require.NoError(suite.T(), err, "Failed to get user")

	require.Equal(suite.T(), "admin", promotedUser.Role)

}
func (suite *UserRepoSuite) TestGetUsers() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser4",
		Password: "testpassword",
		Role:     "user",
	}
	_, err := suite.repository.CreateUser(user)
	require.NoError(suite.T(), err, "Failed to create user")
	users, err := suite.repository.GetUsers()
	require.NoError(suite.T(), err, "Failed to get users")

	require.Equal(suite.T(), 1, len(users))
}
func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserRepoSuite))
}
