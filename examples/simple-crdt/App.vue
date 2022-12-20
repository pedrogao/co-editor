<template>
  <span>Multiline message is:</span>
  <textarea v-model="message" @input="onChange" class="input-area" placeholder="add multiple lines"></textarea>
  <button @click="draw">绘制</button>
  <button @click="clear">清除</button>
</template>
  
<script>
import { defineComponent } from "vue";
import { ClientState, socket } from './client';
import { loadActionsInfo } from './lib/dt';
import { draw, clear } from './tree';

socket.on('connect', onConnect);

function onConnect() {
  console.log('connect ' + socket.id)
}

export default defineComponent({
  name: 'app',
  created() {
    socket.on('server-actions', (actionsInfo) => {
      const actions = loadActionsInfo(actionsInfo)
      for (const action of actions) {
        this.clientState.serverAddAction(action)
      }
      this.updateMessage()
      draw(this.clientState.tree)
    })
    socket.on('client-init', (info) => {
      console.log('client-init')
      const { uid, timeStamp, actionsInfo } = info
      this.clientState = new ClientState(timeStamp, uid)
      window.clientState = this.clientState
      const actions = loadActionsInfo(actionsInfo)
      for (const action of actions) {
        this.clientState.serverAddAction(action)
      }
      this.updateMessage()
      draw(this.clientState.tree)
    })
  },
  data() {
    return {
      start: 0,
      end: 0,
      clientState: null,
      message: ''
    }
  },
  methods: {
    updateMessage() {
      const input = document.getElementsByClassName('input-area')[0]
      const start = input && input.selectionStart
      const end = input && input.selectionEnd
      this.message = this.clientState.toString()
      this.$nextTick(() => {
        input.setSelectionRange(start, end)
      });
    },
    draw() {
      clear()
      draw(this.clientState.tree)
    },
    clear() {
      clear()
    },
    onChange(value) {
      console.log('value: ', value);
      this.clientState.update(this.message)
      draw(this.clientState.tree)
      this.updateMessage()
    }
  },
})
</script>
  
<style scoped>
.input-area {
  display: block;
  width: 80%;
  height: 100px;
  font-size: 1rem;
  margin: auto;
  line-height: 1.5;
  color: #495057;
  background-color: #fff;
  background-clip: padding-box;
  border: 1px solid #ced4da;
}

h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>