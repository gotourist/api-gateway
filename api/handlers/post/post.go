package post

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"

	"github.com/iman_task/api-gateway/api/models"
	collectpb "github.com/iman_task/api-gateway/genproto/collect"
	postpb "github.com/iman_task/api-gateway/genproto/post"
	"github.com/iman_task/api-gateway/pkg/logger"
	"github.com/iman_task/api-gateway/pkg/utils"
)

func (p *PostHandler) ListPosts(c *gin.Context) {

	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: errStr,
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(p.config.CtxTimeout))
	defer cancel()

	response, err := p.serviceManager.PostService().ListPosts(
		ctx, &postpb.ListPostsRequest{
			Limit:  params.Limit,
			Page:   params.Page,
			Search: params.Search,
			Filter: params.Filters,
		},
	)

	if err == nil && response != nil && response.Code == 0 {
		c.JSON(http.StatusOK, response)
		return
	}

	if response != nil && response.Code == 3 {
		c.JSON(http.StatusBadRequest, convertPostErrorMessages(response.Errors))
		return
	}

	utils.ResponseError(p.log, err, c)
}

func (p *PostHandler) DetailPost(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: []string{"Invalid parameter"},
		})
		p.log.Error("Invalid post id parameter", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(p.config.CtxTimeout))
	defer cancel()

	response, err := p.serviceManager.PostService().DetailPost(
		ctx, &postpb.DetailPostRequest{
			Id: id,
		},
	)

	if err == nil && response != nil && response.Code == 0 {
		c.JSON(http.StatusOK, response)
		return
	}

	if response != nil && response.Code == 3 {
		c.JSON(http.StatusBadRequest, convertPostErrorMessages(response.Errors))
		return
	}

	utils.ResponseError(p.log, err, c)
}

func (p *PostHandler) UpdatePost(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: []string{"Invalid parameter"},
		})
		p.log.Error("Invalid post id parameter", logger.Error(err))
		return
	}

	var (
		updatePostRequest models.UpdatePostRequest
	)

	err = c.ShouldBindJSON(&updatePostRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: []string{"Invalid request data"},
		})
		p.log.Error("Error binding json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(p.config.CtxTimeout))
	defer cancel()

	response, err := p.serviceManager.PostService().UpdatePost(
		ctx, &postpb.UpdatePostRequest{
			Id:    id,
			Title: updatePostRequest.Title,
			Body:  updatePostRequest.Body,
		},
	)

	if err == nil && response != nil && response.Code == 0 {
		c.JSON(http.StatusOK, response)
		return
	}

	if response != nil && response.Code == 3 {
		c.JSON(http.StatusBadRequest, convertPostErrorMessages(response.Errors))
		return
	}

	utils.ResponseError(p.log, err, c)
}

func (p *PostHandler) DeletePost(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationError{
			NonFieldErrors: []string{"Invalid parameter"},
		})
		p.log.Error("Invalid post id parameter", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(p.config.CtxTimeout))
	defer cancel()

	response, err := p.serviceManager.PostService().DeletePost(
		ctx, &postpb.DeletePostRequest{
			Id: id,
		},
	)

	if err == nil && response != nil && response.Code == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	if response != nil && response.Code == 3 {
		c.JSON(http.StatusBadRequest, convertPostErrorMessages(response.Errors))
		return
	}

	utils.ResponseError(p.log, err, c)
}

func (p *PostHandler) CollectPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(p.config.CtxTimeout))
	defer cancel()

	response, err := p.serviceManager.CollectService().CollectPosts(
		ctx, &collectpb.CollectPostsRequest{},
	)

	if err == nil && response != nil && response.Code == 0 {
		c.JSON(http.StatusOK, response)
		return
	}

	if response != nil && response.Code == 3 {
		c.JSON(http.StatusBadRequest, convertCollectErrorMessages(response.Errors))
		return
	}

	utils.ResponseError(p.log, err, c)
}

func (p *PostHandler) CheckCollectStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(p.config.CtxTimeout))
	defer cancel()

	response, err := p.serviceManager.CollectService().CheckStatus(
		ctx, &collectpb.CheckStatusRequest{},
	)

	if err == nil && response != nil && response.Code == 0 {
		c.JSON(http.StatusOK, response)
		return
	}

	if response != nil && response.Code == 3 {
		c.JSON(http.StatusBadRequest, convertCollectErrorMessages(response.Errors))
		return
	}

	utils.ResponseError(p.log, err, c)
}
