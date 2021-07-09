import { UrlBodyType, UrlParamsType } from "../dtos/url.dto";
import { RouteHandler } from "fastify";
import shortid from "shortid";

const redirectToShortUrl: RouteHandler<{ Params: UrlParamsType }> = async (
  request,
  reply
) => {
  const shortId = String(request.params.shortid);

  const found = await request.prisma.url.findUnique({
    where: { shortId },
    select: { url: true },
  });

  if (!found?.url) return reply.notFound("Url has not found !");

  await request.prisma.url.update({
    where: { shortId },
    data: { clicks: { increment: 1 } },
  });

  return reply.redirect(found.url);
};

const getShortUrlStatus: RouteHandler<{ Params: UrlParamsType }> = async (
  request,
  reply
) => {
  const shortId = String(request.params.shortid);

  const found = await request.prisma.url.findUnique({
    where: { shortId },
  });

  if (!found) return reply.notFound("Url has not found !");

  return reply.send(found);
};

const createShortUrl: RouteHandler<{ Body: UrlBodyType }> = async (
  request,
  reply
) => {
  const url = String(request.body.url);

  if (!url) return reply.badRequest("Url must be in body !");

  const newShortUrl = await request.prisma.url.create({
    data: { url, shortId: shortid.generate() },
  });

  return reply.send({ shortId: newShortUrl.shortId });
};

export default { redirectToShortUrl, getShortUrlStatus, createShortUrl };
