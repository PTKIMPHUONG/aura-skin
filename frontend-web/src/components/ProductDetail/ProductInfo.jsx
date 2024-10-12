import React, { useState, useContext } from "react";
import {
  Typography,
  Rating,
  Button,
  Box,
  Chip,
  TextField,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { CartContext } from "../../context/CartContext";

function ProductInfo({ product, selectedVariant }) {
  const [quantity, setQuantity] = useState(1);
  const { addToCart } = useContext(CartContext);
  const navigate = useNavigate();

  const price = selectedVariant ? selectedVariant.price : product.default_price;
  const discountPercentage = 0; // Tính toán giảm giá nếu có

  const handleBuyNow = () => {
    const itemToAdd = selectedVariant
      ? { ...selectedVariant, product_name: product.product_name }
      : product;
    addToCart({ ...itemToAdd, quantity });
    navigate("/order-confirmation", {
      state: {
        cartItems: [{ ...itemToAdd, quantity }],
        totalPrice: price * quantity,
      },
    });
  };

  return (
    <Box sx={{ maxWidth: "90%" }}>
      <Typography variant="h5" fontWeight="bold" gutterBottom>
        {product.product_name}
      </Typography>
      <Box display="flex" alignItems="center" my={1}>
        <Rating value={0} readOnly size="small" />
        <Typography variant="body2" ml={1}>
          (0 đánh giá)
        </Typography>
        <Typography variant="body2" ml={1}>
          0 đã bán
        </Typography>
      </Box>
      <Box display="flex" alignItems="center" my={1}>
        <Typography variant="h4" color="error" fontWeight="bold">
          {price.toLocaleString("vi-VN")}đ
        </Typography>
        {discountPercentage > 0 && (
          <Chip
            label={`-${discountPercentage}%`}
            color="success"
            size="small"
            sx={{ ml: 2 }}
          />
        )}
      </Box>
      <Box display="flex" alignItems="center" my={2}>
        <Typography variant="body1" mr={2}>
          Số lượng:
        </Typography>
        <TextField
          type="number"
          value={quantity}
          onChange={(e) =>
            setQuantity(Math.max(1, parseInt(e.target.value) || 1))
          }
          InputProps={{ inputProps: { min: 1 } }}
          size="small"
          sx={{ width: "100px" }}
        />
      </Box>
      <Box mt={4}>
        <Button
          variant="contained"
          color="primary"
          fullWidth
          sx={{ height: "50px", mb: 2 }}
          onClick={handleBuyNow}
        >
          Mua Online
        </Button>
        <Button variant="outlined" fullWidth sx={{ height: "50px" }}>
          Thêm vào giỏ hàng
        </Button>
      </Box>
    </Box>
  );
}

export default ProductInfo;
