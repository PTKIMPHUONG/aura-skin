import React, { useState } from "react";
import {
  Box,
  TextField,
  Button,
  Grid,
  Avatar,
  Radio,
  RadioGroup,
  FormControlLabel,
  Input,
} from "@mui/material";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import authService from "../../services/AuthService";

const defaultAvatar = require("../../assets/images/defaultImageUser.png"); // Thêm đường dẫn đến ảnh mặc định

const UserProfileForm = ({ profileData, onUpdateProfile }) => {
  const [formData, setFormData] = useState(profileData);
  const [selectedFile, setSelectedFile] = useState(null);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
  };

  const handleFileUpload = async () => {
    if (selectedFile) {
      try {
        // Giả sử AuthService có phương thức uploadProfilePicture
        const response = await authService.uploadProfilePicture(selectedFile);
        if (response.success) {
          setFormData((prevData) => ({
            ...prevData,
            imageUser: response.data.imageUrl,
          }));
        }
      } catch (error) {
        console.error("Error uploading profile picture:", error);
        // Hiển thị thông báo lỗi
      }
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    onUpdateProfile(formData);
  };

  return (
    <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
      <Grid container spacing={2}>
        <Grid item xs={12} md={8}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Tên người dùng"
            name="username"
            autoComplete="username"
            value={formData.username || ""}
            onChange={handleChange}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email"
            name="email"
            autoComplete="email"
            value={formData.email || ""}
            onChange={handleChange}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="phoneNumber"
            label="Số điện thoại"
            name="phoneNumber"
            autoComplete="tel"
            value={formData.phoneNumber || ""}
            onChange={handleChange}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Cập nhật thông tin
          </Button>
        </Grid>
        <Grid
          item
          xs={12}
          md={4}
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar
            src={formData.imageUser || defaultAvatar}
            sx={{ width: 150, height: 150, mb: 2 }}
          />
          <Input
            accept="image/*"
            id="contained-button-file"
            type="file"
            onChange={handleFileChange}
            style={{ display: "none" }}
          />
          <label htmlFor="contained-button-file">
            <Button
              variant="contained"
              component="span"
              startIcon={<CloudUploadIcon />}
              sx={{ mt: 2 }}
            >
              Tải ảnh mới
            </Button>
          </label>
          {selectedFile && (
            <Button
              onClick={handleFileUpload}
              variant="outlined"
              sx={{ mt: 2 }}
            >
              Cập nhật ảnh đại diện
            </Button>
          )}
        </Grid>
      </Grid>
    </Box>
  );
};

export default UserProfileForm;
