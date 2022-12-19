/**
 * LWWSet-Set
 *
 * Operation-based Last-Write-Wins Set CRDT
 * 基于op的LWW(最后写胜出)的集合
 *
 * See base class CmRDT-Set.js for the rest of the API
 *
 * Sources:
 * "A comprehensive study of Convergent and Commutative Replicated Data Types"
 * http://hal.upmc.fr/inria-00555588/document, "Figure 8: LWW-Set (state-based)"
 */

import { CmRDTSet } from ".";

export class LWWSet extends CmRDTSet {
  _resolveValueState(added, removed, compareFunc) {
    // Sort both sets with the given comparison function
    // or use "distance" sort by default
    compareFunc = compareFunc ? compareFunc : (a, b) => (a || 0) - (b || 0);
    const sortedAdded = Array.from(added).sort(compareFunc).reverse();
    const sortedRemoved = Array.from(removed).sort(compareFunc).reverse();
    // If the latest add operation is greater or equal than latest remove operation,
    // we include it in the state
    return compareFunc(sortedAdded[0], sortedRemoved[0]) > -1;
  }

  static from(json) {
    return new LWWSet(json.values);
  }
}
