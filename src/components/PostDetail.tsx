import { useParams } from 'react-router-dom';
import { useQuery, useSubscription } from '@apollo/client';
import { GET_POST_WITH_COMMENTS } from '../graphql/queries';
import { COMMENTS_SUBSCRIPTION } from '../graphql/subscriptions';
import CommentForm from './CommentForm';

export default function PostDetail() {
  const { id } = useParams<{ id: string }>();
  const { data, loading, error, refetch } = useQuery(GET_POST_WITH_COMMENTS, {
    variables: { id },
  });

  useSubscription(COMMENTS_SUBSCRIPTION, {
    variables: { postId: id },
    onData: () => refetch(),
  });

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p>Ошибка: {error.message}</p>;

  const post = data.post;

  return (
    <div>
      <h2>{post.title}</h2>
      <p>{post.content}</p>
      <h3>Комментарии</h3>
      {post.comments.map((comment: any) => (
        <div key={comment.id}>
          <p><b>{comment.author}</b>: {comment.text}</p>
          {comment.replies.map((reply: any) => (
            <p key={reply.id} style={{ marginLeft: 20 }}>
              ↳ <b>{reply.author}</b>: {reply.text}
            </p>
          ))}
        </div>
      ))}
      <CommentForm postId={id} />
    </div>
  );
}