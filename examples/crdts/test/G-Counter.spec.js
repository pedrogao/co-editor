import { GCounter } from "..";

describe("value", () => {
  test("returns the count", () => {
    const counter = new GCounter("A");
    expect(counter.value).toBe(0);
  });

  test("increments the count by 42", () => {
    const counter = new GCounter("A");
    counter.increment();
    counter.increment(42);
    expect(counter.value).toBe(43);
  });
});

describe("merge", () => {
  test("merges two counters with same id", () => {
    const counter1 = new GCounter("A");
    const counter2 = new GCounter("A");
    counter1.increment();
    counter2.increment();
    counter1.merge(counter2);
    expect(counter1.value).toBe(1);
  });

  test("merges two counters with same values", () => {
    const counter1 = new GCounter("A");
    const counter2 = new GCounter("B");
    counter1.increment();
    counter2.increment();
    counter1.merge(counter2);
    counter2.merge(counter1);
    expect(counter1.value).toBe(2);
    expect(counter2.value).toBe(2);
  });

  test("merges four different counters", () => {
    const counter1 = new GCounter("A");
    const counter2 = new GCounter("B");
    const counter3 = new GCounter("C");
    const counter4 = new GCounter("D");
    counter1.increment();
    counter2.increment();
    counter3.increment();
    counter4.increment();
    counter1.merge(counter2);
    counter1.merge(counter3);
    counter1.merge(counter4);
    expect(counter1.value).toBe(4);
  });

  test("doesn't overwrite its own value on merge", () => {
    const counter1 = new GCounter("A");
    const counter2 = new GCounter("B");
    counter1.increment();
    counter2.increment();
    counter1.merge(counter2);
    counter2.merge(counter1);
    counter1.increment();
    counter1.merge(counter2);
    expect(counter1.value).toBe(3);
  });

  test("doesn't overwrite others' values on merge", () => {
    const counter1 = new GCounter("A");
    const counter2 = new GCounter("B");
    counter1.increment();
    counter2.increment();
    counter1.merge(counter2);
    counter2.merge(counter1);
    counter1.increment();
    counter2.increment();
    counter1.merge(counter2);
    expect(counter1.value).toBe(4);
  });
});
