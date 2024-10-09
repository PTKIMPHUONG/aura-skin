import React from "react";
import {
  Grid,
  Card,
  CardMedia,
  CardContent,
  Typography,
  Box,
} from "@mui/material";
import { Link } from "react-router-dom";
import mockCategories from "../../data/mockCategories";

const ProductCategories = () => {
  return (
    <Box sx={{ my: 4 }}>
      <Typography variant="h5" sx={{ mb: 3, textAlign: "center" }}>
        DANH MỤC HOT
      </Typography>
      <Grid container spacing={2}>
        {mockCategories.map((category) => (
          <Grid item xs={6} sm={4} md={3} lg={2} key={category.id}>
            <Card
              component={Link}
              to={`/category/${category.id}`}
              sx={{
                height: "100%",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
                textAlign: "center",
                boxShadow: "none",
                textDecoration: "none", // Loại bỏ gạch chân mặc định của Link
                color: "inherit", // Giữ màu chữ mặc định
                "&:hover": {
                  boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
                },
              }}
            >
              <CardMedia
                component="img"
                sx={{
                  width: "80%",
                  height: "auto",
                  objectFit: "contain",
                  aspectRatio: "1 / 1",
                }}
                image={category.image}
                alt={category.name}
              />
              <CardContent>
                <Typography variant="body2" component="div">
                  {category.name}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default ProductCategories;
