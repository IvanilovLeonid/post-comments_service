import { useMutation } from '@apollo/client';
import { CREATE_POST, GET_ALL_POSTS } from '../graphql/queries';
import { useState } from 'react';

export default function CreatePost() {
    const [createPost] = useMutation(CREATE_POST, {
        refetchQueries: [{ query: GET_ALL_POSTS }],
    });

    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        await createPost({
            variables: {
                input: {
                    title,
                    content,
                    author: 'Фронтенд', // можно сделать поле
                    allowComments: true,
                },
            },
        });
        setTitle('');
        setContent('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <h2>Создать пост</h2>
            <input
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Заголовок"
            />
            <br />
            <textarea
                value={content}
                onChange={(e) => setContent(e.target.value)}
                placeholder="Содержание"
            />
            <br />
            <button type="submit">Создать</button>
        </form>
    );
}
