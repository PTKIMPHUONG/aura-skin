import React from "react";
import { Typography, Box, Avatar, Rating } from "@mui/material";

function ProductReviews({ reviews }) {
  if (!reviews || reviews.length === 0) {
    return null;
  }
  return (
    <Box>
      <Typography variant="h5">Đánh giá sản phẩm</Typography>
      {reviews.map((review, index) => (
        <Box key={index} mt={2}>
          <Box display="flex" alignItems="center">
            <Avatar src={review.userAvatar} />
            <Box ml={2}>
              <Typography variant="subtitle1">{review.userName}</Typography>
              <Rating value={review.rating} readOnly size="small" />
            </Box>
          </Box>
          <Typography variant="body2" mt={1}>
            {review.comment}
          </Typography>
        </Box>
      ))}
    </Box>
  );
}

export default ProductReviews;
