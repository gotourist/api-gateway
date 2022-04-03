package services

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/iman_task/api-gateway/config"
	postpb "github.com/iman_task/api-gateway/genproto/post"
)

type ServiceManager interface {
	PostService() postpb.PostServiceClient
}

type serviceManager struct {
	postService postpb.PostServiceClient
}

func (s *serviceManager) PostService() postpb.PostServiceClient {
	return s.postService
}

func NewServiceManager(conf *config.Config) (ServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		postService: postpb.NewPostServiceClient(connPost),
	}

	return serviceManager, nil
}
