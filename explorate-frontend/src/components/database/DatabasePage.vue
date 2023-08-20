<script>
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'
import DatabaseMod from './DatabaseMod.vue';

export default {
    name: 'DatabasePage',
    components: {
        FooterMod,
        DatabaseMod,
        HeaderMod
    },
}
</script>

<template>
    <HeaderMod/>
    <div class="relative flex">
        <img class="h-[647px] w-full inline" src="../../assets/IMG_3707.png" alt="cover image">
        <h1 class="absolute top-[170px] left-20 inline-flex justify-start text-white font-bold text-[70px]">Meet our NGOs</h1>
        <p class="absolute top-[250px] pl-1 left-20 inline-flex justify-start text-white font-semibold text-[35px]">Your Journey Starts Here</p>
    </div>
    <!-- <router-link class="underline text-blue-600" to="/database/add">Add a new opportunity</router-link> -->
    <DatabaseMod class="-translate-y-[70px]"/>
    <FooterMod />
</template>