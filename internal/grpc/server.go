package grpc

import (
	"context"
	"github.com/MrTomSawyer/go-kv-storage/internal/models"     // Importing the models package for Entry type
	kv1 "github.com/MrTomSawyer/go-kv-storage/proto/gen/go/kv" // Importing the generated gRPC protobuf package
	"google.golang.org/grpc"                                   // Importing the gRPC package
	"time"                                                     // Importing the time package
)

// KVHandler defines the interface for key-value storage operations.
type KVHandler interface {
	Get(ctx context.Context, key string) (models.Entry, bool)
	Set(ctx context.Context, key string, value string, TTL time.Duration)
	Delete(ctx context.Context, key string)
}

// serverAPI implements the KeyValueStorage gRPC server interface.
type serverAPI struct {
	kv1.UnimplementedKeyValueStorageServer           // Embedding to have forward-compatible implementations
	KVHandler                              KVHandler // The handler for key-value storage operations
}

// Register registers the KeyValueStorage server with the given gRPC server.
func Register(gRPC *grpc.Server, KVHandler KVHandler) {
	kv1.RegisterKeyValueStorageServer(gRPC, &serverAPI{KVHandler: KVHandler})
}

// Get handles the gRPC request to retrieve a value by key.
func (s *serverAPI) Get(ctx context.Context, req *kv1.GetRequest) (*kv1.GetResponse, error) {
	v, found := s.KVHandler.Get(ctx, req.GetKey())
	return &kv1.GetResponse{
		Value: v.Value,
		Found: found,
	}, nil
}

// Set handles the gRPC request to add or update a key-value pair with a specified TTL.
func (s *serverAPI) Set(ctx context.Context, req *kv1.SetRequest) (*kv1.SetResponse, error) {
	s.KVHandler.Set(ctx, req.GetKey(), req.GetValue(), time.Duration(req.GetTtl()))
	return &kv1.SetResponse{
		Success: true,
	}, nil
}

// Delete handles the gRPC request to remove a value by key.
func (s *serverAPI) Delete(ctx context.Context, req *kv1.DeleteRequest) (*kv1.DeleteResponse, error) {
	s.KVHandler.Delete(ctx, req.GetKey())
	return &kv1.DeleteResponse{
		Success: true,
	}, nil
}
