import * as jdiff from "diff";
import { ADD, DELETE } from "./const";

class OldAction {
  constructor(position, infos) {
    this.position = position;
    this.infos = infos;
  }
}

/**
 * diff text
 * @param {string} oldStr
 * @param {string} newStr
 */
export const diff = (oldStr, newStr) => {
  const parts = jdiff.diffChars(oldStr, newStr);
  let index = 0,
    infoMaps = new Map();

  for (const p of parts) {
    let infos = infoMaps.get(index);

    if (p.added === true || p.removed === true) {
      switch (true) {
        case p.added:
          if (!infos) {
            infos = [];
            infoMaps.set(index, infos);
          }
          infos.push({ operate: ADD, chars: p.value });
          break;
        case p.removed:
          if (!infos) {
            infos = [];
            infoMaps.set(index + 1, infos);
          }
          infos.push({ operate: DELETE, chars: p.value });
          break;
        default:
      }
    }

    if (p.removed !== true) {
      index = index + p.count;
    }
  }
  const oldActions = [];
  for (let [position, infos] of infoMaps.entries()) {
    oldActions.push(new OldAction(position, infos));
  }
  return oldActions;
};
