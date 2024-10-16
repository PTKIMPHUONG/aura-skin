import React, { useState, useEffect } from "react";
import {
  Box,
  TextField,
  Button,
  Grid,
  Avatar,
  Input,
  CircularProgress,
} from "@mui/material";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import UserService from "../../services/UserService";
import { useAuth } from "../../context/Authcontext";

const defaultAvatar = require("../../assets/images/defaultImageUser.png");

const UserProfileForm = ({ profileData, onUpdateProfile }) => {
  const [formData, setFormData] = useState(profileData);
  const [selectedFile, setSelectedFile] = useState(null);
  const [previewImage, setPreviewImage] = useState(null);
  const [isUploading, setIsUploading] = useState(false);
  const [isUpdating, setIsUpdating] = useState(false);
  const { user, updateUser } = useAuth();

  useEffect(() => {
    if (selectedFile) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setPreviewImage(reader.result);
      };
      reader.readAsDataURL(selectedFile);
    } else {
      setPreviewImage(null);
    }
  }, [selectedFile]);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      setSelectedFile(file);
    }
  };

  const handleFileUpload = async () => {
    if (selectedFile && user) {
      setIsUploading(true);
      try {
        console.log("Selected file:", selectedFile);
        const response = await UserService.uploadProfilePicture(
          user.id,
          selectedFile
        );
        console.log("Upload response:", response);
        if (response.status === 200) {
          setFormData((prevData) => ({
            ...prevData,
            user_image: response.data.user_image,
          }));
          updateUser({
            ...user,
            user_image: response.data.user_image,
          });
          setPreviewImage(null);
          setSelectedFile(null);
          alert("Cập nhật ảnh đại diện thành công");
        } else {
          throw new Error(response.message || "Có lỗi xảy ra khi tải ảnh lên");
        }
      } catch (error) {
        console.error("Error in handleFileUpload:", error);
        let errorMessage = "Lỗi khi tải ảnh lên: ";
        if (error.response && error.response.data) {
          errorMessage += `${
            error.response.data.message || "Lỗi server không xác định"
          }. Vui lòng thử lại sau hoặc liên hệ hỗ trợ.`;
        } else if (error.request) {
          errorMessage +=
            "Không nhận được phản hồi từ server. Vui lòng kiểm tra kết nối mạng và thử lại.";
        } else {
          errorMessage += "Đã xảy ra lỗi không xác định. Vui lòng thử lại sau.";
        }
        alert(errorMessage);
      } finally {
        setIsUploading(false);
      }
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsUpdating(true);
    try {
      const response = await UserService.updateUserProfile(user.id, formData);
      if (response.success) {
        onUpdateProfile(response.data);
        updateUser(response.data);
      }
    } catch (error) {
      console.error("Error updating user profile:", error);
      // Hiển thị thông báo lỗi
    } finally {
      setIsUpdating(false);
    }
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
            disabled={isUpdating}
          >
            {isUpdating ? <CircularProgress size={24} /> : "Cập nhật thông tin"}
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
            src={previewImage || formData.user_image || defaultAvatar}
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
              Chọn ảnh mới
            </Button>
          </label>
          {selectedFile && (
            <Button
              onClick={handleFileUpload}
              variant="outlined"
              sx={{ mt: 2 }}
              disabled={isUploading}
            >
              {isUploading ? (
                <CircularProgress size={24} />
              ) : (
                "Cập nhật ảnh đại diện"
              )}
            </Button>
          )}
        </Grid>
      </Grid>
    </Box>
  );
};

export default UserProfileForm;
