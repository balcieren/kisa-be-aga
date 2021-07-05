import { darken } from "polished";
import styled from "styled-components";

export const FooterContainer = styled.footer`
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;

  a {
    text-decoration: none;
    color: ${({ theme }) => theme.colors.primary};
    transition: all 0.3s ease;
    :hover {
      font-size: 1.1em;
      color: ${({ theme }) => darken(0.1, theme.colors.primary)};
    }
  }
`;
