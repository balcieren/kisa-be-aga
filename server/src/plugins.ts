import { FastifyPluginAsync } from "fastify";
import fp from "fastify-plugin";

const plugins: FastifyPluginAsync = async (fastify) => {
   fastify.register(import("fastify-cors"));
   fastify.register(import("fastify-swagger"), {
      exposeRoute: true,
      routePrefix: "/",
      openapi: {
         info: {
            title: "KISA BE AGA",
            description: "API document for kısa be aga",
            version: "0.0.1",
         },
      },
   });
   fastify.register(import("fastify-prisma-client"));
   fastify.register(import("fastify-sensible"));
};

export default fp(plugins);
