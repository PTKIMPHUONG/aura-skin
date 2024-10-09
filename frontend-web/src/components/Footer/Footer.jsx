import React from "react";
import { Box, Container, Divider } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import FooterInfo from "./FooterInfo";
import Copyright from "./Copyright";

const Footer = () => {
  const theme = useTheme();

  return (
    <Box sx={{ bgcolor: theme.palette.layout.main, mt: 8, py: 4 }}>
      <Container maxWidth="lg">
        <FooterInfo />
      </Container>
      <Divider
        sx={{
          width: "100%",
          my: 4,
          borderBottomWidth: 3,
          bgcolor: "white",
          borderColor: "white",
        }}
      />
      <Container maxWidth="lg">
        <Copyright />
      </Container>
    </Box>
  );
};

export default Footer;
