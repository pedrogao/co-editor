import type { Quill } from "quill";

export interface CounterOptions {
  ele: HTMLElement;
  unit: string;
}

export class Counter {
  private quill: Quill;
  private options: CounterOptions;

  constructor(quill: Quill, options: CounterOptions) {
    this.quill = quill;
    this.options = options;

    quill.on("text-change", this.update.bind(this));
  }

  calculate() {
    let text = this.quill.getText();
    if (this.options.unit === "word") {
      text = text.trim();
      // Splitting empty text returns a non-empty array
      return text.length > 0 ? text.split(/\s+/).length : 0;
    } else {
      return text.length;
    }
  }

  update() {
    const length = this.calculate();
    const container = this.options.ele;
    container.innerText = length + "";
  }
}
