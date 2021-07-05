declare namespace NodeJS {
  export interface ProcessEnv {
    NODE_ENV: "development" | "production";
    BASE_API_URL: string;
    WEB_URL: string;
  }
}
