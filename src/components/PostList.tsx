import { useQuery } from '@apollo/client';
import { GET_ALL_POSTS } from '../graphql/queries';
import { Link } from 'react-router-dom';

export default function PostList() {
  const { data, loading, error } = useQuery(GET_ALL_POSTS);

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p>Ошибка: {error.message}</p>;

  return (
    <ul>
      {data.posts.edges.map(({ node }: any) => (
        <li key={node.id}>
          <Link to={`/post/${node.id}`}>{node.title}</Link>
        </li>
      ))}
    </ul>
  );
}