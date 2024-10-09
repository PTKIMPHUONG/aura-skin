import React from "react";
import { Grid, Box, Pagination } from "@mui/material";
import ProductCard from "./Product/ProductCard";

const ProductList = ({ products, currentPage, onPageChange }) => {
  console.log("Products in ProductList:", products);
  const productsPerPage = 16; // 4 hàng x 4 cột
  const totalPages = Math.ceil(products.length / productsPerPage);

  const displayedProducts = products.slice(
    (currentPage - 1) * productsPerPage,
    currentPage * productsPerPage
  );

  return (
    <Box>
      <Grid container spacing={4}>
        {" "}
        {/* Khoảng cách giữa các hàng */}
        {Array.from({ length: Math.ceil(displayedProducts.length / 4) }).map(
          (_, rowIndex) => (
            <Grid container item key={rowIndex} spacing={2}>
              {" "}
              {/* Khoảng cách giữa các cột */}
              {displayedProducts
                .slice(rowIndex * 4, (rowIndex + 1) * 4)
                .map((product) => (
                  <Grid item xs={12} sm={6} md={3} key={product.id}>
                    <ProductCard product={product} />
                  </Grid>
                ))}
            </Grid>
          )
        )}
      </Grid>
      <Box sx={{ display: "flex", justifyContent: "center", mt: 8 }}>
        <Pagination
          count={totalPages}
          page={currentPage}
          onChange={onPageChange}
          color="primary"
        />
      </Box>
    </Box>
  );
};

export default ProductList;
