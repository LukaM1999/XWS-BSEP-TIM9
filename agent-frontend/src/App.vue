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

[type="search"] {
  width: 500px !important;
}

.vs-input [type="text"] {
  width: 500px !important;
}

.vs-input .subject {
  width: 100% !important;
}

.vs-select {
  width: 200px !important;
}

.vs-input .user{
  width: 100% !important;
}

h1 {
  color: white;
}

.vs-navbar__item {
  font-size: 1rem !important;
}

.vs-card {
  max-width: 100% !important;
}
</style>
