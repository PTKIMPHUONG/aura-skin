import React, { createContext, useState, useContext, useEffect } from "react";
import authService from "../services/AuthService";

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(null);
  const [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const loadUserFromStorage = () => {
      const storedAuthData = authService.getCurrentUser();
      if (storedAuthData) {
        setUser(storedAuthData.user);
        setToken(storedAuthData.token);
        setIsAdmin(storedAuthData.user.is_admin || false);
      }
    };

    loadUserFromStorage();
  }, []);

  const login = (userData, userToken) => {
    setUser(userData);
    setToken(userToken);
    setIsAdmin(userData.is_admin || false);
    authService.saveUserToStorage({ user: userData, token: userToken });
  };

  const logout = () => {
    setUser(null);
    setToken(null);
    setIsAdmin(false);
    authService.removeUserFromStorage();
  };

  const isAuthenticated = () => {
    return !!user && !!token;
  };

  return (
    <AuthContext.Provider
      value={{ user, token, isAdmin, login, logout, isAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
