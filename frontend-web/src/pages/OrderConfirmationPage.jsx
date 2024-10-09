import React from "react";
import { useLocation, useNavigate } from "react-router-dom";
import OrderConfirmation from "../components/Order/OrderConfirmation";
import mockAddressUsers from "../data/mockAddressUsers";
import { useAuth } from "../context/Authcontext";
import { useCart } from "../context/CartContext";

const OrderConfirmationPage = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const { user } = useAuth();
  const { cartItems, totalPrice } = location.state || {
    cartItems: [],
    totalPrice: 0,
  };
  // const { state } = useLocation();
  // const { cartItems = [], totalPrice = 0 } = state || {};

  // Lấy địa chỉ của user đang đăng nhập
  const userAddresses =
    mockAddressUsers.find((addressUser) => addressUser.userId === user.id)
      ?.addresses || [];

  const handleSubmitOrder = (orderData) => {
    // Xử lý logic đặt hàng ở đây
    console.log("Đơn hàng đã được đặt:", orderData);
    // Chuyển hướng đến trang cảm ơn hoặc trang chi tiết đơn hàng
    navigate("/order-success");
  };

  return (
    <OrderConfirmation
      cartItems={cartItems}
      totalPrice={totalPrice}
      onSubmitOrder={handleSubmitOrder}
      userAddresses={userAddresses}
    />
  );
};

export default OrderConfirmationPage;
