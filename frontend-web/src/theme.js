import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#A877B2",
    },
    secondary: {
      main: "#f50057",
    },
    layout: {
      main: "#C3C8FF",
    },
    text: {
      primary: "#333333",
      secondary: "#666666",
    },
  },
  typography: {
    fontFamily: [
      "Inter",
      "Jost",
      "Asap",
      "Arima",
      "Jomolhari",
      "Asar",
      "-apple-system",
      "BlinkMacSystemFont",
      '"Segoe UI"',
      "Roboto",
      '"Helvetica Neue"',
      "Arial",
      "sans-serif",
      '"Apple Color Emoji"',
      '"Segoe UI Emoji"',
      '"Segoe UI Symbol"',
    ].join(","),
    h1: {
      fontFamily: "Jost, Arial, sans-serif",
    },
    h2: {
      fontFamily: "Asap, Arial, sans-serif",
    },
    body1: {
      fontFamily: "Inter, Arial, sans-serif",
    },
    // Bạn có thể thêm các variant typography khác ở đây
  },
  // Các cấu hình theme khác
  components: {
    MuiAppBar: {
      styleOverrides: {
        root: {
          backgroundColor: "#C3C8FF",
          color: "#333333",
        },
      },
    },
    MuiIconButton: {
      styleOverrides: {
        root: {
          color: "#333333",
        },
      },
    },
  },
});

export default theme;
