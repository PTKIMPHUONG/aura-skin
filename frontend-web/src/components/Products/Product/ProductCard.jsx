import React from "react";
import { Card, CardContent, CardMedia, Typography, Box } from "@mui/material";
import { styled } from "@mui/material/styles";
import { Link } from "react-router-dom";

const StyledCard = styled(Card)(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  height: "100%",
  borderRadius: "16px",
  boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
}));

const StyledCardMedia = styled(CardMedia)({
  height: 0,
  paddingTop: "100%", // Tỷ lệ 1:1
});

const StyledCardContent = styled(CardContent)({
  "& *": {
    fontFamily: "Jost, sans-serif !important",
  },
});

const DiscountBadge = styled(Box)({
  position: "absolute",
  top: "10px",
  right: "10px",
  backgroundColor: "#D10000",
  color: "white",
  padding: "4px 8px",
  borderRadius: "4px",
  fontSize: "0.75rem",
  fontWeight: "bold",
  fontFamily: "Jost, sans-serif",
});

const ProductCard = ({ product }) => {
  return (
    <Link
      to={`/products/product-detail/${product.product_id}`}
      style={{ textDecoration: "none" }}
    >
      <StyledCard>
        <Box sx={{ position: "relative" }}>
          <StyledCardMedia
            image={product.default_image || "đường dẫn ảnh mặc định"}
            title={product.product_name || "Tên sản phẩm"}
          />
        </Box>
        <StyledCardContent>
          <Typography
            gutterBottom
            variant="h6"
            component="div"
            sx={{ fontWeight: 500, fontSize: "20px", textAlign: "center" }}
          >
            {product.product_name || "Tên sản phẩm"}
          </Typography>
          <Box
            display="flex"
            alignItems="center"
            justifyContent="center"
            sx={{ textAlign: "center" }}
          >
            <Typography
              variant="h6"
              color="#000000"
              sx={{ fontSize: "16px", fontWeight: "unset", mr: 1 }}
            >
              {product.default_price
                ? `${product.default_price.toLocaleString("vi-VN")}đ`
                : "Giá không có sẵn"}
            </Typography>
          </Box>
        </StyledCardContent>
      </StyledCard>
    </Link>
  );
};

export default ProductCard;
