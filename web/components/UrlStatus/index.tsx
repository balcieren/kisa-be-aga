import { FC } from "react";
import {
  Container,
  Top,
  Bottom,
  Button,
  TextContainer,
  IconWrapper,
} from "./styles";
import {
  RiLinksFill,
  RiDeleteBin2Fill,
  RiTimeFill,
  RiCursorFill,
  RiArrowDownFill,
  RiArrowUpFill,
} from "react-icons/ri";
import { useCopyToClipboard, useToggle } from "react-use";
import { toast } from "react-toastify";
import {
  CSSTransition,
  SwitchTransition,
  TransitionGroup,
} from "react-transition-group";

type UrlStatusProps = { shortId: string; clicks: number; createdAt: string };
const UrlStatus: FC<UrlStatusProps> = ({ shortId, clicks, createdAt }) => {
  const [state, copyToClipboard] = useCopyToClipboard();
  const [showStatus, setShowStatus] = useToggle(false);
  const copyUrl = () => {
    toast.success("URL clipboard copied", {
      position: "top-center",
      autoClose: 1500,
    });
    copyToClipboard(process.env.WEB_URL + shortId);
  };
  const removeUrl = () => {};

  return (
    <Container>
      <Top>
        <h1>kisabea.ga/{shortId}</h1>
        <div>
          <Button onClick={copyUrl} color="primary">
            <RiLinksFill size={24} />
          </Button>
          <Button
            onClick={setShowStatus}
            showStatus={showStatus}
            color="success"
          >
            <RiArrowDownFill size={24} />
          </Button>
          <Button onClick={removeUrl} color="danger">
            <RiDeleteBin2Fill size={24} />
          </Button>
        </div>
      </Top>
      <CSSTransition
        in={showStatus}
        classNames="show-animation"
        timeout={300}
        unmountOnExit
      >
        <Bottom>
          <div className="box">
            <IconWrapper color="danger">
              <RiCursorFill size={32} />
            </IconWrapper>
            <TextContainer>
              <h1>{clicks}</h1>
              <p>People clicked</p>
            </TextContainer>
          </div>
          <div className="box">
            <IconWrapper color="warning">
              <RiTimeFill size={32} />
            </IconWrapper>
            <TextContainer>
              <h1>{createdAt}</h1>
              <p>It was created at</p>
            </TextContainer>
          </div>
        </Bottom>
      </CSSTransition>
    </Container>
  );
};

export default UrlStatus;
