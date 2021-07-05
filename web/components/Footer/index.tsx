import { FC } from "react";
import { FooterContainer } from "./styles";

const Footer: FC = () => {
  return (
    <FooterContainer>
      <p>
        Created by{" "}
        <a target="_blank" href="https://github.com/UrduX">
          Eren Balcı
        </a>
      </p>
    </FooterContainer>
  );
};

export default Footer;
