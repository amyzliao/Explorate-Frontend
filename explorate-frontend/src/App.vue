<template>
  <router-view />
  <!-- <TestPage/>
  <HomePage/>
  <DatabasePage/>
  <ResourcePage/>
  <AboutPage/> -->
</template>

<script>
// import HomePage from './components/HomePage.vue'
// import TestPage from './components/TestPage.vue'
// import AboutPage from './components/AboutPage.vue'
// import DatabasePage from './components/DatabasePage.vue'
// import ResourcePage from './components/ResourcePage.vue'

export default {
  name: 'App',
  components: {
    // HomePage
    // AboutPage,
    // TestPage,
    // DatabasePage,
    // ResourcePage
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
}
body {
  padding:0;
  margin:0;
}
.form {
  padding: 20px;
}
.form input {
  width: 300px;
  height: 40px;
  padding-left: 20px;
  display: block;
  margin-bottom: 30px;
  margin-right: auto;
  margin-left: auto;
  border: 2px solid orangered;
}
.form button {
  width: 320px;
  height: 40px;
  border: 2px solid orangered;
  background: orangered;
  color: white;
  cursor: pointer;
}
</style>
