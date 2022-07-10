<template>
  <div style="padding: 6rem">
    <vs-dialog v-model="showModal" prevent-close not-close ref="jobOffer" width="600px">
      <template #header>
        <div class="col d-flex justify-content-end" style="margin-left: 9rem; margin-top: 0.5rem">
          <h4 modal-title>
            <b>Add job offer</b>
          </h4>
        </div>

        <div class="col d-flex justify-content-end">
          <button type="button" class="close" style="float:right;" @click="closeModal()" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>

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
        <div class="row justify-content-center mt-4">
          <div class="col-10 d-flex justify-content-center mt-2" style="margin-left: 2rem">
            <form v-on:submit.prevent="addSkill">
              <div class="row justify-content-start">
                <div class="col-5 d-flex justify-content-center">
                  <vs-input label-placeholder="Skill" v-model="skill"/>
                </div>
                <div class="col-5 d-flex justify-content-end" style="margin-left: inherit">
                  <vs-button color="dark" type="filled" :disabled="skill.length <= 0 ">Add</vs-button>
                </div>
              </div>
            </form>
          </div>
        </div>
        <div class="row justify-content-start">
          <div class="col justify-content-start" style="margin-left: 3.5rem;">
            <table>
              <th v-if="skills.length > 0" style="text-align: center; border-color: rosybrown;
                border-bottom-style: solid;
                border-bottom-width: thin;
                height: 20px;">
                <p>Skills</p>
              </th>
              <tr v-for="item in skills" v-bind:key="item">
                <td style="height: 10px; width: 50rem; max-width: 50rem; minwidth: 50rem; overflow: hidden; text-align: left">
                  <p>{{ item }}</p>
                </td>
                <td style="height: 10px; width: 100px; max-width: 100px; minwidth: 100px; overflow: hidden;">
                  <vs-button icon @click="removeSkill(item)">
                    <i class='bx bx-minus'></i>
                  </vs-button>
                </td>
              </tr>
            </table>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <div class="row justify-content-end">
            <div class="col d-flex justify-content-end">
              <vs-button @click="addJobOffer" class="btn-primary" style="background: #be1d7b">Add job offer</vs-button>
            </div>
          </div>
        </div>
      </template>
    </vs-dialog>
    <div class="row justify-content-end">
      <div class="col d-flex justify-content-end mt-4">
        <vs-button @click="openModal" class="vs-button--size-large" color="##be1d7b"><strong>Add job offer</strong></vs-button>
      </div>
    </div>
    <h1>My job offers</h1>
    <div class="row">
      <div class="col-4 mt-4" v-for="jobOffer in myJobs" :key="jobOffer.id">
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
            <div class="row mb-2">
              <div class="col">
                <h5 class="text-lg-start" v-if="jobOffer.company == undefined">Employer</h5>
                <h5 class="text-lg-start" v-if="jobOffer.company != undefined">Company</h5>
                <p class="text-lg-start mt-2" v-if="jobOffer.company != undefined" style="font-size: large">{{jobOffer.company}}</p>
                <p class="text-lg-start mt-2" v-if="jobOffer.company == undefined" style="font-size: large ">{{profile.firstName}} {{profile.lastName}}</p>
              </div>
            </div>
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
            <div class="row mt-3 mb-4">
              <div class="col">
                <h5 class="text-lg-start">Skills</h5>
                <div class="row justify-content-start">
                  <div class="col justify-content-start">
                    <table>
                      <tr v-for="item in jobOffer.skills" v-bind:key="item">
                        <td style="overflow: hidden; text-align: left">
                          <p style="font-size: large">{{ item }}</p>
                        </td>
                      </tr>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </vs-card>
      </div>
    </div>
  </div>
</template>



<script>
import axios from 'axios';
import moment from "moment";

