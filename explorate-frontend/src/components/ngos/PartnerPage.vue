<script>
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'

export default {
    name: 'PartnerPage',
    components: {
        FooterMod,
        HeaderMod
    },
    computed: {
        view() {
            let user = localStorage.getItem("user-info");
            if (user) {
                const name = JSON.parse(user).account;
                return name == 'volunteer' ? 'HeaderModVolunteer' : 'HeaderModNgo'
            } else {
                return 'HeaderMod'
            }
        }
    }
}
</script>

<template>
    <HeaderMod/>
    <div class="bg-[#E7513F] mx-11 rounded-[50px] h-[760px]">
        <div class="grid grid-cols-2 h-full">
            <div class="text-[#F2DFC1] text-[64px] flex items-center justify-start px-[150px]">
                <div class="flex flex-col justify-start">
                    <h1 class="w-[400px] font-bold text-left">
                    Spread Your Mission to More Volunteers
                    </h1>
                    <p class="text-[32px] text-left">
                        Volunteer coordination streamlined
                    </p>
                    <a href="#">
                        <h1 class="bg-[#E9D1AD] text-black font-bold text-xl inline-block px-14 py-7 rounded-[47px]">
                            Let's Connect
                        </h1>
                    </a>
                </div>
            </div>
            <div class="flex justify-center items-center">
                <img src="../../assets/turtle.png" alt="turtle pic">
            </div>
        </div>
    </div>
    <div class="flex flex-col justify-center mt-6">
        <h1 class="font-bold text-[50px] w-[1020px] self-center">
            Free up your capacity to generate impact on the ground
        </h1>
        <p class="text-[25px] w-[950px] self-center">
            Explorate serves to be a solution to the disconnect between NGOs and volunteers. 
            Let your impactful work be seen and reach reliable volunteers. Meet other NGOs working in your field from across the world. 
        </p>
    </div>
    <div class="grid grid-cols-2 my-10 gap-[140px]">
        <div class="flex justify-end items-center">
            <img class="rounded-[20px]" src="../../assets/Photo.jpg" alt="background">
        </div>
        <div class="text-[22px] w-[440px] flex flex-col justify-between text-left h-[280px] self-center">
            <p class="text-[#E7513F] font-bold">Access Universal Talents and Form Global Connections</p>
            <p class="text-[#797979]">
                Connect with people who resonate with your cause and other impact-driven organizations working in your field. 
                Increase your exposure on a global level
            </p>
        </div>
    </div>
    <div class="grid grid-cols-3 my-10 ml-20">
        <div class="w-[785px] h-[640px] bg-[#E9D1AD] bg-opacity-40 rounded-[50px] pl-[30px] flex flex-col justify-center">
            <div class="w-[600px]">
                <h1 class="font-bold text-[50px] text-[#242331] text-left">Let Us Bring Volunteers to Your Door</h1>
                <div class="text-[#797979] text-[22px] text-left">
                    <p class="my-5">Explorate is set out to facilitate impact driven programs that empower and bridge both parties.</p>
                    <p class="mt-5">We build a database for you to feature your volunteer needs and implement filters to ensure that the opportunities are shared with the volunteers who align with your missions the most.</p>
                </div>
            </div>
        </div>
        <div class="translate-x-[120px] flex items-center">
            <img src="../../assets/screenshot.png" alt="screenshot">
        </div>
        <div class="flex justify-end">
            <div class="w-[313px] h-[640px] bg-[#5D90CD] rounded-l-[50px]"></div>
        </div>
    </div>
    <div class="flex flex-col gap-[89px]">
        <h1 class="font-bold text-[50px]">Recruiting volunteers in less than an hour</h1>
        <div class="flex justify-center gap-10">
            <div class="border border-[#E7513F] rounded-[40px] flex flex-col items-center pt-[80px] pb-[30px] gap-y-[10px] w-[382px]">
                <img class="h-[135px] w-[68px] mb-12" src="../../assets/Vector.png" alt="one">
                <h1 class="font-bold text-3xl">Express your interest</h1>
                <p class="text-[#797979] text-[22px] h-[150px]">Fill out a quick interest form to tell <br> us more about your organization</p>
                <p class="font-bold text-[18px] text-[#5961F9] underline italic">2 minutes</p>
            </div>
            <div class="border border-[#E7513F] rounded-[40px] flex flex-col items-center pt-[80px] pb-[30px] gap-y-[10px] w-[382px]">
                <img class="h-[135px] w-[86px] mb-12" src="../../assets/Vector-1.png" alt="two">
                <h1 class="font-bold text-3xl">Connect Face-to Face</h1>
                <p class="text-[#797979] text-[22px] h-[150px]">Once we identify a fit, we’ll invite <br> you to an interview for us to learn more about each other.</p>
                <p class="font-bold text-[18px] text-[#5961F9] underline italic">30 minutes</p>
            </div>
            <div class="border border-[#E7513F] rounded-[40px] flex flex-col items-center pt-[80px] pb-[30px] gap-y-[10px] w-[382px]">
                <img class="h-[135px] w-[88px] mb-12" src="../../assets/Vector-2.png" alt="three">
                <h1 class="font-bold text-3xl">Apply to Be a Partner</h1>
                <p class="text-[#797979] text-[22px] h-[150px]">Let us know who you are looking for to attract the talents that fit your needs. Leave the searching to us from here on. </p>
                <p class="font-bold text-[18px] text-[#5961F9] underline italic">20 minutes</p>
            </div>
        </div>
    </div>
    <a href="#">
        <h1 class="mt-20 mb-32 bg-[#E7513F] text-white font-bold text-xl inline-block px-14 py-7 rounded-[47px]">
            Count me in!
        </h1>
    </a>
    <FooterMod />
</template>

