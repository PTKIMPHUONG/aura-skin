import React from "react";
import { Grid, Box, Typography } from "@mui/material";
import ProductCard from "../Products/Product/ProductCard";

const HomeProductList = ({ title, products }) => {
  console.log("HomeProductList received:", { title, products });

  if (!Array.isArray(products)) {
    console.error("Products is not an array:", products);
    return <Typography>Không có sản phẩm nào để hiển thị.</Typography>;
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom sx={{ textAlign: "center" }}>
        {title}
      </Typography>
      {products.length > 0 ? (
        <Grid container spacing={2}>
          {products.map((product) => (
            <Grid item xs={6} sm={3} key={product.id}>
              <ProductCard product={product} />
            </Grid>
          ))}
        </Grid>
      ) : (
        <Typography>Không có sản phẩm nào.</Typography>
      )}
    </Box>
  );
};

export default HomeProductList;
