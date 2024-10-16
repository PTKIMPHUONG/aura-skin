import React, { useState, useEffect } from "react";
import { Box, IconButton } from "@mui/material";
import ArrowBackIosNewIcon from "@mui/icons-material/ArrowBackIosNew";
import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos";

function ProductImageGallery({ images, selectedImage, onImageSelect }) {
  const [startIndex, setStartIndex] = useState(0);
  const visibleImages = 4; // Số lượng hình ảnh hiển thị cùng lúc

  useEffect(() => {
    if (selectedImage && images.some((img) => img.image === selectedImage)) {
      onImageSelect(selectedImage);
    }
  }, [selectedImage, images, onImageSelect]);

  if (!images || images.length === 0) {
    return null;
  }

  const handleScroll = (direction) => {
    if (direction === "left" && startIndex > 0) {
      setStartIndex(startIndex - 1);
    } else if (
      direction === "right" &&
      startIndex < images.length - visibleImages
    ) {
      setStartIndex(startIndex + 1);
    }
  };

  return (
    <Box>
      <Box mb={2} sx={{ maxWidth: "100%", margin: "0 auto" }}>
        <img
          src={selectedImage || (images[0] && images[0].image)}
          alt="Main product"
          style={{ width: "100%", height: "auto", objectFit: "contain" }}
        />
      </Box>
      <Box sx={{ position: "relative", display: "flex", alignItems: "center" }}>
        <IconButton
          onClick={() => handleScroll("left")}
          sx={{ position: "absolute", left: 0, zIndex: 1 }}
          disabled={startIndex === 0}
        >
          <ArrowBackIosNewIcon />
        </IconButton>
        <Box sx={{ display: "flex", overflow: "hidden", margin: "0 30px" }}>
          {images
            .slice(startIndex, startIndex + visibleImages)
            .map((image, index) => (
              <Box
                key={startIndex + index}
                sx={{
                  width: "25%",
                  flexShrink: 0,
                  padding: "0 4px",
                }}
              >
                <img
                  src={image.thumbnail}
                  alt={`Product view ${startIndex + index + 1}`}
                  style={{
                    width: "100%",
                    height: "auto",
                    cursor: "pointer",
                    objectFit: "cover",
                    border:
                      image.image === selectedImage
                        ? "2px solid #1976d2"
                        : "none",
                  }}
                  onClick={() => onImageSelect(image.image)}
                />
              </Box>
            ))}
        </Box>
        <IconButton
          onClick={() => handleScroll("right")}
          sx={{ position: "absolute", right: 0, zIndex: 1 }}
          disabled={startIndex >= images.length - visibleImages}
        >
          <ArrowForwardIosIcon />
        </IconButton>
      </Box>
    </Box>
  );
}

export default ProductImageGallery;
