import React, { useState, useEffect } from "react";
import {
  Box,
  Button,
  Typography,
  CircularProgress,
  Toolbar,
} from "@mui/material";
import { Add as AddIcon } from "@mui/icons-material";
import AdminHeader from "../../components/Header/AdminHeader";
import AdminSidebar from "../../components/Sidebar/SidebarAdmin";
import ProductList from "../../components/Admin/AdminProducts/ProductList";
import ProductService from "../../services/ProductService";

const ProductsPage = () => {
  const [products, setProducts] = useState([]); // Khởi tạo là mảng rỗng
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    try {
      setLoading(true);
      const response = await ProductService.getAllProducts();
      console.log("Fetched products:", response.data.data); // Kiểm tra dữ liệu
      // Kiểm tra xem response.data có phải là mảng không
      if (Array.isArray(response.data.data)) {
        setProducts(response.data.data); // Đảm bảo là mảng
      } else {
        setProducts([]); // Nếu không phải mảng, gán mảng rỗng
        setError("Dữ liệu không hợp lệ.");
      }
      setLoading(false);
    } catch (err) {
      console.error("Error fetching products:", err);
      setError("Không thể tải danh sách sản phẩm. Vui lòng thử lại sau.");
      setLoading(false);
    }
  };

  const handleEditProduct = (product) => {
    console.log("Editing product:", product);
  };

  const handleDeleteProduct = async (productId) => {
    try {
      await ProductService.deleteProduct(productId);
      setProducts(products.filter((p) => p.product_id !== productId)); // Sử dụng product_id
    } catch (error) {
      console.error("Error deleting product:", error);
    }
  };

  return (
    <Box sx={{ display: "flex" }}>
      <AdminHeader />
      <AdminSidebar />
      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <Toolbar />
        {loading ? (
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            height="100vh"
          >
            <CircularProgress />
          </Box>
        ) : error ? (
          <Typography color="error">{error}</Typography>
        ) : (
          <>
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
                mb: 2,
              }}
            >
              <Typography variant="h5">Quản lý sản phẩm</Typography>
              <Button variant="contained" startIcon={<AddIcon />}>
                Thêm sản phẩm mới
              </Button>
            </Box>
            <ProductList
              products={products}
              onEditProduct={handleEditProduct}
              onDeleteProduct={handleDeleteProduct}
            />
          </>
        )}
      </Box>
    </Box>
  );
};

export default ProductsPage;
