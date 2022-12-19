export class LamportClock {
  constructor(id, time) {
    this.id = id;
    this.time = time || 0;
  }

  tick() {
    return new LamportClock(this.id, this.time + 1);
  }

  merge(clock) {
    this.time = Math.max(this.time, clock.time);
    return new LamportClock(this.id, this.time);
  }

  clone() {
    return new LamportClock(this.id, this.time);
  }

  static isEqual(a, b) {
    return a.id == b.id && a.time == b.time;
  }

  static compare(a, b) {
    if (!a || !a.time) a = { time: 0 };
    if (!b || !b.time) b = { time: 0 };

    // time优先
    var dist = a.time - b.time;
    // 如果 time一致，那么 id大的则优先
    if (dist === 0 && a.id !== b.id) return a.id < b.id ? -1 : 1;

    return dist;
  }
}
