import { FastifySchema } from "fastify";
import { urlBodyValidate, urlParamsValidate } from "./validates";

export const redirectToShortUrlSchema: FastifySchema = { params: urlParamsValidate };

export const getShortUrlStatusSchema: FastifySchema = { params: urlParamsValidate };

export const createShortUrlSchema: FastifySchema = { body: urlBodyValidate };
