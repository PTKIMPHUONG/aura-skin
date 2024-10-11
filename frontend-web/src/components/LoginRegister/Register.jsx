import React, { useState } from "react";
import AuthForm from "./AuthForm";
import authService from "../../services/AuthService";

const Register = () => {
  const [error, setError] = useState("");

  const fields = [
    {
      name: "username",
      label: "Tên người dùng",
      type: "text",
      autoComplete: "username",
    },
    {
      name: "email",
      label: "Địa chỉ Email",
      type: "email",
      autoComplete: "email",
    },
    {
      name: "password",
      label: "Mật khẩu",
      type: "password",
      autoComplete: "new-password",
    },
    {
      name: "phone_number",
      label: "Số điện thoại",
      type: "tel",
      autoComplete: "tel",
    },
  ];

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const username = formData.get("username");
    const email = formData.get("email");
    const password = formData.get("password");
    const phone_number = formData.get("phone_number");

    try {
      const result = await authService.register(
        username,
        email,
        password,
        phone_number
      );
      if (result.success) {
        alert("Đăng ký thành công! Đang chuyển hướng đến trang đăng nhập...");
        setTimeout(() => {
          window.location.href = "/login";
        }, 2000);
      } else {
        setError(result.message);
      }
    } catch (error) {
      console.error("Registration error:", error);
      setError("Có lỗi xảy ra khi đăng ký");
    }
  };

  return (
    <AuthForm
      title="Đăng ký"
      fields={fields}
      submitText="Đăng ký"
      alternativeText="Đã có tài khoản?"
      alternativeLink="/login"
      onSubmit={handleSubmit}
      error={error}
    />
  );
};

export default Register;
