import React, { useEffect } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../../context/Authcontext";

const ProtectedRoute = () => {
  const { isAuthenticated, user, token } = useAuth();

  useEffect(() => {
    console.log("ProtectedRoute - User:", user, "Token:", token);
  }, [user, token]);

  if (!isAuthenticated()) {
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
};

export default ProtectedRoute;
