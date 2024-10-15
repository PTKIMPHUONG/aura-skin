import api from "./api";

const CategoryService = {
  getAllCategories: async () => {
    try {
      const response = await api.get("/categories");
      return response.data;
    } catch (error) {
      console.error("Error fetching all categories:", error);
      throw error;
    }
  },
  getCategoryById: async (categoryId) => {
    try {
      const response = await api.get(`/categories/${categoryId}`);
      return response.data;
    } catch (error) {
      console.error("Error fetching category by ID:", error);
      throw error;
    }
  },
};

export default CategoryService;
