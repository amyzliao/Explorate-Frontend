<script setup>
import { onMounted } from 'vue'
import { initFlowbite } from 'flowbite'
onMounted(() => {
    initFlowbite();
})
</script>

<script>
import { userName, userImg, userEmail } from '@/firebase';
export default {
    data() {
        return {
            username: '',
            userimg: '',
            useremail: '',
        }
    },
    methods: {
        loadUser() {
            this.username = userName()
            this.userimg = userImg()
            this.useremail = userEmail()
            console.log('loadUser')
            console.log(this.username + this.userimg + this.useremail)
        }
    },
    async mounted() {
        this.loadUser()
    }
}
</script>

<template>
    <div class="flex md:order-2 gap-1">
        <button id="dropdownAvatarNameButton" data-dropdown-toggle="dropdownAvatarName"
            class="flex h-9 bg-[#ffebcd] pr-3 items-center text-sm font-medium text-gray-900 rounded-full hover:bg-[#e1ba7e] md:mr-0 focus:bg-[#e1ba7e] dark:text-white"
            type="button">
            <span class="sr-only">Open user menu</span>
            <img class="w-8 h-8 mr-2 rounded-full" :src="userimg" alt="user photo">
            {{ username }}
            <svg class="w-2.5 h-2.5 ml-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                viewBox="0 0 10 6">
                <path stroke="currentColor" stroke-linecap="roun    d" stroke-linejoin="round" stroke-width="2"
                    d="m1 1 4 4 4-4" />
            </svg>
        </button>

        <!-- Dropdown menu -->
        <div id="dropdownAvatarName"
            class="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 dark:divide-gray-600">
            <div class="px-4 py-3 text-sm text-gray-900 dark:text-white">
                <div class="font-medium ">volunteer</div>
                <div class="truncate">{{ useremail }}</div>
            </div>
            <ul class="py-2 text-sm text-gray-700 dark:text-gray-200"
                aria-labelledby="dropdownInformdropdownAvatarNameButtonationButton">
                <li>
                    <a href="#"
                        class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Dashboard</a>
                </li>
                <li>
                    <router-link :to="'/volunteers/profile'"
                        class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Profile</router-link>
                </li>
                <li>
                    <a href="#"
                        class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Settings</a>
                </li>
            </ul>
            <div class="py-2">
                <router-link :to="'/signout'"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white">Sign
                    out</router-link>
            </div>
        </div>

    </div>
</template>