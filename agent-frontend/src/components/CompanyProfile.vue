<template>
 <div style="overflow: hidden;">
   <vs-navbar :color="'guide'" fixed shadow center-collapsed v-model="active">
     <template #left>
       <div class="row">
         <div class="col">
           <img @click="$router.push('/')" style="cursor: pointer" src="/logo.png" width="100" height="100" alt="">
         </div>
         <div class="col align-self-center">
           <p @click="$router.push('/')"
             style="font-family: 'Bauhaus 93'; margin-bottom: 0rem;
             margin-left:-2rem ;  font-size: xxx-large; color: #048ce3; cursor: pointer">
             AGENTY</p>
         </div>
       </div>
     </template>
     <template #right>
       <vs-navbar-item :to="`/company-profile/${companyName}/overview`"
                       :active="active === 'overview'" id="overview">
         Overview
       </vs-navbar-item>
       <vs-navbar-item :to="`/company-profile/${companyName}/comments`"
                       :active="active === 'comments'" id="comments">
         Comments
       </vs-navbar-item>
       <vs-navbar-item :to="`/company-profile/${companyName}/salaries`"
                       :active="active === 'salaries'" id="salaries">
         Salaries
       </vs-navbar-item>
       <vs-navbar-item :to="`/company-profile/${companyName}/interviews`"
                       :active="active === 'interviews'" id="interviews">
         Interviews
       </vs-navbar-item>
       <vs-navbar-item :to="`/company-profile/${companyName}/job-offers`"
                       :active="active === 'jobOffers'" id="jobOffers">
         Job offers
       </vs-navbar-item>
     </template>
   </vs-navbar>
   <router-view></router-view>
 </div>
</template>

<script>
import axios from "axios";

export default {
  name: "CompanyProfile",
  data() {
    return {
      companyName: null,
      company: null,
      active: 'overview',
    }
  },
  async mounted() {
    this.companyName = this.$route.params.companyName;
    await this.getCompany();
  },
  methods: {
    async getCompany() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.companyName}`).then(response => {
        this.company = response.data;
      });
    }
  },
}
</script>

<style scoped>

</style>
