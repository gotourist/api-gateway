package services

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/iman_task/api-gateway/config"
	collectpb "github.com/iman_task/api-gateway/genproto/collect"
	postpb "github.com/iman_task/api-gateway/genproto/post"
)

type ServiceManager interface {
	PostService() postpb.PostServiceClient
	CollectService() collectpb.GoServiceClient
}

type serviceManager struct {
	postService    postpb.PostServiceClient
	collectService collectpb.GoServiceClient
}

func (s *serviceManager) PostService() postpb.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CollectService() collectpb.GoServiceClient {
	return s.collectService
}

func NewServiceManager(conf *config.Config) (ServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connCollect, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CollectServiceHost, conf.CollectServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		postService:    postpb.NewPostServiceClient(connPost),
		collectService: collectpb.NewGoServiceClient(connCollect),
	}

	return serviceManager, nil
}
