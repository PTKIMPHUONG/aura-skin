import React from "react";
import { Grid, Box, Typography } from "@mui/material";
import ProductCard from "../Products/Product/ProductCard";

const HomeProductList = ({ title, products }) => {
  return (
    <Box>
      <Typography
        mt={10}
        variant="h4"
        gutterBottom
        sx={{ textAlign: "center" }}
      >
        {title}
      </Typography>
      <Grid container mt={2} spacing={2}>
        {products.slice(0, 8).map((product) => (
          <Grid item xs={6} sm={3} key={product.id}>
            <ProductCard product={product} />
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default HomeProductList;
