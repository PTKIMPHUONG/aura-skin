import React, { useState } from "react";
import {
  AppBar,
  Toolbar,
  InputBase,
  IconButton,
  Typography,
  Box,
  Divider,
} from "@mui/material";
import { styled, alpha, useTheme } from "@mui/material/styles";
import SearchIcon from "@mui/icons-material/Search";
import ShoppingBagOutlinedIcon from "@mui/icons-material/ShoppingBagOutlined";
import FavoriteBorderOutlinedIcon from "@mui/icons-material/FavoriteBorderOutlined";
import PersonOutlineOutlinedIcon from "@mui/icons-material/PersonOutlineOutlined";
import LocationOnOutlinedIcon from "@mui/icons-material/LocationOnOutlined";
import Navigation from "./Navbar/Navigation"; // Thêm import này
import { Link, useNavigate } from "react-router-dom";
import { Avatar, Menu, MenuItem } from "@mui/material";
import { useAuth } from "../../context/Authcontext";
import FaceIcon from "@mui/icons-material/Face"; // Import FaceIcon

// Điều chỉnh Search component
const Search = styled("div")(({ theme }) => ({
  position: "relative",
  borderRadius: theme.shape.borderRadius,
  backgroundColor: alpha(theme.palette.common.white, 0.15),
  "&:hover": {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  marginLeft: "auto",
  marginRight: "auto",
  width: "60%",
  display: "flex",
  alignItems: "center",
}));

// Điều chỉnh SearchIconWrapper
const SearchIconWrapper = styled("div")(({ theme }) => ({
  position: "absolute",
  right: 0,
  top: 0,
  height: "100%",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
}));

// Điều chỉnh StyledInputBase
const StyledInputBase = styled(InputBase)(({ theme }) => ({
  color: "inherit",
  width: "100%",
  backgroundColor: "white",
  "& .MuiInputBase-input": {
    padding: "12px",
    paddingRight: `calc(1em + ${theme.spacing(4)})`, // Thêm paddingRight
    transition: theme.transitions.create("width"),
    width: "100%",
    boxShadow: "0px 4px 4px 0px rgba(0, 0, 0, 0.25)",
  },
}));

// Điều chỉnh StyledToolbar
const StyledToolbar = styled(Toolbar)(({ theme }) => ({
  display: "flex",
  justifyContent: "space-between",
  alignItems: "center",
  fontFamily: "Inter",
  padding: theme.spacing(1, 2),
  minHeight: "100px !important", // Sử dụng !important để ghi đè
  [`${theme.breakpoints.up("xs")} and (orientation: landscape)`]: {
    minHeight: "100px !important",
  },
  [theme.breakpoints.up("sm")]: {
    minHeight: "100px !important",
  },
}));

function Header() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  const [anchorEl, setAnchorEl] = useState(null);

  const theme = useTheme();

  const iconStyle = {
    color: theme.palette.primary.main,
    "& .MuiSvgIcon-root": { fontSize: 32 },
  };

  const handleSearch = () => {
    // Xử lý logic tìm kiếm ở đây
    console.log("Tìm kiếm được kích hoạt");
  };

  const handleMenuOpen = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleAccountDetails = () => {
    handleMenuClose();
    navigate("/user/profile");
  };

  const handleLogout = () => {
    handleMenuClose();
    logout();
    navigate("/");
  };

  // Thêm các hàm xử lý cho location, cart và favorite
  const handleLocationClick = () => {
    console.log("Location icon clicked");
    navigate("/user/addresses");
  };

  const handleCartClick = () => {
    console.log("Cart icon clicked");
    navigate("/user/cart");
  };

  const handleFavoriteClick = () => {
    console.log("Favorite icon clicked");
    navigate("/user/favorites");
  };

  return (
    <AppBar position="static" color="primary" elevation={0}>
      <StyledToolbar disableGutters>
        {" "}
        {/* Thêm prop disableGutters */}
        <Box sx={{ display: "flex", alignItems: "center", flexGrow: 1 }}>
          <Search sx={{ marginLeft: "auto", marginRight: "11%" }}>
            <StyledInputBase
              placeholder="Tìm sản phẩm, danh mục hoặc thương hiệu mong muốn..."
              inputProps={{ "aria-label": "search" }}
            />
            <SearchIconWrapper>
              <IconButton
                sx={iconStyle}
                onClick={handleSearch}
                aria-label="search button"
              >
                <SearchIcon />
              </IconButton>
            </SearchIconWrapper>
          </Search>
        </Box>
        <Box
          sx={{ display: "flex", marginRight: "48px", alignItems: "center" }}
        >
          <Typography variant="body2" sx={{ fontSize: "16px", padding: "8px" }}>
            Hotline: 0899787933
          </Typography>
          <Typography
            sx={{
              mx: 4,
              fontSize: "32px", // Tăng kích thước chữ để làm dấu dài hơn
              fontWeight: "100",
              color: theme.palette.primary.main, // Sử dụng màu primary
              userSelect: "none", // Ngăn người dùng chọn dấu này
            }}
          >
            │ {/* Sử dụng ký tự đường thẳng dọc Unicode */}
          </Typography>
          <Box sx={{ display: "flex" }}>
            <IconButton sx={iconStyle} onClick={handleLocationClick}>
              <LocationOnOutlinedIcon />
            </IconButton>
            <IconButton sx={iconStyle} onClick={handleCartClick}>
              <ShoppingBagOutlinedIcon />
            </IconButton>
            <IconButton sx={iconStyle} onClick={handleFavoriteClick}>
              <FavoriteBorderOutlinedIcon />
            </IconButton>
            {user ? (
              <>
                <IconButton onClick={handleMenuOpen} sx={iconStyle}>
                  {user.user_image ? (
                    <Avatar
                      src={user.user_image}
                      alt={user.username}
                      sx={{ width: 32, height: 32 }}
                    />
                  ) : (
                    <FaceIcon sx={{ fontSize: 40 }} /> // Sử dụng FaceIcon và điều chỉnh kích thước nếu cần
                  )}
                </IconButton>
                <Menu
                  anchorEl={anchorEl}
                  open={Boolean(anchorEl)}
                  onClose={handleMenuClose}
                >
                  <MenuItem onClick={handleAccountDetails}>
                    Chi tiết tài khoản
                  </MenuItem>
                  <MenuItem onClick={handleLogout}>Đăng xuất</MenuItem>
                </Menu>
              </>
            ) : (
              <IconButton component={Link} to="/login" sx={iconStyle}>
                <PersonOutlineOutlinedIcon />
              </IconButton>
            )}
          </Box>
        </Box>
      </StyledToolbar>
      <Box sx={{ bgcolor: "background.paper", py: 1 }}>
        <Typography
          variant="h4"
          align="center"
          sx={{
            fontFamily: "Arima",
            marginTop: "48px",
            fontWeight: "400",
            fontSize: "48px",
          }}
        >
          AURA SKIN
        </Typography>
      </Box>
      <Box sx={{ bgcolor: "background.paper", py: 1 }}>
        <Navigation textColor="black" />
      </Box>
      <Box
        sx={{
          bgcolor: "background.paper",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          // py: 1,
        }}
      >
        <Divider
          sx={{
            borderColor: "black", // Màu của đường kẻ
            borderBottomWidth: 1, // Độ dày của đường kẻ
            width: "85%", // Chiều rộng 100%
            // my: 2, // Margin top và bottom
            marginBottom: "72px !important",
          }}
        />
      </Box>
    </AppBar>
  );
}

export default Header;
