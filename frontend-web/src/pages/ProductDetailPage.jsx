import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { Container, Grid, Box, CircularProgress } from "@mui/material";
import ProductImageGallery from "../components/ProductDetail/ProductImageGallery";
import ProductInfo from "../components/ProductDetail/ProductInfo";
import ProductDetails from "../components/ProductDetail/ProductDetails";
import ProductReviews from "../components/ProductDetail/ProductReviews";
import {
  RecommendedProducts,
  ViewedProducts,
} from "../components/ProductDetail/RecommentsViewed";
import mockProducts from "../data/mockProducts";
import BreadcrumbNavigation from "../components/BreadcrumbNav/BreadcrumbNavigation";

function ProductDetailPage() {
  const [product, setProduct] = useState(null);
  const location = useLocation();

  useEffect(() => {
    console.log("ProductDetailPage mounted");
    const searchParams = new URLSearchParams(location.search);
    const productId = searchParams.get("id");
    console.log("Product ID:", productId);

    const foundProduct = mockProducts.find(
      (p) => p.id.toString() === productId
    );
    console.log("Found product:", foundProduct);

    if (foundProduct) {
      setProduct(foundProduct);
    } else {
      console.error("Không tìm thấy sản phẩm");
    }
  }, [location]);

  if (!product) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ px: { xs: 2, sm: 3, md: 4 } }}>
      <BreadcrumbNavigation product={product} />
      <Grid container spacing={4} sx={{ mt: 2 }}>
        <Grid item xs={12} md={6}>
          <Box sx={{ pr: { md: 2 } }}>
            <ProductImageGallery images={product.images} />
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <Box sx={{ pl: { md: 2 } }}>
            <ProductInfo product={product} />
          </Box>
        </Grid>
      </Grid>

      <Box mt={4}>
        <ProductDetails product={product} />
      </Box>

      <Box mt={4}>
        <ProductReviews reviews={product.reviews} />
      </Box>

      <Box mt={4}>
        <RecommendedProducts products={product.recommendedProducts} />
      </Box>

      <Box mt={4}>
        <ViewedProducts products={product.viewedProducts} />
      </Box>
    </Container>
  );
}

export default ProductDetailPage;
