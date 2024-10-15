import React, { useState, useContext } from "react";
import {
  Typography,
  Rating,
  Button,
  Box,
  TextField,
  useTheme,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { CartContext } from "../../context/CartContext";

function ProductInfo({ product, variants, onVariantSelect }) {
  const [quantity, setQuantity] = useState(1);
  const [selectedVariant, setSelectedVariant] = useState(variants[0]);
  const { addToCart } = useContext(CartContext);
  const navigate = useNavigate();
  const theme = useTheme();

  if (!product) {
    return <Typography>Đang tải thông tin sản phẩm...</Typography>;
  }

  const price = selectedVariant?.price || product.default_price;

  const handleBuyNow = () => {
    const itemToAdd = selectedVariant
      ? { ...selectedVariant, product_name: product.product_name }
      : product;
    addToCart({ ...itemToAdd, quantity });
    navigate("/order-confirmation");
  };

  const handleAddToCart = () => {
    const itemToAdd = selectedVariant
      ? { ...selectedVariant, product_name: product.product_name }
      : product;
    addToCart({ ...itemToAdd, quantity });
  };

  const handleVariantSelect = (variant) => {
    setSelectedVariant(variant);
    onVariantSelect(variant.image);
  };

  return (
    <Box sx={{ maxWidth: "100%" }}>
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
      <Typography variant="h4" color="error" fontWeight="bold" my={2}>
        {price}đ
      </Typography>

      <Box my={2}>
        {variants.map((variant) => (
          <Button
            key={variant.variant_id}
            variant={selectedVariant === variant ? "contained" : "outlined"}
            fullWidth
            onClick={() => handleVariantSelect(variant)}
            sx={{
              mb: 1,
              justifyContent: "center",
              textAlign: "center",
              height: "auto",
              padding: "10px",
              whiteSpace: "normal",
              backgroundColor:
                selectedVariant === variant
                  ? theme.palette.primary.main
                  : "transparent",
              color:
                selectedVariant === variant
                  ? "white"
                  : theme.palette.primary.main,
              border: `1px solid ${theme.palette.primary.main}`,
              borderRadius: "4px",
              "&:hover": {
                backgroundColor:
                  selectedVariant === variant ? "#9F6DA8" : "#EFEFEF",
                color:
                  selectedVariant === variant
                    ? "white"
                    : theme.palette.primary.main,
              },
              transition: "background-color 0.3s, color 0.3s",
            }}
          >
            {variant.variant_name}
          </Button>
        ))}
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
          fullWidth
          sx={{
            height: "50px",
            mb: 2,
            backgroundColor: theme.palette.primary.main,
            "&:hover": {
              backgroundColor: theme.palette.primary.dark,
            },
          }}
          onClick={handleBuyNow}
        >
          MUA NGAY ONLINE
        </Button>
        <Button
          variant="outlined"
          fullWidth
          sx={{
            height: "50px",
            color: theme.palette.primary.main,
            borderColor: theme.palette.primary.main,
            "&:hover": {
              borderColor: theme.palette.primary.dark,
              color: theme.palette.primary.dark,
            },
          }}
          onClick={handleAddToCart}
        >
          THÊM VÀO GIỎ
        </Button>
      </Box>
    </Box>
  );
}

export default ProductInfo;
