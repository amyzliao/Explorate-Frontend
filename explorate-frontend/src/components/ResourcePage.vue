<template>
    <HeaderMod />
    <h1>Resource Page</h1>

    <FooterMod />
</template>

<script>
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'ResourcePage',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>