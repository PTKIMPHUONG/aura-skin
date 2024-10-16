import React from "react";
import { Grid, Box, Typography } from "@mui/material";
import VariantCard from "../../components/Variants/Variant/VariantCard"; // Import VariantCard

const FavoriteProductsList = ({ favorites }) => {
  return (
    <Box sx={{ flexGrow: 1, padding: 2 }}>
      <Typography variant="h5" gutterBottom>
        Danh sách sản phẩm yêu thích
      </Typography>
      <Grid container spacing={2}>
        {favorites.map((variant) => (
          <Grid item xs={12} sm={6} md={4} lg={3} key={variant.variant_id}>
            <VariantCard variant={variant} /> {/* Sử dụng VariantCard */}
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default FavoriteProductsList;
