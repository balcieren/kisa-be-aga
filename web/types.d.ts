import { ThemeColors } from "./theme/colors";

declare module "styled-components" {
  export interface DefaultTheme {
    colors: ThemeColors;
  }
}
