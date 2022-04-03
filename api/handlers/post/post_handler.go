package post

import (
	"github.com/iman_task/api-gateway/api/models"
	configPkg "github.com/iman_task/api-gateway/config"
	postpb "github.com/iman_task/api-gateway/genproto/post"
	"github.com/iman_task/api-gateway/pkg/logger"
	"github.com/iman_task/api-gateway/services"
)

// PostHandlerConfig ...
type PostHandlerConfig struct {
	Conf           configPkg.Config
	Logger         logger.Logger
	ServiceManager services.ServiceManager
}

// PostHandler ...
type PostHandler struct {
	config         configPkg.Config
	log            logger.Logger
	serviceManager services.ServiceManager
}

// New ...
func New(config *PostHandlerConfig) *PostHandler {
	return &PostHandler{
		config:         config.Conf,
		log:            config.Logger,
		serviceManager: config.ServiceManager,
	}
}

func convertErrorMessages(errors []*postpb.Error) []*models.Error {
	var errorMessages []*models.Error
	for _, e := range errors {
		errorMessages = append(errorMessages, &models.Error{
			Field:    e.Field,
			Messages: e.Messages,
			Errors:   convertErrorMessages(e.Errors),
		})
	}
	return errorMessages
}
