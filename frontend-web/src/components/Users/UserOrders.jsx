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
            <Paper key={order.id} sx={{ p: 2, mb: 2 }}>
              <Grid container spacing={2}>
                <Grid item xs={12} sm={6}>
                  <Typography variant="subtitle1">
                    Đơn hàng #{order.id}
                  </Typography>
                  <Typography variant="body2">
                    Ngày đặt: {order.date}
                  </Typography>
                </Grid>
                <Grid item xs={12} sm={6}>
                  <Typography variant="subtitle1" align="right">
                    Trạng thái: {order.status}
                  </Typography>
                  <Typography variant="body2" align="right">
                    Tổng tiền: {order.total.toLocaleString()} đ
                  </Typography>
                </Grid>
                {order.items.map((item) => (
                  <Grid item xs={12} key={item.productId}>
                    <Box sx={{ display: "flex", alignItems: "center" }}>
                      <img
                        src={item.image}
                        alt={item.name}
                        style={{ width: 50, height: 50, marginRight: 10 }}
                      />
                      <Box>
                        <Typography variant="body1">{item.name}</Typography>
                        <Typography variant="body2">
                          Số lượng: {item.quantity}
                        </Typography>
                      </Box>
                    </Box>
                  </Grid>
                ))}
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
