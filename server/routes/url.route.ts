import { FastifyPluginCallback } from "fastify";
import urlController from "../controllers/url.controller";

const urlRoute: FastifyPluginCallback = (fastify, options, done) => {
  fastify
    .get("/:shortid", {}, urlController.redirectToShortUrl)
    .get("/:shortid/status", {}, urlController.getShortUrlStatus)
    .post("/", {}, urlController.createShortUrl);
  done();
};

export default urlRoute;
