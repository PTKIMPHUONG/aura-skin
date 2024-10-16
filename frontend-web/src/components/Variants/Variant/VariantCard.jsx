import React from "react";
import {
  Card,
  CardMedia,
  CardContent,
  Typography,
  CardActionArea,
  Box,
} from "@mui/material";
import { Link, useNavigate } from "react-router-dom"; // Thay đổi ở đây
import ProductVariantService from "../../../services/ProductVariantService"; // Import service

const VariantCard = ({ variant }) => {
  const navigate = useNavigate(); // Sử dụng useNavigate

  const handleVariantClick = async () => {
    try {
      const response = await ProductVariantService.getProductByVariantId(
        variant.variant_id
      );
      const productId = response.data.product_id; // Giả sử API trả về product_id
      navigate(`/products/product-detail/${productId}`); // Điều hướng đến trang sản phẩm
    } catch (error) {
      console.error("Error fetching product ID:", error);
    }
  };

  return (
    <Card sx={{ height: "100%", display: "flex", flexDirection: "column" }}>
      <CardActionArea
        onClick={handleVariantClick} // Gọi hàm khi nhấp vào variant
        sx={{ flexGrow: 1, display: "flex", flexDirection: "column" }}
      >
        <CardMedia
          component="img"
          height="200"
          image={variant.thumbnail || "/placeholder-image.jpg"}
          alt={variant.variant_name}
        />
        <CardContent
          sx={{
            flexGrow: 1,
            display: "flex",
            flexDirection: "column",
            justifyContent: "space-between",
          }}
        >
          <Box>
            <Typography
              gutterBottom
              variant="subtitle1"
              component="div"
              sx={{
                overflow: "hidden",
                textOverflow: "ellipsis",
                display: "-webkit-box",
                WebkitLineClamp: 2,
                WebkitBoxOrient: "vertical",
                lineHeight: "1.2em",
                height: "2.4em",
              }}
            >
              {variant.variant_name}
            </Typography>
          </Box>
          <Typography variant="body2" color="text.secondary">
            {variant.price
              ? `${variant.price.toLocaleString("vi-VN")} ₫`
              : "Giá không có sẵn"}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
};

export default VariantCard;
