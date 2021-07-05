import { darken, rgba } from "polished";
import styled from "styled-components";

export const InputWrapper = styled.div`
  position: relative;
  display: flex;
  align-items: center;
  background: #ffffff;
  box-shadow: 0px 25px 50px rgba(0, 0, 0, 0.1);
  border-radius: 30px;
  padding: 0rem 0.25rem;
  height: 3.75rem;
  transition: all 0.3s ease;

  :focus-within {
    input {
      ::placeholder {
        transition: transform 0.3s ease, opacity 0.3s ease;
        transform: translateX(-5rem);
        opacity: 0;
      }
    }
  }
  input {
    background: transparent;
    outline: none;
    border: none;
    width: 100%;
    height: 3.75rem;
    font-weight: 300;
    font-size: 1.1em;
    line-height: 24px;
    ::placeholder {
      color: ${rgba("#000000", 0.6)};
    }
  }
  button {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    outline: none;
    border: none;
    cursor: pointer;
    color: white;
    font-size: 1.1rem;
    background: ${({ theme }) => theme.colors.primary};
    padding: 1rem 1.5rem;
    border-radius: 30px;
    box-shadow: 0px 25px 50px ${({ theme }) => rgba(theme.colors.primary, 0.3)};
    transition: transform 0.3s ease, background-color 0.3s ease;
    :hover {
      background-color: ${({ theme }) => darken(0.1, theme.colors.primary)};
      transform: translateY(-0.1rem);
    }
  }
`;

export const IconWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  color: black;
  width: 4rem;
`;
