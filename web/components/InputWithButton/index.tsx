import { DetailedHTMLProps, InputHTMLAttributes, forwardRef } from "react";
import { IconWrapper, InputWrapper } from "./styles";
import { RiLinksFill } from "react-icons/ri";

type InputWithButtonProps = DetailedHTMLProps<
  InputHTMLAttributes<HTMLInputElement>,
  HTMLInputElement
> & { buttonText: string; buttonOnClick?: () => void };

const InputWithButton = forwardRef<HTMLInputElement, InputWithButtonProps>(
  ({ buttonText, buttonOnClick, ...other }, ref) => {
    return (
      <InputWrapper>
        <IconWrapper>
          <RiLinksFill size={32} />
        </IconWrapper>
        <input autoComplete="off" ref={ref} {...other} />
        <button onClick={buttonOnClick}>{buttonText}</button>
      </InputWrapper>
    );
  }
);

export default InputWithButton;
