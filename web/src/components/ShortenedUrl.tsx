import {
  Box,
  Text,
  IconButton,
  ButtonGroup,
  useClipboard,
  useToast,
} from '@chakra-ui/react';
import { FiLink, FiTrash } from 'react-icons/fi';

type Props = {
  shortId: string;
};

const ShortenedUrl: React.FC<Props> = ({ shortId }) => {
  const toast = useToast();

  const { onCopy } = useClipboard(shortId);

  const copyToClipboard = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    onCopy();

    toast({
      title: 'Hata',
      description: 'URL panoya kopyalandı',
      duration: 3000,
      status: 'success',
    });
  };

  return (
    <Box boxShadow='2xl' maxW='3xl' mx='auto' p='6' my='6' rounded='md'>
      <Box display='flex' justifyContent='space-between' alignItems='center'>
        <Text fontWeight='semibold' fontSize='xl'>
          {shortId}
        </Text>
        <ButtonGroup spacing={3}>
          <IconButton
            aria-label='copy clipboard'
            icon={<FiLink />}
            colorScheme='blue'
            onClick={copyToClipboard}
          />
          <IconButton
            aria-label='copy clipboard'
            icon={<FiTrash />}
            colorScheme='red'
            disabled
          />
        </ButtonGroup>
      </Box>
    </Box>
  );
};

export default ShortenedUrl;
