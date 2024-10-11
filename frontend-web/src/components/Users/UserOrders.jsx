import React, { useState } from "react";
import { Box, Tabs, Tab, Typography, Grid, Paper } from "@mui/material";

const UserOrders = ({ orders }) => {
  const [value, setValue] = useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  const filterOrders = (status) => {
    if (status === "Tất cả") return orders;
    return orders.filter((order) => order.status === status);
  };

  const statuses = [
    "Tất cả",
    "Chờ xác nhận",
    "Chờ lấy hàng",
    "Đang giao",
    "Đã giao",
    "Đã hủy",
  ];

  return (
    <Box sx={{ width: "100%" }}>
      <Typography variant="h5" gutterBottom>
        Quản lý đơn hàng
      </Typography>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={value}
          onChange={handleChange}
          aria-label="order status tabs"
        >
          {statuses.map((status, index) => (
            <Tab label={status} key={index} />
          ))}
        </Tabs>
      </Box>
      {statuses.map((status, index) => (
        <TabPanel value={value} index={index} key={index}>
          {filterOrders(status).map((order) => (
            <Paper key={order.order_id} sx={{ p: 2, mb: 2 }}>
              <Grid container spacing={2}>
                <Grid item xs={12} sm={6}>
                  <Typography variant="subtitle1">
                    Đơn hàng #{order.order_id}
                  </Typography>
                  <Typography variant="body2">
                    Ngày đặt: {new Date(order.created_at).toLocaleDateString()}
                  </Typography>
                </Grid>
                <Grid item xs={12} sm={6}>
                  <Typography variant="subtitle1" align="right">
                    Trạng thái: {order.status}
                  </Typography>
                  <Typography variant="body2" align="right">
                    Tổng tiền: {order.total_amount.toLocaleString()} đ
                  </Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography variant="body2">
                    Người nhận: {order.recipient_name}
                  </Typography>
                  <Typography variant="body2">
                    Số điện thoại: {order.contact_number}
                  </Typography>
                  <Typography variant="body2">
                    Địa chỉ: {order.address_line}, {order.ward},{" "}
                    {order.district}, {order.province}, {order.country}
                  </Typography>
                  <Typography variant="body2">
                    Phí vận chuyển: {order.delivery_fee.toLocaleString()} đ
                  </Typography>
                </Grid>
              </Grid>
            </Paper>
          ))}
        </TabPanel>
      ))}
    </Box>
  );
};

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
}

export default UserOrders;
