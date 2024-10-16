import api from "./api";

const ProductService = {
  getAllProducts: async () => {
    try {
      const response = await api.get("/products");
      return response; // Trả về toàn bộ response, không chỉ response.data
    } catch (error) {
      console.error("Error fetching all products:", error);
      throw error;
    }
  },

  getProductById: async (id) => {
    if (!id) {
      throw new Error("Product ID is undefined");
    }
    try {
      const response = await api.get(`/products/${id}`);
      return response.data;
    } catch (error) {
      console.error(`Error fetching product with id ${id}:`, error);
      throw error;
    }
  },

  getVariantsByProductId: async (productId) => {
    try {
      const response = await api.get(`/products/${productId}/product-variants`);
      return response.data.data; // Trả về mảng variants từ thuộc tính data
    } catch (error) {
      console.error(
        `Error fetching variants for product id ${productId}:`,
        error
      );
      throw error;
    }
  },

  getVariantsByProductName: async (productName) => {
    try {
      const response = await api.get(
        `/products/${productName}/product-variants`
      );
      return response.data;
    } catch (error) {
      console.error(
        `Error fetching variants for product name ${productName}:`,
        error
      );
      throw error;
    }
  },

  createProduct: async (productData) => {
    try {
      const response = await api.post("/products/create", productData);
      return response.data;
    } catch (error) {
      console.error("Error creating product:", error);
      throw error;
    }
  },

  updateProduct: async (id, productData) => {
    try {
      const response = await api.put(`/products/update/${id}`, productData);
      return response.data;
    } catch (error) {
      console.error(`Error updating product with id ${id}:`, error);
      throw error;
    }
  },

  deleteProduct: async (id) => {
    try {
      const response = await api.delete(`/products/delete/${id}`);
      return response.data;
    } catch (error) {
      console.error(`Error deleting product with id ${id}:`, error);
      throw error;
    }
  },

  getProductsByCategory: async (categoryId) => {
    try {
      const response = await api.get(`/categories/${categoryId}/products`);
      return response.data;
    } catch (error) {
      console.error("Error fetching products by category:", error);
      throw error;
    }
  },

  getFeaturedProducts: async () => {
    try {
      const response = await api.get("/products");
      // Có thể lọc hoặc lấy một số sản phẩm ��ầu tiên làm sản phẩm nổi bật
      return response.data.slice(0, 4);
    } catch (error) {
      console.error("Error fetching featured products:", error);
      throw error;
    }
  },

  getNewProducts: async () => {
    try {
      const response = await api.get("/products/new");
      return response.data;
    } catch (error) {
      console.error("Error fetching new products:", error);
      throw error;
    }
  },
};

export default ProductService;
