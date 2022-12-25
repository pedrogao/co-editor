import type { Quill } from "quill";

export interface CounterOptions {
  wordElement?: HTMLElement;
  wordUnit?: string;
  lineElement?: HTMLElement;
  lineUnit?: string;
}

export class Counter {
  private quill: Quill;
  private options: CounterOptions;

  constructor(quill: Quill, options: CounterOptions) {
    this.quill = quill;
    this.options = options;

    quill.on("text-change", this.update.bind(this));
  }

  // refer: https://juejin.cn/post/7029090052877582349
  update() {
    const { wordElement, wordUnit, lineElement, lineUnit } = this.options;
    if (!wordElement && !lineElement) {
      return;
    }

    let zhWordCount = 0,
      lineCount = 0;

    const text = this.quill.getText();
    for (let i = 0; i < text.length; i++) {
      const c = text.charAt(i);
      // 基本汉字
      if (c.match(/[\u4e00-\u9fa5]/)) {
        zhWordCount++;
      }
      // 基本汉字补充
      else if (c.match(/[\u9FA6-\u9fcb]/)) {
        zhWordCount++;
      } else if (c === "\n") {
        lineCount++;
      }
    }

    if (wordElement) {
      wordElement.innerText =
        zhWordCount + (wordUnit !== undefined ? wordUnit : "");
    }
    if (lineElement) {
      lineElement.innerText =
        lineCount + (lineUnit !== undefined ? lineUnit : "");
    }
  }
}
