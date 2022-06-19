<template>
  <div style="padding: 6rem;">
    <vs-dialog v-model="showModal" ref="comment" width="600px">
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
            <vs-select  v-model="seniority" label-placeholder="Seniority">
              <vs-option label="Junior" value="Junior">Junior</vs-option>
              <vs-option label="Medior" value="Medior">Medior</vs-option>
              <vs-option label="Senior" value="Senior">Senior</vs-option>
            </vs-select>
          </div>
        </div>
        <div class="row mt-4 justify-content-center">
          <div class="col-5">
            <vs-select  v-model="engagement" label-placeholder="Engagement">
              <vs-option label="Full-time" value="Full-time">Full-time</vs-option>
              <vs-option label="Part-time" value="Part-time">Part-time</vs-option>
              <vs-option label="Freelance" value="Freelance">Freelance</vs-option>
              <vs-option label="Internship" value="Internship">Internship</vs-option>
            </vs-select>
          </div>
          <div class="col-5 d-inline-flex">
            <div class="row">
              <div class="col">
                <vs-radio v-model="currentlyEmployed" label="Current employee" :val="1">Current employee</vs-radio>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <vs-radio v-model="currentlyEmployed" label="Former employee" :val="0">Former employee</vs-radio>
              </div>
            </div>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10">
            <vs-input v-model="subject" id="subject" class="subject" label-placeholder="Subject" type="text"></vs-input>
          </div>
        </div>
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-center">
            <textarea class="vs-input" maxlength="3000" style="width: 100%" v-model="content" placeholder="Comment"/>
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
              <vs-button @click="addComment" class="btn-primary">Add comment</vs-button>
            </div>
          </div>
        </div>
      </template>
    </vs-dialog>
    <div v-if="role === 'USER'" class="row justify-content-end">
      <div class="col d-flex justify-content-end">
        <vs-button @click="openModal" class="vs-button--size-large" color="#7dcdec"><strong>Add comment</strong></vs-button>
      </div>
    </div>
    <div class="row">
      <div class="col-4 mt-4" v-for="comment in comments" :key="comment.id">
        <vs-card style="padding-left: 5%;">
          <template #title>
            <div class="row">
              <div class="col">
                <p class="text-lg-end">{{comment.dateCreated | formatDate}} </p>
                <h3 class="text-start" style="font-weight: bold; font-size: x-large">{{comment.subject}}</h3>
              </div>
            </div>
            <div class="row mt-1">
              <div class="col">
                <star-rating :increment="0.5"
                             :max-rating="5"
                             :star-size="25"
                             :read-only="true"
                             :rating="comment.rating"
                             :round-start-rating="false">
                </star-rating>
              </div>
            </div>
          </template>
          <template #text>
            <div class="row mt-3 mb-4">
              <div class="col">
                <p class="text-lg-start" style="font-size: large">{{comment.content}}</p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-end"><i>{{comment.position}}</i></p>
              </div>
            </div>
            <div class="row">
              <div class="col">
                <p class="text-lg-end"><i>{{comment.engagement}}
                  ({{comment.currentlyEmployed ? 'Current employee': 'Former employee'}})</i></p>
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
import StarRating from "vue-star-rating";

export default {
  name: "Comments",
  data() {
    return {
      comments: [],
      showModal: false,
      role: "",
      content: "",
      seniority: "",
      position: "",
      engagement: "",
      currentlyEmployed: 1,
      rating: 3,
      subject: "",
      dateCreated: "",
      industries: [],
      positions: [],
    }
  },
  components:{
    StarRating,
  },
  async mounted() {
    this.$parent.active = 'comments';
    await this.getComments();
    await this.getPositions();
    this.role = this.$store.getters.user?.role?.authority;
  },
  methods: {
    async getComments() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}/comment`).then(response => {
        this.comments = response.data;
      });
    },
    async getIndustries(){
      await axios.get(`${process.env.VUE_APP_BACKEND}/industries.json`,).then(response => {
        this.industries = response.data?.industries
      });
    },
    async getPositions(){
      await axios.get(`${process.env.VUE_APP_BACKEND}/positions.json`,).then(response => {
        this.positions = response.data?.positions.flatMap(position => position.value)
      });
    },
    resetCommentForm() {
      this.content = "";
      this.subject = "";
      this.position = "";
      this.engagement = "";
      this.currentlyEmployed = 1;
      this.seniority = "";
      this.rating = 3;
    },
    openModal(){
      this.showModal = true;
      this.$nextTick(() => {
        const element = document.getElementById("vs-input--subject");
        element.style.width = "100%";
      });
    },
    async addComment(){
      const comment = {
        content: this.content,
        position: `${this.position} (${this.seniority})` ,
        engagement: this.engagement,
        currentlyEmployed: Boolean(this.currentlyEmployed),
        rating: this.rating,
        subject: this.subject,
        companyName: this.$route.params.companyName,
      };
      const loading = this.$vs.loading({
        container: this.$refs.comment,
        color: 'primary',
        scale: 0.6,
        center: true
      });
      await axios.post(`${process.env.VUE_APP_BACKEND}/company/comment`, comment).then(response => {
        loading.close()
        this.$vs.notification({
          color: 'success',
          title: 'Success',
          text: 'Comment added successfully',
          position: 'top-right',
        });
        this.comments.push(response.data);
        this.showModal = false;
        this.resetCommentForm();
      }).catch(error => {
        loading.close()
        this.$vs.notification({
          color: 'danger',
          title: 'Error',
          text: 'Error adding comment',
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
