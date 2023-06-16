<template>
    <table border="1">
        <tr class="col-title">
            <td>NGO</td>
            <td>Contact</td>
            <td>Title</td>
            <td>Description</td>
            <td>Location</td>
            <td>Website</td>
            <td>Functions</td>
        </tr>
        <tr v-for="opp in oppCollection" :key="opp.id">
            <td>{{ opp.name }}</td>
            <td>{{ opp.contact }}</td>
            <td>{{ opp.title }}</td>
            <td>{{ opp.description }}</td>
            <td>{{ opp.location }}</td>
            <td>{{ opp.website }}</td>
            <td>
                <router-link :to="'/edit/'+opp.id">Edit</router-link>
                <button v-on:click="deleteThisOpp(opp.id)">Delete</button>
            </td>
        </tr>
    </table>
</template>

<script>
import { useLoadOpps, deleteOpp } from '@/firebase'

export default {
    name: 'DatabaseMod',
    data() {
        return {
            oppCollection: []
        }
    },
    methods: {
        async loadData() {
            let result = await useLoadOpps()
            this.oppCollection = result
            console.warn(this.oppCollection)
        },
        async deleteThisOpp(id) {
            let result = await deleteOpp(id)
            console.warn(result)
            if (result == 200) {
                this.loadData()
            }
        }
    },
    async mounted() {
        this.loadData()
    }
}
</script>

<style>
table {
    margin: 10px;
}
td {
    width: 160px;
    height: 40px;
}
.col-title {
    color: blue;
}
</style>