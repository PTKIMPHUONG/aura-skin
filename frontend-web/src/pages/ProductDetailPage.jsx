import React, { useEffect, useState } from "react";
import { useParams, useLocation } from "react-router-dom";
import { Container, Grid, Box, CircularProgress } from "@mui/material";
import ProductImageGallery from "../components/ProductDetail/ProductImageGallery";
import ProductInfo from "../components/ProductDetail/ProductInfo";
import ProductDetails from "../components/ProductDetail/ProductDetails";
import ProductReviews from "../components/ProductDetail/ProductReviews";
import {
  RecommendedProducts,
  ViewedProducts,
} from "../components/ProductDetail/RecommentsViewed";
import BreadcrumbNavigation from "../components/BreadcrumbNav/BreadcrumbNavigation";
import ProductService from "../services/ProductService";

function ProductDetailPage() {
  const [product, setProduct] = useState(null);
  const [variants, setVariants] = useState([]);
  const [selectedVariantImage, setSelectedVariantImage] = useState(null);
  const { id } = useParams();
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const productId = id || searchParams.get("id");

  useEffect(() => {
    const fetchProductData = async () => {
      if (!productId) {
        console.error("Product ID is undefined");
        return;
      }
      try {
        const response = await ProductService.getProductById(productId);
        setProduct(response.data);
        const variantsData = await ProductService.getVariantsByProductId(
          productId
        );
        setVariants(Array.isArray(variantsData) ? variantsData : []);
      } catch (error) {
        console.error("Error fetching product data:", error);
      }
    };
    fetchProductData();
  }, [productId]);

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

  const allImages = [product.default_image, ...variants.map((v) => v.image)];

  return (
    <Container maxWidth="lg" sx={{ px: { xs: 2, sm: 3, md: 4 } }}>
      <BreadcrumbNavigation product={product} />
      <Grid container spacing={4} sx={{ mt: 2 }}>
        <Grid item xs={12} md={6}>
          <Box sx={{ pr: { md: 2 } }}>
            <ProductImageGallery
              images={allImages}
              selectedVariantImage={selectedVariantImage}
            />
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <Box sx={{ pl: { md: 2 } }}>
            <ProductInfo
              product={product}
              variants={variants}
              onVariantSelect={(image) => setSelectedVariantImage(image)}
            />
          </Box>
        </Grid>
      </Grid>

      <Box mt={4}>
        <ProductDetails product={product} />
      </Box>

      <Box mt={4}>
        <ProductReviews reviews={product.reviews || []} />
      </Box>

      <Box mt={4}>
        <RecommendedProducts products={[]} />
      </Box>

      <Box mt={4}>
        <ViewedProducts products={[]} />
      </Box>
    </Container>
  );
}

export default ProductDetailPage;
