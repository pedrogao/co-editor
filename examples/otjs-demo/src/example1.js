const { Changeset, transform } = require("changesets");

const operation0 = Changeset.create().insert("钢铁侠").end();
const operation1 = Changeset.create().insert("雷神").end();

console.log(operation0, operation1);
