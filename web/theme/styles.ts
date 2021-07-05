import { createGlobalStyle } from "styled-components";

export const GlobalStyle = createGlobalStyle`

*{
  margin: 0rem;
  padding: 0rem;
  box-sizing: border-box;
}

body{
  font-family: 'Poppins', sans-serif;
  background: ${(props) => props.theme.colors.background};
}

button{
 user-select: none;
}

.Toastify__toast {
  border-radius: 1rem;
}
.Toastify__toast--rtl {
}
.Toastify__toast--dark {
}
.Toastify__toast--default {
}
.Toastify__toast--info {
  background: ${({ theme }) => theme.colors.primary};
}
.Toastify__toast--success {
  background: ${({ theme }) => theme.colors.success};
}
.Toastify__toast--warning {
  background: ${({ theme }) => theme.colors.warning};
}
.Toastify__toast--error {
  background: ${({ theme }) => theme.colors.danger};
}
.Toastify__toast-body {
}

`;
