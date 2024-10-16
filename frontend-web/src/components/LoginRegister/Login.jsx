import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import AuthForm from "./AuthForm";
import authService from "../../services/AuthService";
import { useAuth } from "../../context/Authcontext";

const Login = () => {
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { login } = useAuth();

  // Thêm định nghĩa cho fields ở đây
  const fields = [
    {
      name: "email",
      label: "Email Address",
      type: "email",
      autoComplete: "email",
    },
    {
      name: "password",
      label: "Password",
      type: "password",
      autoComplete: "current-password",
    },
  ];

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const email = formData.get("email");
    const password = formData.get("password");
    try {
      const result = await authService.login(email, password);
      console.log("Login result:", result);
      if (result.success) {
        await login(result.user, result.token);
        console.log("User after login:", result.user);
        if (result.user.isAdmin) {
          console.log("User is admin, redirecting to admin page");
          navigate("/admin");
        } else {
          console.log("User is not admin, redirecting to home page");
          navigate("/");
        }
      } else {
        setError(result.message);
      }
    } catch (error) {
      console.error("Login error:", error);
      setError("Có lỗi xảy ra khi đăng nhập");
    }
  };

  return (
    <AuthForm
      title="Đăng nhập"
      fields={fields}
      submitText="Đăng nhập"
      alternativeText="Chưa có tài khoản?"
      alternativeLink="/register"
      onSubmit={handleSubmit}
      error={error}
    />
  );
};

export default Login;
