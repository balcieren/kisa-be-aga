import { useState } from 'react';
import {
  Box,
  Text,
  Input,
  Button,
  InputGroup,
  InputLeftElement,
  chakra,
  useToast,
} from '@chakra-ui/react';
import { FiLink } from 'react-icons/fi';
import Template from '../components/Template';

import axiosInstance from '../utils/api';
import ShortenedUrl from '../components/ShortenedUrl';
import { GetServerSideProps } from 'next';

type Props = { host: string };
const Homepage: React.FC<Props> = ({ host }) => {
  const [url, setUrl] = useState('');
  const [shortId, setShortId] = useState<string>('');
  const toast = useToast();

  const shortenUrl = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const res = await axiosInstance.post('/url', { url });
      setShortId(res.data.shortId);
    } catch (error) {
      console.log('catch err:: ', error.message);
      return toast({
        title: 'Hata',
        description: error.message,
        duration: 3000,
        status: 'error',
      });
    }
  };

  return (
    <Template>
      <Box>
        <Text fontWeight='bold' fontSize='4xl' mb='4'>
          URL kısalt be aga
        </Text>
        <chakra.form display='flex' onSubmit={shortenUrl}>
          <InputGroup>
            <InputLeftElement
              pointerEvents='none'
              children={<FiLink color='gray' />}
            />
            <Input
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder='url gir bakem'
              roundedRight='none'
              _focus={{ outline: 'none' }}
            />
          </InputGroup>
          <Button colorScheme='blue' roundedLeft='none' type='submit'>
            Kısalt
          </Button>
        </chakra.form>
      </Box>

      {shortId !== '' && <ShortenedUrl shortId={`${host}/${shortId}`} />}
    </Template>
  );
};

export const getServerSideProps: GetServerSideProps = async (context) => {
  const {
    req: {
      headers: { host },
    },
  } = context;

  return {
    props: {
      host,
    },
  };
};

export default Homepage;
