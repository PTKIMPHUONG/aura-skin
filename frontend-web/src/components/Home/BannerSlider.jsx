import React, { useState, useEffect } from "react";
import { Box, Typography, Grid, Container } from "@mui/material";
import bannerImage from "../../assets/images/Banner/8d9219c8fdb373e261e8ba267bb151a0.png";

const BannerSlider = () => {
  const [timeLeft, setTimeLeft] = useState({
    days: 6,
    hours: 12,
    minutes: 53,
    seconds: 36,
  });

  useEffect(() => {
    const timer = setInterval(() => {
      setTimeLeft((prevTime) => {
        if (prevTime.seconds > 0) {
          return { ...prevTime, seconds: prevTime.seconds - 1 };
        } else if (prevTime.minutes > 0) {
          return { ...prevTime, minutes: prevTime.minutes - 1, seconds: 59 };
        } else if (prevTime.hours > 0) {
          return {
            ...prevTime,
            hours: prevTime.hours - 1,
            minutes: 59,
            seconds: 59,
          };
        } else if (prevTime.days > 0) {
          return {
            ...prevTime,
            days: prevTime.days - 1,
            hours: 23,
            minutes: 59,
            seconds: 59,
          };
        } else {
          clearInterval(timer);
          return prevTime;
        }
      });
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  return (
    <Container maxWidth="lg">
      <Box
        sx={{
          height: 300,
          width: "100%",
          backgroundImage: `url(${bannerImage})`,
          backgroundSize: "cover",
          backgroundPosition: "center",
          position: "relative",
          borderRadius: 2,
          overflow: "hidden",
        }}
      >
        <Box
          sx={{
            position: "absolute",
            top: 10,
            left: 10,
            color: "white",
            textAlign: "left",
          }}
        >
          <Typography
            variant="h5"
            sx={{
              textShadow: "0px 4px 4px rgba(0, 0, 0, 0.25)",
              mb: 1,
              fontWeight: "bold",
            }}
          >
            ĐẶT NGAY RƯỚC QUÀ TO
          </Typography>
          <Grid
            container
            spacing={1}
            sx={{
              display: "flex",
              justifyContent: "center",
              textAlign: "center",
              mb: 1,
            }}
          >
            {Object.entries(timeLeft).map(([key, value]) => (
              <Grid item key={key}>
                <Box sx={{ bgcolor: "#0075FF", p: 0.5, borderRadius: 1 }}>
                  <Typography variant="body2">
                    {value.toString().padStart(2, "0")}
                  </Typography>
                  <Typography variant="caption">{key.toUpperCase()}</Typography>
                </Box>
              </Grid>
            ))}
          </Grid>
          <Typography
            variant="caption"
            sx={{
              textAlign: "center",
              position: "relative",
              left: "20%",
              mb: 0.5,
            }}
          >
            Số lượng giới hạn chỉ còn:
          </Typography>
          <Typography
            variant="body1"
            sx={{
              position: "relative",
              left: "40%",
              mb: 1,
              color: "secondary.main",
            }}
          >
            470
          </Typography>
          <Box
            sx={{
              bgcolor: "primary.main",
              p: 0.5,
              borderRadius: 1,
              display: "inline-block",
              position: "relative",
              left: "24%",
            }}
          >
            <Typography variant="caption">QUÀ ĐỘC QUYỀN</Typography>
          </Box>
        </Box>
      </Box>
    </Container>
  );
};

export default BannerSlider;
