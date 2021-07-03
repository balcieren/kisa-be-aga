import Fastify, { FastifyInstance } from "fastify";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "fastify-prisma-client";
import fastifyCors from "fastify-cors";
import urlRoute from "./routes/url.route";
import schemaValidator from "./middlewares/validator.middleware";

const buildFastify = () => {
  const fastify: FastifyInstance = Fastify({
    logger: { level: "info", prettyPrint: true },
  });
  fastify
    .register(fastifyCors)
    .register(fastifyPrisma)
    .register(fastifySensible)
    .register(urlRoute, { prefix: "/api/url" })
    .setValidatorCompiler(schemaValidator);

  fastify.get("/", async () => {
    return "KISA BE AGA !";
  });

  return fastify;
};

buildFastify().listen(3001);

export default buildFastify;
