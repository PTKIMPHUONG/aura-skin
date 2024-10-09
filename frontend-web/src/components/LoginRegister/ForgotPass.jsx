import React from "react";
import AuthForm from "./AuthForm";

const ForgotPassword = () => {
  const fields = [
    {
      name: "email",
      label: "Email Address",
      type: "email",
      autoComplete: "email",
    },
  ];

  const handleSubmit = (event) => {
    event.preventDefault();
    // Xử lý quên mật khẩu
  };

  return (
    <AuthForm
      title="Forgot Password"
      fields={fields}
      submitText="Reset Password"
      alternativeText="Remember your password?"
      alternativeLink="/login"
      onSubmit={handleSubmit}
    />
  );
};

export default ForgotPassword;
