import Head from "next/head";
import Link from "next/link";
import React from "react";
import { Container } from "@mui/material";

import Navbar from "../components/Navbar/navbar";

const Layout = ({ children, router }) => {
  return (
    <>
      <Head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>Git Status</title>
      </Head>
      <Navbar path={router.asPath} />
      <Container maxWidth="md">{children}</Container>
    </>
  );
};

export default Layout;
