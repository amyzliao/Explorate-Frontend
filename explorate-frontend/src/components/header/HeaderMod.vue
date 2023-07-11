<script setup>
import { onMounted } from 'vue'
import { initFlowbite } from 'flowbite'
onMounted(() => {
    initFlowbite();
})
</script>

<script>
import SignedOutMod from './SignedOutMod.vue'
import SignedInMod from './SignedInMod.vue'
export default {
    components: {
        SignedOutMod,
        SignedInMod
    },
    computed: {
        view() {
            let user = localStorage.getItem("user-info");
            if (user) {
                return 'SignedInMod'
            } else {
                return 'SignedOutMod'
            }
        }
    }
}
</script>

<template>
    <nav class="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 left-0 border-b border-gray-200 dark:border-gray-600">
        <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
            <router-link :to="'/'" class="flex items-center">
                <img src="../../../src/assets/explorate-logo.png" class="h-8 mr-3" alt="Explorate Logo">
                <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Explorate</span>
            </router-link>
            <!-- not sure what this does -->
            <button data-collapse-toggle="navbar-dropdown" type="button"
                class="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
                aria-controls="navbar-dropdown" aria-expanded="false">
                <span class="sr-only">Open main menu</span>
                <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M1 1h15M1 7h15M1 13h15" />
                </svg>
            </button>
            <!-- <SignedOutMod/> -->
            <!-- <SignedInMod /> -->
            <component :is="view" />
            <div class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1" id="navbar-dropdown">
                <ul
                    class="flex flex-col p-4 md:p-0 mt-4 font-medium border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                    <li>
                        <router-link :to="'/'"
                            class="block py-2 pl-3 pr-4 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-primary md:p-0 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700"
                            aria-current="page">Home</router-link>
                    </li>
                    <li>
                        <button id="dropdownNavbarLink" data-dropdown-toggle="dropdownNavbar"
                            class="flex items-center justify-between w-full py-2 pl-3 pr-4 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-primary md:p-0 md:w-auto dark:text-white md:dark:hover:text-blue-500 dark:focus:text-white dark:border-gray-700 dark:hover:bg-gray-700 md:dark:hover:bg-transparent">Volunteers
                            <svg class="w-2.5 h-2.5 ml-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                                fill="none" viewBox="0 0 10 6">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="m1 1 4 4 4-4" />
                            </svg></button>
                        <!-- Dropdown menu -->
                        <div id="dropdownNavbar"
                            class="z-10 hidden font-normal bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 dark:divide-gray-600">
                            <ul class="py-2 text-sm text-gray-700 dark:text-gray-400" aria-labelledby="dropdownLargeButton">
                                <li>
                                    <router-link :to="'/database'"
                                        class="block px-4 py-2 hover:bg-primary_hover dark:hover:bg-gray-600 dark:hover:text-white">Database</router-link>
                                </li>
                                <li>
                                    <router-link :to="'/volunteers/resources'"
                                        class="block px-4 py-2 hover:bg-primary_hover dark:hover:bg-gray-600 dark:hover:text-white">Resources</router-link>
                                </li>
                            </ul>
                        </div>
                    </li>
                    <li>
                        <router-link :to="'/partners'"
                            class="block py-2 pl-3 pr-4 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-primary md:p-0 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">NGO
                            Partners</router-link>
                    </li>
                    <li>
                        <router-link :to="'/about'"
                            class="block py-2 pl-3 pr-4 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-primary md:p-0 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">About
                            Us</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>