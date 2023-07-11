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
        <tr v-for="opp in oppCollection" :key="opp.Org_ID">
            <td>{{ opp.Org_name.String }}</td>
            <td>{{ opp.Contact.String }}</td>
            <td>{{ opp.Impact_space.String }}</td>
            <td>{{ opp.Mission_statement.String }}</td>
            <td>{{ opp.Org_location.String }}</td>
            <td>{{ opp.Website.String }}</td>
            <td v-if="loggedIn()">
                <button v-on:click="goToEdit(opp.Org_ID)">Edit</button>
                <br/>
                <button v-on:click="deleteThisOpp(opp.id)">Delete</button>
            </td>
            <!-- <h1>oppcollection</h1>
                <p>{{ oppCollection }}</p> -->
        </tr>
    </table>
</template>

<script>
import { dbDelete } from '@/firebase';
import axios from 'axios';

export default {
    name: 'DatabaseMod',
    data() {
        return {
            oppCollection: []
        }
    },
    methods: {
        async loadData() {
            // let result = dbUseLoad('opportunities')
            let result = await axios.get("http://localhost:3000/ngo_opps")
            this.oppCollection = result.data
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