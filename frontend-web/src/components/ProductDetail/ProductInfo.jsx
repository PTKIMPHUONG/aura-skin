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

function ProductInfo({ product }) {
  const [quantity, setQuantity] = useState(1);
  const { addToCart } = useContext(CartContext);
  const navigate = useNavigate();

  const discountPercentage = Math.round(
    ((product.originalPrice - product.price) / product.originalPrice) * 100
  );

  const handleBuyNow = () => {
    console.log("handleBuyNow called", product, quantity);
    addToCart({ ...product, quantity });
    navigate("/order-confirmation", {
      state: {
        cartItems: [{ ...product, quantity }],
        totalPrice: product.price * quantity,
      },
    });
  };

  return (
    <Box sx={{ maxWidth: "90%" }}>
      <Typography variant="h5" fontWeight="bold" gutterBottom>
        {product.name}
      </Typography>
      <Box display="flex" alignItems="center" my={1}>
        <Rating value={product.rating || 0} readOnly size="small" />
        <Typography variant="body2" ml={1}>
          ({product.reviews?.length || 0} đánh giá)
        </Typography>
        <Typography variant="body2" ml={1}>
          {product.soldCount || 0} đã bán
        </Typography>
      </Box>
      <Box display="flex" alignItems="center" my={1}>
        <Typography variant="h4" color="error" fontWeight="bold">
          {product.price.toLocaleString("vi-VN")}đ
        </Typography>
        <Typography
          variant="body1"
          color="text.secondary"
          sx={{ textDecoration: "line-through" }}
          ml={2}
        >
          {product.originalPrice.toLocaleString("vi-VN")}đ
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
