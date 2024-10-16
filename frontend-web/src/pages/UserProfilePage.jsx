import React, { useState, useEffect } from "react";
import { Box, Container, Typography, Snackbar } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import UserProfileForm from "../components/Users/UserProfileForm";
import { useAuth } from "../context/Authcontext";
import UserService from "../services/UserService";

const UserProfilePage = () => {
  const { user, updateUser } = useAuth();
  const [profileData, setProfileData] = useState(null);
  const [snackbar, setSnackbar] = useState({ open: false, message: "" });

  useEffect(() => {
    if (user) {
      setProfileData(user);
    }
  }, [user]);

  const handleUpdateProfile = async (updatedData) => {
    try {
      const response = await UserService.updateUserProfile(
        user.id,
        updatedData
      );
      if (response.success) {
        setProfileData(response.data);
        updateUser(response.data);
        setSnackbar({ open: true, message: "Cập nhật thông tin thành công" });
      }
    } catch (error) {
      console.error("Error updating user profile:", error);
      setSnackbar({
        open: true,
        message: "Có lỗi xảy ra khi cập nhật thông tin",
      });
    }
  };

  if (!profileData) {
    return <Typography>Đang tải...</Typography>;
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Box sx={{ display: "flex" }}>
        <SidebarUser />
        <Box sx={{ flexGrow: 1, ml: 4 }}>
          <Typography variant="h5" gutterBottom>
            Thông tin tài khoản
          </Typography>
          <UserProfileForm
            profileData={profileData}
            onUpdateProfile={handleUpdateProfile}
          />
        </Box>
      </Box>
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={() => setSnackbar({ ...snackbar, open: false })}
        message={snackbar.message}
      />
    </Container>
  );
};

export default UserProfilePage;
