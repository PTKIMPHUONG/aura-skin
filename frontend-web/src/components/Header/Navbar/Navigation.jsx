import React, { useState, useRef, useEffect } from "react";
import {
  Box,
  Button,
  IconButton,
  Popper,
  Paper,
  MenuList,
  MenuItem,
  Grow,
  ClickAwayListener,
} from "@mui/material";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import mockBrands from "../../../data/mockBrands";
import mockCategories from "../../../data/mockCategories";

const Navigation = () => {
  const navigate = useNavigate();
  const [openMenu, setOpenMenu] = useState(null);
  const [isInitialized, setIsInitialized] = useState(false);
  const timeoutRef = useRef();
  const buttonRefs = useRef({});

  useEffect(() => {
    const menuItems = [
      "danhMucSanPham",
      "chamSocDa",
      "trangDiem",
      "thuongHieu",
    ];
    menuItems.forEach((menu) => {
      buttonRefs.current[menu] = React.createRef();
    });
    setIsInitialized(true);

    return () => {
      clearTimeout(timeoutRef.current);
    };
  }, []);

  const handleMouseEnter = (menu) => {
    clearTimeout(timeoutRef.current);
    setOpenMenu(menu);
  };

  const handleMouseLeave = () => {
    timeoutRef.current = setTimeout(() => {
      setOpenMenu(null);
    }, 0);
  };

  const handleClickAway = () => {
    setOpenMenu(null);
  };

  const buttonStyle = {
    color: "black", // Đặt màu chữ thành đen
    margin: "0px 20px",
    fontFamily: "Asap",
    "&:hover": {
      backgroundColor: "transparent", // Giữ nền trong suốt khi hover
      color: "black", // Giữ màu chữ đen khi hover
    },
  };

  const handleCategoryClick = (categoryId) => {
    navigate(`/category/${categoryId}`);
  };

  const renderMenu = (menu, items, label) => {
    const isOpen = openMenu === menu;

    return (
      <Box
        onMouseEnter={() => handleMouseEnter(menu)}
        onMouseLeave={handleMouseLeave}
      >
        <Button
          ref={buttonRefs.current[menu]}
          sx={buttonStyle}
          endIcon={<KeyboardArrowDownIcon />}
        >
          {label}
        </Button>
        {isInitialized && (
          <Popper
            open={isOpen}
            anchorEl={buttonRefs.current[menu]?.current}
            role={undefined}
            transition
            disablePortal={false}
            placement="bottom-start"
            style={{ zIndex: 1300 }} // Thêm z-index cao hơn
          >
            {({ TransitionProps, placement }) => (
              <Grow
                {...TransitionProps}
                style={{
                  transformOrigin:
                    placement === "bottom" ? "center top" : "center bottom",
                }}
              >
                <Paper elevation={3}>
                  <ClickAwayListener onClickAway={handleClickAway}>
                    <MenuList
                      autoFocusItem={isOpen}
                      id={`menu-list-grow-${menu}`}
                      onMouseEnter={() => handleMouseEnter(menu)}
                      onMouseLeave={handleMouseLeave}
                    >
                      {items.map((item) => (
                        <MenuItem
                          key={item.id || item.path}
                          component={RouterLink}
                          to={
                            menu === "danhMucSanPham"
                              ? `/category/${item.id}`
                              : `/${menu}/${item.id || item.path}`
                          }
                          onClick={() => setOpenMenu(null)}
                        >
                          {item.name || item.label}
                        </MenuItem>
                      ))}
                    </MenuList>
                  </ClickAwayListener>
                </Paper>
              </Grow>
            )}
          </Popper>
        )}
      </Box>
    );
  };

  const chamSocDaItems = [
    { label: "Sản phẩm 1", path: "san-pham-1" },
    { label: "Sản phẩm 2", path: "san-pham-2" },
  ];

  const trangDiemItems = [
    { label: "Sản phẩm A", path: "san-pham-a" },
    { label: "Sản phẩm B", path: "san-pham-b" },
  ];

  return (
    <Box
      sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        my: 2,
      }}
    >
      <IconButton component={RouterLink} to="/" sx={{ color: "black", mr: 1 }}>
        <HomeOutlinedIcon sx={{ fontSize: "40px" }} />
      </IconButton>
      {isInitialized && (
        <>
          {renderMenu("danhMucSanPham", mockCategories, "DANH MỤC SẢN PHẨM")}
          <Button component={RouterLink} to="/products" sx={buttonStyle}>
            DANH SÁCH SẢN PHẨM
          </Button>
          <Button
            component={RouterLink}
            to="/san-pham-ban-chay"
            sx={buttonStyle}
          >
            SẢN PHẨM BÁN CHẠY
          </Button>
          <Button component={RouterLink} to="/new" sx={buttonStyle}>
            NEW
          </Button>
          {renderMenu("chamSocDa", chamSocDaItems, "CHĂM SÓC DA")}
          {renderMenu("trangDiem", trangDiemItems, "TRANG ĐIỂM")}
          {renderMenu("thuongHieu", mockBrands, "THƯƠNG HIỆU")}
        </>
      )}
    </Box>
  );
};

export default Navigation;
