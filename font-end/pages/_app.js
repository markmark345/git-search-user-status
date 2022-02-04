import Layout from "../layouts/layout";
import "../styles/globals.css";

const MyApp = ({ Component, pageProps, router }) => {
  return (
    <Layout router={router}>
      <Component {...pageProps} key={router.route} />
    </Layout>
  );
};

export default MyApp;
