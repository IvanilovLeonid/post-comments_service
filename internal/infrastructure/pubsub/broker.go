package pubsub

import (
	"context"
	"social-comments/internal/core/domain"
	"sync"
)

type CommentEvent struct {
	PostID  string
	Comment *domain.Comment
}

type Broker struct {
	mu          sync.RWMutex
	subscribers map[string][]chan *CommentEvent
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]chan *CommentEvent),
	}
}

func (b *Broker) Subscribe(ctx context.Context, postID string) <-chan *CommentEvent {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan *CommentEvent, 1)
	b.subscribers[postID] = append(b.subscribers[postID], ch)

	go func() {
		<-ctx.Done()
		b.unsubscribe(postID, ch)
	}()

	return ch
}

func (b *Broker) Publish(postID string, comment *domain.Comment) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	event := &CommentEvent{PostID: postID, Comment: comment}
	for _, ch := range b.subscribers[postID] {
		ch <- event
	}
}

func (b *Broker) unsubscribe(postID string, ch chan *CommentEvent) {
	b.mu.Lock()
	defer b.mu.Unlock()

	subscribers := b.subscribers[postID]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			b.subscribers[postID] = append(subscribers[:i], subscribers[i+1:]...)
			close(ch)
			return
		}
	}
}
