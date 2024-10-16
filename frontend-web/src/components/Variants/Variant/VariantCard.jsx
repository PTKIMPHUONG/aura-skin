import React from "react";
import {
  Card,
  CardMedia,
  CardContent,
  Typography,
  CardActionArea,
  Box,
} from "@mui/material";
import { Link } from "react-router-dom";

const VariantCard = ({ variant }) => {
  return (
    <Card sx={{ height: "100%", display: "flex", flexDirection: "column" }}>
      <CardActionArea
        component={Link}
        to={`/product/${variant.product_id}/variant/${variant.variant_id}`}
        sx={{ flexGrow: 1, display: "flex", flexDirection: "column" }}
      >
        <CardMedia
          component="img"
          height="200"
          image={variant.thumbnail || "/placeholder-image.jpg"}
          alt={variant.variant_name}
        />
        <CardContent
          sx={{
            flexGrow: 1,
            display: "flex",
            flexDirection: "column",
            justifyContent: "space-between",
          }}
        >
          <Box>
            <Typography
              gutterBottom
              variant="subtitle1"
              component="div"
              sx={{
                overflow: "hidden",
                textOverflow: "ellipsis",
                display: "-webkit-box",
                WebkitLineClamp: 2,
                WebkitBoxOrient: "vertical",
                lineHeight: "1.2em",
                height: "2.4em",
              }}
            >
              {variant.variant_name}
            </Typography>
          </Box>
          <Typography variant="body2" color="text.secondary">
            {variant.price
              ? `${variant.price.toLocaleString("vi-VN")} ₫`
              : "Giá không có sẵn"}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
};

export default VariantCard;
