import Fastify, { FastifyInstance } from "fastify";
import fastifyCors from "fastify-cors";
import fastifySwagger from "fastify-swagger";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "fastify-prisma-client";
import urlRoute from "./routes/url.route";

const buildFastify = () => {
  const fastify: FastifyInstance = Fastify({
    logger: { level: "info", prettyPrint: true },
  });

  fastify
    .register(fastifyCors)
    .register(fastifySwagger, {
      exposeRoute: true,
      routePrefix: "/docs",
      staticCSP: true,
      swagger: {
        info: {
          title: "KISA BE AGA",
          description: "Testing the KISA BE AGA swagger API",
          version: "0.1.0",
        },
      },
    })
    .register(fastifyPrisma, { logLevel: "debug" })
    .register(fastifySensible)
    .register(urlRoute, { prefix: "/api/url" });

  fastify.get("/", async (request, reply) => {
    reply.send({ message: "KISA BE AGA !" });
  });

  return fastify;
};

buildFastify().listen(3001);

export default buildFastify;
