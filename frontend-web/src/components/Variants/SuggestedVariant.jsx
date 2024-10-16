import React, { useState, useEffect } from "react";
import { Box, Typography, Grid } from "@mui/material";
import VariantCard from "../Variants/Variant/VariantCard";
import ProductVariantService from "../../services/ProductVariantService";

const SuggestedVariants = ({ selectedVariantId }) => {
  const [suggestedVariants, setSuggestedVariants] = useState([]);

  useEffect(() => {
    const fetchSuggestedVariants = async () => {
      if (selectedVariantId) {
        console.log("Fetching suggested variants for:", selectedVariantId);
        try {
          const response = await ProductVariantService.getRelatedVariants(
            selectedVariantId
          );
          console.log("API response:", response);
          setSuggestedVariants(response.data || []);
        } catch (error) {
          console.error("Error fetching suggested variants:", error);
        }
      }
    };

    fetchSuggestedVariants();
  }, [selectedVariantId]);

  console.log("Suggested variants:", suggestedVariants);

  if (suggestedVariants.length === 0) {
    console.log("No suggested variants to display");
    return null;
  }

  return (
    <Box mt={4}>
      <Typography variant="h5" gutterBottom>
        Sản phẩm tương tự
      </Typography>
      <Grid container spacing={2}>
        {suggestedVariants.map((variant) => (
          <Grid item xs={6} sm={4} md={3} key={variant.variant_id}>
            <VariantCard variant={variant} />
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default SuggestedVariants;
