<script>
import { dbGetUser, signInWithGoogle } from '@/firebase'
import HeaderMod from '../header/HeaderMod.vue';
import FooterMod from '../FooterMod.vue'

export default {
    name: 'LogInPage',
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
            let result = dbGetUser(this.email, this.password).then((res) => {
                const user = res.docs[0].data()

                if (user) {
                    localStorage.setItem("user-info", JSON.stringify(user))
                    this.$router.push({ name: 'HomePage' })
                }

            }).catch((error) => {
                console.log(error)
            })
            if (result.value) {
                localStorage.setItem("user-info", this.email)
                this.$router.push({ name: 'HomePage' })
            }
        },
        googleSignIn() {
            signInWithGoogle().then((result) => {
                var user = result.user;

                const uid = user.uid

                if (user) {
                    localStorage.setItem("user-info", JSON.stringify(uid))
                    this.$router.push({ name: 'HomePage' })
                }

            }).catch((error) => {
                console.log(error)
            })
        }
    }
}
</script>

<template>
    <HeaderMod />
    <div
        class="w-full max-w-sm p-4 mt-24 mx-auto mb-16 bg-primary border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700">
        <form class="space-y-6" action="#">
            <h5 class="text-xl font-bold text-white dark:text-white">Sign in to Explorate</h5>
            <div>
                <label for="email" class="block mb-2 text-sm font-medium text-white dark:text-white">Your email</label>
                <input type="email" name="email" id="email"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                    placeholder="name@explorate.com" required v-model="email">
            </div>
            <div>
                <label for="password" class="block mb-2 text-sm font-medium text-white dark:text-white">Your
                    password</label>
                <input type="password" name="password" id="password" placeholder="••••••••"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                    required v-model="password">
            </div>
            <div class="flex items-start">
                <div class="flex items-start">
                    <div class="flex items-center h-5">
                        <input id="remember" type="checkbox" value=""
                            class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
                            required>
                    </div>
                    <label for="remember" class="ml-2 text-sm font-medium text-white dark:text-gray-300">Remember
                        me</label>
                </div>
                <a href="#" class="ml-auto text-sm text-blue-700 hover:underline dark:text-blue-500">Forgot Password?</a>
            </div>
            <button type="submit"
                class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                v-on:click="login">
                Sign in</button>
            <button type="submit"
                class="w-full text-white bg-secondary_yellow hover:bg-[#ae8f61] focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                v-on:click="googleSignIn">
                Sign in with Google</button>
            <div class="text-sm font-medium text-white dark:text-gray-300">
                Not registered? <router-link :to="'/signup'" class="text-blue-700 hover:underline dark:text-blue-500">Create account</router-link>
            </div>
        </form>
    </div>


    <!-- old -->
    <div class="signin">
        <img class="logo" src="../../assets/-Insert_image_here-.svg.png" />
        <h1>Sign In</h1>
        <input type="text" placeholder="Email" v-model="email" />
        <input type="password" placeholder="Password" v-model="password" />
        <button class="signinpagebutton" v-on:click="login">Sign In</button>
        <button class="google" v-on:click="googleSignIn">Sign In with Google</button>
        <a href="#">Having trouble logging in?</a>
        <button class="signinpagesignupbutton">Sign Up</button>
    </div>
    <FooterMod />
</template>

<style>
.signin a {
    text-decoration: none;
    color: purple;
    font-size: 12px;
}
</style>