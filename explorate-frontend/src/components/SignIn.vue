<template>
    <HeaderMod />
    <div class="signin">
        <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
        <h1>Sign In</h1>
        <input type="text" placeholder="Email" v-model="email"/>
        <input type="password" placeholder="Password" v-model="password"/>
        <button class="signinpagebutton" v-on:click="login">Sign In</button>
        <button class="google">Sign Up with Google</button>
        <a href="#">Having trouble logging in?</a>
        <button class="signinpagesignupbutton">Sign Up</button>
    </div>
    <FooterMod />
</template>

<script>
// import { query, where, collection } from 'firebase/firestore'
import { dbGetUser } from '@/firebase'
import HeaderMod from './HeaderMod.vue'
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
        HeaderMod,
        FooterMod
    },
    methods: {
        async login() {
            let result = dbGetUser('users', this.email, this.password).then((res) => {
                const user = res.docs[0].data()

                if (user) {
                    console.log(res.docs[0].data())
                    localStorage.setItem("user-info", JSON.stringify(res.docs[0].data()))
                    this.$router.push({name: 'HomePage'})
                }
                
            }).catch((error) => {
                console.log(error)
            })
            // console.log(result)
            if (result.value) {
                localStorage.setItem("user-info", this.email)
                this.$router.push({name: 'HomePage'})
            }
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