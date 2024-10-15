import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { Container, Typography, Grid, Box } from "@mui/material";
import Sidebar from "../components/Sidebar/SidebarProducts";
import ProductList from "../components/Products/ProductList";
import ProductService from "../services/ProductService";
import CategoryService from "../services/CategoryService";

const CategoryProductListPage = () => {
  const { categoryId } = useParams();
  const [products, setProducts] = useState([]);
  const [category, setCategory] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        setError(null);
        if (categoryId) {
          const [categoryResponse, productsResponse] = await Promise.all([
            CategoryService.getCategoryById(categoryId),
            ProductService.getProductsByCategory(categoryId),
          ]);
          setCategory(categoryResponse);
          setProducts(productsResponse.data);
        } else {
          // Xử lý trường hợp không có categoryId nếu cần
        }
        setLoading(false);
      } catch (error) {
        setError("Có lỗi xảy ra khi tải dữ liệu");
        setLoading(false);
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [categoryId]);

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
  };

  if (loading) return <Typography>Đang tải...</Typography>;
  if (error) return <Typography>Lỗi: {error}</Typography>;

  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        {category ? `Sản phẩm: ${category.category_name}` : "Tất cả sản phẩm"}
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={3}>
          <Sidebar
            categoryName={category ? category.category_name : "Tất cả sản phẩm"}
          />
        </Grid>
        <Grid item xs={12} md={9}>
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
