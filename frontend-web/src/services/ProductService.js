import React from "react";
import ProductCard from "../components/Products/Product/ProductCard.js";
import "../styles/ProductList.css";

async function getProducts() {
  const response = await fetch("/api/products");
  const data = await response.json();
  return data;
}

export default getProducts;
