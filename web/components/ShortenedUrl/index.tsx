import { FC } from "react";
import { Container } from "./styles";
import { RiLinksFill } from "react-icons/ri";
import { toast } from "react-toastify";
import { useCopyToClipboard } from "react-use";
import { CSSTransition, Transition } from "react-transition-group";

type ShortenedUrlProps = { shortId?: string };
const ShortenedUrl: FC<ShortenedUrlProps> = ({ shortId }) => {
  const [state, copyToClipboard] = useCopyToClipboard();

  const copyUrl = () => {
    toast.success("URL clipboard copied", {
      position: "top-center",
      autoClose: 1500,
    });
    copyToClipboard(process.env.WEB_URL + shortId);
  };
  return (
    !!shortId && (
      <Container>
        <h1>kisabea.ga/{shortId}</h1>
        <button onClick={copyUrl}>
          <RiLinksFill size={24} />
        </button>
      </Container>
    )
  );
};

export default ShortenedUrl;
