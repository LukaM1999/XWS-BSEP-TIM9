<template>
  <div class="row" style="padding: 5%">
    <vs-dialog v-if="company" v-model="showModal" ref="comment" width="600px">
      <template #header>
        <h4>
          <b>Add company</b>
        </h4>
      </template>
      <div class="con-form">
        <div class="row mt-4 justify-content-center">
          <div class="col-5">
            <vs-input label-placeholder="Name" v-model="company.name"></vs-input>
          </div>
          <div class="col-5">
            <vs-input label-placeholder="Email" v-model="company.email"></vs-input>
          </div>
        </div>
        <div class="row mt-4 justify-content-center">
          <div class="col-5">
            <vs-input label-placeholder="Web site" v-model="company.website"></vs-input>
          </div>
          <div class="col-5">
            <vs-input label-placeholder="Phone" v-model="company.phone"></vs-input>
          </div>
        </div>
        <div class="row mt-4 justify-content-center">
          <div class="col-10">
            <vs-input label-placeholder="Address" id="address" v-model="company.address"></vs-input>
          </div>
        </div>
        <div class="row mt-4 justify-content-center">
          <div class="col-5">
            <vs-input label-placeholder="City" v-model="company.city"></vs-input>
          </div>
          <div class="col-5">
            <vs-input label-placeholder="Country" v-model="company.country"></vs-input>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10">
            <textarea class="vs-input" placeholder="Description" style="width:100%; height: 8rem;" v-model="company.description" id="description"  type="text"></textarea>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-start">
            <vs-select label-placeholder="Size" v-model="company.size"  :multiple="false">
              <vs-option value="<20" label="<20">&lt;20</vs-option>
              <vs-option value="20-50" label="20-50">20-50</vs-option>
              <vs-option value="51-100" label="51-100">51-100</vs-option>
              <vs-option value="101-250" label="101-250">101-250</vs-option>
              <vs-option value="251-500" label="251-500">251-500</vs-option>
              <vs-option value="501-1000" label="501-1000">501-1000</vs-option>
              <vs-option value="1000+" label="1000+">1000+</vs-option>
            </vs-select>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <div class="row justify-content-end">
            <div class="col d-flex justify-content-end">
              <vs-button @click="addCompany" class="btn-primary">Save</vs-button>
            </div>
          </div>
        </div>
      </template>
    </vs-dialog>
    <h1 class="text-lg-center mt-1" style="">My companies</h1>
    <div class="row mt-1 justify-content-end">
      <div class="col d-flex justify-content-end">
        <vs-button @click="openModal" class="vs-button--size-large" color="#7dcdec"><strong>Add company</strong></vs-button>
      </div>
    </div>
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
  name: "OwnerCompanies",
  data() {
    return {
      companies: [],
      company: {
        name: "",
        email: "",
        website: "",
        phone: "",
        address: "",
        city: "",
        country: "",
        description: "",
        size: "",
        industry: "",
        ownerUsername: ""
      },
      companySearch: "",
      showModal: false,
    }
  },
  async mounted() {
    await this.getAllCompanies();
    this.$parent.active = 'my-companies';
    this.company.ownerUsername = this.$store.getters.user?.username;
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
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/company/owner/${this.$store.getters.user?.username}`);
      if (response.status === 200) {
        this.companies = response.data;
      }
    },
    openModal(){
      this.showModal = true;
      this.$nextTick(() => {
        const element = document.getElementById("vs-input--address");
        element.style.width = "100%";
      });
    },
    async addCompany(){
      await axios.post(`${process.env.VUE_APP_BACKEND}/company`, this.company)
      .then(response => {
        if (response.status === 201) {
          this.showModal = false;
          this.companies.push(response.data);
          this.$vs.notification({
            title: "Success",
            text: "Company added successfully!",
            color: "success",
            position: "top-right"
          });
        }
      })
      .catch(error => {
        this.$vs.notification({
          title: "Error",
          text: "Error adding company",
          color: "danger",
          position: "top-right"
        });
        throw error;
      });
    },

  }
}
</script>

<style scoped>

</style>
