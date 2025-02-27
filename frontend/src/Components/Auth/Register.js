import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { register, login as apiLogin } from '../../api/api'; // Переименовываем login из API
import { useAuth } from '../../context/AuthContext';

import '../../styles/Auth/Register.css';

const Register = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const { setAuthToken, setAuthUser } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      // Регистрируем пользователя
      await register(username, email, password);

      // Выполняем логин после успешной регистрации
      const loginResponse = await apiLogin(email, password);

      // Сохраняем токен и данные пользователя
      const { token, user } = loginResponse.data;
      localStorage.setItem('token', token);
      setAuthToken(token);
      setAuthUser(user);

      // Перенаправляем на главную страницу
      navigate('/');
    } catch (error) {
      setError('Не удалось зарегистрироваться или войти. Пожалуйста, попробуйте снова.');
      console.error('Ошибка при регистрации или входе:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="register-container">
      <form className="register-form" onSubmit={handleSubmit}>
        <h2>Регистрация</h2>
        {error && <p className="error-message">{error}</p>}
        <input
          type="text"
          placeholder="Имя пользователя"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          autoComplete="username"
          className="form-input"
        />
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          autoComplete="email"
          className="form-input"
        />
        <input
          type="password"
          placeholder="Пароль"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          autoComplete="new-password"
          className="form-input"
        />
        <button type="submit" className="register-button" disabled={loading}>
          {loading ? 'Регистрация...' : 'Зарегистрироваться'}
        </button>
        <p className="login-link">
          Уже есть аккаунт? <Link to="/login">Войти здесь</Link>
        </p>
      </form>
    </div>
  );
};

export default Register;
