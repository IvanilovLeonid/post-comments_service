import { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CREATE_POST } from '../graphql/mutations';

export default function CreatePostForm() {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [author, setAuthor] = useState('');
  const [createPost] = useMutation(CREATE_POST);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await createPost({
      variables: {
        input: { title, content, author, allowComments: true },
      },
    });
    setTitle('');
    setContent('');
    setAuthor('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input value={title} onChange={e => setTitle(e.target.value)} placeholder="Заголовок" required />
      <textarea value={content} onChange={e => setContent(e.target.value)} placeholder="Содержание" required />
      <input value={author} onChange={e => setAuthor(e.target.value)} placeholder="Автор" required />
      <button type="submit">Создать пост</button>
    </form>
  );
}