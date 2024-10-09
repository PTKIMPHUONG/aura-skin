import React from "react";
import { Container, Box } from "@mui/material";
import BannerSlider from "../components/Home/BannerSlider";
import ProductCategories from "../components/Home/HomeProductCategories";
import HomeProductList from "../components/Home/HomeProductList";
import PromotionalBanner from "../components/Home/PromotionalBanner";
import BrandLogos from "../components/Home/BrandLogos";
import mockProducts from "../data/mockProducts";

const HomePage = () => {
  // Chọn 8 sản phẩm ngẫu nhiên cho featuredProducts
  const featuredProducts = mockProducts
    .sort(() => 0.5 - Math.random())
    .slice(0, 8);

  // Chọn 8 sản phẩm khác cho newProducts
  const newProducts = mockProducts
    .filter((product) => !featuredProducts.includes(product))
    .sort(() => 0.5 - Math.random())
    .slice(0, 8);

  return (
    <>
      <Box
        sx={{
          maxWidth: "1400px",
          margin: "0 auto",
          padding: "0 8px", // Thêm padding để tránh dính sát vào mép khi màn hình nhỏ
        }}
      >
        <BannerSlider />
      </Box>
      <Container maxWidth="lg">
        <Box my={10}>
          <ProductCategories />
        </Box>
        <Box my={10}>
          <HomeProductList
            title="Sản phẩm nổi bật"
            products={featuredProducts}
          />
        </Box>
        <Box my={10}>
          <PromotionalBanner />
        </Box>
        <Box my={10}>
          <HomeProductList title="Sản phẩm mới" products={newProducts} />
        </Box>
        <Box my={10}>
          <BrandLogos />
        </Box>
      </Container>
    </>
  );
};

export default HomePage;
