import React from "react";
import {
  Typography,
  Box,
  Table,
  TableBody,
  TableCell,
  TableRow,
} from "@mui/material";

function ProductDetails({ product }) {
  if (!product) {
    return <Typography>Đang tải thông tin sản phẩm...</Typography>;
  }

  return (
    <Box>
      <Typography variant="h6" fontWeight="bold" mb={2}>
        Chi tiết sản phẩm
      </Typography>
      <Table>
        <TableBody>
          <TableRow>
            <TableCell component="th" scope="row">
              Tên sản phẩm
            </TableCell>
            <TableCell>{product.product_name}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Dung tích
            </TableCell>
            <TableCell>{product.capacity}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Xuất xứ thương hiệu
            </TableCell>
            <TableCell>{product.origin}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Sản xuất tại
            </TableCell>
            <TableCell>{product.manufactured_in}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Hạn sử dụng
            </TableCell>
            <TableCell>{product.expiration_date}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Thành phần
            </TableCell>
            <TableCell>{product.ingredients}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Cách sử dụng
            </TableCell>
            <TableCell>{product.usage}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Bảo quản
            </TableCell>
            <TableCell>{product.storage}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Đối tượng sử dụng
            </TableCell>
            <TableCell>{product.target_customers}</TableCell>
          </TableRow>
        </TableBody>
      </Table>
      <Typography variant="h6" fontWeight="bold" mt={4} mb={2}>
        Mô tả sản phẩm
      </Typography>
      <Typography variant="body1">{product.description}</Typography>
      <Typography variant="h6" fontWeight="bold" mt={4} mb={2}>
        Đặc tính nổi bật
      </Typography>
      <Typography variant="body1">{product.features}</Typography>
    </Box>
  );
}

export default ProductDetails;
