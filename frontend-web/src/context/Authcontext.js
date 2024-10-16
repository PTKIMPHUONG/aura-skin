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
        setIsAdmin(storedAuthData.user.isAdmin || false);
      }
    };

    loadUserFromStorage();
  }, []);

  const login = async (user, token) => {
    console.log("Logging in user:", user);
    setUser(user);
    setToken(token);
    setIsAdmin(user.isAdmin); // Đảm bảo đang sử dụng đúng tên trường
    // Lưu thông tin vào localStorage nếu cần
  };

  const logout = () => {
    setUser(null);
    setToken(null);
    setIsAdmin(false);
    authService.logout();
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
