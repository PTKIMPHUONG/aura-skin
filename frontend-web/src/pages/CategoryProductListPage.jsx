import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { Container, Typography, Grid, Box } from "@mui/material";
import Sidebar from "../components/Sidebar/SidebarProducts";
import ProductList from "../components/Products/ProductList";
import mockProducts from "../data/mockProducts";
import mockCategories from "../data/mockCategories";

const CategoryProductListPage = () => {
  const { categoryId } = useParams();
  const [products, setProducts] = useState([]);
  const [category, setCategory] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);

  useEffect(() => {
    const selectedCategory = categoryId
      ? mockCategories.find((cat) => cat.id === parseInt(categoryId))
      : null;
    setCategory(selectedCategory);

    console.log("Selected Category:", selectedCategory);

    const filteredProducts = selectedCategory
      ? mockProducts.filter(
          (product) => product.category === selectedCategory.name
        )
      : mockProducts;

    console.log("Filtered Products:", filteredProducts);

    setProducts(filteredProducts);
  }, [categoryId]);

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
  };

  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        Sản phẩm
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={3}>
          <Sidebar categoryName={category ? category.name : null} />
        </Grid>
        <Grid item xs={12} md={9}>
          <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            mb={4}
          >
            {/* Thêm các tùy chọn sắp xếp nếu cần */}
          </Box>
          <ProductList
            products={products}
            currentPage={currentPage}
            onPageChange={handlePageChange}
          />
        </Grid>
      </Grid>
    </Container>
  );
};

export default CategoryProductListPage;
