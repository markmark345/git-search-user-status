import Layout from "../layouts/layout";
import ContextProvider from "../context/ContextProvider ";
import "../styles/globals.css";

const MyApp = ({ Component, pageProps, router }) => {
  return (
    <ContextProvider>
      <Layout router={router}>
        <Component {...pageProps} key={router.route} />
      </Layout>
    </ContextProvider>
  );
};

export default MyApp;
