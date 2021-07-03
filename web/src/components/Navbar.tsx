import Link from 'next/link';
import {
  Box,
  Button,
  Container,
  Stack,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import ThemeSwitcher from './ThemeSwitcher';

function NavbarLink({ url, name }: { url: string; name: string }) {
  const router = useRouter();

  const activePage = url === router.pathname;

  const activeLinkColor = useColorModeValue(
    activePage ? 'blackAlpha.600' : 'blackAlpha.800',
    activePage ? 'whiteAlpha.600' : 'whiteAlpha.800'
  );

  return (
    <Link href={url}>
      <Button color={activeLinkColor} variant='ghost'>
        {name}
      </Button>
    </Link>
  );
}

const Navbar = () => {
  return (
    <Box as='header' py='2' bg='blackAlpha.100' mb='10'>
      <Container maxWidth='6xl'>
        <Stack
          as='nav'
          direction={['column', 'row']}
          spacing='3'
          alignItems='center'
          justifyContent='space-between'
        >
          <Box display='flex' alignItems='center'>
            <Text fontWeight='bold' mr='4'>KISA BE AGA</Text>
            <NavbarLink name='Anasayfa' url='/' />
          </Box>
          <ThemeSwitcher />
        </Stack>
      </Container>
    </Box>
  );
};

export default Navbar;
