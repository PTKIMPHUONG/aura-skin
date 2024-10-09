import React from "react";
import { Box, Typography, TextField, Button, Link } from "@mui/material";

const AuthForm = ({
  title,
  fields,
  submitText,
  alternativeText,
  alternativeLink,
  onSubmit,
  error,
}) => {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        maxWidth: 400,
        margin: "auto",
        padding: 3,
        borderRadius: 2,
        boxShadow: 3,
        backgroundColor: "white",
      }}
    >
      <Typography component="h1" variant="h5">
        {title}
      </Typography>
      <Box component="form" onSubmit={onSubmit} noValidate sx={{ mt: 1 }}>
        {fields.map((field) => (
          <TextField
            key={field.name}
            margin="normal"
            required
            fullWidth
            id={field.name}
            label={field.label}
            name={field.name}
            autoComplete={field.autoComplete}
            type={field.type}
          />
        ))}
        {error && (
          <Typography color="error" align="center">
            {error}
          </Typography>
        )}
        <Button
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          {submitText}
        </Button>
        <Typography variant="body2" align="center">
          {alternativeText}{" "}
          <Link href={alternativeLink} variant="body2">
            {alternativeText === "Forgot password?"
              ? "Reset here"
              : "Click here"}
          </Link>
        </Typography>
      </Box>
    </Box>
  );
};

export default AuthForm;
