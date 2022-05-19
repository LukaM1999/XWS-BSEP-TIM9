<template>
<div class="row">
  <div class="col">
    <vs-button primary @click="getAllUsers()">Get all users</vs-button>
    <vs-button primary @click="getProfile()">Get profile</vs-button>
    <vs-button primary @click="searchProfile()">Search profile</vs-button>
  </div>
</div>

</template>

<script>

import axios from 'axios';

export default {
  name: "HomePage",
  methods: {
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
    }
  }
}
</script>

<style scoped>

</style>