export default {
  name: "MyJobOffers",
  data() {
    return {
      active: 'guide',
      currentUserId: null,
      profile: null,
      createPostDialog: false,
      text: '',
      image: '',
      link: '',
      isSearching: false,
      search: '',
      usersFound: [],
      showModal: false,
      jobOffers: [],
      myJobs: [],
      jobs: [],
      position: "",
      seniority: "",
      description: "",
      criteria: "",
      positions: [],
      searchEmpty: true,
      skills: [],
      skill: "",
      newSkillId: 0,
      employer: "",
    }
  },
  async beforeMount() {
    await this.getProfile()
    await this.getPositions();
    await this.getMyJobs();
  },
  async mounted() {
    if ('serviceWorker' in navigator) {
      navigator.serviceWorker.register('/OneSignalSDKWorker.js').then(function (registration) {
      }).catch(function (e) {
        console.log('SW registration failed with error:', e);
      });
    }
    await this.$OneSignal.init({
      appId: process.env.VUE_APP_ONESIGNAL_APP_ID,
      autoResubscribe: true,
      promptOptions: {
        slidedown: {
          prompts: [
            {
              type: "category",
              autoPrompt: true,
              text: {
                actionMessage: "Would you like to receive notifications?",
                acceptButton: "Allow",
                cancelButton: "Cancel",

                negativeUpdateButton: "Cancel",
                positiveUpdateButton: "Save Preferences",
                updateMessage: "Update your push notification subscription preferences.",
              },
              delay: {
                timeDelay: 5
              },
              categories: [
                {
                  tag: "connections",
                  label: "New connection"
                },
                {
                  tag: "messages",
                  label: "New messages"
                },
                {
                  tag: "posts",
                  label: "New posts",
                },
              ]
            }
          ]
        }
      }
    });
    this.$store.commit('setOneSignalToken', process.env.VUE_APP_ONESIGNAL_TOKEN)
    await this.$OneSignal.setExternalUserId(this.profile.id)
    await this.$OneSignal.sendTag('user_id', this.profile.id);
    await this.$OneSignal.sendTag('full_name', this.profile.firstName + ' ' + this.profile.lastName);
  },
  methods : {
    logOut() {
      this.$store.commit('setToken', null);
      this.$store.commit('setUser', null);
      this.$router.push('/');
    },
    async getAllUsers() {
      const loading = this.$vs.loading();
      const response = await axios.get(process.env.VUE_APP_BACKEND + '/security/user').catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting all users',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      loading.close();
      this.$vs.notification({
        title: 'Success',
        text: JSON.stringify(response.data),
        color: 'success',
        position: 'top-right',
        duration: 10000
      });
    },
    async getProfile() {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${this.$store.getters.user?.id}`)
      if (response.data) {
        this.profile = response.data.profile
      }
    },
    async searchProfile() {
      this.checkIfSearchEmpty()
      if (this.search === '') return
      this.isSearching = true
      const response = await axios.get(process.env.VUE_APP_BACKEND + `/profile?search=${this.search}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error searching profiles',
          color: 'danger',
          position: 'top-right'
        });
        this.isSearching = false
        throw error;
      });
      this.isSearching = false
      this.usersFound = response.data.profiles
      //this.changeRoute('/user/search')
    },
    checkIfSearchEmpty() {
      if (this.search === '') {
        this.usersFound = []
        this.searchEmpty = true
      } else
        this.searchEmpty = false
    },
    changeRoute(path) {
      this.$router.replace(path)
    },
    openCreatePostDialog() {
      this.createPostDialog = true
    },
    openModal(){
      this.showModal = true;
    },
    closeModal(){
      this.showModal = false;
      this.resetJobOfferForm();
    },
    async createPost() {
      const newPost = {
        profile: {
          id: this.profile.id,
          firstName: this.profile.firstName,
          lastName: this.profile.lastName
        },
        createdAt: moment().format(),
        content: {
          text: this.text,
          image: this.image,
          links: [this.link]
        }
      }
      const loading = this.$vs.loading();
      const response = await axios.post(process.env.VUE_APP_BACKEND + '/post', newPost).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error while creating post',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      loading.close();
      this.$vs.notification({
        title: 'Success',
        text: 'New post successfully created!',
        color: 'success',
        position: 'top-right',
      });
      this.createPostDialog = false
      await this.sendNotification()
      this.$router.replace('/user/my-posts')
      this.resetCreatePostDialog()
    },
    async sendNotification(){
      const connections = await this.getMyConnections()
      for(let c of connections){
        const notification = {
          app_id: process.env.VUE_APP_ONESIGNAL_APP_ID,
          contents: { en: `New post from ${this.profile?.firstName} ${this.profile?.lastName}\n\n${this.text.substring(0, 100)}${this.text.length > 100 ? '...' : ''}` },
          url: 'https://localhost:7777/user/posts',
          filters: [
            {field: "tag", key: "posts", relation: "=", value: 1},
            {field: "tag", key: "user_id", relation: "=", value: c.subjectId === this.profile.id ? c.issuerId : c.subjectId }
          ]
        }
        await axios.post('https://onesignal.com/api/v1/notifications', notification)
      }
    },
    resetCreatePostDialog() {
      this.text = ''
      this.image = ''
      this.link = ''
    },
    viewProfile(id){
      this.search = "";
      this.searchEmpty = true;
      this.$router.push({name: 'profileInfo', params: { id: id }})
    },
    resetSearch(){
      this.search = "";
      this.searchEmpty = true;
    },
    async addJobOffer(){
      const jobOfferRequest = {
        jobOffer: {
          profileId: this.profile.id,
          company: "",
          position: `${this.position} (${this.seniority})`,
          description: this.description,
          criteria: this.criteria,
          skills: this.skills,
          createdAt: moment().format(),
        }
      };
      const loading = this.$vs.loading({
        container: this.$refs.jobOffer,
        color: 'primary',
        scale: 0.6,
        center: true
      });
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/job-offer`, jobOfferRequest)
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
      await this.getMyJobs();
      this.showModal = false;
      this.resetJobOfferForm();
    },
    async getPositions(){
      await axios.get(process.env.VUE_APP_FRONTEND +'/positions.json').then(response => {
        this.positions = response.data?.positions.flatMap(position => position.value)
      });
    },
    resetJobOfferForm() {
      this.position = "";
      this.description = "";
      this.criteria = "";
      this.seniority = "";
      this.skill = "";
      this.skills = [];
    },
    addSkill() {
      this.skills.push(this.skill);
      this.skill = "";
    },
    removeSkill(s){
      for(let i=0; i<this.skills.length; i++){
        if(this.skills[i] == s){
          this.skills.splice(i, 1);
          this.showModal = true;
        }
      }
    },
    async getMyJobs() {
      const myJobsRequest = {
        profileId: this.profile.id,
      };
      await axios.get(`${process.env.VUE_APP_BACKEND}/job-offer/profile/${this.profile.id}`, myJobsRequest)
        .then(response => {
          this.myJobs = response.data.jobOffers;
        })
    },
  },
}

</script>

<style scoped>

</style>
