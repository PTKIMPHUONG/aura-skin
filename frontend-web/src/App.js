import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import theme from "./theme";
import Header from "./components/Header/Header";
import ProductListPage from "./pages/ProductListPage";
import ProductDetailPage from "./pages/ProductDetailPage";
import Footer from "./components/Footer/Footer";
import HomePage from "./pages/HomePage";
import Login from "./components/LoginRegister/Login";
import Register from "./components/LoginRegister/Register";
import ForgotPassword from "./components/LoginRegister/ForgotPass";
import ResetPassword from "./components/LoginRegister/ResetPass";
import { AuthProvider } from "./context/Authcontext";
import UserFavoritesPage from "./pages/UserFavoritesPage";
import ProtectedRoute from "./components/Auth/ProtectedRoute";
import UserProfilePage from "./pages/UserProfilePage";
import UserCartPage from "./pages/UserCartPage";
import UserOrdersPage from "./pages/UserOrdersPage";
import UserAddressesPage from "./pages/UserAddressesPage";
import OrderConfirmationPage from "./pages/OrderConfirmationPage";
import CategoryProductListPage from "./pages/CategoryProductListPage";
import { CartProvider } from "./context/CartContext";
function App() {
  return (
    <AuthProvider>
      <CartProvider>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Router>
            <Header />
            <Routes>
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/forgot-password" element={<ForgotPassword />} />
              <Route path="/reset-password" element={<ResetPassword />} />
              <Route path="/" element={<HomePage />} />
              <Route path="/products" element={<ProductListPage />} />
              <Route
                path="/category/:categoryId"
                element={<CategoryProductListPage />}
              />
              <Route
                path="/products/product-detail"
                element={<ProductDetailPage />}
              />
              <Route element={<ProtectedRoute />}>
                <Route path="/user/profile" element={<UserProfilePage />} />
                <Route path="/user/cart" element={<UserCartPage />} />
                <Route path="/user/orders" element={<UserOrdersPage />} />
                <Route path="/user/addresses" element={<UserAddressesPage />} />
                {/* <Route path="/user/coupons" element={<UserCouponsPage />} /> */}
                <Route path="/user/favorites" element={<UserFavoritesPage />} />
                <Route
                  path="/order-confirmation"
                  element={<OrderConfirmationPage />}
                />
              </Route>
            </Routes>
            <Footer />
          </Router>
        </ThemeProvider>
      </CartProvider>
    </AuthProvider>
  );
}

export default App;
