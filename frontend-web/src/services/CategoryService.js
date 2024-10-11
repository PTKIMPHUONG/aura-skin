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
};

export default CategoryService;
