import { FastifySchemaCompiler } from "fastify";

const yupOptions = {
  strict: false,
  abortEarly: false,
  stripUnknown: true,
  recursive: true,
};

const schemaValidator: FastifySchemaCompiler<any> =
  ({ schema }) =>
  (data) => {
    try {
      const result = schema.validateSync(data, yupOptions);
      return { value: result };
    } catch (error) {
      return { error };
    }
  };

export default schemaValidator;
