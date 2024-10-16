import React from "react";
import { Box, Typography, Grid } from "@mui/material";
import VariantCard from "./Variant/VariantCard";

const HomeSuggestList = ({ title, variants }) => {
  return (
    <Box>
      <Typography variant="h4" gutterBottom sx={{ textAlign: "center" }}>
        {title}
      </Typography>
      <Grid container spacing={2}>
        {variants.map((variant) => (
          <Grid item xs={6} sm={3} key={variant.variant_id}>
            <VariantCard variant={variant} />
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default HomeSuggestList;
