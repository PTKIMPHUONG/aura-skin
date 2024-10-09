import React from "react";
import { Grid, Typography, Box, Link } from "@mui/material";
import { Facebook, Instagram, Twitter } from "@mui/icons-material";
import GoogleIcon from "@mui/icons-material/Google";

const FooterInfo = () => {
  return (
    <Grid container spacing={4}>
      <Grid item xs={12} sm={6} md={3}>
        <Typography
          variant="h6"
          color="black"
          gutterBottom
          sx={{ fontFamily: "Jomolhari", mb: 2 }}
        >
          THÔNG TIN LIÊN HỆ
        </Typography>
        <Typography
          variant="body2"
          color="black"
          sx={{ fontFamily: "Asar", mb: 2 }}
        >
          <Box
            component="span"
            sx={{ display: "flex", alignItems: "center", mb: 2 }}
          >
            <Box component="span" sx={{ mr: 2 }}>
              📞
            </Box>{" "}
            0123456789
          </Box>
          <Box
            component="span"
            sx={{ display: "flex", alignItems: "center", mb: 2 }}
          >
            <Box component="span" sx={{ mr: 1 }}>
              ✉️
            </Box>{" "}
            auraskin@cosmetic.com
          </Box>
          <Box
            component="span"
            sx={{ display: "flex", alignItems: "center", lineHeight: 2 }}
          >
            <Box component="span" sx={{ mr: 1 }}>
              📍
            </Box>{" "}
            140 Lê Trọng Tấn, Phường Tây Thạnh, Quận Tân Phú, TP.HCM
          </Box>
        </Typography>
      </Grid>
      <Grid item xs={12} sm={6} md={3}>
        <Typography
          variant="h6"
          color="black"
          gutterBottom
          sx={{ fontFamily: "Jomolhari", mb: 2 }}
        >
          DANH MỤC
        </Typography>
        <Box component="ul" sx={{ listStyle: "none", padding: 0, margin: 0 }}>
          {[
            "Trang chủ",
            "Sản phẩm bán chạy",
            "Sản phẩm mới",
            "Chăm sóc da",
            "Trang điểm",
            "Thương hiệu",
          ].map((item) => (
            <Box
              component="li"
              key={item}
              sx={{ fontFamily: "Asar", color: "black", mb: 1.5 }}
            >
              {item}
            </Box>
          ))}
        </Box>
      </Grid>
      <Grid item xs={12} sm={6} md={3}>
        <Typography
          variant="h6"
          color="black"
          gutterBottom
          sx={{ fontFamily: "Jomolhari", mb: 2 }}
        >
          DỊCH VỤ
        </Typography>
        <Box component="ul" sx={{ listStyle: "none", padding: 0, margin: 0 }}>
          {[
            "Hỗ trợ",
            "Chính sách đổi trả",
            "Chính sách giao hàng",
            "Chứng chỉ đại lý chính hãng",
            "Tuyển dụng",
            "Liên hệ",
          ].map((item) => (
            <Box
              component="li"
              key={item}
              sx={{ fontFamily: "Asar", color: "black", mb: 1.5 }}
            >
              {item}
            </Box>
          ))}
        </Box>
      </Grid>
      <Grid item xs={12} sm={6} md={3}>
        <Typography
          variant="h6"
          color="black"
          gutterBottom
          sx={{ fontFamily: "Jomolhari", mb: 2 }}
        >
          THEO DÕI CHÚNG TÔI
        </Typography>
        <Box textAlign="center">
          <Link href="https://www.facebook.com/" sx={{ color: "black", mr: 2 }}>
            <Facebook />
          </Link>
          <Link
            href="https://www.instagram.com/"
            sx={{ color: "black", mr: 2 }}
          >
            <Instagram />
          </Link>
          <Link href="https://www.twitter.com/" sx={{ color: "black", mr: 2 }}>
            <Twitter />
          </Link>
          <Link href="https://www.google.com/" sx={{ color: "black" }}>
            <GoogleIcon />
          </Link>
        </Box>
      </Grid>
    </Grid>
  );
};

export default FooterInfo;
