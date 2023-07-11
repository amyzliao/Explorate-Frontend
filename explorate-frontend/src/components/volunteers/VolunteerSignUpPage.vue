<script>
import { dbCreate, dbGetUser, signInWithGoogle } from '@/firebase';
import FooterMod from '../FooterMod.vue';
import HeaderMod from '../header/HeaderMod.vue';
// import HeaderModSignedOut from '../HeaderModSignedOut.vue';
export default {
    name: 'VolunteerSignUpPage',
    data() {
        return {
            name: '',
            email: '',
            account: 'volunteer',
            password: ''
        }
    },
    components: {
        HeaderMod,
        FooterMod
    },
    methods: {
        async signUp() {
            let result = dbCreate('users',
                {
                    name: this.name,
                    contact: this.email,
                    password: this.password,
                    account: this.account
                })

            if (result == 201) {
                dbGetUser(this.email, this.password).then((res) => {
                    let userId = res.docs[0].id
                    dbCreate('volunteers',
                        {
                            uid: userId,
                            name: this.name,
                            contact: this.email
                        })
                })
                this.$router.push({ name: 'HomePage' });
            }
        },
        googleSignUp() {
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
    <h1
        class="mt-28 text-4xl font-extrabold leading-none tracking-tight text-primary md:text-5xl lg:text-6xl dark:text-white">
        Volunteer Sign Up</h1>
    <form class="w-[30%] mx-auto mt-5 bg-primary rounded-sm px-10 pb-10 pt-7 mb-16">
        <div class="mb-6">
            <label for="email" class="block mb-2 text-sm font-medium text-white dark:text-white">Your name</label>
            <input type="email" id="email"
                class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light"
                placeholder="John Doe" required v-model="name">
        </div>
        <div class="mb-6">
            <label for="email" class="block mb-2 text-sm font-medium text-white dark:text-white">Your email</label>
            <input type="email" id="email"
                class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light"
                placeholder="name@explorate.com" required v-model="email">
        </div>
        <div class="mb-6">
            <label for="password" class="block mb-2 text-sm font-medium text-white dark:text-white">Your password</label>
            <input type="password" id="password"
                class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light"
                required>
        </div>
        <div class="mb-6">
            <label for="repeat-password" class="block mb-2 text-sm font-medium text-white dark:text-white">Repeat
                password</label>
            <input type="password" id="repeat-password"
                class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light"
                required v-model="password">
        </div>
        <div class="flex mx-auto justify-center mb-6">
            <div class="flex items-center h-5">
                <input id="terms" type="checkbox" value=""
                    class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-2 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
                    required>
            </div>
            <label for="terms" class="ml-2 text-sm font-medium text-white dark:text-gray-300">I agree with the <a
                    href="#" class="text-blue-600 hover:underline dark:text-blue-500">terms and conditions</a>
            </label>
        </div>
        <button type="submit"
            class="w-[100%] flex justify-center mb-3 mx-auto text-white bg-secondary_blue hover:bg-[#498fae] focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-lg px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
            v-on:click="signUp">
            Create my account</button>
        <button
            class="w-[100%] flex justify-center mx-auto text-white bg-secondary_yellow hover:bg-[#ae8f61] focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-lg px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
            v-on:click="googleSignUp">
            Sign up with Google</button>
    </form>

    <!-- old -->
    <div class="signup">
        <img class="logo" src="../../assets/-Insert_image_here-.svg.png" />
        <h1>Sign Up</h1>
        <button class="google" v-on:click="googleSignUp">Sign Up with Google</button>
        <input type="text" placeholder="Full Name" v-model="name" />
        <input type="text" placeholder="Email" v-model="email" />
        <input type="password" placeholder="Password" v-model="password" />
        <button class="signuppagebutton" v-on:click="signUp">Sign Up</button>
        <button class="signuppagesigninbutton">Sign In</button>
    </div>
    <FooterMod />
</template>

