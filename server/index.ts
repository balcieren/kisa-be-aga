import Fastify, { FastifyInstance } from "fastify";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "./plugins/prisma.plugin";
import urlRoute from "./routes/url.route";
import { PrismaClient } from "@prisma/client";

const fastify: FastifyInstance = Fastify({ logger: { prettyPrint: true } });

fastify
  .register(fastifyPrisma)
  .register(fastifySensible)
  .register(urlRoute, { prefix: "/api/url" });

fastify.get("/", async () => {
  return "KISA BE AGA 🤣";
});

fastify.listen(3000).catch((err) => {
  fastify.log.error(err);
});

declare module "fastify" {
  interface FastifyRequest {
    prisma: PrismaClient;
  }
}
