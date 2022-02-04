import Head from "next/head";
import Link from "next/link";
import React from "react";
import {Container} from '@mui/material';

const Layout = ({ children, router }) => {
  return (
    <>
      <Head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>Git Status</title>
      </Head>
      <Container maxWidth="lg">{children}</Container>
    </>
  );
};

export default Layout;
