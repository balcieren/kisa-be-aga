import fastifyBuild from "../index";

test("GET `api/url/123` - Redirect to Url'", () => {
  const fastify = fastifyBuild();
  return fastify
    .inject()
    .get("/api/url/123")
    .end((err, res) => {
      expect(res.statusCode).toEqual(302);
    });
});

test("GET `api/url/123/status` - Get Url's Status'", () => {
  const fastify = fastifyBuild();
  return fastify
    .inject()
    .get("/api/url/123/status")
    .end((err, res) => {
      expect(res.statusCode).toEqual(200);
    });
});

test("POST `api/url` - Create Short Url'", () => {
  const fastify = fastifyBuild();
  return fastify
    .inject()
    .post("/api/url")
    .body({ url: "https://www.npmjs.com/package/fastify-prisma-client" })
    .end((err, res) => {
      expect(res.statusCode).toEqual(200);
      expect(res.body).toHaveProperty(["shortid"]);
    });
});
