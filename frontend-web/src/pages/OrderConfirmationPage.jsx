import React, { useState, useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import OrderConfirmation from "../components/Order/OrderConfirmation";
import { useAuth } from "../context/Authcontext";
import { useCart } from "../context/CartContext";
import OrderService from "../services/OrderService";
import UserService from "../services/UserService";
import { Snackbar, Alert } from "@mui/material";

const OrderConfirmationPage = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const { user } = useAuth();
  const { cartItems, totalPrice } = location.state || {
    cartItems: [],
    totalPrice: 0,
  };
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState("");
  const [snackbarSeverity, setSnackbarSeverity] = useState("success");
  const [userAddresses, setUserAddresses] = useState([]);

  useEffect(() => {
    const fetchUserAddresses = async () => {
      try {
        const response = await UserService.getUserAddresses(user.id);
        setUserAddresses(response.data);
      } catch (error) {
        console.error("Error fetching user addresses:", error);
      }
    };
    fetchUserAddresses();
  }, [user.id]);

  const handleSubmitOrder = async (orderData) => {
    try {
      const response = await OrderService.createOrder(orderData);
      if (response.success) {
        setSnackbarMessage("Đặt hàng thành công!");
        setSnackbarSeverity("success");
        setOpenSnackbar(true);
        setTimeout(() => navigate("/order-success"), 2000);
      } else {
        throw new Error(response.message || "Đặt hàng không thành công");
      }
    } catch (error) {
      setSnackbarMessage(
        error.message || "Đặt hàng không thành công. Vui lòng thử lại."
      );
      setSnackbarSeverity("error");
      setOpenSnackbar(true);
    }
  };

  return (
    <>
      <OrderConfirmation
        cartItems={cartItems}
        totalPrice={totalPrice}
        onSubmitOrder={handleSubmitOrder}
        userAddresses={userAddresses}
      />
      <Snackbar
        open={openSnackbar}
        autoHideDuration={6000}
        onClose={() => setOpenSnackbar(false)}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert
          onClose={() => setOpenSnackbar(false)}
          severity={snackbarSeverity}
          sx={{ width: "100%" }}
        >
          {snackbarMessage}
        </Alert>
      </Snackbar>
    </>
  );
};

export default OrderConfirmationPage;
