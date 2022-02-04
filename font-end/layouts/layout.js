import Head from "next/head";
import Link from "next/link";
import React from "react";

const Layout = ({ chlidren, router }) => {
    return (
        <div className="relative py-14 px-4 md:px-6 2xl:px-20 2xl:container 2xl:mx-auto">
            {children}
        </div>
    )
}

export default Layout