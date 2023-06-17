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

<style>
.signup {
    background-color: gray;
    width: 500px;
    margin: auto;
    padding-bottom: 20px;
}

.signup input {
    width: 300px;
    height: 40px;
    padding-left: 20px;
    display: block;
    margin-bottom: 30px;
    margin-right: auto;
    margin-left: auto;
}

.signup button {
    height: 40px;
    cursor: pointer;
    border: none;
    display: block;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 30px;
    width: 326px;
    font-size: 16px;
}

.signuppagebutton {
    border-radius: 20px;
    background-color: #e7513f;
    color: white;
}

.google {
    background-color: white;
    border-radius: 10px;

}

.signuppagesigninbutton {
    border-radius: 20px;
    background-color: white;
}
</style>
