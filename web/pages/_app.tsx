import { FC } from "react";
import { AppProps } from "next/app";
import Head from "next/head";
import { ThemeProvider } from "styled-components";
import { colors, GlobalStyle } from "../theme";
import AppLayout from "../components/AppLayout";
import Content from "../components/Content";
import Footer from "../components/Footer";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const App: FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <ThemeProvider theme={{ colors }}>
      <Head>
        <title>Kısa Be Aga</title>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin="true"
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap"
          rel="stylesheet"
        />
      </Head>
      <GlobalStyle />
      <ToastContainer />
      <AppLayout>
        <Content>
          <Component {...pageProps} />
        </Content>
        <Footer />
      </AppLayout>
    </ThemeProvider>
  );
};

export default App;
