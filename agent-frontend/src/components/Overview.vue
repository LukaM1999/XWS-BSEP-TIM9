<template>
<div class="row justify-content-center" style="padding: 6rem">
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
    }
  },
  components:{
    StarRating,
  },
  async mounted() {
    this.$parent.active = 'overview';
    await this.getCompany();
  },
  methods: {
    async getCompany() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}`).then(response => {
        this.company = response.data;
      });
    }
  },
}
</script>

<style scoped>

</style>
