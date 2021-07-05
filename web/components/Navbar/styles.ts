import styled from "styled-components";

export const NavbarContainer = styled.nav`
  position: relative;
  display: grid;
  align-items: center;
  justify-items: stretch;
  grid-auto-flow: column;
  grid-template-columns: 1fr;
  .box1 {
    grid-column: span 2;
  }
  .box2 {
    grid-column: span 1;
  }
`;
