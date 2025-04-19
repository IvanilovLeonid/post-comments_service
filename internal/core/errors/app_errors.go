package apperrors

import "errors"

var (
	ErrPostNotFound      = errors.New("post not found")
	ErrPostNotFoundByID  = errors.New("post not found by ID")
	ErrGetAllPosts       = errors.New("error getting all posts")
	ErrCommentNotFound   = errors.New("comment not found")
	ErrContentTooLong    = errors.New("content exceeds maximum length")
	ErrInvalidPagination = errors.New("invalid pagination parameters")
	ErrCommentsDisabled  = errors.New("comments are disabled for this post")
	GetCommentsByPost    = errors.New("error getting comments by post")
	ErrRepliesNotFound   = errors.New("replies not found")
	ErrCreatingComment   = errors.New("error creating comment")
	ErrCreatingPost      = errors.New("error creating post")
)

const (
	MaxContentLength        = 2000
	InitialCommentsCapacity = 100
	InitialPostsCapacity    = 10
)
