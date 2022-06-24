<template>
  <div>
    <div class="center examplex">
      <vs-navbar target-scroll="#padding-scroll-content" style="background-color: lavenderblush;" padding-scroll
                 center-collapsed v-model="active">
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
          <vs-button @click="logOut()">Log out</vs-button>
        </template>
      </vs-navbar>
    </div>
    <div class="row">
      <div class="col">
        <router-view></router-view>
      </div>
    </div>
  </div>

</template>

<script>

import axios from 'axios';

export default {
  name: "HomePage",
  data() {
    return {
      active: 'guide',
    }
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
      const loading = this.$vs.loading();
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${this.$store.getters.user?.id}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting user',
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
    }
  }
}
</script>

<style scoped>

</style>
