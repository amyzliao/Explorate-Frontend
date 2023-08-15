<template>
    <footer class="flex justify-between gap-[48px] bg-[#E7513F] dark:bg-gray-900 text-white py-[48px] px-[80px] h-[212px]">
        <div class="flex items-center">
            <router-link :to="'/'">
                <img src="../assets/explorate-logo-white.png" class="h-[95px]" alt="Explorate Logo">
            </router-link>
            <div class="flex p-2 bg-white ml-[58px] p-[12px] gap-[12px] rounded-[14px] w-auto">
                <input type="text" id="first_name" class="w-[290px] bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block px-[28px] py-[14px] dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Enter your email">
                <button type="submit" class="w-[163px] text-white bg-[#5FBAE1] hover:bg-blue-500 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm sm:w-auto px-[28px] pt-[14px] pb-[10px] text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    Subscribe Now
                </button>
            </div>
        </div>
        <div class="gap-[48px]">
            <ul class="text-white dark:text-gray-400 font-medium flex gap-[16px] py-[12px] px-[8px] text-[16px]">
                <li>
                    <router-link :to="'#'" class="hover:underline ">Eleven</router-link>
                </li>
                <li>
                    <router-link :to="'#'" class="hover:underline">Twelve</router-link>
                </li>
                <li>
                    <router-link :to="'#'" class="hover:underline ">Thirteen</router-link>
                </li>
                <li>
                    <router-link :to="'#'" class="hover:underline">Fourteen</router-link>
                </li>
                <li>
                    <router-link :to="'#'" class="hover:underline ">Fifteen</router-link>
                </li>
                <li>
                    <router-link :to="'#'" class="hover:underline">Sixteen</router-link>
                </li>
            </ul>
            <div class="flex px-[8px] mt-4 space-x-5 sm:mt-0 gap-[16px] justify-end">
                <a href="#" class="text-white hover:text-gray-900 dark:hover:text-white">
                    <svg class="w-[24px] h-[24px]" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 14">
                        <path fill-rule="evenodd" d="M19.7 3.037a4.26 4.26 0 0 0-.789-1.964 2.84 2.84 0 0 0-1.984-.84c-2.767-.2-6.926-.2-6.926-.2s-4.157 0-6.928.2a2.836 2.836 0 0 0-1.983.84A4.225 4.225 0 0 0 .3 3.038a30.148 30.148 0 0 0-.2 3.206v1.5c.01 1.071.076 2.142.2 3.206.094.712.363 1.39.784 1.972.604.536 1.38.837 2.187.848 1.583.15 6.731.2 6.731.2s4.161 0 6.928-.2a2.844 2.844 0 0 0 1.985-.84 4.27 4.27 0 0 0 .787-1.965c.124-1.064.19-2.135.2-3.206V6.243a30.672 30.672 0 0 0-.202-3.206ZM8.008 9.59V3.97l5.4 2.819-5.4 2.8Z" clip-rule="evenodd"/>
                    </svg>
                    <span class="sr-only">Youtube page</span>
                </a>
                <a href="#" class="text-white hover:text-gray-900 dark:hover:text-white">
                    <svg class="w-[24px] h-[24px]" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 8 19">
                        <path fill-rule="evenodd" d="M6.135 3H8V0H6.135a4.147 4.147 0 0 0-4.142 4.142V6H0v3h2v9.938h3V9h2.021l.592-3H5V3.591A.6.6 0 0 1 5.592 3h.543Z" clip-rule="evenodd"/>
                    </svg>
                    <span class="sr-only">Facebook page</span>
                </a>
                <a href="#" class="text-white hover:text-gray-900 dark:hover:text-white">
                    <svg class="w-[24px] h-[24px]" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 17">
                        <path fill-rule="evenodd" d="M20 1.892a8.178 8.178 0 0 1-2.355.635 4.074 4.074 0 0 0 1.8-2.235 8.344 8.344 0 0 1-2.605.98A4.13 4.13 0 0 0 13.85 0a4.068 4.068 0 0 0-4.1 4.038 4 4 0 0 0 .105.919A11.705 11.705 0 0 1 1.4.734a4.006 4.006 0 0 0 1.268 5.392 4.165 4.165 0 0 1-1.859-.5v.05A4.057 4.057 0 0 0 4.1 9.635a4.19 4.19 0 0 1-1.856.07 4.108 4.108 0 0 0 3.831 2.807A8.36 8.36 0 0 1 0 14.184 11.732 11.732 0 0 0 6.291 16 11.502 11.502 0 0 0 17.964 4.5c0-.177 0-.35-.012-.523A8.143 8.143 0 0 0 20 1.892Z" clip-rule="evenodd"/>
                    </svg>
                    <span class="sr-only">Twitter page</span>
                </a>
                <a href="#" class="text-white hover:text-gray-900 dark:hover:text-white">
                    <svg class="w-[24px] h-[24px]" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                        <path fill-rule="evenodd" d="M8 0C5.829 0 5.556.01 4.703.048 3.85.088 3.269.222 2.76.42a3.917 3.917 0 0 0-1.417.923A3.927 3.927 0 0 0 .42 2.76C.222 3.268.087 3.85.048 4.7.01 5.555 0 5.827 0 8.001c0 2.172.01 2.444.048 3.297.04.852.174 1.433.372 1.942.205.526.478.972.923 1.417.444.445.89.719 1.416.923.51.198 1.09.333 1.942.372C5.555 15.99 5.827 16 8 16s2.444-.01 3.298-.048c.851-.04 1.434-.174 1.943-.372a3.916 3.916 0 0 0 1.416-.923c.445-.445.718-.891.923-1.417.197-.509.332-1.09.372-1.942C15.99 10.445 16 10.173 16 8s-.01-2.445-.048-3.299c-.04-.851-.175-1.433-.372-1.941a3.926 3.926 0 0 0-.923-1.417A3.911 3.911 0 0 0 13.24.42c-.51-.198-1.092-.333-1.943-.372C10.443.01 10.172 0 7.998 0h.003zm-.717 1.442h.718c2.136 0 2.389.007 3.232.046.78.035 1.204.166 1.486.275.373.145.64.319.92.599.28.28.453.546.598.92.11.281.24.705.275 1.485.039.843.047 1.096.047 3.231s-.008 2.389-.047 3.232c-.035.78-.166 1.203-.275 1.485a2.47 2.47 0 0 1-.599.919c-.28.28-.546.453-.92.598-.28.11-.704.24-1.485.276-.843.038-1.096.047-3.232.047s-2.39-.009-3.233-.047c-.78-.036-1.203-.166-1.485-.276a2.478 2.478 0 0 1-.92-.598 2.48 2.48 0 0 1-.6-.92c-.109-.281-.24-.705-.275-1.485-.038-.843-.046-1.096-.046-3.233 0-2.136.008-2.388.046-3.231.036-.78.166-1.204.276-1.486.145-.373.319-.64.599-.92.28-.28.546-.453.92-.598.282-.11.705-.24 1.485-.276.738-.034 1.024-.044 2.515-.045v.002zm4.988 1.328a.96.96 0 1 0 0 1.92.96.96 0 0 0 0-1.92zm-4.27 1.122a4.109 4.109 0 1 0 0 8.217 4.109 4.109 0 0 0 0-8.217zm0 1.441a2.667 2.667 0 1 1 0 5.334 2.667 2.667 0 0 1 0-5.334z" clip-rule="evenodd"/>
                    </svg>
                    <span class="sr-only">Instagram Account</span>
                </a>
                <a href="#" class="text-white hover:text-gray-900 dark:hover:text-white">
                    <svg class="w-[24px] h-[24px]" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 15 15">
                        <path fill-rule="evenodd" d="M7.979 5v1.586a3.5 3.5 0 0 1 3.082-1.574C14.3 5.012 15 7.03 15 9.655V15h-3v-4.738c0-1.13-.229-2.584-1.995-2.584-1.713 0-2.005 1.23-2.005 2.5V15H5.009V5h2.97ZM3 2.487a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/>
                        <path d="M3 5.012H0V15h3V5.012Z"/>
                    </svg>
                    <span class="sr-only">Dribbble account</span>
                </a>
            </div>
            <span class="text-[14px] text-white sm:text-center dark:text-gray-400 flex justify-end my-[19px] px-[8px]">© 2023 <a href="https://flowbite.com/" class="hover:underline">Explorate™</a>. All Rights Reserved.</span>
        </div>
    </footer>
</template>


<script>
export default {
    name: 'FooterMod'
}
</script>
