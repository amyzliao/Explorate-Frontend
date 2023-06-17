<template>
    <div class="nav">
        <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
        <button class="signup">Sign Up</button>
        <button class="signin">Sign In</button>
        <router-link to="/database">Database (temp)</router-link>
        <!-- <router-link to="/about">About Us</router-link> -->
        <router-link to="/">About Us</router-link>
        <router-link to="/resources">NGO Partner</router-link>
        <router-link to="/">Volunteer</router-link>
    </div>
</template>

<script>
export default {
    name: 'HeaderMod',
}
</script>

<style>
.nav {
    background-color: white;
    overflow: hidden;
}

.logo {
    height: 50px;
    margin-left: 10px;
    margin-top: 10px;
    float: left;
}

.nav a {
    float: right;
    text-decoration: none;
    color: #212121;
    padding: 20px;
}

.nav a:hover {
    background-color: #f1f1f1;
}

.nav button {
    float: right;
    text-decoration: none;
    padding-top: 10px;
    padding-bottom: 10px;
    margin-top: 10px;
    margin-bottom: 10px;
    border: 1px solid #212121;
    border-radius: 10px;
    font-family: "Inter";
}

.nav button:hover {
    cursor: pointer;
}

.signin:hover {
    background-color: #f1f1f1;
}

.signup:hover {
    background-color: #414141;
}

.signin {
    float: right;
    text-decoration: none;
    padding: 20px;
    background-color: white;
    color: #212121;
    margin-left: 20px;
}

.signup {
    float: right;
    text-decoration: none;
    padding: 20px;
    background-color: #212121;
    color: white;
    margin-left: 15px;
    margin-right: 20px;
}
</style>