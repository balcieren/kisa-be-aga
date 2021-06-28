import { FastifyPluginCallback } from "fastify";
import urlController from "../controllers/url.controller";
import {
  createShortUrlSchema,
  getShortUrlStatusSchema,
  redirectToShortUrlSchema,
} from "../schemas/url.schema";

const urlRoute: FastifyPluginCallback = (fastify, options, done) => {
  fastify
    .get(
      "/:shortid",
      { schema: redirectToShortUrlSchema },
      urlController.redirectToShortUrl
    )
    .get(
      "/:shortid/status",
      { schema: getShortUrlStatusSchema },
      urlController.getShortUrlStatus
    )
    .post("/", { schema: createShortUrlSchema }, urlController.createShortUrl);
  done();
};

export default urlRoute;
