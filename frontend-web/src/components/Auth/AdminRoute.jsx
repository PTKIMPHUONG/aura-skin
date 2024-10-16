import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../../context/Authcontext";

const AdminRoute = () => {
  const { user, isAuthenticated } = useAuth();
  console.log("AdminRoute - User:", user);
  console.log("AdminRoute - Is authenticated:", isAuthenticated());
  console.log("AdminRoute - Is admin:", user?.isAdmin);

  if (!isAuthenticated()) {
    return <Navigate to="/login" />;
  }

  if (!user.isAdmin) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
};

export default AdminRoute;
