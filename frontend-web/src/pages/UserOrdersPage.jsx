import React from "react";
import { Box, Container } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import UserOrders from "../components/Users/UserOrders";
import { useAuth } from "../context/Authcontext";
import mockOrders from "../data/mockOrders";

const UserOrdersPage = () => {
  const { user } = useAuth();
  const userOrders =
    mockOrders.find((order) => order.userId === user.id)?.orders || [];

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Box sx={{ display: "flex" }}>
        <SidebarUser />
        <Box sx={{ flexGrow: 1, ml: 4 }}>
          <UserOrders orders={userOrders} />
        </Box>
      </Box>
    </Container>
  );
};

export default UserOrdersPage;
