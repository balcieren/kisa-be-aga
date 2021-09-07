import { FastifySchema } from "fastify";
import { urlBodyValidate, urlParamsValidate } from "./validates";

export const redirectToShortUrlSchema: FastifySchema = { tags: ["url"], params: urlParamsValidate };

export const getShortUrlStatusSchema: FastifySchema = { tags: ["url"], params: urlParamsValidate };

export const createShortUrlSchema: FastifySchema = { tags: ["url"], body: urlBodyValidate };
