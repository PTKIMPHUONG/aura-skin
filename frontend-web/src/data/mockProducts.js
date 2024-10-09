// src/data/mockProducts.js
const mockProducts = Array(32)
  .fill()
  .map((_, index) => ({
    id: index + 1,
    name: `Son Merzy V6 Blue Dream ${index + 1}`,
    price: 169000,
    originalPrice: 299000,
    description:
      "Son môi đẹp và bền màu, giữ màu lâu trôi, không gây khô môi. Phù hợp với mọi loại da.",
    urlImage: require("../assets/images/clio_copy_fef881b3ab6a4e2eba206f50726cbbd9_1024x1024.webp"),
    images: [
      require("../assets/images/clio_copy_fef881b3ab6a4e2eba206f50726cbbd9_1024x1024.webp"),
      require("../assets/images/m11_6d8271cbcf1949a4a77d2cfdeeed3403_1024x1024.webp"),
      require("../assets/images/home_category_5_medium.webp"),
      require("../assets/images/home_category_11_medium.webp"),
      require("../assets/images/m11_6d8271cbcf1949a4a77d2cfdeeed3403_1024x1024.webp"),
      require("../assets/images/home_category_5_medium.webp"),
      require("../assets/images/home_category_11_medium.webp"),
      require("../assets/images/m11_6d8271cbcf1949a4a77d2cfdeeed3403_1024x1024.webp"),
      require("../assets/images/home_category_5_medium.webp"),
      require("../assets/images/home_category_11_medium.webp"),
    ],
    texture: "Kem lì",
    volume: "4.5g",
    brand: "Merzy",
    origin: "Hàn Quốc",
    madeIn: "Hàn Quốc",
    expirationDate: "36 tháng kể từ ngày sản xuất",
    reviews: [
      {
        id: 1,
        userName: "Người dùng 1",
        rating: 5,
        comment: "Sản phẩm rất tốt!",
      },
      {
        id: 2,
        userName: "Người dùng 2",
        rating: 4,
        comment: "Màu đẹp, giá hợp lý",
      },
    ],
    recommendedProducts: [1, 2, 3, 4].map((i) => ({
      id: i,
      name: `Son Merzy V6 Blue Dream ${i}`,
      price: 169000,
      urlImage: require("../assets/images/clio_copy_fef881b3ab6a4e2eba206f50726cbbd9_1024x1024.webp"),
    })),
    viewedProducts: [5, 6, 7, 8].map((i) => ({
      id: i,
      name: `Son Merzy V6 Blue Dream ${i}`,
      price: 169000,
      urlImage: require("../assets/images/clio_copy_fef881b3ab6a4e2eba206f50726cbbd9_1024x1024.webp"),
    })),
    category: "Son môi",
  }));

export default mockProducts;
