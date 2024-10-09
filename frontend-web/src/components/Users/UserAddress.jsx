import React from "react";
import { Box, Typography, Paper, Button, Grid } from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";

const UserAddress = ({ addresses }) => {
  return (
    <Box sx={{ width: "100%" }}>
      <Typography variant="h5" gutterBottom>
        Quản lý địa chỉ
      </Typography>
      <Button variant="contained" color="primary" sx={{ mb: 2 }}>
        Thêm địa chỉ mới
      </Button>
      {addresses.map((address) => (
        <Paper key={address.id} sx={{ p: 2, mb: 2 }}>
          <Grid container spacing={2} alignItems="center">
            <Grid item xs={12} sm={8}>
              <Typography variant="subtitle1">{address.fullName}</Typography>
              <Typography variant="body2">{address.phone}</Typography>
              <Typography variant="body2">{address.address}</Typography>
              {address.isDefault && (
                <Typography variant="body2" color="primary">
                  Địa chỉ mặc định
                </Typography>
              )}
            </Grid>
            <Grid item xs={12} sm={4}>
              <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
                <Button startIcon={<EditIcon />} sx={{ mr: 1 }}>
                  Sửa
                </Button>
                <Button startIcon={<DeleteIcon />} color="error">
                  Xóa
                </Button>
              </Box>
            </Grid>
          </Grid>
        </Paper>
      ))}
    </Box>
  );
};

export default UserAddress;
