package comment

import (
	"context"
	"errors"
	"fmt"
	"social-comments/internal/core/domain"
	"social-comments/internal/core/errors"
	"social-comments/internal/core/repository"
	"social-comments/pkg/logging"
	"social-comments/pkg/utils/pagination"
)

type Service struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
	logger      logger.Logger
}

func NewService(
	commentRepo repository.CommentRepository,
	postRepo repository.PostRepository,
	logger logger.Logger,
) *Service {
	return &Service{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		logger:      logger,
	}
}

func (s *Service) CreateComment(ctx context.Context, req domain.CreateCommentRequest) (*domain.Comment, error) {
	if err := validateCommentRequest(req); err != nil {
		s.logger.Error("validation error: %v", err)
		return &domain.Comment{}, err
	}
	// TODO
	// Проверяем, разрешены ли комментарии к посту
	//post, err := s.commentRepo.GetByPostID(ctx, req.PostID)
	//if err != nil {
	//	s.logger.Error("failed to get post: %v", err)
	//	return domain.Comment{}, apperrors.ErrPostNotFound
	//}
	//if !post.AllowComments {
	//	return domain.Comment{}, apperrors.ErrCommentsDisabled
	//}

	// Создаем комментарий через репозиторий
	comment := domain.Comment{
		Author:   req.Author,
		Text:     req.Text,
		PostID:   req.PostID,
		ParentID: req.ParentID,
	}

	createdComment, err := s.commentRepo.Create(ctx, comment)
	if err != nil {
		s.logger.Error("failed to create comment: %v", err)
		return &domain.Comment{}, apperrors.ErrCreatingComment
	}

	return &createdComment, nil
}

func (s *Service) GetCommentsByPostID(ctx context.Context, postID, page, pageSize int) ([]domain.Comment, error) {
	if postID <= 0 {
		return nil, errors.New("invalid post ID")
	}

	paginator := pagination.NewPaginator(page, pageSize)
	//if !paginator.IsValid() {
	//	return nil, errors.New("invalid pagination parameters")
	//}

	comments, err := s.commentRepo.GetByPostID(ctx, postID, paginator.Limit(), paginator.Offset())
	if err != nil {
		s.logger.Error("failed to get comments: %v", err)
		return nil, apperrors.GetCommentsByPost
	}
	//var result []*domain.Comment
	//for i := range comments {
	//	result = append(result, &comments[i])
	//}
	//return result, nil
	return comments, nil
}

func (s *Service) GetCommentReplies(ctx context.Context, commentID int) ([]*domain.Comment, error) {
	// Проверка на валидность идентификатора
	if commentID <= 0 {
		return nil, errors.New("invalid comment ID: must be greater than zero")
	}

	// Получение ответов на комментарий из репозитория
	replies, err := s.commentRepo.GetReplies(ctx, commentID)
	if err != nil {
		// Логируем ошибку с контекстом
		s.logger.Error("failed to get replies for comment ID %d: %v", commentID, err)
		// Возвращаем специфическую ошибку, чтобы клиенты могли обрабатывать её
		return nil, apperrors.ErrRepliesNotFound
	}

	replyPointers := make([]*domain.Comment, len(replies))
	for i := range replies {
		replyPointers[i] = &replies[i]
	}

	// Return converted slice
	return replyPointers, nil
}

func validateCommentRequest(req domain.CreateCommentRequest) error {
	if req.Author == "" {
		return errors.New("comment author cannot be empty")
	}
	if req.Text == "" {
		return errors.New("comment text cannot be empty")
	}
	if len(req.Text) > apperrors.MaxContentLength {
		return fmt.Errorf("comment exceeds maximum length of %d", apperrors.MaxContentLength)
	}
	if req.PostID <= 0 {
		return errors.New("invalid post ID")
	}
	return nil
}
