import { FastifyPluginAsync } from "fastify";
import {
   createShortUrlHandler,
   getShortUrlStatusHandler,
   redirectToShortUrlHandler,
} from "./handlers";
import { createShortUrlSchema, getShortUrlStatusSchema, redirectToShortUrlSchema } from "./schemas";

const urlRoute: FastifyPluginAsync = async (fastify) => {
   fastify.get("/:shortid", { schema: redirectToShortUrlSchema }, redirectToShortUrlHandler);
   fastify.get("/:shortid/status", { schema: getShortUrlStatusSchema }, getShortUrlStatusHandler);
   fastify.post("/", { schema: createShortUrlSchema }, createShortUrlHandler);
};

export default urlRoute;
