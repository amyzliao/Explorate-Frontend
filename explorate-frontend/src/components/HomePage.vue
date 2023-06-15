<template>
    <HeaderMod />
    <h1>Home Page</h1>

    <FooterMod />
</template>

<script>
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'HomePage',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>