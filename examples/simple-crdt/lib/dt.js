import { diff } from "./util";
import { ADD, DELETE, OPERATE_MAP } from "./const";

export class TimestampManager {
  constructor(timestamp = 0) {
    this.timestamp = timestamp;
  }

  getTimestamp() {
    return ++this.timestamp;
  }

  setTimestamp(timestamp) {
    this.timestamp = timestamp;
  }

  updateTimestamp(newTimestamp) {
    this.timestamp = Math.max(newTimestamp, this.timestamp);
  }
}

export class IdManager {
  constructor() {
    this.id = 0;
  }

  getId() {
    return ++this.id;
  }
}

export class Id {
  constructor(uid, timestamp) {
    this.uid = uid;
    this.timestamp = timestamp;
  }

  /**
   * @param {Id} id
   */
  bigger(id) {
    if (this.timestamp > id.timestamp) {
      return true;
    }
    if (this.timestamp === id.timestamp && this.uid > id.uid) {
      return true;
    }

    return false;
  }

  small(id) {
    return !this.bigger(id);
  }

  /**
   * @param {Id} id
   */
  equal(id) {
    return this.timestamp === id.timestamp && this.uid === id.uid;
  }

  toString() {
    return `U${this.uid}@T${this.timestamp}`;
  }

  toJSON() {
    const { uid, timestamp } = this;
    return { uid, timestamp };
  }
}

export class Action {
  constructor(preId, id, operate, value) {
    this.preId = preId;
    this.id = id;
    this.operate = operate;
    this.value = value;
  }

  toString() {
    return `${this.id.toString()}\t ${
      OPERATE_MAP[this.operate]
    }\t ${this.preId.toString()}\t ${this.value}`;
  }

  toJSON() {
    const { preId, id, operate, value } = this;
    return {
      operate,
      preId,
      value,
      id,
    };
  }
}

export class Word {
  constructor(id, value, preId, operate) {
    this.id = id;
    this.isDelete = false;
    this.preId = preId;
    this.operate = operate;
    this.value = value;
  }

  delete() {
    this.isDelete = true;
  }
}

export class Info {
  constructor(operate, chars) {
    this.operate = operate;
    this.chars = chars;
  }
}

export class Node {
  constructor(word) {
    this.word = word;
    this.nextNode = [];
  }

  addNext(node) {
    if (!this.hasNext(node)) {
      this.nextNode = this.nextNode.sort((a, b) =>
        a.word.id.bigger(b.word.id) ? 1 : -1
      );
      return true;
    }
    return false;
  }

  toJSON() {
    const { id, preId, operate, value } = this.word;
    return {
      id,
      preId,
      operate,
      value,
    };
  }

  hasNext(node) {
    for (const n of this.nextNode) {
      if (n.word.id.equal(node.word.id)) return true;
    }
    return false;
  }

  getNodeByposition(position) {
    if (position <= 0) {
      return this.word.id;
    }
    for (let n of this.nextNode) {
      if (!n.word.isDelete) {
        position = position - 1;
      }
      const re = n.getNodeByposition(position);
      if (re !== null) return re;
    }
    return null;
  }

  getStr() {
    const str = this.word.isDelete ? "" : this.word.value;
    if (this.nextNode.length === 0) {
      return str;
    }
    return str + this.nextNode.map((n) => n.getStr()).join("");
  }
}

export class State {
  constructor(timestamp, uid) {
    this.timestampManager = new TimestampManager(timestamp);
    this.uid = uid;
    this.notes = new Map();
    const id = new Id(0, 0);
    const word = new Word(id, "", id, ADD);
    this.tree = new Node(word); // 因果树
    this.notes.set(id.toString(), this.tree);
  }

  createId() {
    return new Id(this.uid, this.timestampManager.getTimestamp());
  }

  addOldActions(oldActions) {
    const actions = [];
    let preId = null;
    for (let oldAction of oldActions) {
      for (let info of oldAction.infos) {
        let i = 0;
        const chars = info.chars.split("");
        if (info.operate === DELETE) {
          i = chars.length - 1;
          chars.reverse();
        }
        for (let char of chars) {
          preId = this.getPreId(oldAction.position + 1);
          if (info.operate === DELETE) {
            i--;
          } else {
            i++;
          }
          const id = this.createId();
          const action = new Action(preId, id, info.operate, char);
          this.addAction(action);
          actions.push(action);
        }
      }
    }
    return actions;
  }

  update(newStr) {
    const oldStr = this.toString();
    const oldActions = diff(oldStr, newStr);
    this.addOldActions(oldActions);
    return this.toString();
  }

  toString() {
    return this.tree.getStr();
  }

  getPreId(position) {
    return this.tree.getNodeByposition(position);
  }

  addAction(action) {
    if (this.notes.get(action.id.toString())) {
      return;
    }

    console.log("action: ", action.toString());
    const node = this.notes.get(action.preId.toString());
    switch (node.operate) {
      case ADD:
        const id = action.id;
        const word = new Word(
          action.id,
          action.value,
          action.preId,
          action.operate
        );
        const newNode = new Node(word);
        this.notes.set(id.toString(), newNode);
        node.addNext(newNode); // 后续节点
        break;
      case DELETE:
        node.word.delete();
        break;
      default:
        break;
    }
    this.timestampManager.updateTimestamp(action.id.timestamp);
  }
}

/**
 * @param {Id} id
 */
const loadId = (id) => {
  if (!id) {
    return null;
  }

  const { uid, timestamp } = id;
  return new Id(uid, timestamp);
};

export const loadActionsInfo = (actionsInfo) => {
  return JSON.parse(actionsInfo).map((action) => {
    const id = loadId(action.id);
    const preId = loadId(action.preId);
    const { operate, value } = action;
    return new Action(preId, id, operate, value);
  });
};
