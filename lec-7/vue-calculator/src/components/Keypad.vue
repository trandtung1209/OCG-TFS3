<template>
  <div class="keypad">
    <div class="key key--grey clear">AC</div>
    <div class="key key--grey del">DEL</div>
    <div class="key key--grey operator">%</div>
    <div class="key key--yellow operator">÷</div>
    <div class="key number">7</div>
    <div class="key number">8</div>
    <div class="key number">9</div>
    <div class="key key--yellow operator">×</div>
    <div class="key number">4</div>
    <div class="key number">5</div>
    <div class="key number">6</div>
    <div class="key key--yellow operator">-</div>
    <div class="key number">1</div>
    <div class="key number">2</div>
    <div class="key number">3</div>
    <div class="key key--yellow operator">+</div>
    <div class="key key--long number">0</div>
    <div class="key number dot">.</div>
    <div class="key key--yellow equal">=</div>
  </div>
</template>

<script>
import Button from "./Button.vue";

export default {
  name: "Keypad",
  component: { Button },
};
</script>

<style scoped>
.keypad {
  display: grid;
  grid-template-columns: repeat(4, 70px);
  grid-template-rows: repeat(5, 70px);
  padding: 1rem;
  grid-gap: 1rem;
}

.key {
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;
  background-color: #505050;
}

.key:hover {
  cursor: pointer;
  filter: brightness(1.3);
}

.key--long {
  grid-column: span 2;
  border-radius: 5rem;
}

.key--grey {
  background-color: #d4d4d2;
  color: #1c1c1c;
}

.key--yellow {
  background-color: #ff9500;
}

.operator--active {
  color: #ff9500;
  background-color: #fff;
}
</style>
