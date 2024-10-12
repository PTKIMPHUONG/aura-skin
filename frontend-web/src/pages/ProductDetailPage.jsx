import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Container, Grid, Box, CircularProgress, Button } from "@mui/material";
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
  const [selectedVariant, setSelectedVariant] = useState(null);
  const { id } = useParams();

  useEffect(() => {
    const fetchProductData = async () => {
      try {
        const productData = await ProductService.getProductById(id);
        setProduct(productData);
        const variantsData = await ProductService.getVariantsByProductId(id);
        setVariants(variantsData);
        if (variantsData.length > 0) {
          setSelectedVariant(variantsData[0]);
        }
      } catch (error) {
        console.error("Error fetching product data:", error);
      }
    };
    fetchProductData();
  }, [id]);

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
            <ProductImageGallery
              images={
                selectedVariant
                  ? selectedVariant.description_images
                  : [product.default_image]
              }
            />
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <Box sx={{ pl: { md: 2 } }}>
            <ProductInfo product={product} selectedVariant={selectedVariant} />
            <Box mt={2}>
              {variants.map((variant) => (
                <Button
                  key={variant.variant_id}
                  variant={
                    selectedVariant &&
                    selectedVariant.variant_id === variant.variant_id
                      ? "contained"
                      : "outlined"
                  }
                  onClick={() => setSelectedVariant(variant)}
                  sx={{ mr: 1, mb: 1 }}
                >
                  {variant.variant_name}
                </Button>
              ))}
            </Box>
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
