import React, { useState } from "react";
import {
  Box,
  Typography,
  Grid,
  Button,
  IconButton,
  TextField,
  Pagination,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import AddIcon from "@mui/icons-material/Add";
import RemoveIcon from "@mui/icons-material/Remove";
import { useNavigate } from "react-router-dom";

const ProductCart = ({ cartItems, onQuantityChange, onRemoveItem }) => {
  const [page, setPage] = useState(1);
  const itemsPerPage = 5;
  const totalPages = Math.ceil(cartItems.length / itemsPerPage);

  const totalPrice = cartItems.reduce(
    (sum, item) => sum + item.price * item.quantity,
    0
  );

  const handlePageChange = (event, value) => {
    setPage(value);
  };

  const displayedItems = cartItems.slice(
    (page - 1) * itemsPerPage,
    page * itemsPerPage
  );

  const navigate = useNavigate();

  return (
    <Box>
      <Typography variant="h5" gutterBottom>
        Giỏ hàng
      </Typography>
      {cartItems.length === 0 ? (
        <Typography>Giỏ hàng của bạn đang trống.</Typography>
      ) : (
        <>
          {displayedItems.map((item) => (
            <Grid
              container
              key={item.productId}
              spacing={2}
              sx={{ mb: 2, alignItems: "center" }}
            >
              <Grid item xs={2}>
                <img
                  src={item.image}
                  alt={item.name}
                  style={{ width: "100%", height: "auto" }}
                />
              </Grid>
              <Grid item xs={4}>
                <Typography>{item.name}</Typography>
              </Grid>
              <Grid item xs={2}>
                <Typography>{item.price.toLocaleString()} đ</Typography>
              </Grid>
              <Grid item xs={3}>
                <Box sx={{ display: "flex", alignItems: "center" }}>
                  <IconButton
                    onClick={() =>
                      onQuantityChange(item.productId, item.quantity - 1)
                    }
                  >
                    <RemoveIcon />
                  </IconButton>
                  <TextField
                    value={item.quantity}
                    onChange={(e) =>
                      onQuantityChange(
                        item.productId,
                        parseInt(e.target.value) || 0
                      )
                    }
                    inputProps={{ min: 1, style: { textAlign: "center" } }}
                    sx={{ width: "50px" }}
                  />
                  <IconButton
                    onClick={() =>
                      onQuantityChange(item.productId, item.quantity + 1)
                    }
                  >
                    <AddIcon />
                  </IconButton>
                </Box>
              </Grid>
              <Grid item xs={1}>
                <IconButton onClick={() => onRemoveItem(item.productId)}>
                  <DeleteIcon />
                </IconButton>
              </Grid>
            </Grid>
          ))}
          <Box sx={{ display: "flex", justifyContent: "center", mt: 2, mb: 2 }}>
            <Pagination
              count={totalPages}
              page={page}
              onChange={handlePageChange}
            />
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "flex-end",
              alignItems: "center",
              mt: 2,
            }}
          >
            <Typography variant="h6" sx={{ mr: 2 }}>
              Tổng cộng: {totalPrice.toLocaleString()} đ
            </Typography>
            <Button
              variant="contained"
              color="primary"
              onClick={() =>
                navigate("/order-confirmation", {
                  state: { cartItems, totalPrice },
                })
              }
            >
              Mua Hàng
            </Button>
          </Box>
        </>
      )}
    </Box>
  );
};

export default ProductCart;
