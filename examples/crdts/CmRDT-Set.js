/**
 * CmRDT-Set
 *
 * Base Class for Operation-Based Set CRDT. Provides a Set interface.
 * 基于op的CRTD集合
 *
 * Operations are described as:
 *
 *   Operation = Tuple3(value : Any, added : Set, removed : Set)
 *
 * This class is meant to be used as a base class for
 * Operation-Based CRDTs that can be derived from Set
 * semantics and which calculate the state (values)
 * based on a set of operations.
 *
 *
 * Sources:
 * "A comprehensive study of Convergent and Commutative Replicated Data Types"
 * http://hal.upmc.fr/inria-00555588/document
 */

import { OperationTuple3 } from "./utils";

export class CmRDTSet extends Set {
  constructor(iterable, options) {
    super();
    this._values = new Set();
    this._operations = iterable ? iterable.map(OperationTuple3.from) : [];
    this._options = options || {};
  }

  values() {
    const shouldIncludeValue = (e) =>
      this._resolveValueState(e.added, e.removed, this._options.compareFunc);
    const getValue = (e) => e.value;

    const state = this._operations.filter(shouldIncludeValue).map(getValue);
    return new Set(state).values()
  }

  has(value) {
    return new Set(this.values()).has(value);
  }

  hasAll(values) {
    const contains = (e) => this.has(e);
    return values.every(contains);
  }

  add(value, tag) {
    // If the value is not in the set yet
    if (!this._values.has(value)) {
      // Create an operation for the value and apply it to this set
      const addOperation = OperationTuple3.create(value, [tag], null);
      this._applyOperation(addOperation);
    } else {
      // If the value is in the set, add a tag to its added set
      this._findOperationsFor(value).map((val) => val.added.add(tag));
    }
  }

  remove(value, tag) {
    this._findOperationsFor(value).map((e) => e.removed.add(tag));
  }

  merge(other) {
    other._operations.forEach((operation) => {
      const value = operation.value;
      if (!this._values.has(value)) {
        const op = OperationTuple3.create(
          value,
          operation.added,
          operation.removed
        );
        this._applyOperation(op);
      } else {
        this._findOperationsFor(value).map((op) => {
          operation.added.forEach((e) => op.added.add(e));
          operation.removed.forEach((e) => op.removed.add(e));
        });
      }
    });
  }

  toJSON() {
    const values = this._operations.map((e) => {
      return {
        value: e.value,
        added: Array.from(e.added),
        removed: Array.from(e.removed),
      };
    });
    return { values: values };
  }

  toArray() {
    return Array.from(this.values());
  }

  isEqual(other) {
    return CmRDTSet.isEqual(this, other);
  }

  _resolveValueState(added, removed, compareFunc) {
    // By default, if there's an add operation present,
    // and there are no remove operations, we include
    // the value in the set
    return added.size > 0 && removed.size === 0;
  }

  _applyOperation(operationTuple3) {
    this._operations.push(operationTuple3);
    this._values.add(operationTuple3.value);
  }

  _findOperationsFor(value) {
    let operations = [];
    if (this._values.has(value)) {
      const isForValue = (e) => e.value === value;
      const notNull = (e) => e !== undefined;
      operations = [this._operations.find(isForValue)].filter(notNull);
    }
    return operations;
  }

  static from(json) {
    return new CmRDTSet(json.values);
  }

  static isEqual(a, b) {
    return a.toArray().length === b.toArray().length && a.hasAll(b.toArray());
  }

  static difference(a, b) {
    const otherDoesntInclude = (x) => !b.has(x);
    const difference = new Set([...a.values()].filter(otherDoesntInclude));
    return difference;
  }
}
