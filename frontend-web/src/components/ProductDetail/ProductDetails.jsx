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
        </TableBody>
      </Table>
      <Typography variant="body1" mt={2}>
        {product.description}
      </Typography>
    </Box>
  );
}

export default ProductDetails;
