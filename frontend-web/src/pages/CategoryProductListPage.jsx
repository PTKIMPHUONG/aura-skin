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
          const categoryResponse = await CategoryService.getCategoryById(
            categoryId
          );
          setCategory(categoryResponse.data);

          const productsResponse = await ProductService.getProductsByCategory(
            categoryId
          );
          setProducts(productsResponse.data);
        } else {
          const productsResponse = await ProductService.getAllProducts();
          setProducts(productsResponse.data);
        }
        setLoading(false);
      } catch (error) {
        setError("Có lỗi xảy ra khi tải dữ liệu");
        setLoading(false);
        console.error("Error fetching data:", error);
        // Xử lý lỗi ở đây, ví dụ: hiển thị thông báo lỗi
      }
    };

    fetchData();
  }, [categoryId]);

  const handlePageChange = async (event, value) => {
    setCurrentPage(value);
    try {
      const response = await ProductService.getProductsByCategory(
        categoryId,
        value
      );
      setProducts(response.data);
    } catch (error) {
      console.error("Error fetching products for page:", error);
    }
  };

  if (loading) return <Typography>Đang tải...</Typography>;
  if (error) return <Typography>Lỗi: {error}</Typography>;

  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        Sản phẩm
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={3}>
          <Sidebar category_id={category ? category.category_id : null} />
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
