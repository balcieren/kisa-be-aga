import { Static } from "@sinclair/typebox";
import { urlBodyValidate, urlParamsValidate } from "./validates";

export type UrlParams = Static<typeof urlParamsValidate>;
export type UrlBody = Static<typeof urlBodyValidate>;
