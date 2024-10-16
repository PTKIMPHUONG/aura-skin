import React, { useState, useCallback } from "react";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  Grid,
  Autocomplete,
  CircularProgress,
  Typography,
} from "@mui/material";
import debounce from "lodash/debounce";

const AddressModal = ({ open, onClose, onSave }) => {
  const [address, setAddress] = useState({
    recipient_name: "",
    contact_number: "",
    address_line: "",
    ward: "",
    district: "",
    province: "",
    country: "",
  });
  const [suggestions, setSuggestions] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleChange = (e) => {
    setAddress({ ...address, [e.target.name]: e.target.value });
  };

  const handleAddressSearch = useCallback(
    debounce(async (value) => {
      if (value.length > 3) {
        setLoading(true);
        setError("");
        try {
          const response = await fetch(
            `https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(
              value
            )}`
          );
          const data = await response.json();
          if (data.length === 0) {
            setError("Không tìm thấy địa chỉ phù hợp");
          }
          setSuggestions(
            data.map((item) => ({
              label: item.display_name,
              value: item,
            }))
          );
        } catch (error) {
          console.error("Error fetching address suggestions:", error);
          setError("Có lỗi xảy ra khi tìm kiếm địa chỉ");
        } finally {
          setLoading(false);
        }
      }
    }, 300),
    []
  );

  const handleAddressSelect = (event, newValue) => {
    if (newValue) {
      setAddress({
        ...address,
        address_line: newValue.value.display_name,
        ward: newValue.value.address.suburb || "",
        district:
          newValue.value.address.city_district ||
          newValue.value.address.town ||
          "",
        province: newValue.value.address.state || "",
        country: newValue.value.address.country || "",
      });
    }
  };

  const handleSave = () => {
    onSave(address);
    onClose();
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>Thêm địa chỉ mới</DialogTitle>
      <DialogContent>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Họ và tên người nhận"
              name="recipient_name"
              value={address.recipient_name}
              onChange={handleChange}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Số điện thoại"
              name="contact_number"
              value={address.contact_number}
              onChange={handleChange}
            />
          </Grid>
          <Grid item xs={12}>
            <Autocomplete
              freeSolo
              options={suggestions}
              onInputChange={(event, newValue) => handleAddressSearch(newValue)}
              onChange={handleAddressSelect}
              renderInput={(params) => (
                <TextField
                  {...params}
                  label="Tìm kiếm địa chỉ"
                  fullWidth
                  InputProps={{
                    ...params.InputProps,
                    endAdornment: (
                      <>
                        {loading ? (
                          <CircularProgress color="inherit" size={20} />
                        ) : null}
                        {params.InputProps.endAdornment}
                      </>
                    ),
                  }}
                />
              )}
            />
            {error && (
              <Typography color="error" variant="caption">
                {error}
              </Typography>
            )}
          </Grid>
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Địa chỉ cụ thể"
              name="address_line"
              value={address.address_line}
              onChange={handleChange}
            />
          </Grid>
          <Grid item xs={6}>
            <TextField
              fullWidth
              label="Phường/Xã"
              name="ward"
              value={address.ward}
              onChange={handleChange}
            />
          </Grid>
          <Grid item xs={6}>
            <TextField
              fullWidth
              label="Quận/Huyện"
              name="district"
              value={address.district}
              onChange={handleChange}
            />
          </Grid>
          <Grid item xs={6}>
            <TextField
              fullWidth
              label="Tỉnh/Thành phố"
              name="province"
              value={address.province}
              onChange={handleChange}
            />
          </Grid>
          <Grid item xs={6}>
            <TextField
              fullWidth
              label="Quốc gia"
              name="country"
              value={address.country}
              onChange={handleChange}
            />
          </Grid>
        </Grid>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Hủy</Button>
        <Button onClick={handleSave} variant="contained" color="primary">
          Lưu
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default AddressModal;
