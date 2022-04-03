package docs

import (
	"github.com/iman_task/api-gateway/api/models"
	postpb "github.com/iman_task/api-gateway/genproto/post"
)

//swagger:route GET /v1/posts/list/ post postsList
// ---
// Get posts list
//
// Get posts list
//
// responses:
// 200: postsListResponse
// 400: badRequest
// 500: internalService

// swagger:response postsListResponse
type swagPostsListResponse struct {
	// in:body
	Body postpb.ListPostsResponse
}

// swagger:route GET /v1/posts/detail/{id}/ post postDetailRequest
// ---
// Get post detail
//
// Get post detail
//
// responses:
//   200: postDetailResponse
//   400: badRequest
//   500: internalService

// swagger:parameters postDetailRequest updatePostRequest deletePostRequest
type swagPostID struct {
	// post id
	// in:path
	// required: true
	// type: integer
	ID int64 `json:"id"`
}

// swagger:response postDetailResponse
type swagPostDetailResponse struct {
	// in:body
	Body postpb.DetailPostResponse
}

// swagger:route PUT /v1/posts/update/{id}/ post updatePostRequest
// ---
// Update post
//
// Update post
//
// responses:
//   200: updatePostResponse
//   400: badRequest
//   500: internalService

// swagger:parameters updatePostRequest
type swagProductUpdateData struct {
	//in:body
	Body models.UpdatePostRequest
}

// swagger:response updatePostResponse
type swagUpdatePostResponse struct {
	// in:body
	Body postpb.UpdatePostResponse
}

// swagger:route DELETE /v1/posts/delete/{id}/ post deletePostRequest
// ---
// Delete post
//
// Delete post
//
// responses:
// 204:
// 400: badRequest
// 500: internalService
