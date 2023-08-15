<script setup>
import { onMounted } from 'vue'
import { initFlowbite } from 'flowbite'
onMounted(() => {
    initFlowbite();
})
</script>

<template>
    <div class="flex md:order-2 gap-10 px-10">
        <router-link :to="'/signin'">
            <button type="button"
                class="text-primary h-[48px] w-[103px] bg-white border-primary border-2 hover:bg-red-100 focus:ring-2 focus:outline-none focus:ring-blue-300 font-medium rounded-[15px] text-sm px-4 py-1 text-center font-roboto">
                Log In</button>
        </router-link>
        <router-link :to="'/signup'">
            <button type="button"
                class="text-white h-[48px] bg-primary border-primary border-2 hover:bg-primary_hover hover:border-primary_hover focus:ring-2 focus:outline-none focus:ring-blue-300 font-medium rounded-[15px] text-sm px-4 text-center">
                Join Community</button>
        </router-link>
    </div>
</template>