import React from "react";
import { Grid, Box, Pagination, Typography } from "@mui/material";
import ProductCard from "./Product/ProductCard";

const ProductList = ({ products, currentPage, onPageChange }) => {
  if (!Array.isArray(products)) {
    console.error("Products is not an array:", products);
    return <Typography>Không có sản phẩm nào để hiển thị.</Typography>;
  }

  const productsPerPage = 16; // 4 hàng x 4 cột
  const totalPages = Math.ceil(products.length / productsPerPage);

  const displayedProducts = products.slice(
    (currentPage - 1) * productsPerPage,
    currentPage * productsPerPage
  );

  if (displayedProducts.length === 0) {
    return <Typography>Không có sản phẩm nào để hiển thị.</Typography>;
  }

  return (
    <Box>
      <Grid container spacing={4}>
        {Array.from({ length: Math.ceil(displayedProducts.length / 4) }).map(
          (_, rowIndex) => (
            <Grid container item key={rowIndex} spacing={2}>
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
      {totalPages > 1 && (
        <Box sx={{ display: "flex", justifyContent: "center", mt: 8 }}>
          <Pagination
            count={totalPages}
            page={currentPage}
            onChange={onPageChange}
            color="primary"
          />
        </Box>
      )}
    </Box>
  );
};

export default ProductList;
