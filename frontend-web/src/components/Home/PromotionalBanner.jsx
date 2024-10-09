import React from "react";
import { Box, Typography, Button } from "@mui/material";

const PromotionalBanner = () => {
  return (
    <Box
      sx={{ bgcolor: "error.main", color: "white", p: 4, textAlign: "center" }}
    >
      <Typography variant="h4" gutterBottom>
        Ưu đãi đặc biệt
      </Typography>
      <Typography variant="body1" gutterBottom>
        Giảm giá lên đến 50% cho tất cả sản phẩm
      </Typography>
      <Button variant="contained" color="secondary" sx={{ mt: 2 }}>
        Mua ngay
      </Button>
    </Box>
  );
};

export default PromotionalBanner;
