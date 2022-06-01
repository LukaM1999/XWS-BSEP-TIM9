<template>
  <div style="padding: 6rem">
    <div>
      <div class="row">
        <div class="col-4 mt-4" v-for="interview in interviews" :key="interview.id">
          <vs-card style="padding-left: 5%;">
            <template #title>
              <div class="row">
                <div class="col">
                  <p class="text-lg-end">{{interview.dateCreated | formatDate}} </p>
                  <h3 class="text-start" style="font-weight: bold; font-size: x-large">{{interview.subject}}</h3>
                </div>
              </div>
              <div class="row mt-1">
                <div class="col">
                  <star-rating :increment="0.5"
                               :max-rating="5"
                               :star-size="25"
                               :read-only="true"
                               :rating="interview.rating"
                               :round-start-rating="false">
                  </star-rating>
                </div>
              </div>
            </template>
            <template #text>
              <div class="row mt-3 mb-4">
                <div class="col">
                  <h5 class="text-lg-start">HR interview</h5>
                  <p class="text-lg-start mt-2" style="font-size: large">{{interview.hr}}</p>
                </div>
              </div>
              <div class="row mt-3 mb-4">
                <div class="col">
                  <h5 class="text-lg-start">Technical interview</h5>
                  <p class="text-lg-start mt-2" style="font-size: large">{{interview.technical}}</p>
                </div>
              </div>
              <div class="row">
                <div class="col">
                  <p class="text-lg-end"><i>{{interview.position}} ({{interview.acceptedOffer ? 'Accepted offer': 'Declined offer'}})</i></p>
                </div>
              </div>
              <div class="row">
                <div class="col">
                  <p class="text-lg-end"><i>Interview from year {{interview.year}}</i></p>
                </div>
              </div>
            </template>
          </vs-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import StarRating from "vue-star-rating";
import axios from "axios";

export default {
  name: "Interviews",
  data() {
    return {
      interviews: [],
    }
  },
  components:{
    StarRating,
  },
  mounted() {
    this.$parent.active = 'interviews';
    this.getInterviews();
  },
  methods: {
    async getInterviews() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}/interview`).then(response => {
        this.interviews = response.data;
      });
    },
  },

}
</script>

<style scoped>

</style>
