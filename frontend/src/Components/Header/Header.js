import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import Notifications from './Notifications';
import Menu from './Menu';  // Import the new Menu component

import '../../styles/Header/Header.css'; 

const Header = () => {
  const { isAuthenticated, user } = useAuth()

  return (
    <header className="header">
      <nav className="header-nav">
        <div className="header-left">
          <Link to="/" className="header-link">
              Blogs
          </Link>
        </div>
        <div className="header-right">
          {isAuthenticated && user && (
            <>
              <Notifications userId={user.id} />
              <Menu user={user} />
            </>
          )}
        </div>
      </nav>
    </header>
  );
};

export default Header;
