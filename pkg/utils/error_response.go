package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/iman_task/api-gateway/api/models"
	"github.com/iman_task/api-gateway/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func ResponseError(log logger.Logger, err error, c *gin.Context) {
	st, _ := status.FromError(err)

	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, models.ObjectNotFoundError{
			NonFieldErrors: []string{st.Message()},
		})
		return
	}

	if st.Code() == codes.InvalidArgument {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: []string{st.Message()},
		})
		return
	}

	if st.Code() == codes.PermissionDenied {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: []string{st.Message()},
		})
		return
	}

	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.InternalServerError{
			NonFieldErrors: []string{"Oops, something went wrong."},
		})
		log.Error("Service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusInternalServerError, models.InternalServerError{
		NonFieldErrors: []string{"Oops, something went wrong."},
	})
	log.Error("Internal Server Error received", logger.Error(err))
}
