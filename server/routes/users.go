package routes

import (
	"context"
	"fmt"
	"time"

	authUser "server/auth"
	models "server/models"
	users "server/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usersCollection *mongo.Collection = OpenCollection(Client, "users")

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func GetUser(user *authUser.LoginRequest) (*users.User, error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	userFromDB := bson.M{}
	dbTmp := &models.DBUser{}
	authUser := &users.User{}
	filter := bson.D{{Key: "email", Value: user.Email}}

	err := usersCollection.FindOne(ctx, filter).Decode(&userFromDB)
	if err != nil {
		fmt.Println(err)
		defer cancel()
		return nil, err
	} else {
		bsonBytes, _ := bson.Marshal(userFromDB)
		bson.Unmarshal(bsonBytes, dbTmp)
		authUser.Id = dbTmp.ID.Hex()
		authUser.Email = *dbTmp.Email
		authUser.Password = *dbTmp.HashedPassword
		errHash := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(user.Password))
		if errHash != nil {
			fmt.Println(errHash)
			defer cancel()

			return nil, errHash
		}
	}

	defer cancel()
	return authUser, nil
}

func CreateUser(user *users.CreateUserRequest) bool {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	userFromDB := bson.M{}
	filter := bson.D{{Key: "email", Value: user.User.Email}}

	err := usersCollection.FindOne(context.Background(), filter).Decode(&userFromDB)
	if err == nil {
		fmt.Println(err)
		defer cancel()
		return false

	} else {
		fmt.Println("User creation check", err)
	}

	hashedPassword, hashError := HashPassword(user.User.Password)

	if hashError != nil {
		fmt.Println(hashError)
		defer cancel()

		return false
	}

	dbUser := models.DBUser{
		ID:             primitive.NewObjectID(),
		Email:          &user.User.Email,
		HashedPassword: &hashedPassword,
	}

	_, insertErr := usersCollection.InsertOne(ctx, dbUser)
	if insertErr != nil {
		fmt.Println(insertErr)
		defer cancel()
		return false
	}

	defer cancel()
	return true
}

func GetUsers() []*users.User {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var usersDB []bson.M

	cursor, err := usersCollection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)

		defer cancel()

		return nil
	}

	if err = cursor.All(ctx, &usersDB); err != nil {
		fmt.Println(err)

		defer cancel()

		return nil
	}

	defer cancel()

	responseUsers := []*users.User{}

	for _, user := range usersDB {
		bsonBytes, _ := bson.Marshal(user)
		respTmp := &users.User{}
		dbTmp := &models.DBUser{}
		bson.Unmarshal(bsonBytes, dbTmp)
		respTmp.Id = dbTmp.ID.Hex()
		respTmp.Email = *dbTmp.Email
		respTmp.Password = *dbTmp.HashedPassword
		responseUsers = append(responseUsers, respTmp)

	}
	return responseUsers
}

func DeleteUser(user *users.DeleteUserRequest) bool {

	userID := user.Id
	docID, _ := primitive.ObjectIDFromHex(userID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	_, err := usersCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		fmt.Println(err)

		defer cancel()

		return false
	}

	defer cancel()

	return true

}
