import Head from 'next/head';
import { Container } from '@chakra-ui/react';
import Navbar from './Navbar';

type Props = {
  title?: string;
};

const Template: React.FC<Props> = ({ children, title }) => {
  return (
    <>
      <Head>
        <title>{title || 'KISA BE AGA'}</title>
      </Head>
      <Navbar />
      <Container maxWidth='6xl' my='4' py='8'>
        {children}
      </Container>
    </>
  );
};

export default Template;
