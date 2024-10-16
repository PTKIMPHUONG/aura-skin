import React from "react";
import {
  Box,
  Avatar,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import {
  Person,
  ShoppingBag,
  LocationOn,
  LocalOffer,
  Favorite,
  ExitToApp,
  ShoppingCart,
} from "@mui/icons-material";
import { Link, useNavigate, useLocation } from "react-router-dom";
import { useAuth } from "../../context/Authcontext";

const SidebarUser = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  const menuItems = [
    { icon: <Person />, text: "Thông tin", path: "/user/profile" },
    { icon: <ShoppingBag />, text: "Quản lý đơn hàng", path: "/user/orders" },
    { icon: <ShoppingCart />, text: "Giỏ hàng", path: "/user/cart" },
    { icon: <LocationOn />, text: "Địa chỉ", path: "/user/addresses" },
    { icon: <LocalOffer />, text: "Mã giảm giá", path: "/user/coupons" },
    { icon: <Favorite />, text: "Sản phẩm yêu thích", path: "/user/favorites" },
  ];

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  console.log("User image URL:", user.user_image);

  return (
    <Box sx={{ width: 240, borderRight: "1px solid #e0e0e0", height: "100%" }}>
      <Box
        sx={{
          p: 2,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Avatar src={user.user_image} sx={{ width: 80, height: 80, mb: 1 }} />
        <Typography variant="subtitle1">{user.username}</Typography>
        <Typography variant="body2" color="text.secondary">
          {user.email}
        </Typography>
      </Box>
      <List>
        {menuItems.map((item) => (
          <ListItem
            button={true}
            key={item.text}
            component={Link}
            to={item.path}
            selected={location.pathname === item.path}
            sx={{
              color:
                location.pathname === item.path
                  ? "primary.main"
                  : "text.primary",
              "& .MuiListItemIcon-root": {
                color:
                  location.pathname === item.path
                    ? "primary.main"
                    : "text.primary",
              },
              "& .MuiListItemText-primary": {
                fontWeight: location.pathname === item.path ? "bold" : "normal",
              },
              "&:hover": {
                backgroundColor: "action.hover",
              },
            }}
          >
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.text} />
          </ListItem>
        ))}
        <ListItem
          button
          onClick={handleLogout}
          sx={{
            color: "text.primary",
            "& .MuiListItemIcon-root": {
              color: "text.primary",
            },
          }}
        >
          <ListItemIcon>
            <ExitToApp />
          </ListItemIcon>
          <ListItemText primary="Đăng xuất" />
        </ListItem>
      </List>
    </Box>
  );
};

export default SidebarUser;
