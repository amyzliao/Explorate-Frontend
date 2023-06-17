<template>
    <HeaderMod />
    <div class="home">
        <div class="left">
            <h1>Global Connections<br>Driven by<br>Sustainable Impact</h1>
            <p>Discover without limitation. <br> Impact without assumption.</p>
            <button class="find">Find a program</button>
            <button class="partner">Partner with us</button>
        </div>
        <div class="right">
            <img class="empty" src="../assets/-Insert_image_here-.svg.png"/>
        </div>
    </div>
    <div>
        <h2>Make the world a better place together</h2>
        <p>Get the community and resources you need all in one place</p>
        <div class="cards">
            <div class="volunteercard">
                Become a volunteer
            </div>
            <div class="ngocard">
                Partner with us
            </div>
            <div class="expeditioncard">
                Join an expedition
            </div>
        </div>
    </div>
    <div class="perspective">
        <div class="text">
            <h2>A New Perspective Before You <br> Volunteer</h2>
            <p>For decades international volunteers have relied mostly on marketing project info from volunteering organizations. 
                Here, you’ll learn about real experience from veteran volunteers and the dos and don’ts to determine which program is right for you! </p>
        </div>
        <div class="image">
            <img class="empty" src="../assets/-Insert_image_here-.svg.png"/>
        </div>
    </div>
    <FooterMod />
</template>

<script>
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'HomePage',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>

<style>
.home {
    padding-top: 50px;
    overflow: hidden;
    margin-bottom: 20px;
    padding-bottom: 50px;
}

.left {
    float: left;
    width: 50%;
    text-align: center;
}

.left h1 {
    font-size: 55px;
    font-weight: normal;
    text-align: left;
    padding-bottom: 10px;
    margin-left: 90px;
}

.left p {
    text-align: left;
    padding-bottom: 20px;
    font-size: 20px;
    margin-left: 90px;
}

.right {
    float: right;
    width: 50%;
    background-color: green;
}

.left button {
    height: 55px;
    width: 240px;
    border-radius: 6px;
    font-family: "Inter";
    font-size: 14px;
    font-weight: bold;
    border: 2px solid #e7513f;
    float: left;
    cursor: pointer;
    margin-left: 90px;
}

.find {
    background-color: #e7513f;
    color: white;
}

.partner {
    background-color: white;
    color: #e7513f;
    margin-left: 30px;
}

.cards {
    display: flex;
    justify-content: center;
    gap: 20px;
    margin-bottom: 50px;
}

.volunteercard {
    background-color: gray;
    height: 600px;
    width: 300px;
}

.ngocard {
    background-color: gray;
    height: 600px;
    width: 300px;
}

.expeditioncard {
    background-color: gray;
    height: 600px;
    width: 300px;
}

.perspective {
    overflow: hidden;
}

.text {
    float: left;
    width: 50%;
    text-align: left;
}

.text p, .text h2 {
    padding-right: 400px;
    margin-left: 90px;
}

.image {
    float:right;
    width: 50%;
    background-color: green;
}
</style>