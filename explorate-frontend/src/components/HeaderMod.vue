<template>
    <div class="nav">
        <button class="signup">Sign Up</button>
        <button class="signin">Sign In</button>
        <router-link to="/database">Database (temp)</router-link>
        <router-link to="/">Volunteer</router-link>
        <router-link to="/about">NGO Partner</router-link>
        <router-link to="/resources">About Us</router-link>
        <router-link to="/">Home (temp)</router-link>
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
    padding: 10px;
    overflow: hidden;
}
.nav a {
    float: right;
    text-decoration: none;
    color: #212121;
    padding: 10px;
    margin-left: 5px;
}
.nav a:hover {
    background-color: orangered;
}

.nav button {
    float: right;
    text-decoration: none;
    padding-top: 10px;
    padding-bottom: 10px;
    border: 1px solid #212121;
    border-radius: 10px;
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
    margin-left: 8px;
}
</style>