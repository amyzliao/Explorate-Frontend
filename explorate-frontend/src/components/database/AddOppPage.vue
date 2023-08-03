<template>
    <HeaderMod />
    <h1
        class="mt-44 mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
        Add an NGO to the database</h1>
    <div className="big-form">
        <!-- <input placeholder="ID" v-model="ID"> -->
        <input placeholder="org_name" v-model="org_name">
        <input placeholder="org_location" v-model="org_location">
        <input placeholder="mission statement" v-model="mission_statement">
        <input placeholder="impact space" v-model="impact_space">
        <input placeholder="un_goals" v-model="un_goals">
        <input placeholder="projects" v-model="projects">
        <input placeholder="values" v-model="values">
        <input placeholder="benefits" v-model="benefits">
        <input placeholder="ethical criteria" v-model="ethical_criteria">
        <input placeholder="ethical details" v-model="ethical_criteria_details">
        <input placeholder="volunteer_location" v-model="volunteer_location">
        <input placeholder="volunteer skills" v-model="volunteer_skills">
        <input placeholder="skill projects" v-model="skill_projects">
        <input placeholder="language" v-model="language_requirements">
        <input placeholder="volunteer qualities" v-model="volunteer_qualities">
        <input placeholder="work_env" v-model="work_env">
        <input placeholder="accomodations" v-model="accomodations">
        <input placeholder="cost" v-model="cost">
        <input placeholder="website" v-model="website">
        <input placeholder="contact" v-model="contact">
        <input placeholder="photos" v-model="photos">
        <input placeholder="refer_others" v-model="refer_others">
        <input placeholder="additional info" v-model="additional_info">
        <input placeholder="finding_volunteers" v-model="finding_volunteers">
        <input placeholder="logo" v-model="org_logo">
        <input placeholder="duration" v-model="program_duration">
        <input placeholder="timeframe" v-model="project_timeframe">
        <button v-on:click="addOpp" type="button"
            class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">
            Add NGO
        </button>
    </div>
    <FooterMod />
</template>

<script>
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'
import axios from 'axios'

export default {
    name: 'AddOppPage',
    components: {
        HeaderMod,
        FooterMod
    },
    data() {
        return {
            // ID: '',
            org_name: '',
            org_location: '',
            mission_statement: '',
            impact_space: '',
            un_goals: '',
            projects: '',
            values: '',
            benefits: '',
            ethical_criteria: '',
            ethical_criteria_details: '',
            volunteer_location: '',
            volunteer_skills: '',
            skill_projects: '',
            language_requirements: '',
            volunteer_qualities: '',
            work_env: '',
            accomodations: '',
            cost: '',
            website: '',
            contact: '',
            photos: '',
            refer_others: '',
            additional_info: '',
            finding_volunteers: '',
            org_logo: '',
            program_duration: '',
            project_timeframe: '',
        }
    },
    methods: {
        async addOpp() {
            axios.post("http://localhost:3000/ngo_opps", {
                // ID: parseInt(this.ID),
                Org_name: this.org_name,
                Org_location: this.org_location,
                Mission_statement: this.mission_statement,
                Impact_space: this.impact_space,
                Un_goals: this.un_goals,
                Projects: this.projects,
                Values: this.values,
                Benefits: this.benefits,
                Ethical_criteria: this.ethical_criteria,
                Ethical_criteria_details: this.ethical_criteria_details,
                Volunteer_location: this.volunteer_location,
                Volunteer_skills: this.volunteer_skills,
                Skill_projects: this.skill_projects,
                Language_requirements: this.language_requirements,
                Volunteer_qualities: this.volunteer_qualities,
                Work_env: this.work_env,
                Accomodations: this.accomodations,
                Cost: this.cost,
                Website: this.website,
                Contact: this.contact,
                Photos: this.photos,
                Refer_others: this.refer_others,
                Additional_info: this.additional_info,
                Finding_volunteers: this.finding_volunteers,
                Org_logo: this.org_logo,
                Program_duration: this.program_duration,
                Project_timeframe: this.project_timeframe
            }).then(result => {
                console.warn("result", result)
                if (result.status == 200) {
                    this.$router.push({ name: 'DatabasePage' });
                }
            })
        }
    }
}
</script>