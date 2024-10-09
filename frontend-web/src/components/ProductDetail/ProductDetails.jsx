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
            <TableCell>{product.name}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Chất son
            </TableCell>
            <TableCell>{product.texture}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Dung tích
            </TableCell>
            <TableCell>{product.volume}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Thương hiệu
            </TableCell>
            <TableCell>{product.brand}</TableCell>
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
            <TableCell>{product.madeIn}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell component="th" scope="row">
              Hạn sử dụng
            </TableCell>
            <TableCell>{product.expirationDate}</TableCell>
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
