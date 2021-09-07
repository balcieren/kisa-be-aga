import Fastify, { FastifyInstance } from "fastify";

const buildFastify = () => {
   const fastify: FastifyInstance = Fastify({
      logger: { prettyPrint: true },
   });

   fastify.register(import("./plugins"));

   fastify.register(import("./routes"), { prefix: "/api" });

   return fastify;
};

buildFastify().listen(1923);

export default buildFastify;
