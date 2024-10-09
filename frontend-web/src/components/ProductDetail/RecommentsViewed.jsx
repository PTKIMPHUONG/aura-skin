import React from "react";
import { Typography, Grid } from "@mui/material";
import ProductCard from "../Products/Product/ProductCard"; // Giả sử bạn đã có component ProductCard

function ProductList({ title, products }) {
  if (!products || products.length === 0) {
    return null;
  }
  return (
    <div>
      <Typography variant="h6" gutterBottom>
        {title}
      </Typography>
      <Grid container spacing={2}>
        {products.map((product) => (
          <Grid item xs={6} sm={4} md={3} key={product.id}>
            <ProductCard product={product} />
          </Grid>
        ))}
      </Grid>
    </div>
  );
}

export function RecommendedProducts({ products }) {
  return <ProductList title="Sản phẩm đề xuất" products={products} />;
}

export function ViewedProducts({ products }) {
  return <ProductList title="Sản phẩm đã xem" products={products} />;
}
