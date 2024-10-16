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
import AddressModal from "../Modals/AddressModal";
import UserService from "../../services/UserService";
import { useAuth } from "../../context/Authcontext";

const OrderConfirmation = ({
  cartItems,
  totalPrice,
  onSubmitOrder,
  userAddresses,
}) => {
  const [selectedAddressId, setSelectedAddressId] = useState(
    userAddresses[0]?.id || ""
  );
  const [openAddressModal, setOpenAddressModal] = useState(false);
  const { user } = useAuth();

  const handleAddressChange = (event) => {
    setSelectedAddressId(event.target.value);
  };

  const selectedAddress =
    userAddresses.find((address) => address.id === selectedAddressId) || {};

  const handleOpenAddressModal = () => {
    setOpenAddressModal(true);
  };

  const handleCloseAddressModal = () => {
    setOpenAddressModal(false);
  };

  const handleSaveNewAddress = async (newAddress) => {
    try {
      const response = await UserService.addUserAddress(user.id, newAddress);
      // Cập nhật danh sách địa chỉ sau khi thêm mới
      // Bạn có thể cần cập nhật state ở component cha hoặc gọi lại API để lấy danh sách địa chỉ mới
    } catch (error) {
      console.error("Error adding new address:", error);
    }
  };

  const shippingFee = 30000;
  const totalPayment = totalPrice + shippingFee;

  return (
    <Box sx={{ maxWidth: 1200, margin: "auto", mt: 4 }}>
      <Typography variant="h5" gutterBottom>
        Đặt hàng
      </Typography>

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
              {address.recipient_name} - {address.address_line}
            </MenuItem>
          ))}
        </Select>
        <Button
          variant="outlined"
          color="primary"
          onClick={handleOpenAddressModal}
          sx={{ mb: 2 }}
        >
          Thêm địa chỉ mới
        </Button>
        <Grid container spacing={2}>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Họ và tên"
              value={selectedAddress.recipient_name || ""}
              InputProps={{ readOnly: true }}
              variant="outlined"
            />
          </Grid>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Số điện thoại"
              value={selectedAddress.contact_number || ""}
              InputProps={{ readOnly: true }}
              variant="outlined"
            />
          </Grid>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Địa chỉ"
              value={
                `${selectedAddress.address_line}, ${selectedAddress.ward}, ${selectedAddress.district}, ${selectedAddress.province}, ${selectedAddress.country}` ||
                ""
              }
              InputProps={{ readOnly: true }}
              variant="outlined"
            />
          </Grid>
        </Grid>
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
                      src={item.thumbnail}
                      alt={item.variant_name}
                      style={{ width: 100, marginRight: 10 }}
                    />
                    <Box>
                      <Typography variant="body1">
                        {item.variant_name}
                      </Typography>
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

      <AddressModal
        open={openAddressModal}
        onClose={handleCloseAddressModal}
        onSave={handleSaveNewAddress}
      />
    </Box>
  );
};

export default OrderConfirmation;
