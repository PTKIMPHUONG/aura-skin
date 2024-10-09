import React, { useState, useEffect } from "react";
import { Box, Container } from "@mui/material";
import SidebarUser from "../components/Sidebar/SidebarUser";
import ProductCart from "../components/Users/ProductCart";
import { useAuth } from "../context/Authcontext";
import mockCarts from "../data/mockCarts";

const UserCartPage = () => {
  const { user } = useAuth();
  const [cartItems, setCartItems] = useState([]);

  useEffect(() => {
    if (user) {
      const userCart = mockCarts.find((cart) => cart.userId === user.id);
      if (userCart) {
        setCartItems(userCart.items);
      }
    }
  }, [user]);

  const handleQuantityChange = (id, newQuantity) => {
    setCartItems((prevItems) =>
      prevItems.map((item) =>
        item.productId === id
          ? { ...item, quantity: Math.max(1, newQuantity) }
          : item
      )
    );
  };

  const handleRemoveItem = (id) => {
    setCartItems((prevItems) =>
      prevItems.filter((item) => item.productId !== id)
    );
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Box sx={{ display: "flex" }}>
        <SidebarUser />
        <Box sx={{ flexGrow: 1, ml: 4 }}>
          <ProductCart
            cartItems={cartItems}
            onQuantityChange={handleQuantityChange}
            onRemoveItem={handleRemoveItem}
          />
        </Box>
      </Box>
    </Container>
  );
};

export default UserCartPage;
