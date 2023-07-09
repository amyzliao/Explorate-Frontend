<template>
    <div class="nav">
        <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
        <button class="logout" v-on:click="logout">Logout</button>
        <router-link to="/database">Database</router-link>
        <router-link to="/">About Us</router-link>
        <router-link to="/resources">NGO Partner</router-link>
        <router-link to="/about">NGO Signed In</router-link>
    </div>
</template>

<script>
import { userID, googleSignOut } from '@/firebase'
export default {
    name: 'HeaderModNgo',
    methods: {
        logout() {
            const googleUID = userID()

            if (googleUID) {
                googleSignOut()
            }

            localStorage.clear();
            this.$router.push({name: 'SignInPage'})
        }
    }
}
</script>

<style>
.nav {
    background-color: white;
    overflow: hidden;
}

.nav img {
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

.signinbutton:hover {
    background-color: #f1f1f1;
}

.signupbutton:hover {
    background-color: #414141;
}
</style>