import React from "react";
import { Breadcrumbs, Link, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";

function BreadcrumbNavigation({ product }) {
  return (
    <Breadcrumbs
      aria-label="breadcrumb"
      sx={{ position: "relative", top: -50 }}
    >
      <Link component={RouterLink} to="/" color="inherit">
        Trang chủ
      </Link>
      <Link component={RouterLink} to="/products" color="inherit">
        Sản phẩm
      </Link>
      <Typography color="text.primary">{product.product_name}</Typography>
    </Breadcrumbs>
  );
}

export default BreadcrumbNavigation;
