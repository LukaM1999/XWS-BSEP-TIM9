<template>
  <div class="row" style="padding: 5%">
    <h1 class="text-lg-center" style="">Company requests</h1>
    <div class="row mt-2 justify-content-center">
      <div class="col-12 d-flex justify-content-center">
        <vs-input type="search" v-model="companySearch" warn label-placeholder="Search companies..." />
      </div>
    </div>
    <div class="col-4 mt-4" v-for="company in filteredCompanies" :key="company.id">
      <vs-card style="padding-left: 5%;">
        <template #title>
          <h3>{{company.name}}</h3>
        </template>
        <template #text>
          <p class="text-lg-start">Owner: <strong>{{company.ownerUsername}}</strong></p>
          <p class="text-lg-start">Address: <strong>{{company.country}} {{company.city}}, {{company.address}}</strong></p>
          <p class="text-lg-start">Employees: <strong>{{company.size}}</strong></p>
          <p class="text-lg-start">Industry: <strong>{{company.industry}}</strong></p>
          <br>
          <p class="text-lg-start">{{company.description}}</p>
          <br>
          <div class="row justify-content-end">
            <div class="col d-flex justify-content-end">
              <vs-button @click="declineCompany(company)" dark size="lg">Decline</vs-button>
            </div>
            <div class="col d-flex justify-content-start">
              <vs-button @click="approveCompany(company)" primary size="lg">Approve</vs-button>
            </div>
          </div>
        </template>
      </vs-card>
    </div>
  </div>

</template>


<script>
import axios from "axios";

export default {
  name: "Requests",
  data() {
    return {
      companies: [],
      companySearch: "",
    }
  },
  async mounted() {
    await this.getAllCompanies();
    this.$parent.active = 'requests';
  },
  computed: {
    filteredCompanies() {
      return this.companies.filter(company => {
        return company.name.toLowerCase().includes(this.companySearch.toLowerCase())
      })
    },
  },
  methods: {
    async getAllCompanies() {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/admin/company`);
      if (response.status === 200) {
        this.companies = response.data;
      }
    },
    async approveCompany(company) {
      const loading = this.$vs.loading()
      await axios.patch(`${process.env.VUE_APP_BACKEND}/admin/company/${company.name}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Failed to approve company',
          color: 'danger',
          position: 'top-right',
          time: 5000
        });
        loading.close()
        throw error
      });
      this.$vs.notification({
        title: 'Success',
        text: 'Company approved!',
        color: 'success',
        position: 'top-right',
        time: 5000
      });
      company.approved = true;
      loading.close()
      this.companies = this.companies.filter(c => c.id !== company.id);
    },
    async declineCompany(company) {
      const loading = this.$vs.loading()
      await axios.delete(`${process.env.VUE_APP_BACKEND}/admin/company/${company.name}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Failed to decline company',
          color: 'danger',
          position: 'top-right',
          time: 5000
        });
        loading.close()
        throw error
      });
      this.$vs.notification({
        title: 'Success',
        text: 'Company declined successfully',
        color: 'success',
        position: 'top-right',
        time: 5000
      });
      loading.close()
      this.companies = this.companies.filter(c => c.id !== company.id);
    },
  }
}
</script>

<style scoped>

</style>
