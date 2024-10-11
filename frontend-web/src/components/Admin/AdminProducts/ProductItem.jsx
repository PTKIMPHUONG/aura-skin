import React from "react";
import { TableCell, TableRow, Button, IconButton } from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";

const ProductItem = ({ product, onEdit, onDelete }) => {
  return (
    <TableRow>
      <TableCell>{product.product_id}</TableCell>
      <TableCell>
        <img
          src={product.default_image}
          alt={product.product_name}
          style={{ width: 50, height: 50, objectFit: "cover" }}
        />
      </TableCell>
      <TableCell>{product.product_name}</TableCell>
      <TableCell>{product.brand}</TableCell>
      <TableCell>{product.default_price.toLocaleString()} đ</TableCell>
      <TableCell>{product.is_active ? "Còn hàng" : "Hết hàng"}</TableCell>
      <TableCell>
        <IconButton color="primary" onClick={() => onEdit(product)}>
          <EditIcon />
        </IconButton>
        <IconButton color="error" onClick={() => onDelete(product.product_id)}>
          <DeleteIcon />
        </IconButton>
      </TableCell>
    </TableRow>
  );
};

export default ProductItem;
