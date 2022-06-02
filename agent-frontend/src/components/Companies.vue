<template>
  <div class="row" style="padding: 5%">
    <h1 class="text-lg-center" style="">Companies</h1>
    <div class="row mt-2 justify-content-center">
      <div class="col-12 d-flex justify-content-center">
        <vs-input type="search" v-model="companySearch" warn label-placeholder="Search companies..." />
      </div>
    </div>
    <div class="col-4 mt-4" v-for="company in filteredCompanies" :key="company.id">
      <vs-card @click="$router.push({name: 'company-profile',  params: {companyName: company.name}})" style="padding-left: 5%;">
        <template #title>
          <h3>{{company.name}}</h3>
        </template>
        <template #text>
          <p class="text-lg-start">Address: <strong>{{company.country}} {{company.city}}, {{company.address}}</strong></p>
          <p class="text-lg-start">Employees: <strong>{{company.size}}</strong></p>
          <p class="text-lg-start">Industry: <strong>{{company.industry}}</strong></p>
          <br>
          <p class="text-lg-start">{{company.description}}</p>
        </template>
      </vs-card>
    </div>
  </div>

</template>

<script>
import axios from "axios";
export default {
  name: "Companies",
  data() {
    return {
      companies: [],
      companySearch: "",
    }
  },
  async mounted() {
    await this.getAllCompanies();
    this.$parent.active = 'companies';
  },
  computed: {
    //Search by company name
    filteredCompanies() {
      return this.companies.filter(company => {
        return company.name.toLowerCase().includes(this.companySearch.toLowerCase())
      })
    },
  },
  methods: {
    async getAllCompanies() {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/company`);
      if (response.status === 200) {
        this.companies = response.data;
      }
    },
  }
}

</script>

<style scoped>

</style>
