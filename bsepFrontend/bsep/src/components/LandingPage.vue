<template>
  <div>
    <div class="row mb-3">
      <div class="col-md-4"></div>
      <div class="col-md-4 d-flex justify-content-center">
        <vs-input color="primary" placeholder="username" v-model="username"></vs-input>
      </div>
      <div class="col-md-4"></div>
    </div>
    <div class="row mb-3">
      <div class="col-md-4"></div>
      <div class="col-md-4 d-flex justify-content-center">
        <vs-input color="primary" type="password" placeholder="password"/>
      </div>
      <div class="col-md-4"></div>
    </div>
    <div class="row">
      <div class="col-md-4"></div>
      <div class="col-md-4 d-flex justify-content-center">
        <vs-button @click="login" color="primary" type="filled">Sign in</vs-button>
      </div>
      <div class="col-md-4"></div>
    </div>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: "LandingPage",
  data() {
    return {
      username: "",

    }
  },
  methods: {
    async login() {
      if(this.username === "") return;
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/user/login/${this.username}`);
      if(response.data){
        console.table(response.data);
      }
    }
  }

}
</script>

<style>

</style>