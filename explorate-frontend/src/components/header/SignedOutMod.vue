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
                class="text-[#E7513F] h-[48px] bg-white border-[#E7513F] border-2 hover:bg-red-100 focus:ring-2 focus:outline-none focus:ring-blue-300 font-medium rounded-[15px] px-8 py-4 font-roboto flex items-center">
                Log In</button>
        </router-link>
        <router-link :to="'/signup'">
            <button type="button"
                class="text-white h-[48px] bg-[#E7513F] border-[#E7513F] border-2 hover:bg-primary_hover hover:border-primary_hover focus:ring-2 focus:outline-none focus:ring-blue-300 font-medium rounded-[15px] px-4 font-roboto flex items-center">
                Join Community</button>
        </router-link>
    </div>
</template>