import { darken, rgba } from "polished";
import styled, { css } from "styled-components";

export const Container = styled.div`
  position: relative;
  background: #ffffff;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1.5rem;
  box-shadow: 0px 25px 50px rgba(0, 0, 0, 0.1);
  border-radius: 30px;
  padding: 1.5rem;
`;

export const Top = styled.div`
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  h1 {
    font-size: 1.5rem;
    font-weight: 2rem;
    font-style: normal;
    font-weight: normal;
  }
  div {
    display: flex;
    gap: 0.25rem;
  }
`;

export type ButtonProps = {
  color: "danger" | "primary" | "success";
  showStatus?: boolean;
};
export const Button = styled.button<ButtonProps>`
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  outline: none;
  border: none;
  cursor: pointer;
  color: white;
  background: ${({ theme, color }) => theme.colors[color]};
  height: 2rem;
  width: 2rem;
  border-radius: 0.625rem;
  box-shadow: 0px 25px 50px
    ${({ theme, color }) => rgba(theme.colors[color], 0.3)};
  transition: transform 0.3s ease, background-color 0.3s ease;
  :hover {
    background-color: ${({ theme, color }) => darken(0.1, theme.colors[color])};
    transform: translateY(-0.1rem);
  }
`;

export const Bottom = styled.div`
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .box {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.25rem;
    transition: transform 0.3s ease, opacity 0.3s ease;
  }

  // enter from
  &.show-animation-enter {
    .box {
      opacity: 0;
      transform: scale(0);
    }
  }

  // enter to
  &.show-animation-enter-active {
    .box {
      opacity: 1;
      transform: scale(1);
    }
  }

  // exit from
  &.show-animation-exit {
    .box {
      opacity: 1;
      transform: scale(1);
    }
  }

  // exit to
  &.show-animation-exit-active {
    .box {
      opacity: 0;
      transform: scale(0);
    }
  }
`;

export type IconWrapperProps = { color: "danger" | "warning" };
export const IconWrapper = styled.div<IconWrapperProps>`
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 100%;
  height: 4rem;
  width: 4rem;
  color: ${({ theme, color }) => theme.colors[color]};
  background: ${({ theme, color }) => rgba(theme.colors[color], 0.41)};
  box-shadow: 0px 25px 50px
    ${({ theme, color }) => rgba(theme.colors[color], 0.3)};
`;

export const TextContainer = styled.div`
  display: flex;
  flex-direction: column;

  h1 {
    font-style: normal;
    font-weight: 500;
    font-size: 18px;
    color: black;
    line-height: 1rem;
  }
  p {
    font-style: normal;
    font-weight: 300;
    font-size: 12px;
    color: black;
  }
`;
