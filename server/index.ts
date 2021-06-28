import Fastify, { FastifyInstance } from "fastify";
import { PrismaClient } from "@prisma/client";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "./plugins/prisma.plugin";
import urlRoute from "./routes/url.route";
import schemaValidator from "./middlewares/validator.middleware";

const fastify: FastifyInstance = Fastify({ logger: { prettyPrint: true } });

fastify
  .register(fastifyPrisma)
  .register(fastifySensible)
  .register(urlRoute, { prefix: "/api/url" })
  .setValidatorCompiler(schemaValidator);

fastify.get("/", async () => {
  return "KISA BE AGA 🤣";
});

fastify.listen(3000).catch((err) => fastify.log.error(err));

declare module "fastify" {
  interface FastifyRequest {
    prisma: PrismaClient;
  }
}
