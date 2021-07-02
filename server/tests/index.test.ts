import fastifyBuild from "../index";

test("GET `/` - Index'", () => {
  const fastify = fastifyBuild();
  return fastify
    .inject()
    .get("/")
    .end((err, res) => {
      expect(res.statusCode).toEqual(200);
      expect(res.body).toBe("KISA BE AGA !");
    });
});
