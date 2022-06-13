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
        <vs-input color="primary" type="password" v-model="password" placeholder="password"/>
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
      password: ""
    }
  },
  mounted() {
    this.$store.commit("setUser", null)
    this.$store.commit("setToken", null)
  },
  methods: {
    isLoginValid() {
      return this.username.length > 0 && this.password.length > 0;
    },
    async login() {
      if (!this.isLoginValid()) {
        return;
      }
      const loading = this.$vs.loading();
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/auth/login`, {
        username: this.username,
        password: this.password
      }).catch(error => {
        this.$vs.notification({
          title: "Error",
          text: "Invalid username/password",
          color: "danger",
          position: "top-right"
        });
        loading.close()
        throw error
      });
      loading.close()
      this.$store.commit("setToken", response.data?.accessToken);
      this.$store.commit("setUser", response.data?.user);
      await this.$router.push(`/home`);
    },
  }
}
</script>

<style>

</style>