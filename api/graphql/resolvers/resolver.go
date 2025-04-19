package resolvers

import (
	"errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	apperrors "social-comments/internal/core/errors"
	"social-comments/internal/core/ports"
	"social-comments/pkg/logging"
)

type Resolver struct {
	PostService    ports.PostService
	CommentService ports.CommentService
	Logger         logger.Logger
}

func (r *Resolver) wrapGQLError(err error) *gqlerror.Error {
	var appErr *apperrors.APIError
	if errors.As(err, &appErr) {
		return &gqlerror.Error{
			Message:    appErr.Message,
			Extensions: appErr.Extensions(),
		}
	}
	return &gqlerror.Error{
		Message: "internal server error",
	}
}
