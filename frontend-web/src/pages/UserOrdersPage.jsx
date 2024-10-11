import React, { useState, useEffect } from "react";
import { Box, Container, Typography } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import UserOrders from "../components/Users/UserOrders";
import { useAuth } from "../context/Authcontext";
import OrderService from "../services/OrderService";

const UserOrdersPage = () => {
  const { user } = useAuth();
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchOrders = async () => {
      if (user && user.id) {
        console.log("Fetching orders for user ID:", user.id);
        try {
          setLoading(true);
          const response = await OrderService.getOrdersByUserId(user.id);
          console.log("API response:", response);
          setOrders(response.data || []);
        } catch (err) {
          console.error("Error details:", err.response);
          setError("Không thể tải lịch sử đơn hàng. Vui lòng thử lại sau.");
        } finally {
          setLoading(false);
        }
      } else {
        console.log("User or user ID is not available"); // Thêm log này
        setLoading(false);
      }
    };

    fetchOrders();
  }, [user]);

  if (loading) return <Typography>Đang tải...</Typography>;
  if (error) return <Typography color="error">{error}</Typography>;

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Box sx={{ display: "flex" }}>
        <SidebarUser />
        <Box sx={{ flexGrow: 1, ml: 4 }}>
          <UserOrders orders={orders} />
        </Box>
      </Box>
    </Container>
  );
};

export default UserOrdersPage;
