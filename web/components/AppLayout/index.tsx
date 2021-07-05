import { FC } from "react";
import { AppContainer } from "./styles";

const AppLayout: FC = ({ children }) => {
  return <AppContainer>{children}</AppContainer>;
};

export default AppLayout;
