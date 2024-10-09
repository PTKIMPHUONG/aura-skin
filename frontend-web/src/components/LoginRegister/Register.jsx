import React from "react";
import AuthForm from "./AuthForm";

const Register = () => {
  const fields = [
    {
      name: "username",
      label: "Username",
      type: "text",
      autoComplete: "username",
    },
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
      autoComplete: "new-password",
    },
    {
      name: "confirmPassword",
      label: "Confirm Password",
      type: "password",
      autoComplete: "new-password",
    },
  ];

  const handleSubmit = (event) => {
    event.preventDefault();
    // Xử lý đăng ký
  };

  return (
    <AuthForm
      title="Register"
      fields={fields}
      submitText="Register"
      alternativeText="Already have an account?"
      alternativeLink="/login"
      onSubmit={handleSubmit}
    />
  );
};

export default Register;
