import { darken, rgba } from "polished";
import styled, { keyframes } from "styled-components";

const fade = keyframes`
  from {
    opacity: 0;
    transform: scale(0);
  }

  to {
   opacity: 1;
   transform: scale(1);
  }
`;

export const Container = styled.div`
  position: relative;
  display: flex;
  gap: 1rem;
  justify-content: center;
  align-items: center;
  height: 100px;
  width: auto;
  background: #ffffff;
  box-shadow: 0px 25px 50px rgba(0, 0, 0, 0.08);
  border-radius: 30px;
  padding: 1rem;
  h1 {
    font-style: normal;
    font-weight: normal;
    font-size: 1.5em;
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
    background: ${({ theme, color }) => theme.colors.primary};
    padding: 0.45rem;
    border-radius: 0.625rem;
    box-shadow: 0px 25px 50px ${({ theme }) => rgba(theme.colors.primary, 0.3)};
    transition: transform 0.3s ease, background-color 0.3s ease;
    :hover {
      background-color: ${({ theme }) => darken(0.1, theme.colors.primary)};
      transform: translateY(-0.1rem);
    }
  }

  animation: ${fade} 0.5s ease;
`;
