<template>
    <HeaderMod/>
    <h1>Add a new program to the database</h1>
    <div class="form">
        <input type="text" name="name" v-model="name" placeholder="NGO Name" />
        <input type="text" name="contact" v-model="contact" placeholder="NGO Contact" />
        <input type="text" name="title" v-model="title" placeholder="Program Title" />
        <input type="text" name="description" v-model="description" placeholder="Program Description" />
        <input type="text" name="location" v-model="location" placeholder="Program Location" />
        <input type="text" name="website" v-model="website" placeholder="Link to website" />
        <button v-on:click="addOpp">Add new program</button>
    </div>
    <FooterMod/>
</template>

<script>
import { createOpp } from '@/firebase'
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'AddOppPage',
    components: {
        HeaderMod,
        FooterMod
    },
    data() {
        return {
            name:'',
            contact:'',
            title:'',
            description:'',
            location:'',
            website:''
        }
    },
    methods: {
        async addOpp() {
            // console.log("addOpp called")
            let result = await createOpp({
                name: this.name,
                contact: this.contact,
                title: this.title,
                description: this.description,
                location: this.location,
                website: this.website
            });
            console.warn(result)
            if (result == 201) {
                this.$router.push({name:'DatabasePage'});
            }
        }
    }
}
</script>