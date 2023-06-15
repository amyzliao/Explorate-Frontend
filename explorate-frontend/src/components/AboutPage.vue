<template>
    <HeaderMod />
    <h1>About Page</h1>

    <FooterMod />
</template>

<script>
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'AboutPage',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>