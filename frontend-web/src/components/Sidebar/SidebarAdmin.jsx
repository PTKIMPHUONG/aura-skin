import React from "react";
import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Toolbar,
} from "@mui/material";
import {
  Dashboard,
  Inventory,
  Category,
  People,
  ShoppingCart,
  Storage,
} from "@mui/icons-material";
import { Link } from "react-router-dom";

const drawerWidth = 240;

const menuItems = [
  { text: "Dashboard", icon: <Dashboard />, path: "/admin" },
  { text: "Products", icon: <Inventory />, path: "/admin/products" },
  { text: "Categories", icon: <Category />, path: "/admin/categories" },
  { text: "Customers", icon: <People />, path: "/admin/customers" },
  { text: "Orders", icon: <ShoppingCart />, path: "/admin/orders" },
  { text: "Stocks", icon: <Storage />, path: "/admin/stocks" },
];

const SidebarAdmin = () => {
  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: "border-box" },
      }}
    >
      <Toolbar />
      <List>
        {menuItems.map((item) => (
          <ListItem button key={item.text} component={Link} to={item.path}>
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.text} />
          </ListItem>
        ))}
      </List>
    </Drawer>
  );
};

export default SidebarAdmin;
