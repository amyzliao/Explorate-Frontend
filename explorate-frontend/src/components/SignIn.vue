<template>
    <HeaderModSignedOut />
    <div class="signin">
        <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
        <h1>Sign In</h1>
        <input type="text" placeholder="Email" v-model="email"/>
        <input type="password" placeholder="Password" v-model="password"/>
        <button class="signinpagebutton" v-on:click="login">Sign In</button>
        <button class="google" v-on:click="googleSignIn">Sign In with Google</button>
        <a href="#">Having trouble logging in?</a>
        <button class="signinpagesignupbutton">Sign Up</button>
    </div>
    <FooterMod />
</template>

<script>
import { dbGetUser, signInWithGoogle } from '@/firebase'
import HeaderModSignedOut from './HeaderModSignedOut.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'SignIn',
    data() {
        return {
            email: '',
            password: ''
        }
    },
    components: {
        HeaderModSignedOut,
        FooterMod
    },
    methods: {
        async login() {
            let result = dbGetUser(this.email, this.password).then((res) => {
                const user = res.docs[0].data()

                if (user) {
                    localStorage.setItem("user-info", JSON.stringify(user))
                    this.$router.push({name: 'HomePage'})
                }
                
            }).catch((error) => {
                console.log(error)
            })
            if (result.value) {
                localStorage.setItem("user-info", this.email)
                this.$router.push({name: 'HomePage'})
            }
        },
        googleSignIn() {
            signInWithGoogle().then((result) => {
                var user = result.user;

                const uid = user.uid

                if (user) {
                    localStorage.setItem("user-info", JSON.stringify(uid))
                    this.$router.push({name: 'HomePage'})
                }

                }).catch((error) => {
                    console.log(error)
                })
        }
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