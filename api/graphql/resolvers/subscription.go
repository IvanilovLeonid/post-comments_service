package resolvers

import (
	"context"
	"fmt"
	"social-comments/api/graphql/generated"
	"social-comments/internal/core/domain"
	"social-comments/internal/infrastructure/pubsub"
)

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *domain.Comment, error) {
	// Получаем брокера из контекста или резолвера
	broker := r.Broker
	if broker == nil {
		return nil, fmt.Errorf("broker not initialized")
	}

	// Создаем канал для GraphQL
	commentsChan := make(chan *domain.Comment)

	// Подписываемся на события
	eventChan := broker.Subscribe(ctx, postID)

	go func() {
		defer close(commentsChan)
		for {
			select {
			case event := <-eventChan:
				if event != nil {
					commentsChan <- event.Comment
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return commentsChan, nil
}
