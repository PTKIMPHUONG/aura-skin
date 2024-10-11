import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../../context/Authcontext";

const AdminRoute = () => {
  const { user } = useAuth();

  if (!user) {
    return <Navigate to="/login" />;
  }

  if (!user.is_admin) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
};

export default AdminRoute;
