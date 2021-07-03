import { Box, Text } from '@chakra-ui/react';
import Template from '../components/Template';
import axiosInstance from '../utils/api';
import { GetServerSideProps } from 'next';

const ShortId = () => {
  return (
    <Template>
      <Box>
        <Text fontWeight='bold' fontSize='4xl' mb='4'>
          URL kısalt be aga
        </Text>
      </Box>
    </Template>
  );
};

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { shortId } = context.params;
  try {
    await axiosInstance.get(`/url/${shortId}/status`);
    context.res
      .writeHead(301, {
        Location: `${process.env.NEXT_PUBLIC_API_URL}/url/${shortId}`,
      })
      .end();
  } catch (error) {
    console.log('catch err:: ', error.message);
    return { notFound: true };
  }
  return {
    props: {},
  };
};

export default ShortId;
