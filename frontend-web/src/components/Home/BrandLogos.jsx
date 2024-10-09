import React from "react";
import { Grid, Box, Typography, Card, CardActionArea } from "@mui/material";
import { Link } from "react-router-dom";
import mockBrands from "../../data/mockBrands"; // hoặc "../../data/mockBrands"

const BrandLogos = () => {
  return (
    <Box sx={{ my: 4 }}>
      <Typography variant="h5" sx={{ mt: 10, mb: 3, textAlign: "center" }}>
        THƯƠNG HIỆU
      </Typography>
      <Grid container spacing={3} justifyContent="center" alignItems="center">
        {mockBrands.map((brand) => (
          <Grid item xs={6} sm={3} key={brand.id}>
            <Card
              elevation={2}
              sx={{
                height: "100%",
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                transition: "0.3s",
                "&:hover": {
                  transform: "scale(1.05)",
                },
              }}
            >
              <CardActionArea
                component={Link}
                to={`/brand/${brand.id}`}
                sx={{
                  height: "100%",
                  display: "flex",
                  flexDirection: "column",
                  justifyContent: "center",
                  alignItems: "center",
                  padding: 2,
                }}
              >
                <Box
                  component="img"
                  src={brand.logo}
                  alt={brand.name}
                  sx={{
                    maxWidth: "80%",
                    maxHeight: "80px",
                    objectFit: "contain",
                  }}
                  onError={(e) => {
                    e.target.onerror = null;
                    e.target.src = "/path/to/fallback/image.png"; // Đường dẫn đến hình ảnh mặc định
                  }}
                />
                <Typography
                  variant="subtitle1"
                  sx={{ mt: 2, textAlign: "center" }}
                >
                  {brand.name}
                </Typography>
              </CardActionArea>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default BrandLogos;
