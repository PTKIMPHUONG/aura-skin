import React, { useState, useEffect } from "react";
import { Box, Typography, Paper, Button, Grid } from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import UserService from "../../services/UserService";
import { useAuth } from "../../context/Authcontext";

const UserAddress = () => {
  const [addresses, setAddresses] = useState([]);
  const { user } = useAuth();

  useEffect(() => {
    const fetchAddresses = async () => {
      try {
        const response = await UserService.getUserAddresses(user.id);
        setAddresses(response.data);
      } catch (error) {
        console.error("Error fetching addresses:", error);
      }
    };
    fetchAddresses();
  }, [user.id]);

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
              <Typography variant="subtitle1">
                {address.recipient_name}
              </Typography>
              <Typography variant="body2">{address.contact_number}</Typography>
              <Typography variant="body2">{`${address.address_line}, ${address.ward}, ${address.district}, ${address.province}, ${address.country}`}</Typography>
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
