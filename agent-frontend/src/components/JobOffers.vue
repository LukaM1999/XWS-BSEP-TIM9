<template>
  <div style="padding: 6rem">
    <vs-dialog v-model="showModal" ref="jobOffer" width="600px">
      <template #header>
        <h4 class="not-margin">
          <b>Add job offer</b>
        </h4>
      </template>
      <div class="con-form">
        <div class="row mt-2 justify-content-center">
          <div class="col-5">
            <vs-select  v-model="position" label-placeholder="Position">
              <vs-option v-for="position in positions" :label="position" :value="position" :key="position">
                {{position}}
              </vs-option>
            </vs-select>
          </div>
          <div class="col-5">
            <vs-select  v-model="seniority" label-placeholder="Seniority">
              <vs-option label="Junior" value="Junior">Junior</vs-option>
              <vs-option label="Medior" value="Medior">Medior</vs-option>
              <vs-option label="Senior" value="Senior">Senior</vs-option>
            </vs-select>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-center">
            <textarea class="vs-input" maxlength="3000" style="width: 100%" v-model="description"
                      placeholder="Write job description and everyday activities..."/>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-center">
            <textarea class="vs-input" maxlength="3000" style="width: 100%" v-model="criteria"
                      placeholder="Write job criteria..."/>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <div class="row justify-content-end">
            <div class="col d-flex justify-content-end">
              <vs-button @click="addJobOffer" class="btn-primary">Add job offer</vs-button>
            </div>
          </div>
        </div>
      </template>
    </vs-dialog>
    <div v-if="role === 'COMPANY_OWNER'" class="row justify-content-end">
      <div class="col d-flex justify-content-end">
        <vs-button @click="openModal" class="vs-button--size-large" color="#7dcdec"><strong>Add job offer</strong></vs-button>
      </div>
    </div>
    <div class="row">
      <div class="col-4 mt-4" v-for="jobOffer in jobOffers" :key="jobOffer.id">
        <vs-card style="padding-left: 5%;">
          <template #title>
            <div class="row">
              <div class="col">
                <p class="text-lg-end">{{jobOffer.createdAt | formatDate}} </p>
                <h3 class="text-start" style="font-weight: bold; font-size: x-large">{{jobOffer.position}}</h3>
              </div>
            </div>
          </template>
          <template #text>
            <div class="row mt-3 mb-4">
              <div class="col">
                <h5 class="text-lg-start">Description</h5>
                <p class="text-lg-start mt-2" style="font-size: large">{{jobOffer.description}}</p>
              </div>
            </div>
            <div class="row mt-3 mb-4">
              <div class="col">
                <h5 class="text-lg-start">Criteria</h5>
                <p class="text-lg-start mt-2" style="font-size: large">{{jobOffer.criteria}}</p>
              </div>
            </div>
            <div class="row justify-content-end">
              <div class="col d-flex justify-content-end">
                <vs-button v-if="role === 'COMPANY_OWNER'" @click="promoteJobOffer(jobOffer)" :disabled="jobOffer.promoted" class="btn-primary">Promote</vs-button>
              </div>
            </div>
          </template>
        </vs-card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";

export default {
  name: "JobOffers",
  data() {
    return {
      jobOffers: [],
      showModal: false,
      role: "",
      position: "",
      seniority: "",
      description: "",
      criteria: "",
      positions: [],
    }
  },
  mounted() {
    this.$parent.active = 'jobOffers';
    this.role = this.$store.getters.user?.role?.authority;
    this.getPositions();
    this.getJobOffers();
  },
  methods: {
    async getJobOffers() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}/jobOffer`).then(response => {
        this.jobOffers = response.data;
      });
    },
    resetJobOfferForm() {
      this.position = "";
      this.description = "";
      this.criteria = "";
      this.seniority = "";
    },
    openModal(){
      this.showModal = true;
    },
    async getPositions(){
      await axios.get(`${process.env.VUE_APP_BACKEND}/positions.json`,).then(response => {
        this.positions = response.data?.positions.flatMap(position => position.value)
      });
    },
    async addJobOffer(){
      const jobOffer = {
        companyName: this.$route.params.companyName,
        position: `${this.position} (${this.seniority})`,
        description: this.description,
        criteria: this.criteria,
      };
      const loading = this.$vs.loading({
        container: this.$refs.jobOffer,
        color: 'primary',
        scale: 0.6,
        center: true
      });
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/company/jobOffer`, jobOffer)
      .catch(error => {
        loading.close()
        this.$vs.notification({
          color: 'danger',
          title: 'Error',
          text: 'Error adding job offer',
          position: 'top-right',
        });
        throw error;
      });
      loading.close()
      this.$vs.notification({
        color: 'success',
        title: 'Success',
        text: 'Job offer added successfully!',
        position: 'top-right',
      });
      this.jobOffers.push(response.data);
      this.showModal = false;
      this.resetJobOfferForm();
    },
    async promoteJobOffer(jobOffer){
      const promotionRequest = {
        jobOffer: {
          id: 0,
          company: jobOffer.companyName,
          profileId: null,
          position: jobOffer.position,
          description: jobOffer.description,
          criteria: jobOffer.criteria,
          createdAt: moment(jobOffer.createdAt).toISOString(),
        },
        username: this.$store.getters.user?.dislinktUsername,
        token: this.$store.getters.user?.dislinktToken,
      };
      const loading = this.$vs.loading({
        container: this.$refs.jobOffer,
        color: 'primary',
        scale: 0.6,
        center: true
      });
      await axios.post(`${process.env.VUE_APP_DISLINKT}/job-offer/promote-job`, promotionRequest)
        .catch(error => {
          loading.close()
          this.$vs.notification({
            color: 'danger',
            title: 'Error',
            text: 'Error promoting job offer',
            position: 'top-right',
          });
          throw error;
        });
      await axios.patch(`${process.env.VUE_APP_BACKEND}/company/${jobOffer.companyName}/jobOffer/${jobOffer.id}`)
        .catch(error => {
          loading.close()
          this.$vs.notification({
            color: 'danger',
            title: 'Error',
            text: 'Error promoting job offer',
            position: 'top-right',
          });
          throw error;
        });
      loading.close()
      this.$vs.notification({
        color: 'success',
        title: 'Success',
        text: 'Job offer promoted successfully!',
        position: 'top-right',
      });
      jobOffer.promoted = true;
    }
  },

}
</script>

<style scoped>

</style>
