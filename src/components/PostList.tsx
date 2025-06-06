import { useQuery } from '@apollo/client';
import { GET_ALL_POSTS } from '../graphql/queries';

export default function PostList() {
    const { data, loading, error } = useQuery(GET_ALL_POSTS);

    if (loading) return <p>Загрузка...</p>;
    if (error) return <p>Ошибка: {error.message}</p>;

    console.log('Received data:', data); // Проверьте структуру данных

    // Добавьте проверку наличия данных
    if (!data?.posts?.edges) return <p>Нет данных</p>;

    return (
        <div>
            <h2>Список постов</h2>
            <ul>
                {data.posts.edges.map(({ node }: any) => (
                    <li key={node.id}>{node.title}</li>
                ))}
            </ul>
        </div>
    );
}
