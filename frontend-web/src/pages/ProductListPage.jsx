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
import ProductService from "../services/ProductService";

const ProductListPage = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [sortBy, setSortBy] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await ProductService.getAllProducts();
        // Giả sử API trả về một đối tượng có thuộc tính 'data' chứa mảng sản phẩm
        const productsData = response.data.data || [];
        setProducts(Array.isArray(productsData) ? productsData : []);
        setLoading(false);
      } catch (err) {
        console.error("Error fetching products:", err);
        setError("Có lỗi xảy ra khi tải danh sách sản phẩm");
        setProducts([]);
        setLoading(false);
      }
    };

    fetchProducts();
  }, []);

  useEffect(() => {
    const searchParams = new URLSearchParams(location.search);
    const page = parseInt(searchParams.get("page")) || 1;
    setCurrentPage(page);
  }, [location.search]);

  const handleSortChange = (event) => {
    setSortBy(event.target.value);
    // Thêm logic sắp xếp sản phẩm ở đây nếu cần
  };

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
    navigate(`/products?page=${value}`);
  };

  if (loading) return <div>Đang tải...</div>;
  if (error) return <div>{error}</div>;

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
                onChange={(e) => setSortBy(e.target.value)}
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
            products={products}
            currentPage={currentPage}
            onPageChange={(event, value) => setCurrentPage(value)}
          />
        </Grid>
      </Grid>
    </Container>
  );
};

export default ProductListPage;
