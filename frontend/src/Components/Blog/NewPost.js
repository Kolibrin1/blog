import React, { useState } from 'react';
import '../../styles/Blog/NewPost.css';
import { createPost } from '../../api/api'; // Импортируем функцию API

const NewPost = ({ onPostCreated }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!title || !content) {
      setErrorMessage('Both fields are required!');
      return;
    }

    try {
      const response = await createPost(title, content); // Отправляем данные на сервер
      setErrorMessage('');
      setSuccessMessage('Post created successfully!');
      onPostCreated(response.data); // Передаём созданный пост родительскому компоненту
      setTitle('');
      setContent('');
    } catch (error) {
      console.error('Failed to create post:', error);
      setErrorMessage('Failed to create post. Please try again later.');
    }
  };

  return (
    <div className="new-post-container">
      <div className="new-post-card">
        <h2>Create New Post</h2>
        <form className="new-post-form" onSubmit={handleSubmit}>
          <input
            type="text"
            className="new-post-input"
            placeholder="Post Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <textarea
            className="new-post-textarea"
            placeholder="Post Content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
          <button type="submit" className="new-post-button">
            Submit
          </button>
          {errorMessage && <div className="error-message">{errorMessage}</div>}
          {successMessage && (
            <div className="success-message">{successMessage}</div>
          )}
        </form>
      </div>
    </div>
  );
};

export default NewPost;
