import React, { useState } from "react";
import {
  Box,
  Typography,
  Checkbox,
  FormControlLabel,
  Button,
  TextField,
  Accordion,
  AccordionSummary,
  AccordionDetails,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { styled } from "@mui/material/styles";

const TitleSidebar = {
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
  fontSize: "22px",
  fontFamily: "Inter",
  fontWeight: "400",
  width: "243px",
  height: "63px",
  borderRadius: "4px",
  boxShadow: "0px 1px 3px 1px rgba(0, 0, 0, 0.15) inset",
};

const StyledSidebar = styled(Box)(({ theme }) => ({
  padding: theme.spacing(2),
  backgroundColor: "#FFFFFF",
  borderRadius: theme.shape.borderRadius,
}));

const StyledAccordion = styled(Accordion)(({ theme }) => ({
  borderBottom: "1px solid #C5C5C5",
  "&.MuiAccordion-root": {
    boxShadow: "none",
    "&:before": {
      display: "none",
    },
  },
  "& .MuiAccordionSummary-root": {
    minHeight: "auto",
    "&.Mui-expanded": {
      minHeight: "auto",
    },
  },
}));

const StyledAccordionDetails = {
  borderTop: "1px solid #C5C5C5",
};

const PriceRangeInput = styled(Box)(({ theme }) => ({
  display: "flex",
  justifyContent: "space-between",
  alignItems: "center",
  marginBottom: theme.spacing(1),
}));

const StyledFormControlLabel = styled(FormControlLabel)(({ theme }) => ({
  display: "block",
  marginBottom: theme.spacing(1),
}));

function Sidebar({ categoryName }) {
  const [expandedPanels, setExpandedPanels] = useState({
    panel1: true,
    panel2: true,
    panel3: true,
  });

  const handleChange = (panel) => (event, isExpanded) => {
    setExpandedPanels((prev) => ({
      ...prev,
      [panel]: isExpanded,
    }));
  };

  return (
    <StyledSidebar>
      <Typography sx={TitleSidebar} variant="h6" gutterBottom>
        {categoryName || "Tất cả"}
      </Typography>

      <StyledAccordion
        expanded={expandedPanels.panel1}
        onChange={handleChange("panel1")}
      >
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography>Danh Mục Sản Phẩm</Typography>
        </AccordionSummary>
        <AccordionDetails sx={StyledAccordionDetails}>
          <StyledFormControlLabel control={<Checkbox />} label="Mắt" />
          <StyledFormControlLabel control={<Checkbox />} label="Mặt" />
          <StyledFormControlLabel control={<Checkbox />} label="Môi" />
          <StyledFormControlLabel
            control={<Checkbox />}
            label="Dụng cụ trang điểm"
          />
          <StyledFormControlLabel control={<Checkbox />} label="Phụ kiện" />
        </AccordionDetails>
      </StyledAccordion>

      <StyledAccordion
        expanded={expandedPanels.panel2}
        onChange={handleChange("panel2")}
      >
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography>Giá</Typography>
        </AccordionSummary>
        <AccordionDetails sx={StyledAccordionDetails}>
          <PriceRangeInput>
            <TextField size="small" defaultValue="0" />
            <Box component="span" mx={1}>
              —
            </Box>
            <TextField size="small" defaultValue="100,000,000" />
          </PriceRangeInput>
        </AccordionDetails>
      </StyledAccordion>

      <StyledAccordion
        expanded={expandedPanels.panel3}
        onChange={handleChange("panel3")}
      >
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography>Thương Hiệu</Typography>
        </AccordionSummary>
        <AccordionDetails sx={StyledAccordionDetails}>
          <StyledFormControlLabel control={<Checkbox />} label="Merzy" />
          <StyledFormControlLabel control={<Checkbox />} label="Black Rouge" />
          <StyledFormControlLabel control={<Checkbox />} label="Klairs" />
        </AccordionDetails>
      </StyledAccordion>

      <Button
        variant="contained"
        fullWidth
        sx={{
          backgroundColor: "primary.main",
          opacity: "0.9",
          "&:hover": { backgroundColor: "primary.main", opacity: "1" },
          marginTop: 2,
        }}
      >
        ÁP DỤNG
      </Button>
    </StyledSidebar>
  );
}

export default Sidebar;
