import { FastifySchema } from "fastify";
import * as yup from "yup";

export const redirectToShortUrlSchema: FastifySchema = {
  params: yup.object({ shortid: yup.string().required() }).required(),
};

export const getShortUrlStatusSchema: FastifySchema = {
  params: yup.object({ shortid: yup.string().required() }).required(),
};

export const createShortUrlSchema: FastifySchema = {
  body: yup.object({ url: yup.string().url().required() }).required(),
};
