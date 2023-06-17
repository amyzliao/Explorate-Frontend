<template>
    <HeaderMod />
    <div class="signin">
        <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
        <h1>Sign In</h1>
        <input type="text" placeholder="Email"/>
        <input type="password" placeholder="Password"/>
        <button class="signinpagebutton">Sign In</button>
        <button class="google">Sign Up with Google</button>
        <a href="#">Having trouble logging in?</a>
        <button class="signinpagesignupbutton">Sign Up</button>
    </div>
    <FooterMod />
</template>

<script>
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'
export default {
    name: 'SignIn',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>

<style>
.signin a {
    text-decoration: none;
    color: purple;
    font-size: 12px;
}
</style>