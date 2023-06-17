<template>
    <HeaderMod />
        <div class="signup">
            <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
            <h1>Sign Up</h1>
            <input type="text" placeholder="Full Name"/>
            <input type="text" placeholder="Email"/>
            <input type="password" placeholder="Password"/>
            <button>Sign Up</button>
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
    display: block;
}

.signup input {
    width: 300px;
    height:40px;
    padding-left: 20px;
    display: block;
    margin-bottom: 30px;
    margin-right: auto;
    margin-left: auto;
    border: 1px solid skyblue;
}

.signup button {
    width: 320px;
    height: 40px;
    border: 1px solid skyblue;
    background: skyblue;
    color: white;
    cursor: pointer;
}
</style>
