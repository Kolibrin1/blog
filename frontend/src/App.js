import React, { useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import PrivateRoute from './PrivateRoute/PrivateRoute';

import Login from './components/Auth/Login';
import Logout from './components/Auth/Logout';
import Register from './components/Auth/Register';

import Header from './components/Header/Header';
import MainPage from './components/MainPage/MainPage';
import Profile from './components/Profile/Profile';
import PostPage from './components/PostPage/PostPage';

function App() {
  const location = useLocation();

  // Убираем или добавляем класс на body в зависимости от текущей страницы
  useEffect(() => {
    if (location.pathname === '/login' || location.pathname === '/register') {
      document.body.classList.remove('with-header');
    } else {
      document.body.classList.add('with-header');
    }
  }, [location]);

  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />

          {/* Показываем Header только если не на страницах логина и регистрации */}
          {location.pathname !== '/login' && location.pathname !== '/register' && <Header />}

          <Route
            path="/"
            element={
              <PrivateRoute>
                <MainPage />
              </PrivateRoute>
            }
          />

          <Route
            path="/logout"
            element={
              <PrivateRoute>
                <Logout />
              </PrivateRoute>
            }
          />

          <Route
            path="/profile"
            element={
              <PrivateRoute>
                <Profile />
              </PrivateRoute>
            }
          />

          <Route
            path="/profile/:username"
            element={
              <PrivateRoute>
                <Profile />
              </PrivateRoute>
            }
          />

          <Route
            path="/post/:postID"
            element={
              <PrivateRoute>
                <PostPage />
              </PrivateRoute>
            }
          />
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
