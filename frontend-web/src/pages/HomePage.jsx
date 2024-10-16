import React, { useState, useEffect } from "react";
import { Container, Box, CircularProgress, Typography } from "@mui/material";
import BannerSlider from "../components/Home/BannerSlider";
import ProductCategories from "../components/Home/HomeProductCategories";
import HomeProductList from "../components/Home/HomeProductList";
import HomeSuggestList from "../components/Variants/HomeSuggestList";
import PromotionalBanner from "../components/Home/PromotionalBanner";
import BrandLogos from "../components/Home/BrandLogos";
import ProductService from "../services/ProductService";
import ProductVariantService from "../services/ProductVariantService";
import { useAuth } from "../context/Authcontext";

const HomePage = () => {
  const [featuredProducts, setFeaturedProducts] = useState([]);
  const [newProducts, setNewProducts] = useState([]);
  const [suggestedVariants, setSuggestedVariants] = useState([]);
  const { user } = useAuth();
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        setIsLoading(true);
        const response = await ProductService.getAllProducts();

        if (response && response.data && response.data.data) {
          const productsData = response.data.data;
          if (Array.isArray(productsData)) {
            setFeaturedProducts(productsData.slice(0, 4));
            setNewProducts(productsData.slice(4, 8));
          } else {
            console.error("Products data is not an array:", productsData);
            setError("Dữ liệu sản phẩm không hợp lệ");
          }
        } else {
          console.error("Invalid response structure:", response);
          setError("Cấu trúc phản hồi không hợp lệ");
        }
      } catch (error) {
        console.error("Error fetching products:", error);
        setError("Không thể tải danh sách sản phẩm. Vui lòng thử lại sau.");
      } finally {
        setIsLoading(false);
      }
    };

    const fetchSuggestedVariants = async () => {
      if (user && user.id) {
        try {
          const response =
            await ProductVariantService.getSuggestVariantsForUser(user.id);
          setSuggestedVariants(response.data || []);
        } catch (error) {
          console.error("Error fetching suggested variants:", error);
        }
      }
    };

    fetchProducts();
    fetchSuggestedVariants();
  }, [user]);

  return (
    <>
      <Box sx={{ maxWidth: "1400px", margin: "0 auto", padding: "0 8px" }}>
        <BannerSlider />
      </Box>
      <Container maxWidth="lg">
        <Box my={10}>
          <ProductCategories />
        </Box>
        {isLoading ? (
          <CircularProgress />
        ) : error ? (
          <Typography color="error">{error}</Typography>
        ) : (
          <>
            {featuredProducts.length > 0 && (
              <Box my={10}>
                <HomeProductList
                  title="Sản phẩm nổi bật"
                  products={featuredProducts}
                />
              </Box>
            )}
            <Box my={10}>
              <PromotionalBanner />
            </Box>
            {newProducts.length > 0 && (
              <Box my={10}>
                <HomeProductList title="Sản phẩm mới" products={newProducts} />
              </Box>
            )}
            {user && suggestedVariants.length > 0 && (
              <Box my={10}>
                <HomeSuggestList
                  title="Sản phẩm đề xuất cho bạn"
                  variants={suggestedVariants}
                />
              </Box>
            )}
          </>
        )}
        <Box my={10}>
          <BrandLogos />
        </Box>
      </Container>
    </>
  );
};

export default HomePage;
