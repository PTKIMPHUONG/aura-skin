import React, { useEffect, useState } from "react";
import { useParams, useLocation } from "react-router-dom";
import {
  Container,
  Grid,
  Box,
  CircularProgress,
  Typography,
} from "@mui/material";
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
import ProductVariantService from "../services/ProductVariantService";
import SuggestedVariants from "../components/Variants/SuggestedVariant";

function ProductDetailPage() {
  const [product, setProduct] = useState(null);
  const [variants, setVariants] = useState([]);
  const [selectedVariantImage, setSelectedVariantImage] = useState(null);
  const [relatedProducts, setRelatedProducts] = useState([]);
  const { id } = useParams();
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const productId = id || searchParams.get("id");
  const [selectedVariantId, setSelectedVariantId] = useState(null);
  const [allImages, setAllImages] = useState([]);

  useEffect(() => {
    const fetchProductData = async () => {
      if (!id) {
        console.error("Product ID is undefined");
        return;
      }
      try {
        const response = await ProductService.getProductById(id);
        setProduct(response.data);
        console.log("Product data:", response.data);

        const variantsData = await ProductService.getVariantsByProductId(id);
        console.log("Variants data:", variantsData);

        // Kiểm tra nếu variantsData không phải là null
        const variantsWithThumbnails = Array.isArray(variantsData)
          ? variantsData.map((variant) => ({
              ...variant,
              thumbnail: variant.thumbnail || variant.image,
            }))
          : []; // Nếu không có variant, gán mảng rỗng

        setVariants(variantsWithThumbnails);
        console.log("Variants with thumbnails:", variantsWithThumbnails);

        // Tạo mảng allImages bao gồm hình ảnh mặc định và thumbnail của các biến thể
        const allImagesArray = [
          {
            image: response.data.default_image,
            thumbnail: response.data.default_image,
          },
          ...variantsWithThumbnails.map((v) => ({
            image: v.thumbnail,
            thumbnail: v.thumbnail,
          })),
        ].filter((img) => img.image && img.thumbnail);

        setAllImages(allImagesArray);
        console.log("All images array:", allImagesArray);

        // Set default image
        if (variantsWithThumbnails.length > 0) {
          setSelectedVariantImage(variantsWithThumbnails[0].image);
          // Fetch related products
          const relatedProductsData =
            await ProductVariantService.getRelatedVariants(
              variantsWithThumbnails[0].id
            );
          setRelatedProducts(relatedProductsData);
          setSelectedVariantId(variantsWithThumbnails[0].variant_id);
        } else {
          // Nếu không có variant, sử dụng default_image
          setSelectedVariantImage(response.data.default_image);
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
              images={allImages}
              selectedImage={selectedVariantImage || product.default_image}
              onImageSelect={(image) => setSelectedVariantImage(image)}
            />
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <Box sx={{ pl: { md: 2 } }}>
            <ProductInfo
              product={product}
              variants={variants}
              onVariantSelect={(variant) => {
                setSelectedVariantId(variant.variant_id);
                setSelectedVariantImage(variant.image);
              }}
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
        <RecommendedProducts products={relatedProducts} />
      </Box>

      <Box mt={4}>
        <ViewedProducts products={[]} />
      </Box>

      <Box mt={4}>
        <SuggestedVariants selectedVariantId={selectedVariantId} />
      </Box>
    </Container>
  );
}

export default ProductDetailPage;
