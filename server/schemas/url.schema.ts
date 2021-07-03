import { createYupSchema } from "fastify-yup-schema";

export const redirectToShortUrlSchema = createYupSchema((yup) => ({
  params: yup.object({ shortid: yup.string().required() }).required(),
}));

export const getShortUrlStatusSchema = createYupSchema((yup) => ({
  params: yup.object({ shortid: yup.string().required() }).required(),
}));

export const createShortUrlSchema = createYupSchema((yup) => ({
  body: yup.object({ url: yup.string().url().required() }).required(),
}));
