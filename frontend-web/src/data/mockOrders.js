const mockOrders = [
  {
    userId: 1,
    orders: [
      {
        id: 1,
        date: "2023-05-01",
        status: "Chờ xác nhận",
        items: [
          {
            productId: 1,
            name: "Son Merzy V6 Blue Dream",
            price: 169000,
            quantity: 1,
            image: require("../assets/images/clio_copy_fef881b3ab6a4e2eba206f50726cbbd9_1024x1024.webp"),
          },
        ],
        total: 169000,
      },
      {
        id: 2,
        date: "2023-04-28",
        status: "Đã giao",
        items: [
          {
            productId: 2,
            name: "Son Merzy V6 Red Passion",
            price: 179000,
            quantity: 2,
            image: require("../assets/images/clio_copy_fef881b3ab6a4e2eba206f50726cbbd9_1024x1024.webp"),
          },
        ],
        total: 358000,
      },
    ],
  },
];

export default mockOrders;
