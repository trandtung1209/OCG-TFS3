<template>
  <div class="container">
    <Screen :result="msg"></Screen>
    <Keypad></Keypad>
  </div>
</template>

<script>
import Screen from "./components/Screen.vue";
import Keypad from "./components/Keypad.vue";

export default {
  name: "App",
  components: {
    Screen,
    Keypad,
  },
  data() {
    return {
      msg: "123",
    };
  },
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html {
  font-size: 62.5%;
}

body {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(-180deg, #bcc5ce 0%, #929ead 98%),
    radial-gradient(
      at top left,
      rgba(255, 255, 255, 0.3) 0%,
      rgba(0, 0, 0, 0.3) 100%
    );
  background-blend-mode: screen;
}

.container {
  border-radius: 2rem;
  color: #fff;
  display: flex;
  flex-direction: column;
  background-color: #1c1c1c;
  font-family: "Poppins", sans-serif;
  font-size: 2.2rem;
  font-weight: 300;
  box-shadow: 5px 12px 20px rgb(36 37 38 / 13%);
}
</style>
