import Fastify, { FastifyInstance } from "fastify";
import fastifySensible from "fastify-sensible";
import fastifyPrisma from "fastify-prisma-client";
import urlRoute from "./routes/url.route";
import schemaValidator from "./middlewares/validator.middleware";

const buildFastify = () => {
  const fastify: FastifyInstance = Fastify({
    logger: { level: "info", prettyPrint: true },
  });
  fastify
    .register(fastifyPrisma)
    .register(fastifySensible)
    .register(urlRoute, { prefix: "/api/url" })
    .setValidatorCompiler(schemaValidator);

  fastify.get("/", async () => {
    return "KISA BE AGA !";
  });

  return fastify;
};

buildFastify().listen(3000);

export default buildFastify;
