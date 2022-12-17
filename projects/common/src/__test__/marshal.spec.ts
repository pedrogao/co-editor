import { marshal, unmarshal } from "../marshaler";
import * as Automerge from "@automerge/automerge";

test("marshal & unmarshal", () => {
  const doc = Automerge.init();
  const b1 = Automerge.save(doc);
  const s = marshal(b1);
  const b2 = unmarshal(s);
  expect(b1).toStrictEqual(b2);
});
