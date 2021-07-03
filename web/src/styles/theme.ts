import { extendTheme } from '@chakra-ui/react';

export default extendTheme({
  components: {
    Button: {
      baseStyle: {
        rounded: 'sm',
        _focus: { boxShadow: 'none' },
      },
    },
  },
});
