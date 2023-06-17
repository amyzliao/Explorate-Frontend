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
@font-face {
  font-family: Roboto, Helvetica, Arial, sans-serif;
  src: local("Inter"), url(./assets/fonts/Inter-VariableFont_slnt\,wght.ttf) format("truetype");
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  font-family: "Inter";
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

.signup, .signin {
    background-color: gray;
    width: 500px;
    margin: auto;
    padding-bottom: 20px;
}

.signup input, .signin input {
    width: 300px;
    height: 40px;
    padding-left: 20px;
    display: block;
    margin-bottom: 30px;
    margin-right: auto;
    margin-left: auto;
}

.signup button, .signin button {
    height: 40px;
    cursor: pointer;
    border: none;
    display: block;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 30px;
    width: 326px;
    font-size: 16px;
    font-family: "Inter";
    margin-top: 10px;
}

.signuppagebutton, .signinpagebutton {
    border-radius: 20px;
    background-color: #e7513f;
    color: white;
}

.google {
    background-color: white;
    border-radius: 10px;

}

.signuppagesigninbutton, .signinpagesignupbutton {
    border-radius: 20px;
    background-color: white;
}
</style>
