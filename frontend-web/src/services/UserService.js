import api from "./api";

const UserService = {
  getUserAddresses: async (userId) => {
    try {
      const response = await api.get(`/user/${userId}/addresses`);
      return response.data;
    } catch (error) {
      throw error.response.data;
    }
  },

  addUserAddress: async (userId, addressData) => {
    try {
      const response = await api.post(`/user/${userId}/addresses`, addressData);
      return response.data;
    } catch (error) {
      throw error.response.data;
    }
  },

  updateUserAddress: async (userId, addressId, addressData) => {
    try {
      const response = await api.put(
        `/user/${userId}/addresses/${addressId}`,
        addressData
      );
      return response.data;
    } catch (error) {
      throw error.response.data;
    }
  },

  deleteUserAddress: async (userId, addressId) => {
    try {
      const response = await api.delete(
        `/user/${userId}/addresses/${addressId}`
      );
      return response.data;
    } catch (error) {
      throw error.response.data;
    }
  },

  uploadProfilePicture: async (userId, file) => {
    try {
      const formData = new FormData();
      formData.append("user_image", file);

      console.log("Uploading file:", file);
      console.log("FormData:", formData);

      const response = await api.post(
        `/user/upload-profile-picture/${userId}`,
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );

      console.log("Upload response:", response);
      return response.data;
    } catch (error) {
      console.error("Error uploading profile picture:", error);
      console.error("Error response:", error.response);
      throw error.response ? error.response.data : error;
    }
  },

  updateUserProfile: async (userId, userData) => {
    try {
      const response = await api.put(`/user/update`, userData);
      return response.data;
    } catch (error) {
      throw error.response.data;
    }
  },

  getUserWishlist: async (userId) => {
    try {
      const response = await api.get(`/user/${userId}/wishlist`);
      return response.data; // Giả sử API trả về danh sách yêu thích
    } catch (error) {
      throw error.response.data;
    }
  },
};

export default UserService;
