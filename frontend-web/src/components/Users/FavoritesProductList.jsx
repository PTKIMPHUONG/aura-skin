import React from "react";
import {
  Grid,
  Card,
  CardMedia,
  CardContent,
  Typography,
  Box,
  Stack,
} from "@mui/material";
import { useNavigate } from "react-router-dom";

const FavoriteProductsList = ({ favorites }) => {
  const navigate = useNavigate();

  const handleProductClick = (productId) => {
    navigate(`/products/product-detail?id=${productId}`);
  };

  const calculateDiscount = (price, originalPrice) => {
    return Math.round(((originalPrice - price) / originalPrice) * 100);
  };

  return (
    <Box sx={{ flexGrow: 1, padding: 2 }}>
      <Typography variant="h5" gutterBottom>
        Danh sách sản phẩm yêu thích
      </Typography>
      <Grid container spacing={2}>
        {favorites.map((product) => (
          <Grid item xs={12} sm={6} md={4} lg={3} key={product.id}>
            <Card
              sx={{
                cursor: "pointer",
                height: "100%",
                position: "relative",
              }}
              onClick={() => handleProductClick(product.id)}
            >
              <Box
                sx={{
                  position: "absolute",
                  top: 10,
                  right: 10,
                  bgcolor: "error.main",
                  color: "white",
                  padding: "2px 6px",
                  borderRadius: "4px",
                  zIndex: 1,
                }}
              >
                -{calculateDiscount(product.price, product.originalPrice)}%
              </Box>
              <CardMedia
                component="img"
                height="200"
                image={product.urlImage}
                alt={product.name}
              />
              <CardContent>
                <Typography gutterBottom variant="body1" component="div">
                  {product.name}
                </Typography>
                <Stack direction="row" spacing={1} alignItems="center">
                  <Typography
                    variant="body1"
                    color="error.main"
                    fontWeight="bold"
                  >
                    {product.price.toLocaleString()}đ
                  </Typography>
                  <Typography
                    variant="body2"
                    color="text.secondary"
                    sx={{ textDecoration: "line-through" }}
                  >
                    {product.originalPrice.toLocaleString()}đ
                  </Typography>
                </Stack>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default FavoriteProductsList;
