import mockProducts from "./mockProducts";
const mockFavorites = [
  {
    id: 1,
    username: "Nguyenminhhoang",
    email: "nguyenminhhoang@gmail.com",
    password: "123",
    imageUser: require("../assets/images/Users/434391559_1576294329886781_8808623104079063343_n.jpg"),
    favorites: mockProducts.slice(0, 6), // Lấy 6 sản phẩm đầu tiên từ mockProducts
  },
  {
    id: 2,
    username: "hoangnguyen",
    email: "hoangnguyen@gmail.com",
    password: "123",
    imageUser: require("../assets/images/Users/434391559_1576294329886781_8808623104079063343_n.jpg"),
    favorites: mockProducts.slice(3, 9), // Lấy 6 sản phẩm tiếp theo từ mockProducts
  },
];

export default mockFavorites;
