import { Static, Type } from "@sinclair/typebox";

export const urlParams = Type.Object({
  shortid: Type.String(),
});
export type UrlParamsType = Static<typeof urlParams>;

export const urlBody = Type.Object({
  url: Type.String({ format: "uri", title: "Url Body" }),
});
export type UrlBodyType = Static<typeof urlBody>;
