import { FastifyPluginAsync } from "fastify";

const routes: FastifyPluginAsync = async (fastify) => {
   fastify.register(import("./modules/url"), { prefix: "/url" });
};

export default routes;
