<template>
  <div>
    <div class="center examplex">
      <vs-navbar target-scroll="#padding-scroll-content" style="background-color: lavenderblush;" padding-scroll
                 center-collapsed v-model="active" height="fixed"
                 not-line>
        <template #left>
          <div class="row">
            <div class="col">
              <img src="/logo.png" width="100" height="100" alt="">
            </div>
            <div class="col align-self-center">
              <p
                style="font-family: 'Bauhaus 93'; margin-bottom: 0rem; margin-left:-2rem ;  font-size: xxx-large; color: #be1d7b">
                DISLINKT</p>
            </div>
          </div>
        </template>
        <div class="row d-flex justify-content-center">
          <div class="col">
            <vs-navbar-item :active="active == 'guide'" @click="resetSearch()" id="guide" to="/user/posts">
              For you
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'myPosts'" @click="resetSearch()" id="myPosts" to="/user/my-posts">
              My posts
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'myJobOffers'" @click="resetSearch()" id="myJobOffers" to="/user/my-job-offers">
              My job offers
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'docs'" @click="resetSearch()" id="docs" to="/user/connections">
              Connections
            </vs-navbar-item>
            <vs-navbar-item :active="active === 'jobOffers'" @click="resetSearch()" id="jobOffers" to="/user/job-offers">
              Job offers
            </vs-navbar-item>
            <vs-navbar-item :active="active === 'messages'" @click="resetSearch()" id="messages" to="/user/messages">
              Messages
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'license'" @click="resetSearch()" id="license" to="/user/profile">
              Profile
            </vs-navbar-item>
          </div>
        </div>
        <template #right>
          <vs-tooltip>
            <vs-button icon flat @click="openCreatePostDialog">
              <i class='bx bx-plus'></i>
            </vs-button>
            <template #tooltip>
              Create new post
            </template>
          </vs-tooltip>
          <div class="center content-inputs">
            <vs-input state="primary" primary v-model="search" @input="searchProfile"
                      placeholder="Search users" :loading="isSearching">
              <template #icon>
                <i class='bx bx-search'></i>
              </template>
            </vs-input>
          </div>
          <vs-button @click="logOut()">Log out</vs-button>
        </template>
      </vs-navbar>
    </div>
    <vs-dialog :prevent-close="true" @close="resetCreatePostDialog" auto-width v-model="createPostDialog">
      <template #header>
        <h4 class="not-margin me-3 ms-3">
          Create post
        </h4>
      </template>
      <div class="con-form" style="display: inline-block;">
        <div class="form-group">
          <div class="d-flex align-items-left">
            <label for="text">Add text</label>
          </div>
          <textarea class="mb-4 form-control vs-input" v-model="text" style="border-radius: 15px" id="text"
                    name="text"></textarea>
        </div>
        <div class="form-group">
          <div class="d-flex align-items-left">
            <label for="img">Add image</label>
          </div>
          <input class="mb-4 form-control" style="border-radius: 15px" type="file" id="img" name="img"/>
        </div>
        <div class="form-group">
          <div class="d-flex align-items-left">
            <label for="link">Add link</label>
          </div>
          <input class="mb-4 form-control" v-model="link" style="border-radius: 15px" type="text" id="link"
                 name="link"/>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <vs-button block @click="createPost">
            Save
          </vs-button>
        </div>
      </template>
    </vs-dialog>
    <div class="row">
      <div class="col">
        <router-view v-if="searchEmpty"></router-view>
        <div v-if="!searchEmpty" style="margin-top: 10%" class="row">
          <div class="col"></div>
          <div class="col" style="border-radius: 15px; background-color: lavenderblush">
            <div :key="i" v-for="(tr, i) in usersFound">
              <div class="row" style="margin-bottom: 5px; margin-top: 5px">
                <div class="col"></div>
                <div class="col-10">
                  <div class="row" style="margin-bottom: 5px; margin-top: 5px">
                    <div class="col d-flex justify-content-center align-self-center">
                      <div class="center con-avatars">
                        <vs-avatar circle primary size="35">
                          <template #text>
                            {{ tr.firstName }} {{ tr.lastName }}
                          </template>
                        </vs-avatar>
                      </div>
                    </div>
                    <div class="col d-flex justify-content-center align-self-end" style="color: black">
                      <vs-navbar-item style="font-size: large; cursor: default">{{ tr.firstName }} {{ tr.lastName }}</vs-navbar-item>
                    </div>
                    <div class="col d-flex justify-content-center">
                      <vs-button @click="viewProfile(tr.id)" dark>Profile</vs-button>
                    </div>
                  </div>
                </div>
                <div class="col"></div>
              </div>
            </div>
          </div>
          <div class="col"></div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>

import axios from 'axios';
import moment from "moment";

export default {
  name: "HomePage",
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
      searchEmpty: true
    }
  },
  async beforeMount() {
    await this.getProfile()
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
  methods: {
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
    async getMyConnections() {
      const response = await axios.get(process.env.VUE_APP_BACKEND + `/connection/${this.$store.getters.user?.id}`)
        .catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error getting connections',
            color: 'danger',
            position: 'top-right'
          });
          throw error;
        });
      return response.data.connections
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
  }
}
</script>

<style scoped>

</style>
