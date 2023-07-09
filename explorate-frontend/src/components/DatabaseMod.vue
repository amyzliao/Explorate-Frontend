<template>
    <table border="1">
        <tr class="col-title">
            <td>NGO</td>
            <td>Contact</td>
            <td>Title</td>
            <td>Description</td>
            <td>Location</td>
            <td>Website</td>
            <td v-if="loggedIn()">Functions</td>
        </tr>
        <tr v-for="opp in oppCollection" :key="opp.id">
            <td>{{ opp.name }}</td>
            <td>{{ opp.contact }}</td>
            <td>{{ opp.title }}</td>
            <td>{{ opp.description }}</td>
            <td>{{ opp.location }}</td>
            <td>{{ opp.website }}</td>
            <td v-if="loggedIn()">
                <button v-on:click="goToEdit(opp.id)">Edit</button>
                <br/>
                <button v-on:click="deleteThisOpp(opp.id)">Delete</button>
            </td>
        </tr>
    </table>
</template>

<script>
import { dbUseLoad, dbDelete } from '@/firebase'

export default {
    name: 'DatabaseMod',
    data() {
        return {
            oppCollection: []
        }
    },
    methods: {
        loadData() {
            let result = dbUseLoad('opportunities')
            this.oppCollection = result
            console.warn(this.oppCollection)
        },
        deleteThisOpp(id) {
            let result = dbDelete('opportunities', id)
            console.warn(result)
            if (result == 200) {
                this.loadData()
            }
        },
        goToEdit(id) {
            const path = "edit/" + id
            this.$router.push(path)
        },
        loggedIn() {
            const user = localStorage.getItem("user-info")
            return user != null
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

td button {
    width: 60px;
    text-align: center;
}
</style>