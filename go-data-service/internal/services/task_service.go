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

// Represents the task service
type TaskService struct {
	db        *mongo.Client
	dbDetails dto.DatabaseDetails
}

// Constructor function
func NewTaskService(db *mongo.Client, dbDetails dto.DatabaseDetails) *TaskService {
	return &TaskService{
		db:        db,
		dbDetails: dbDetails,
	}
}

// Function for creating a new task for a user
func (t *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	log.Println("Handling creating task request")
	coll := t.db.Database(t.dbDetails.DatabaseName).Collection(t.dbDetails.UserCollectionName)

	objId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Printf("Failed to convert: %s, to object id", req.UserId)
		return &pb.CreateTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	// Create new task with its own ObjectID
	newTask := models.Task{
		Id:      primitive.NewObjectID(),
		Title:   req.Task.Title,
		Details: req.Task.Details,
		IsDone:  req.Task.IsDone,
	}

	// Find and update user document
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$push", bson.D{{"tasks", newTask}}}}

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Database error while updating user with new task: %s", err.Error())
		return &pb.CreateTaskResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	if result.ModifiedCount == 0 {
		log.Printf("No document was modified when adding task")
		return &pb.CreateTaskResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
	}

	createdTask := &pb.TaskDTO{
		TaskId:  newTask.Id.Hex(),
		Title:   newTask.Title,
		Details: newTask.Details,
		IsDone:  newTask.IsDone,
	}

	return &pb.CreateTaskResponse{
		Status: pb.DataStatus_Data_Success,
		Task:   createdTask,
	}, nil
}

// Function for getting a task of a user
func (t *TaskService) ReadTask(ctx context.Context, req *pb.ReadTaskRequest) (*pb.ReadTaskResponse, error) {
	log.Println("Handling read task request")

	coll := t.db.Database(t.dbDetails.DatabaseName).Collection(t.dbDetails.UserCollectionName)

	objId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Printf("Failed to create object id from: %s", req.UserId)
		return &pb.ReadTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	taskId, err := primitive.ObjectIDFromHex(req.TaskId)
	if err != nil {
		log.Printf("Failed to convert task id: %s, to object id", req.TaskId)
		return &pb.ReadTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	filter := bson.D{
		{"_id", objId},
		{"tasks._id", taskId},
	}

	var user models.User
	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No matching user or task found")
			return &pb.ReadTaskResponse{Status: pb.DataStatus_Data_No_Task_Found}, nil
		}
		log.Printf("Database error: %s", err.Error())
		return &pb.ReadTaskResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	for _, task := range user.Tasks {
		if task.Id == taskId {
			foundTask := &pb.TaskDTO{
				TaskId:  task.Id.Hex(),
				Title:   task.Title,
				Details: task.Details,
				IsDone:  task.IsDone,
			}
			return &pb.ReadTaskResponse{Status: pb.DataStatus_Data_Success, Task: foundTask}, nil
		}
	}

	return &pb.ReadTaskResponse{Status: pb.DataStatus_Data_No_Task_Found}, nil
}

// Function for reading multiple tasks of a user
func (t *TaskService) ReadMultipleTasks(ctx context.Context, req *pb.ReadMultipleTasksRequest) (*pb.ReadMultipleTasksResponse, error) {
	log.Printf("Handling request to read tasks - Page: %d, Limit: %d", req.Page, req.Limit)

	coll := t.db.Database(t.dbDetails.DatabaseName).Collection(t.dbDetails.UserCollectionName)

	objId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Printf("Failed to convert: %s, to object id", req.UserId)
		return &pb.ReadMultipleTasksResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	filter := bson.D{{"_id", objId}}
	var user models.User
	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No user found in database")
			return &pb.ReadMultipleTasksResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
		}
		log.Printf("Database error on fetching user: %s", err.Error())
		return &pb.ReadMultipleTasksResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	// Apply pagination
	page := req.Page
	limit := req.Limit
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 25
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	tasks := make([]*pb.TaskDTO, 0)

	if int(startIndex) < len(user.Tasks) {
		// Adjust endIndex if it exceeds array bounds
		if int(endIndex) > len(user.Tasks) {
			endIndex = int32(len(user.Tasks))
		}

		for _, task := range user.Tasks[startIndex:endIndex] {
			tasks = append(tasks, &pb.TaskDTO{
				TaskId:  task.Id.Hex(),
				Title:   task.Title,
				Details: task.Details,
				IsDone:  task.IsDone,
			})
		}
	}

	return &pb.ReadMultipleTasksResponse{
		Status: pb.DataStatus_Data_Success,
		Tasks:  tasks,
	}, nil
}

// Function for updating a task
func (t *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	log.Println("Handling update task request")
	coll := t.db.Database(t.dbDetails.DatabaseName).Collection(t.dbDetails.UserCollectionName)

	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Printf("Failed to convert user ID: %s", err)
		return &pb.UpdateTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	taskId, err := primitive.ObjectIDFromHex(req.Task.TaskId)
	if err != nil {
		log.Printf("Failed to convert task ID: %s", err)
		return &pb.UpdateTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	// Update the task directly in the database
	filter := bson.D{
		{"_id", userId},
		{"tasks._id", taskId},
	}

	update := bson.D{{
		"$set", bson.D{
			{"tasks.$.title", req.Task.Title},
			{"tasks.$.details", req.Task.Details},
			{"tasks.$.is_done", req.Task.IsDone},
		},
	}}

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Database error: %s", err)
		return &pb.UpdateTaskResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	if result.MatchedCount == 0 {
		return &pb.UpdateTaskResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
	}

	if result.ModifiedCount == 0 {
		return &pb.UpdateTaskResponse{Status: pb.DataStatus_Data_No_Task_Found}, nil
	}

	return &pb.UpdateTaskResponse{
		Status: pb.DataStatus_Data_Success,
		Task:   req.Task,
	}, nil
}

// Function for deleting a task
func (t *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	log.Printf("Handling delete task request for user: %s, task: %s", req.UserId, req.TaskId)

	coll := t.db.Database(t.dbDetails.DatabaseName).Collection(t.dbDetails.UserCollectionName)

	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Printf("Failed to convert user ID to ObjectID: %v", err)
		return &pb.DeleteTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	taskId, err := primitive.ObjectIDFromHex(req.TaskId)
	if err != nil {
		log.Printf("Failed to convert task ID to ObjectID: %v", err)
		return &pb.DeleteTaskResponse{Status: pb.DataStatus_Data_Internal_Error}, nil
	}

	filter := bson.D{{"_id", userId}}
	update := bson.D{{"$pull", bson.D{{"tasks", bson.D{{"_id", taskId}}}}}}

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Database error while deleting task: %v", err)
		return &pb.DeleteTaskResponse{Status: pb.DataStatus_Data_Database_Error}, nil
	}

	if result.MatchedCount == 0 {
		return &pb.DeleteTaskResponse{Status: pb.DataStatus_Data_No_User_Found}, nil
	}

	if result.ModifiedCount == 0 {
		return &pb.DeleteTaskResponse{Status: pb.DataStatus_Data_No_Task_Found}, nil
	}

	return &pb.DeleteTaskResponse{Status: pb.DataStatus_Data_Success}, nil
}
