<template>
  <router-view />
</template>

<script>
import '@/assets/css/main.css' // (@ in vue points to src/)

export default {
  name: 'App',
  components: {

}
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Quicksand:wght@300;400;500;600;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap');
@font-face {
  font-family: 'Quicksand', san-serif;
  font-family: 'Roboto', san-serif;
  src: url('https://fonts.googleapis.com/css2?family=Quicksand:wght@300;400;500;600;700&display=swap');
}

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  font-family: "Quicksand";
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

.home, .resource, .resource-2 {
    padding-top: 50px;
    overflow: hidden;
    margin-bottom: 20px;
    padding-bottom: 50px;
}

.home-left, .resource-left, .resource-left-2 {
    float: left;
    width: 50%;
    text-align: center;
    color: #383b3d;
}

.home-left h1, .resource-left h1 {
    font-size: 55px;
    font-weight: normal;
    text-align: left;
    padding-bottom: 10px;
    margin-left: 100px;
}

.home-left p, .resource-left p {
    text-align: left;
    padding-bottom: 20px;
    font-size: 20px;
    margin-left: 100px;
}

.home-right, .resource-right {
    float: right;
    width: 50%;
    background-color: green;
    margin: auto;
}

.home-left button, .resource-left button {
    height: 55px;
    width: 240px;
    border-radius: 6px;
    font-family: "Inter";
    font-size: 14px;
    font-weight: bold;
    border: 2px solid #e7513f;
    float: left;
    cursor: pointer;
    margin-left: 100px;
    margin-bottom: 30px;
}

.find, .connect {
    background-color: #e7513f;
    color: white;
}

.logout {
    float: right;
    text-decoration: none;
    padding: 20px;
    background-color: #212121;
    color: white;
    margin-left: 15px;
    margin-right: 20px;
}
</style>
