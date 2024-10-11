import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from "@mui/material";
import ProductItem from "./ProductItem";

const ProductList = ({ products, onEditProduct, onDeleteProduct }) => {
  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Hình ảnh</TableCell>
            <TableCell>Tên sản phẩm</TableCell>
            <TableCell>Thương hiệu</TableCell>
            <TableCell>Giá</TableCell>
            <TableCell>Trạng thái</TableCell>
            <TableCell>Hành động</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {products.map((product) => (
            <ProductItem
              key={product.product_id}
              product={product}
              onEdit={onEditProduct}
              onDelete={onDeleteProduct}
            />
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default ProductList;
