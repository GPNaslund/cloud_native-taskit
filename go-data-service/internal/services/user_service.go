package service

import (
	"context"
	"log"

	"gn222gq.2dv013.a2/internal/dto"
	"gn222gq.2dv013.a2/internal/models"
	pb "gn222gq.2dv013.a2/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Represents the user service
type UserService struct {
	db        *mongo.Client
	dbDetails dto.DatabaseDetails
}

// Constructor function
func NewUserService(db *mongo.Client, dbDetails dto.DatabaseDetails) *UserService {
	return &UserService{
		db:        db,
		dbDetails: dbDetails,
	}
}

// Function for creating a user
func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("Handling create user request")
	coll := u.db.Database(u.dbDetails.DatabaseName).Collection(u.dbDetails.UserCollectionName)

	// Check if username is allready taken
	filter := bson.D{{"username", req.Username}}
	var result models.User
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err == nil {
		log.Println("Username is allready taken")
		return &pb.CreateUserResponse{Status: pb.DataStatus_Data_Invalid_Username}, nil
	}

	// Create a new user
	newUser := models.User{Username: req.Username, Password: req.Password, Tasks: []models.Task{}}
	insertResult, err := coll.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Datbase error on creating a new user: %s", err.Error())
		return &pb.CreateUserResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Failed to convert inserted id to object id")
		return &pb.CreateUserResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}
	id := oid.Hex()

	userDto := &pb.UserDTO{UserId: &id, Username: req.Username, Password: req.Password}
	return &pb.CreateUserResponse{Status: pb.DataStatus_Data_Success, User: userDto}, nil
}

// Function for getting a user
func (u *UserService) ReadUser(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	log.Println("Handling read user request")
	coll := u.db.Database(u.dbDetails.DatabaseName).Collection(u.dbDetails.UserCollectionName)

	filter := bson.D{{"username", req.Username}}
	var result models.User
	err := coll.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No user found with provided username")
			return &pb.ReadUserResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
		}
		log.Printf("Database error on trying to fetch a user: %s", err.Error())
		return &pb.ReadUserResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	id := result.Id.Hex()
	userDto := &pb.UserDTO{UserId: &id, Username: result.Username, Password: result.Password}
	return &pb.ReadUserResponse{Status: pb.DataStatus_Data_Success, User: userDto}, nil
}

// Function for updating a user
func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Println("Handling update user request")
	coll := u.db.Database(u.dbDetails.DatabaseName).Collection(u.dbDetails.UserCollectionName)

	objId, err := primitive.ObjectIDFromHex(*req.User.UserId)
	if err != nil {
		log.Printf("Failed to convert user id: %s, to object id", *req.User.UserId)
		return &pb.UpdateUserResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", bson.D{{"password", req.User.Password}}}}

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Database error on trying to update user: %s", err.Error())
		return &pb.UpdateUserResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	if result.ModifiedCount == 0 {
		log.Println("Database could not update user")
		return &pb.UpdateUserResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
	}

	return &pb.UpdateUserResponse{Status: pb.DataStatus_Data_Success, User: req.User}, nil
}

// Function for deleting a user
func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Println("Handling delete user request")
	coll := u.db.Database(u.dbDetails.DatabaseName).Collection(u.dbDetails.UserCollectionName)
	objId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Printf("Failed to convert user id: %s, to object id", req.UserId)
		return &pb.DeleteUserResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	filter := bson.D{{"_id", objId}}
	result, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		log.Printf("Database error on trying to delete user: %s", err.Error())
		return &pb.DeleteUserResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	if result.DeletedCount == 0 {
		log.Println("Database failed to delete user")
		return &pb.DeleteUserResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
	}

	return &pb.DeleteUserResponse{Status: pb.DataStatus_Data_Success}, nil
}
