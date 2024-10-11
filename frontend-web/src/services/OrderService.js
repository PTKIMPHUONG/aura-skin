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
};

export default OrderService;
