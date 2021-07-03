import Fastify, { FastifyInstance } from "fastify";
import fastifyCors from "fastify-cors";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "fastify-prisma-client";
import { fastifyYupSchema } from "fastify-yup-schema";
import urlRoute from "./routes/url.route";

const buildFastify = () => {
  const fastify: FastifyInstance = Fastify({
    logger: { level: "info", prettyPrint: true },
  });

  fastify
    .register(fastifyCors)
    .register(fastifyPrisma)
    .register(fastifyYupSchema)
    .register(fastifySensible)
    .register(urlRoute, { prefix: "/api/url" });

  fastify.get("/", async () => {
    return "KISA BE AGA !";
  });

  return fastify;
};

buildFastify().listen(3001);

export default buildFastify;
