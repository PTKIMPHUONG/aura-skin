import React, { useState, useEffect } from "react";
import {
  Container,
  Typography,
  Grid,
  Box,
  Button,
  Select,
  MenuItem,
} from "@mui/material";
import { useNavigate, useLocation } from "react-router-dom";
import ProductList from "../components/Products/ProductList";
import Sidebar from "../components/Sidebar/SidebarProducts";
import mockProducts from "../data/mockProducts";

const ProductListPage = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [sortBy, setSortBy] = useState("");
  const [currentPage, setCurrentPage] = useState(1);

  useEffect(() => {
    const searchParams = new URLSearchParams(location.search);
    const page = parseInt(searchParams.get("page")) || 1;
    setCurrentPage(page);
  }, [location.search]);

  const handleSortChange = (event) => {
    setSortBy(event.target.value);
  };

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
    navigate(`/products?page=${value}`);
  };

  // Thêm logic sắp xếp sản phẩm ở đây nếu cần

  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        Danh sách sản phẩm
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={3}>
          <Sidebar />
        </Grid>
        <Grid item xs={12} md={9}>
          <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            mb={4}
          >
            <Box>
              <span>Sắp xếp theo:</span>
              <Button variant="outlined" size="small" sx={{ mx: 1 }}>
                Mới nhất
              </Button>
              <Button variant="outlined" size="small" sx={{ mx: 1 }}>
                Bán chạy
              </Button>
              <Select
                value={sortBy}
                onChange={handleSortChange}
                displayEmpty
                variant="outlined"
                size="small"
                sx={{ mx: 1 }}
              >
                <MenuItem value="">Giá</MenuItem>
                <MenuItem value="asc">Giá tăng dần</MenuItem>
                <MenuItem value="desc">Giá giảm dần</MenuItem>
              </Select>
            </Box>
            <Button variant="outlined" size="small">
              Đặt lại
            </Button>
          </Box>
          <ProductList
            products={mockProducts}
            currentPage={currentPage}
            onPageChange={handlePageChange}
          />
        </Grid>
      </Grid>
    </Container>
  );
};

export default ProductListPage;
