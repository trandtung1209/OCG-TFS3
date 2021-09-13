<template>
  <div class="display">{{ result }}</div>
</template>

<script>
export default {
  name: "Screen",
  props: ["result"],
};
</script>

<style scoped>
.display {
  display: flex;
  flex: 1;
  padding: 3rem 2rem;
  font-size: 6rem;
  justify-content: flex-end;
  align-items: flex-end;
  cursor: default;
  overflow: hidden;
  max-width: 330px;
}
</style>
