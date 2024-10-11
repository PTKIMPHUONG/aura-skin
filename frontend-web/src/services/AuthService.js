import api from "./api";

const authService = {
  login: async (email, password) => {
    try {
      const response = await api.post("/user/login", { email, password });
      if (response.data && response.data.status === 200 && response.data.data) {
        const userData = {
          token: response.data.data.token,
          user: {
            id: response.data.data.data._id,
            username: response.data.data.data.username,
            email: response.data.data.data.email,
            phoneNumber: response.data.data.data.phone_number,
            isAdmin: response.data.data.data.is_admin,
            deliveryAddresses: response.data.data.data.delivery_addresses || [],
            imageUser: response.data.data.data.imageUser || null,
          },
        };
        authService.saveUserToStorage(userData);
        return { success: true, ...userData };
      }
      return {
        success: false,
        message: response.data.message || "Đăng nhập không thành công",
      };
    } catch (error) {
      console.error("Login error:", error);
      return {
        success: false,
        message: error.response?.data?.message || "Có lỗi xảy ra khi đăng nhập",
      };
    }
  },

  logout: () => {
    authService.removeUserFromStorage();
  },

  getCurrentUser: () => {
    const authData = localStorage.getItem("authData");
    return authData ? JSON.parse(authData) : null;
  },

  saveUserToStorage: (userData) => {
    localStorage.setItem("authData", JSON.stringify(userData));
  },

  removeUserFromStorage: () => {
    localStorage.removeItem("authData");
  },

  isAuthenticated: () => {
    return !!authService.getCurrentUser();
  },

  register: async (username, email, password, phone_number) => {
    try {
      const response = await api.post("/user/register", {
        username,
        email,
        password,
        phone_number,
      });
      if (response.data && response.data.status === 200) {
        return { success: true, message: "Đăng ký thành công" };
      }
      return {
        success: false,
        message: response.data.message || "Đăng ký không thành công",
      };
    } catch (error) {
      console.error("Registration error:", error);
      return {
        success: false,
        message: error.response?.data?.message || "Có lỗi xảy ra khi đăng ký",
      };
    }
  },
};

export default authService;
