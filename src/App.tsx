import PostList from './components/PostList';
import CreatePost from './components/CreatePost';

export default function App() {
    return (
        <div style={{ padding: '1rem' }}>
            <h1>📝 Социальные посты</h1>
            <CreatePost />
            <hr />
            <PostList />
        </div>
    );
}
