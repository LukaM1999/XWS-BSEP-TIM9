<template>
  <div style="overflow-x: hidden">
    <vs-navbar :color="'guide'" fixed shadow center-collapsed>
      <template #left>
        <div class="row">
          <div class="col">
            <img @click="$router.push('/')" style="cursor: pointer" src="/logo.png" width="100" height="100" alt="">
          </div>
          <div class="col align-self-center">
            <p @click="$router.push('/')"
              style="font-family: 'Bauhaus 93'; cursor: pointer;
               margin-bottom: 0rem; margin-left:-2rem ;
               font-size: xxx-large; color: #048ce3">
              AGENTY
            </p>
          </div>
        </div>
      </template>
      <template #right>
        <vs-navbar-item v-if="role === 'COMPANY_OWNER'" :to="`/home/my-companies`"
                        :active="active === 'my-companies'" id="my-companies">
          My companies
        </vs-navbar-item>
        <vs-navbar-item :to="`/home/companies`"
                        :active="active === 'companies'" id="companies">
          Companies
        </vs-navbar-item>
        <vs-navbar-item :to="`/home/profile`"
                        :active="active === 'profile'" id="profile">
          Profile
        </vs-navbar-item>
        <vs-button class="vs-button--size-large" @click="$router.push('/')" flat >Log out</vs-button>
      </template>
    </vs-navbar>
    <router-view></router-view>
  </div>

</template>

<script>
import axios from "axios";
export default {
  name: "HomePage",
  data() {
    return {
      active: 'companies',
      companies: [],
      showModal: false,
      role: "",
      industries: [],
      positions: [],
      companySearch: "",
    }
  },
  async mounted() {
    await this.getAllCompanies();
    this.role = this.$store.getters.user?.role?.authority;
  },
  computed:{
    //Search by company name
    filteredCompanies() {
      return this.companies.filter(company => {
        return company.name.toLowerCase().includes(this.companySearch.toLowerCase())
      })
    },
  },
  methods: {
    async getAllCompanies(){
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/company`);
      if(response.status === 200){
        this.companies = response.data;
      }
    },
  }
}
</script>

<style scoped>

</style>
