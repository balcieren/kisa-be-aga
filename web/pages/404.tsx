import { NextPage } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import styled from "styled-components";

const ErrorPage: NextPage = () => {
  const router = useRouter();
  return (
    <Container>
      <Image
        id="image"
        onClick={() => router.push("/")}
        src="/error404.svg"
        height={600}
        width={600}
      />
      <p>Click on the picture to go to the main page</p>
    </Container>
  );
};

export default ErrorPage;

const Container = styled.div`
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 90vh;

  #image {
    cursor: pointer;
    transition: transform 0.3s ease;
    :hover {
      transform: scale(1.1);
    }
  }
`;
