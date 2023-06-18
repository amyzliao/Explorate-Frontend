<template>
    <HeaderMod />
        <div class="signup">
            <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
            <h1>Sign Up</h1>
            <button class="google">Sign Up with Google</button>
            <input type="text" placeholder="Full Name"/>
            <input type="text" placeholder="Email"/>
            <input type="password" placeholder="Password"/>
            <button class="signuppagebutton">Sign Up</button>
            <button class="signuppagesigninbutton">Sign In</button>
        </div>
    <FooterMod/>
</template>

<script>
import FooterMod from './FooterMod.vue';
import HeaderMod from './HeaderMod.vue';
export default {
    name: 'SignUp',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>

