<template>
    <HeaderMod />
    <h1>Database Page</h1>
    <router-link to="/add">Add a new opportunity</router-link>
    <DatabaseMod />
    <FooterMod />
</template>

<script>
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'
import DatabaseMod from './DatabaseMod.vue';

export default {
    name: 'DatabasePage',
    components: {
        HeaderMod,
        FooterMod,
        DatabaseMod
    }
}
</script>