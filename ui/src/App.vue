<template>
  <div id="app" class="container">
    <header-app/>
    <router-view></router-view>
  </div>
</template>

<script>
import menu from "./components/shared/menu"
export default {
  name: 'App',
  components: {
    headerApp :menu
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /*text-align: center;*/
  color: #2c3e50;
  margin-top: 60px;
}
</style>
