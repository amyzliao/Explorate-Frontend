<script>
// import HeaderModVolunteer from '../HeaderModVolunteer.vue'
// import HeaderModNgo from '../HeaderModNgo.vue';
// import HeaderModSignedOut from '../HeaderModSignedOut.vue';
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
    <div class="mt-24 mb-16">
        <h1 class="mx-auto text-5xl font-extrabold text-primary dark:text-white">Database</h1>
        <router-link class="underline text-blue-600" to="/database/add">Add a new opportunity</router-link>
        <DatabaseMod />
    </div>
    <FooterMod />
</template>