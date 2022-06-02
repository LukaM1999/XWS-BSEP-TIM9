<template>
<div class="row justify-content-center" style="padding: 6rem">
  <vs-dialog v-if="company" v-model="showModal" ref="comment" width="600px">
    <template #header>
      <h4 class="not-margin">
        <b>Edit company</b>
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
          <textarea class="vs-input" style="width:100%; height: 8rem;" v-model="company.description" id="description" title="" type="text"></textarea>
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
            <vs-button @click="editCompany" class="btn-primary">Save changes</vs-button>
          </div>
        </div>
      </div>
    </template>
  </vs-dialog>
  <div class="col d-flex justify-content-center">
    <vs-card style="padding-left: 5%;" v-if="company">
      <template #title>
        <h3 style="font-weight: bold; font-size: x-large">{{company.name}}</h3>
      </template>
      <template #text>
        <div class="row justify-content-center mb-3">
          <div class="col d-flex justify-content-center">
            <star-rating :increment="0.5"
                         :max-rating="5"
                         :star-size="25"
                         :read-only="true"
                         :rating="company.rating"
                         :round-start-rating="false">
            </star-rating>
            <div class="row" style="margin-top: 0.25%; margin-left: 1px">
              <div class="col">
                <p style="font-size: small">({{company.ratingCount}} ratings)</p>
              </div>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col border-end border-3">
            <div class="row">
              <div class="col">
               <p>ADDRESS</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-center"><i><strong>{{company.country}} {{company.city}}, {{company.address}}</strong></i></p>
              </div>
            </div>
          </div>
          <div class="col border-end border-3">
            <div class="row">
              <div class="col">
                <p>EMPLOYEES</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-center"><i><strong>{{company.size}}</strong></i></p>
              </div>
            </div>
          </div>
          <div class="col border-end border-3">
            <div class="row">
              <div class="col">
                <p>INDUSTRY</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-center"><i><strong>{{company.industry}}</strong></i></p>
              </div>
            </div>
          </div>
          <div class="col border-end border-3">
            <div class="row">
              <div class="col">
                <p>EMAIL</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-center"><i><strong>{{company.email}}</strong></i></p>
              </div>
            </div>
          </div>
          <div class="col border-end border-3">
            <div class="row">
              <div class="col">
                <p>WEBSITE</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <a :href="`https://${company.website}`">{{company.website}}</a>
              </div>
            </div>
          </div>
          <div class="col border-end border-3">
            <div class="row">
              <div class="col">
                <p>PHONE</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-center"><i><strong>{{company.phone}}</strong></i></p>
              </div>
            </div>
          </div>
          <div class="col">
            <div class="row">
              <div class="col">
                <p>ESTABLISHED</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-center"><i><strong>{{company.yearEstablished}}</strong></i></p>
              </div>
            </div>
          </div>
        </div>
        <br>
        <div class="row">
          <div class="col">
            <p class="text-lg-start">{{company.description}}</p>
          </div>
        </div>
        <div class="row justify-content-end">
          <div class="col d-flex justify-content-end">
            <vs-button v-if="role === 'COMPANY_OWNER'" @click="openModal" class="btn-primary">Edit company</vs-button>
          </div>
        </div>
      </template>
    </vs-card>
  </div>
</div>
</template>

<script>
import axios from "axios";
import StarRating from 'vue-star-rating'

export default {
  name: "Overview",
  data() {
    return {
      company: null,
      role: "",
      showModal: false,
    }
  },
  components:{
    StarRating,
  },
  async mounted() {
    this.$parent.active = 'overview';
    this.role = this.$store.getters.user?.role?.authority;
    await this.getCompany();
  },
  methods: {
    async getCompany() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}`).then(response => {
        this.company = response.data;
      });
    },
    async editCompany() {
      await axios.patch(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}`, this.company).then(response => {
        this.company = response.data;
        this.showModal = false;
        this.$vs.notification({
          title: 'Success',
          text: 'Company updated successfully!',
          color: 'success',
          position: 'top-right'
        });
      }).catch (error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Something went wrong',
          color: 'danger',
          position: 'top-right'
        });
      });
    },
    openModal(){
      this.showModal = true;
      this.$nextTick(() => {
        const element = document.getElementById("vs-input--address");
        element.style.width = "100%";
      });
    },
  },
}
</script>

<style scoped>

</style>
