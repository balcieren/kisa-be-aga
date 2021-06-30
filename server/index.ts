import Fastify, { FastifyInstance } from "fastify";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "fastify-prisma-client";
import urlRoute from "./routes/url.route";
import schemaValidator from "./middlewares/validator.middleware";

const fastify: FastifyInstance = Fastify({ logger: { prettyPrint: true } });

fastify
  .register(fastifyPrisma)
  .register(fastifySensible)
  .register(urlRoute, { prefix: "/api/url" })
  .setValidatorCompiler(schemaValidator);

fastify.get("/", async () => {
  return "hello";
});

fastify.listen(3000).catch((err) => fastify.log.error(err));
