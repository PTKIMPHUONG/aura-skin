import React from "react";
import AuthForm from "./AuthForm";

const ResetPassword = () => {
  const fields = [
    { name: "code", label: "Reset Code", type: "text", autoComplete: "off" },
    {
      name: "newPassword",
      label: "New Password",
      type: "password",
      autoComplete: "new-password",
    },
    {
      name: "confirmPassword",
      label: "Confirm New Password",
      type: "password",
      autoComplete: "new-password",
    },
  ];

  const handleSubmit = (event) => {
    event.preventDefault();
    // Xử lý đặt lại mật khẩu
  };

  return (
    <AuthForm
      title="Reset Password"
      fields={fields}
      submitText="Set New Password"
      alternativeText="Back to login"
      alternativeLink="/login"
      onSubmit={handleSubmit}
    />
  );
};

export default ResetPassword;
