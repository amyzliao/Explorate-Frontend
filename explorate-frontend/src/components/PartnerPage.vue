<template>
    <component :is="view" />
    <h1>Resource Page</h1>
    <div class="resource">
        <div class="resource-left">
            <h1>Spread Your<br>Mission to More<br>Volunteers</h1>
            <p>Volunteer Coordination Streamlined.</p>
            <button class="connect">Let's connect!</button>
        </div>
        <div class="resource-right">
            <img class="empty" src="../assets/-Insert_image_here-.svg.png"/>
        </div>
    </div>
    <div class="partner-page">
        <h2>Free up your capacity to generate impact on the ground</h2>
        <p>
            Explorate serves to be a solution to the disconnect between NGOs and volunteers.<br>
            Let your impactful work be seen and reach reliable volunteers.<br>
            Meet other NGOs working in your field from across the world!
        </p>
        <div class="partner-page-2">
            <img class="empty" src="../assets/-Insert_image_here-.svg.png"/>
            <div class="access-universal-talents-section">
                <div class="access-universal-talents">
                    <h3>Access Universal Talents</h3>
                    <p>
                        At Explorate, you can find the skills, experience, and anything else from across the world to meet your needs to grow your impact.
                    </p>
                </div>
                <div class="form-global-connections">
                    <h3>Form Global Connections</h3>
                    <p>
                        Connect with people who resonate with your cause and other impact-driven organizations working in your field. 
                        Increase your exposure on a global level.
                    </p>
                </div>
            </div>
        </div>
    </div>
    <div class="resource-2">
        <div class="resource-left-2">
            <h3>Let Us Bring Volunteers to Your Door</h3>
            <p>Explorate is set out to facilitate impact-driven programs that empower and bridge both parties.</p>
            <p>We build a database for you to feature your volunteer needs and implement filters to ensure that the 
                opportunities are shared with the volunteers who align with your missions the most.</p>
        </div>
        <div class="resource-right-2">
            <img class="empty" src="../assets/-Insert_image_here-.svg.png"/>
        </div>
    </div>
    <div class="recruiting-volunteers">
        <h3>Recruiting volunteers &lt; 1 hour</h3>
    </div>
    <FooterMod />
</template>

<script>
import HeaderModVolunteer from './HeaderModVolunteer.vue'
import HeaderModNgo from './HeaderModNgo.vue'
import HeaderModSignedOut from './HeaderModSignedOut.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'PartnerPage',
    components: {
        HeaderModVolunteer,
        HeaderModNgo,
        FooterMod,
        HeaderModSignedOut
    },
    computed: {
        view() {
            let user = localStorage.getItem("user-info");
            if (user) {
                const name = JSON.parse(user).account;
                return name == 'volunteer' ? 'HeaderModVolunteer' : 'HeaderModNgo'
            } else {
                return 'HeaderModSignedOut'
            }
        }
    }
}
</script>

<style>
.partner-page, .partner-page-2 {
    overflow: hidden;
    color: #383b3d;
}

.access-universal-talents-section {
    float: right;
    display: block;
}

.access-universal-talents {
    width: 180px;
    height: 160px;
    padding: 3px;
    margin-right: 50px;
    background-color: #eceefa;
    border-radius: 10px;
    margin-bottom: 10px;
}

.form-global-connections {
    width: 180px;
    height: 180px;
    padding: 3px;
    margin-right: 80px;
    background-color: #eceefa;
    border-radius: 10px;
    margin-bottom: 10px;
}

.access-universal-talents h3, .form-global-connections h3 {
    font-size: 15px;
    text-align: left;
    margin-left: 20px;
}

.access-universal-talents p, .form-global-connections p {
    font-size: 13px;
}

.resource-left-2 h3 {
    text-align: left;
    margin-left: 200px;
    font-size: 30px;
    font-weight: 500;
    color: #ea5a39;
}

.resource-left-2 p {
    text-align: left;
    margin-left: 200px;
    font-weight: 300;
}

.recruiting-volunteers h3 {
    color: #ea5a39;
    font-size: 30px;
}
</style>