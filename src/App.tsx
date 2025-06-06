import { Routes, Route } from 'react-router-dom';
import PostList from './components/PostList';
import PostDetail from './components/PostDetail';
import CreatePostForm from './components/CreatePostForm';

function App() {
  return (
    <Routes>
      <Route path="/" element={<><CreatePostForm /><PostList /></>} />
      <Route path="/post/:id" element={<PostDetail />} />
    </Routes>
  );
}

export default App;