import api from "./api";

const OrderService = {
  getOrdersByUserId: async (userId) => {
    try {
      const response = await api.get(`/user/${userId}/order-history`);
      return response.data;
    } catch (error) {
      console.error("Error fetching user orders:", error);
      throw error;
    }
  },
  createOrder: async (orderData) => {
    try {
      const response = await api.post("/order/create", orderData);
      return response.data;
    } catch (error) {
      throw error.response.data;
    }
  },
};

export default OrderService;
