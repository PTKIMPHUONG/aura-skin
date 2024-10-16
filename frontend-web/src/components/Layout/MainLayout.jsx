import React from "react";
import { Outlet } from "react-router-dom";
import Header from "../Header/Header";
import Footer from "../Footer/Footer";

const MainLayout = () => {
  return (
    <>
      <Header />
      <main style={{ minHeight: "calc(100vh - 64px - 100px)" }}>
        <Outlet />
      </main>
      <Footer />
    </>
  );
};

export default MainLayout;
