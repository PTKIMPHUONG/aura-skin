import React from "react";
import { Box, Container } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import UserAddress from "../components/Users/UserAddress";
import { useAuth } from "../context/Authcontext";
import mockAddressUsers from "../data/mockAddressUsers";

const UserAddressesPage = () => {
  const { user } = useAuth();
  const userAddresses =
    mockAddressUsers.find((addressUser) => addressUser.userId === user.id)
      ?.addresses || [];

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Box sx={{ display: "flex" }}>
        <SidebarUser />
        <Box sx={{ flexGrow: 1, ml: 4 }}>
          <UserAddress addresses={userAddresses} />
        </Box>
      </Box>
    </Container>
  );
};

export default UserAddressesPage;
