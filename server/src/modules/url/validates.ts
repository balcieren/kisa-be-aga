import { Type } from "@sinclair/typebox";

export const urlParamsValidate = Type.Object({
   shortid: Type.String(),
});

export const urlBodyValidate = Type.Object({
   url: Type.String({ format: "uri" }),
});
