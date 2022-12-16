// eslint-disable-next-line @typescript-eslint/no-var-requires
const consola = require('consola');

export const logger = consola.create({
  // level: 4,
  reporters: [new consola.JSONReporter()],
  defaults: {},
});
