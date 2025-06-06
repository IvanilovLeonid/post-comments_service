import { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CREATE_COMMENT } from '../graphql/mutations';

export default function CommentForm({ postId }: { postId: string }) {
  const [author, setAuthor] = useState('');
  const [text, setText] = useState('');
  const [createComment] = useMutation(CREATE_COMMENT);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await createComment({ variables: { input: { author, text, postId } } });
    setAuthor('');
    setText('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input value={author} onChange={e => setAuthor(e.target.value)} placeholder="Автор" required />
      <textarea value={text} onChange={e => setText(e.target.value)} placeholder="Комментарий" required />
      <button type="submit">Отправить</button>
    </form>
  );
}