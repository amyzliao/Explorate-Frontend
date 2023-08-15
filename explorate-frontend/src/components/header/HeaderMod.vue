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
    <div class="w-full items-center p-4 flex justify-between bg-white dark:bg-gray-900 w-full top-0 left-0 dark:border-gray-600">
        <router-link :to="'/'">
            <img src="../../assets/explorate-logo.png" class="h-[102px]" alt="Explorate Logo">
        </router-link>
        <div class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1">
            <component :is="view"/>
            <div id="navbar-dropdown">
                <ul class="mx-6 flex items-center flex-col h-[102px] md:p-0 mt-4 font-medium border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700 gap-9">
                    <li>
                        <button id="dropdownNavbarLink" data-dropdown-toggle="dropdownNavbar"
                            class="flex items-center justify-between w-full text-gray-900 rounded hover:bg-gray-100 font-bold">Volunteers
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
                                        class="block px-4 py-2 hover:bg-primary_hover font-bold">Database</router-link>
                                </li>
                                <li>
                                    <router-link :to="'/volunteers/resources'"
                                        class="block px-4 py-2 hover:bg-primary_hover font-bold">Resources</router-link>
                                </li>
                            </ul>
                        </div>
                    </li>
                    <li>
                        <router-link :to="'/partners'"
                            class="block text-gray-900 rounded hover:bg-gray-100 font-bold">NGO
                            Partners</router-link>
                        </li>
                    <li>
                        <router-link :to="'/about'"
                            class="block text-gray-900 rounded hover:bg-gray-100 font-bold">About
                            Us</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>