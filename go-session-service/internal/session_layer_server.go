package internal

import (
	"context"
	"log"
	"time"

	uuid "github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	pb "gn222gq.2dv013.a2/protos"
)

// Represents the grpc session layer server
type SessionLayerServer struct {
	pb.UnimplementedSessionLayerServer
	db  *redis.Client
	ttl time.Duration
}

// Constructor function
func NewSessionLayerServer(db *redis.Client, ttl time.Duration) *SessionLayerServer {
	return &SessionLayerServer{
		db:  db,
		ttl: ttl,
	}
}

// Function for getting a session based on a user id
func (s *SessionLayerServer) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.GetSessionResponse, error) {
	log.Println("Handling get session request")
	val, err := s.db.Get(ctx, req.UserId).Result()
	if err != nil {
		if err == redis.Nil {
			log.Print("No session found")
			return &pb.GetSessionResponse{Status: pb.SessionStatus_Session_No_Session_Found}, nil
		}
		log.Printf("Database error while trying to get session: %s", err.Error())
		return &pb.GetSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
	}
	return &pb.GetSessionResponse{Status: pb.SessionStatus_Session_Success, SessionToken: &val}, nil
}

// Function for creating a new session
func (s *SessionLayerServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	var token string
	for {
		token = uuid.New().String()
		_, err := s.db.Get(ctx, token).Result()
		if err == redis.Nil {
			break
		} else if err != nil {
			log.Printf("Database error while checking for duplicate token uuid: %s", err.Error())
			return &pb.CreateSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
		}
	}

	// Set token -> user id mapping
	err := s.db.Set(ctx, token, req.UserId, s.ttl).Err()
	if err != nil {
		log.Printf("Error while setting token:userid, %s", err.Error())
		return &pb.CreateSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
	}

	// Set user id -> token mapping
	err = s.db.Set(ctx, req.UserId, token, s.ttl).Err()
	if err != nil {
		log.Printf("Error while setting userid:token, %s", err.Error())
		return &pb.CreateSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
	}

	return &pb.CreateSessionResponse{Status: pb.SessionStatus_Session_Success, SessionToken: &token}, nil
}

// Function for deleting a session
func (s *SessionLayerServer) DeleteSession(ctx context.Context, req *pb.DeleteSessionRequest) (*pb.DeleteSessionResponse, error) {
	userId, err := s.db.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("Did not find provided token")
			return &pb.DeleteSessionResponse{Status: pb.SessionStatus_Session_No_Session_Found}, nil
		}
		log.Printf("Database error on trying to get token: %s", err.Error())
		return &pb.DeleteSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
	}

	_, err = s.db.Del(ctx, req.SessionToken).Result()
	if err != nil {
		log.Printf("Database error on deletion of sessiontoken:userid occured: %s", err.Error())
		return &pb.DeleteSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
	}

	_, err = s.db.Del(ctx, userId).Result()
	if err != nil {
		log.Printf("Database error on deletion of userid:sessiontoken: %s", err.Error())
		return &pb.DeleteSessionResponse{Status: pb.SessionStatus_Session_Database_Error}, nil
	}

	return &pb.DeleteSessionResponse{Status: pb.SessionStatus_Session_Success}, nil
}
