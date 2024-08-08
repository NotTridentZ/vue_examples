<template>
  <div>
    <div>
      <div v-for="message in messages" :key="message.id">
        {{ message.text }}
      </div>
    </div>
    <input
      v-model="newMessage"
      @keyup.enter="sendMessage"
      placeholder="Type a message"
    />
  </div>
</template>

<script>
import axios from "axios";
import io from "socket.io-client";

export default {
  data() {
    return {
      messages: [],
      newMessage: "",
      socket: io("http://localhost:8080"),
    };
  },
  created() {
    this.fetchMessages();
    this.socket.on("newMessage", this.addMessage);
  },
  methods: {
    async fetchMessages() {
      const response = await axios.get("http://localhost:8080/messages");
      this.messages = response.data;
    },
    async sendMessage() {
      await axios.post("http://localhost:8080/messages", {
        text: this.newMessage,
      });
      this.newMessage = "";
    },
    addMessage(message) {
      this.messages.push(message);
    },
  },
};
</script>