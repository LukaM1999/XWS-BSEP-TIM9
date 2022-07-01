<template>
  <div>
    <div class="center examplex">
      <vs-navbar target-scroll="#padding-scroll-content" style="background-color: lavenderblush;" padding-scroll
                 center-collapsed v-model="active"
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
            <vs-navbar-item :active="active == 'guide'" id="guide" to="/user/posts">
              For you
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'myPosts'" id="myPosts" to="/user/my-posts">
              My posts
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'docs'" id="docs">
              Connections
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'components'" id="components">
              Job offers
            </vs-navbar-item>
            <vs-navbar-item :active="active === 'messages'" id="messages" to="/user/messages">
              Messages
            </vs-navbar-item>
            <vs-navbar-item :active="active == 'license'" id="license" to="/user/profile">
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
            <textarea class="mb-4 form-control vs-input" v-model="text" style="border-radius: 15px" id="text" name="text"></textarea>
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
            <input class="mb-4 form-control" v-model="link" style="border-radius: 15px" type="text" id="link" name="link"/>
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
        <router-view></router-view>
      </div>
    </div>
  </div>

</template>

<script>

import axios from 'axios';
import OneSignalVue from "onesignal-vue";
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
      link: ''
    }
  },
  async beforeMount(){
    await this.getProfile()
  },
  async mounted() {
    if ('serviceWorker' in navigator) {
      navigator.serviceWorker.register('/OneSignalSDKWorker.js').then(function(registration) {
      }).catch(function(e) {
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

                /* CATEGORY SLIDEDOWN SPECIFIC TEXT */
                negativeUpdateButton:"Cancel",
                positiveUpdateButton:"Save Preferences",
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
      if(response.data){
        this.profile = response.data.profile
        console.log(this.profile)
      }
    },
    async searchProfile() {
      const loading = this.$vs.loading();
      const response = await axios.get(process.env.VUE_APP_BACKEND + '/profile?search=Luka').catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error searching profiles',
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
    changeRoute(path){
      this.$router.replace(path)
    },
    openCreatePostDialog(){
      this.createPostDialog = true
    },
    async createPost(){
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
      this.$router.push('my-posts')
      this.resetCreatePostDialog()
    },
    resetCreatePostDialog(){
      this.text = ''
      this.image = ''
      this.link = ''
    }
  }
}
</script>

<style scoped>

</style>
