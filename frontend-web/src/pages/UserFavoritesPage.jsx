import React from "react";
import { Box, Container, Typography } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import FavoriteProductsList from "../components/Users/FavoritesProductList";
import mockFavorites from "../data/mockFavorites";
import { useAuth } from "../context/Authcontext";

const UserFavoritesPage = () => {
  const { user } = useAuth();

  if (!user) {
    return (
      <Typography>Vui lòng đăng nhập để xem sản phẩm yêu thích</Typography>
    );
  }

  // Tìm thông tin yêu thích của user trong mockFavorites
  const userFavorites = mockFavorites.find(
    (mockUser) => mockUser.email === user.email
  );

  if (!userFavorites) {
    return (
      <Typography>
        Không tìm thấy danh sách yêu thích cho người dùng này
      </Typography>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Box sx={{ display: "flex" }}>
        <SidebarUser />
        <FavoriteProductsList favorites={userFavorites.favorites || []} />
      </Box>
    </Container>
  );
};

export default UserFavoritesPage;
