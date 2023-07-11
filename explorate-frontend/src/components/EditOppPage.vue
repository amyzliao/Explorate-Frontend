<template>
    <HeaderModVolunteer/>
    <h1>Edit opportunity</h1>
    <div class="form">
        <input type="text" name="name" v-model="opp.name" placeholder="NGO Name" />
        <input type="text" name="contact" v-model="opp.contact" placeholder="NGO Contact" />
        <input type="text" name="title" v-model="opp.title" placeholder="Program Title" />
        <input type="text" name="description" v-model="opp.description" placeholder="Program Description" />
        <input type="text" name="location" v-model="opp.location" placeholder="Program Location" />
        <input type="text" name="website" v-model="opp.website" placeholder="Link to website" />
        <button v-on:click="editOpp">Save changes</button>
    </div>
    <FooterMod/>
</template>

<script>
import { dbUpdate } from '@/firebase'
import HeaderModVolunteer from './HeaderModVolunteer.vue'
import FooterMod from './FooterMod.vue'
import axios from 'axios';
export default {
    name: 'EditOppPage',
    components: {
        HeaderModVolunteer,
        FooterMod
    },
    data() {
        return {
            opp: {
                name:'',
                contact:'',
                title:'',
                description:'',
                location:'',
                website:''
            },
            oppId:''
        }
    },
    methods: {
        async editOpp() {
            let result = dbUpdate('opportunities', this.oppId, {
                name: this.opp.name,
                contact: this.opp.contact,
                title: this.opp.title,
                description: this.opp.description,
                location: this.opp.location,
                website: this.opp.website
            });
            console.warn(result)
            if (result == 200) {
                this.$router.push({name:'DatabasePage'});
            }
        }
    },
    async mounted() {
        this.oppId = this.$route.params.id
        // console.log(this.oppId)
        // let result = await dbGet('opportunities', this.oppId)
        let result = await axios.get("http://localhost:3000/ngo_opps/"+this.oppId)
        // console.log("works")
        this.opp = result
        // this.opp.name = result.ID
        console.log(this.opp)
    }
}
</script>