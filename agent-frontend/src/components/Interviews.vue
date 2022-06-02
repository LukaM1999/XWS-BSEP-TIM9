<template>
  <div style="padding: 6rem">
    <vs-dialog v-model="showModal" ref="interview" width="600px">
      <template #header>
        <h4 class="not-margin">
          <b>Add comment</b>
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
            <vs-select  v-model="year" label-placeholder="Year">
              <vs-option v-for="y in years" :label="y" :value="y" :key="y">
                {{y}}
              </vs-option>
            </vs-select>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10">
            <vs-input v-model="subject" id="subject" class="subject" label-placeholder="Subject" type="text"></vs-input>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-center">
            <textarea class="vs-input" maxlength="3000" style="width: 100%" v-model="hr"
                      placeholder="Describe HR interview..."/>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-center">
            <textarea class="vs-input" maxlength="3000" style="width: 100%" v-model="technical"
                      placeholder="Describe technical interview..."/>
          </div>
        </div>
        <div class="row mt-4 justify-content-center">
          <div class="col-5">
            <vs-select id="duration" v-model="duration" label-placeholder="Duration">
              <vs-option v-for="i in 20" :label="i===1 ? `${i} week` : `${i} weeks` " :value="i" :key="i">
                {{i}} {{i===1 ? 'week' : 'weeks'}}
              </vs-option>
            </vs-select>
          </div>
          <div class="col-5">
            <vs-select  v-model="difficulty" label-placeholder="Difficulty">
              <vs-option label="Easy" value="1" key="1">Easy</vs-option>
              <vs-option label="Medium" value="2" key="2">Medium</vs-option>
              <vs-option label="Hard" value="3" key="3">Hard</vs-option>
            </vs-select>
          </div>
        </div>
        <div class="row mt-4 justify-content-start ms-4">
          <div class="col-3 ms-3">
            <vs-switch v-model="acceptedOffer">
              <template #off>
                Declined offer
              </template>
              <template #on>
                Accepted offer
              </template>
            </vs-switch>
          </div>
        </div>
        <div class="row justify-content-start mt-3">
          <div class="col-4">
            <p>Rating</p>
          </div>
        </div>
        <div class="row justify-content-center" style="margin-top: -2%">
          <div class="col-10">
            <star-rating id="rating"
                         :increment="1"
                         :max-rating="5"
                         :star-size="25"
                         rating="3"
                         v-model="rating">
            </star-rating>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <div class="row justify-content-end">
            <div class="col d-flex justify-content-end">
              <vs-button @click="addInterview" class="btn-primary">Add interview</vs-button>
            </div>
          </div>
        </div>
      </template>
    </vs-dialog>
    <div v-if="role === 'USER'" class="row justify-content-end">
      <div class="col d-flex justify-content-end">
        <vs-button @click="openModal" class="vs-button--size-large" color="#7dcdec"><strong>Add interview</strong></vs-button>
      </div>
    </div>
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
</template>

<script>
import StarRating from "vue-star-rating";
import axios from "axios";
import moment from "moment";

export default {
  name: "Interviews",
  data() {
    return {
      interviews: [],
      showModal: false,
      role: "",
      subject: "",
      hr: "",
      technical: "",
      position: "",
      difficulty: 1,
      rating: 3,
      years: [],
      year: "",
      duration: 1,
      acceptedOffer: true,
      positions: [],
    }
  },
  components:{
    StarRating,
  },
  mounted() {
    this.$parent.active = 'interviews';
    this.role = this.$store.getters.user?.role?.authority;
    this.getPositions();
    this.getInterviews();
    for (let i = 0; i < 10; i++) {
      this.years.push(moment().subtract(i, 'years').format('YYYY'));
    }
  },
  methods: {
    async getInterviews() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}/interview`).then(response => {
        this.interviews = response.data;
      });
    },
    resetInterviewForm() {
      this.subject = "";
      this.hr = "";
      this.technical = "";
      this.position = "";
      this.difficulty = 1;
      this.rating = 3;
      this.year = "";
      this.duration = 1;
      this.acceptedOffer = true;
    },
    openModal(){
      this.showModal = true;
      this.$nextTick(() => {
        const element = document.getElementById("vs-input--subject");
        element.style.width = "100%";
      });
    },
    async getPositions(){
      await axios.get(`${process.env.VUE_APP_BACKEND}/positions.json`,).then(response => {
        this.positions = response.data?.positions.flatMap(position => position.value)
      });
    },
    async addInterview(){
      const interview = {
        companyName: this.$route.params.companyName,
        subject: this.subject,
        hr: this.hr,
        technical: this.technical,
        position: this.position,
        difficulty: this.difficulty,
        rating: this.rating,
        year: this.year,
        duration: this.duration,
        acceptedOffer: this.acceptedOffer,
      };
      const loading = this.$vs.loading({
        container: this.$refs.interview,
        color: 'primary',
        scale: 0.6,
        center: true
      });
      await axios.post(`${process.env.VUE_APP_BACKEND}/company/interview`, interview).then(response => {
        loading.close()
        this.$vs.notification({
          color: 'success',
          title: 'Success',
          text: 'Interview added successfully!',
          position: 'top-right',
        });
        this.interviews.push(response.data);
        this.showModal = false;
        this.resetInterviewForm();
      }).catch(error => {
        loading.close()
        this.$vs.notification({
          color: 'danger',
          title: 'Error',
          text: 'Error adding interview',
          position: 'top-right',
        });
        throw error;
      });
    }
  },

}
</script>

<style scoped>

</style>
