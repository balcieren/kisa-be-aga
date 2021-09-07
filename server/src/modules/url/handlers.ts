import { RouteHandler } from "fastify";
import shortid from "shortid";
import { UrlBody, UrlParams } from "./dto";

type RedirectToShortUrl = {
   Params: UrlParams;
};
export const redirectToShortUrlHandler: RouteHandler<RedirectToShortUrl> = async function (
   request,
   reply
) {
   const shortId = String(request.params.shortid);

   const found = await this.prisma.url.findUnique({
      where: { shortId },
      select: { url: true },
   });

   if (!found?.url) return reply.notFound("Url has not found !");

   await this.prisma.url.update({
      where: { shortId },
      data: { clicks: { increment: 1 } },
   });

   return reply.redirect(found.url);
};

type GetShortUrlStatus = {
   Params: UrlParams;
};
export const getShortUrlStatusHandler: RouteHandler<GetShortUrlStatus> = async function (
   request,
   reply
) {
   const shortId = String(request.params.shortid);

   const found = await this.prisma.url.findUnique({
      where: { shortId },
   });

   if (!found) return reply.notFound("Url has not found !");

   return reply.send(found);
};

type CreateShortUrl = { Body: UrlBody };
export const createShortUrlHandler: RouteHandler<CreateShortUrl> = async function (request, reply) {
   const url = String(request.body.url);

   if (!url) return reply.badRequest("Url must be in body !");

   const newShortUrl = await this.prisma.url.create({
      data: { url, shortId: shortid.generate() },
   });

   return reply.send({ shortId: newShortUrl.shortId });
};
