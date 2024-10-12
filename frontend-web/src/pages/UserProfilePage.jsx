import React, { useState, useEffect } from "react";
import { Box, Container, Typography } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import UserProfileForm from "../components/Users/UserProfileForm";
import { useAuth } from "../context/Authcontext";
import AuthService from "../services/AuthService";

const UserProfilePage = () => {
  const { user } = useAuth();
  const [profileData, setProfileData] = useState(null);

  useEffect(() => {
    if (user) {
      setProfileData(user);
    }
  }, [user]);

  const handleUpdateProfile = async (updatedData) => {
    try {
      // Giả sử AuthService có phương thức updateUserProfile
      const response = await AuthService.updateUserProfile(updatedData);
      if (response.success) {
        setProfileData(response.data);
        // Hiển thị thông báo cập nhật thành công
      }
    } catch (error) {
      console.error("Error updating user profile:", error);
      // Hiển thị thông báo lỗi
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
    </Container>
  );
};

export default UserProfilePage;
