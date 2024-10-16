import React, { useEffect, useState } from "react";
import { Box, Container, Typography } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import FavoriteProductsList from "../components/Users/FavoritesProductList";
import UserService from "../services/UserService"; // Import UserService
import { useAuth } from "../context/Authcontext";

const UserFavoritesPage = () => {
  const { user } = useAuth();
  const [favorites, setFavorites] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchUserFavorites = async () => {
      if (!user || !user.id) {
        setError("Người dùng không hợp lệ.");
        setLoading(false);
        return;
      }

      try {
        const userFavorites = await UserService.getUserWishlist(user.id);
        setFavorites(userFavorites.data); // Đảm bảo bạn lấy đúng dữ liệu
      } catch (error) {
        console.error("Error fetching user favorites:", error);
        setError("Không thể tải danh sách yêu thích.");
      } finally {
        setLoading(false);
      }
    };

    fetchUserFavorites();
  }, [user]);

  if (loading) {
    return <Typography>Đang tải...</Typography>;
  }

  if (error) {
    return <Typography color="error">{error}</Typography>;
  }

  if (!favorites.length) {
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
        <Box sx={{ flexGrow: 1, ml: 4 }}>
          <FavoriteProductsList favorites={favorites} />
        </Box>
      </Box>
    </Container>
  );
};

export default UserFavoritesPage;
