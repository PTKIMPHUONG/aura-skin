import React, { useState, useEffect } from "react";
import { Box, Container, Typography } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import UserProfileForm from "../components/Users/UserProfileForm";
import { useAuth } from "../context/Authcontext";
import mockUsers from "../data/mockUsers";

const UserProfilePage = () => {
  const { user } = useAuth();
  const [profileData, setProfileData] = useState(null);

  useEffect(() => {
    if (user) {
      const currentUser = mockUsers.find(
        (mockUser) => mockUser.email === user.email
      );
      if (currentUser) {
        setProfileData(currentUser);
      }
    }
  }, [user]);

  const handleUpdateProfile = (updatedData) => {
    // Xử lý cập nhật thông tin người dùng ở đây
    console.log(updatedData);
    // Sau khi cập nhật thành công, cập nhật state
    setProfileData(updatedData);
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
