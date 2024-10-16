import React, { useState, useEffect } from "react";
import { Box, Typography, CircularProgress, Toolbar } from "@mui/material";
import AdminHeader from "../../components/Admin/AdminHeader";
import AdminSidebar from "../../components/Sidebar/SidebarAdmin";
import UserList from "../../components/Admin/AdminUsers/UserList";
import UserService from "../../services/UserService"; // Giả sử bạn có service này

const AdminUserManagementPage = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    try {
      setLoading(true);
      const response = await UserService.getAllUsers();
      setUsers(response.data);
      setLoading(false);
    } catch (err) {
      console.error("Error fetching users:", err);
      setError("Không thể tải danh sách người dùng. Vui lòng thử lại sau.");
      setLoading(false);
    }
  };

  const handleEditUser = (user) => {
    console.log("Editing user:", user);
    // Implement edit user logic
  };

  const handleDeleteUser = async (userId) => {
    try {
      await UserService.deleteUser(userId);
      setUsers(users.filter((u) => u.id !== userId));
    } catch (error) {
      console.error("Error deleting user:", error);
    }
  };

  return (
    <Box sx={{ display: "flex" }}>
      <AdminHeader />
      <AdminSidebar />
      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <Toolbar />
        {loading ? (
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            height="100vh"
          >
            <CircularProgress />
          </Box>
        ) : error ? (
          <Typography color="error">{error}</Typography>
        ) : (
          <>
            <Typography variant="h5" sx={{ mb: 2 }}>
              Quản lý tài khoản người dùng
            </Typography>
            <UserList
              users={users}
              onEditUser={handleEditUser}
              onDeleteUser={handleDeleteUser}
            />
          </>
        )}
      </Box>
    </Box>
  );
};

export default AdminUserManagementPage;
