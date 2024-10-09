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

  const handleFileUpload = () => {
    if (selectedFile) {
      // Xử lý tải lên file ở đây
      console.log("Tải lên file:", selectedFile);
      // Sau khi tải lên thành công, cập nhật URL ảnh mới
      // setFormData(prevData => ({
      //   ...prevData,
      //   imageUser: newImageUrl
      // }));
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
            id="fullName"
            label="Họ tên"
            name="fullName"
            autoComplete="name"
            value={formData.fullName}
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
            value={formData.email}
            onChange={handleChange}
          />
          <RadioGroup
            row
            aria-label="gender"
            name="gender"
            value={formData.gender}
            onChange={handleChange}
          >
            <FormControlLabel value="Nam" control={<Radio />} label="Nam" />
            <FormControlLabel value="Nữ" control={<Radio />} label="Nữ" />
          </RadioGroup>
          <TextField
            margin="normal"
            fullWidth
            id="birthDate"
            label="Ngày/tháng/năm sinh"
            name="birthDate"
            type="date"
            value={formData.birthDate}
            onChange={handleChange}
            InputLabelProps={{
              shrink: true,
            }}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="phoneNumber"
            label="Số điện thoại"
            name="phoneNumber"
            autoComplete="tel"
            value={formData.phoneNumber}
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
            src={formData.imageUser}
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
