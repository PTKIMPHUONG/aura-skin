import React, { useState } from "react";
import {
  Box,
  Typography,
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Select,
  MenuItem,
  Grid,
  TextField,
} from "@mui/material";

const OrderConfirmation = ({
  cartItems,
  totalPrice,
  onSubmitOrder,
  userAddresses,
}) => {
  const [selectedAddressId, setSelectedAddressId] = useState(
    userAddresses[0]?.id || ""
  );

  const handleAddressChange = (event) => {
    setSelectedAddressId(event.target.value);
  };

  const selectedAddress =
    userAddresses.find((address) => address.id === selectedAddressId) || {};

  const shippingFee = 30000;
  const totalPayment = totalPrice + shippingFee;

  return (
    <Box sx={{ maxWidth: 1200, margin: "auto", mt: 4 }}>
      <Typography variant="h5" gutterBottom>
        Đặt hàng
      </Typography>

      {/* Địa chỉ nhận hàng */}
      <Paper sx={{ p: 2, mb: 3, bgcolor: "#f3f4f6" }}>
        <Typography variant="subtitle1" gutterBottom>
          Địa chỉ nhận hàng
        </Typography>
        <Select
          value={selectedAddressId}
          onChange={handleAddressChange}
          fullWidth
          sx={{ mb: 2 }}
        >
          {userAddresses.map((address) => (
            <MenuItem key={address.id} value={address.id}>
              {address.fullName} - {address.address}
            </MenuItem>
          ))}
        </Select>
        <Grid container spacing={2}>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Họ và tên"
              value={selectedAddress.fullName || ""}
              InputProps={{ readOnly: true }}
              variant="outlined"
            />
          </Grid>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Số điện thoại"
              value={selectedAddress.phone || ""}
              InputProps={{ readOnly: true }}
              variant="outlined"
            />
          </Grid>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Địa chỉ"
              value={selectedAddress.address || ""}
              InputProps={{ readOnly: true }}
              variant="outlined"
            />
          </Grid>
        </Grid>
        <Box sx={{ mt: 1, textAlign: "right" }}>
          <Button href="#" color="primary">
            Sửa
          </Button>
          <Button href="#" color="primary">
            Xóa
          </Button>
        </Box>
      </Paper>

      {/* Danh sách sản phẩm */}
      <TableContainer component={Paper} sx={{ mb: 3 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>
                <Typography
                  variant="subtitle1"
                  sx={{ position: "relative", left: "5%" }}
                >
                  Sản phẩm
                </Typography>
              </TableCell>
              <TableCell align="center">
                <Typography variant="subtitle1">Đơn giá</Typography>
              </TableCell>
              <TableCell align="center">
                <Typography variant="subtitle1">Số lượng</Typography>
              </TableCell>
              <TableCell align="center">
                <Typography variant="subtitle1">Thành tiền</Typography>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {cartItems.map((item) => (
              <TableRow key={item.productId}>
                <TableCell>
                  <Box sx={{ display: "flex", alignItems: "center" }}>
                    <img
                      src={item.image}
                      alt={item.name}
                      style={{ width: 100, marginRight: 10 }}
                    />
                    <Box>
                      <Typography variant="body1">{item.name}</Typography>
                      <Typography variant="body2" color="text.secondary">
                        Phân loại: {item.category}
                      </Typography>
                    </Box>
                  </Box>
                </TableCell>
                <TableCell align="center">
                  <Typography variant="body1">
                    {item.price.toLocaleString()}đ
                  </Typography>
                </TableCell>
                <TableCell align="center">
                  <Typography variant="body1">{item.quantity}</Typography>
                </TableCell>
                <TableCell align="center">
                  <Typography variant="body1">
                    {(item.price * item.quantity).toLocaleString()}đ
                  </Typography>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {/* Phương thức thanh toán và tổng cộng */}
      <Paper sx={{ p: 2, mb: 3 }}>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            mb: 2,
          }}
        >
          <Typography variant="h6">Phương thức thanh toán</Typography>
          <Typography variant="body1">Thanh toán khi nhận hàng</Typography>
        </Box>
        <Box sx={{ borderTop: "1px solid #e0e0e0", pt: 2 }}>
          <Box sx={{ display: "flex", justifyContent: "space-between", mb: 1 }}>
            <Typography variant="body1">Tổng tiền hàng</Typography>
            <Typography variant="body1">
              {totalPrice.toLocaleString()}đ
            </Typography>
          </Box>
          <Box sx={{ display: "flex", justifyContent: "space-between", mb: 1 }}>
            <Typography variant="body1">Phí vận chuyển</Typography>
            <Typography variant="body1">
              {shippingFee.toLocaleString()}đ
            </Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              mt: 1,
              pt: 1,
              borderTop: "1px solid #e0e0e0",
            }}
          >
            <Typography variant="h6">Tổng thanh toán</Typography>
            <Typography variant="h6" sx={{ fontWeight: "bold" }}>
              {totalPayment.toLocaleString()}đ
            </Typography>
          </Box>
        </Box>
      </Paper>

      <Box sx={{ textAlign: "right" }}>
        <Button
          variant="contained"
          color="primary"
          size="large"
          onClick={() =>
            onSubmitOrder({
              shippingInfo: selectedAddress,
              cartItems,
              totalPayment,
            })
          }
        >
          Đặt hàng
        </Button>
      </Box>
    </Box>
  );
};

export default OrderConfirmation;
