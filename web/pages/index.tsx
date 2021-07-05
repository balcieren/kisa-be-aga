import { NextPage } from "next";
import { useState } from "react";
import {
  Container,
  Heading,
  Items,
  Section,
  ShortenedLinksSection,
  Title,
} from "../components/HomePage/styles";
import InputWithButton from "../components/InputWithButton";
import ShortenedUrl from "../components/ShortenedUrl";
import UrlStatus from "../components/UrlStatus";
import { useUrlService } from "../services/url.service";

const HomePage: NextPage = () => {
  const { shortId, shortenUrl } = useUrlService();
  const [inputText, setInputText] = useState<string>("");

  return (
    <Container>
      <div>
        <Heading>KISA BE AGA</Heading>
        <InputWithButton
          onChange={({ target }) => setInputText(target.value)}
          placeholder="https://..."
          buttonText="Shorten"
          buttonOnClick={() => shortenUrl(inputText)}
        />
      </div>
      <Section>
        <ShortenedUrl shortId={shortId} />
      </Section>

      <ShortenedLinksSection>
        <Title>My Shortened Links</Title>
        <Items>
          <UrlStatus shortId="23123" clicks={5} createdAt="1 September 201" />
        </Items>
      </ShortenedLinksSection>
    </Container>
  );
};

export default HomePage;
