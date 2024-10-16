import api from "./api";

const ProductVariantService = {
  getSuggestVariantsForHomepage: async (userId) => {
    try {
      const response = await api.get(
        `/product-variants/suggest/user/${userId}`
      );
      return response.data;
    } catch (error) {
      console.error("API error:", error);
      throw error;
    }
  },

  getRelatedVariants: async (variantId) => {
    try {
      const response = await api.get(
        `/product-variants/suggest/variant/${variantId}`
      );
      return response.data;
    } catch (error) {
      console.error("Error fetching related variants:", error);
      throw error;
    }
  },

  getSuggestVariantsForUser: async (userId) => {
    try {
      const response = await api.get(
        `/product-variants/suggest/user/${userId}`
      );
      return response.data;
    } catch (error) {
      console.error("Error fetching suggested variants for user:", error);
      throw error;
    }
  },

  getProductByVariantId: async (variantId) => {
    try {
      const response = await api.get(`/products/variant/${variantId}`);
      return response.data;
    } catch (error) {
      console.error("Error fetching product by variant ID:", error);
      throw error;
    }
  },
};

export default ProductVariantService;
