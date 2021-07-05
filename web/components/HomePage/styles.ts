import styled from "styled-components";

export const Container = styled.section`
  display: grid;
  grid-template-columns: 60vw;
  row-gap: 2.5rem;
  min-height: 80vh;
  align-items: center;
  justify-content: center;

  @media screen and (max-width: 1440px) {
    grid-template-columns: 70vw;
  }

  @media screen and (max-width: 1280px) {
    grid-template-columns: 90vw;
  }

  @media screen and (max-width: 1024px) {
    grid-template-columns: 90vw;
  }

  @media screen and (max-width: 960px) {
    grid-template-columns: 80vw;
  }
  @media screen and (max-width: 600px) {
    grid-template-columns: 90vw;
  }
`;

export const Section = styled.section`
  display: grid;
  grid-template-columns: 30vw;
  justify-content: center;

  @media screen and (max-width: 960px) {
    grid-template-columns: 1fr;
  }
  @media screen and (max-width: 600px) {
    grid-template-columns: 1fr;
  }
`;

export const ShortenedLinksSection = styled.div`
  display: flex;
  flex-direction: column;
  gap: 1rem;
`;
export const Items = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(25rem, 1fr));
  gap: 0.75rem;
`;

export const Heading = styled.h1`
  font-size: 4.5em;
  text-align: center;
  font-weight: 700;
`;
export const Title = styled.h1`
  text-align: start;
  font-weight: 500;
`;
