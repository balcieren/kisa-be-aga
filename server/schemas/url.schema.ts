import { FastifySchema } from "fastify";
import { urlBody, urlParams } from "../dtos/url.dto";

export const redirectToShortUrlSchema: FastifySchema = { params: urlParams };

export const getShortUrlStatusSchema: FastifySchema = { params: urlParams };

export const createShortUrlSchema: FastifySchema = { body: urlBody };
