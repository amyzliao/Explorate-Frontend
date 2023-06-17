<template>
    <div class="nav">
        <router-link to="/">Home</router-link>
        <router-link to="/about">About</router-link>
        <router-link to="/resources">Resources</router-link>
        <router-link to="/database">Database</router-link>
    </div>
</template>

<script>
export default {
    name: 'HeaderMod',
}
</script>

<style>
.nav {
    background-color: orange;
    padding: 20px;
}
.nav a {
    text-decoration: none;
    color: white;
    padding: 20px;
}
.nav a:hover {
    background-color: orangered;
}
</style>