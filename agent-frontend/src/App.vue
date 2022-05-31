<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {

  },
  mounted() {
    this.$vs.setColor('warn', '#FFFFFF')
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.vs-input {
  width: 500px !important;
}

h1 {
  color: white;
}

.vs-navbar__item {
  font-size: 1rem !important;
}
</style>
