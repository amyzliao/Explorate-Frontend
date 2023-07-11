<script>
// import HeaderModVolunteer from '../HeaderModVolunteer.vue'
// import HeaderModNgo from '../HeaderModNgo.vue'
// import HeaderModSignedOut from './HeaderModSignedOut.vue'
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'
import { userID, googleSignOut } from '@/firebase';

export default {
    name: 'VolunteerSignOutPage',
    components: {
        FooterMod,
        HeaderMod
    },
    methods: {
        logout() {
            if (userID()) {
                googleSignOut()
            }
            localStorage.clear();
            this.$router.push({ name: 'SignInPage' })
        }
    },
}
</script>

<template>
    <HeaderMod />
    <!-- block A: sign out -->
    <div class="mb-16">
        <h1
            class="mt-44 mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
            Are you sure you want to sign out?</h1>
        <p class="mb-6 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400">We hate to see you
            go.</p>
        <a href="#" v-on:click="logout"
            class="inline-flex items-center justify-center px-5 py-3 text-base font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900">
            Sign me out!
            <svg class="w-3.5 h-3.5 ml-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M1 5h12m0 0L9 1m4 4L9 9" />
            </svg>
        </a>
    </div>
    <FooterMod />
</template>

