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
      const dto = {
        commonNameSubject: "retard3",
        nameSubject: "retard3",
        surnameSubject: "retard3",
        emailSubject: "retard3@gmail.com",
        countrySubject: "RS",
        startDate: "2022-07-04",
        endDate: "2023-07-04",
        commonNameIssuer: "retard2",
        nameIssuer: "retard2",
        surnameIssuer: "retard2",
        emailIssuer: "retard2@gmail.com",
        countryIssuer: "RS",
        serialNumberIssuer: "6",
        serialNumberSubject: "1",
        authoritySubject: "root",
        keyUsages: [
          0,
          1,
          2
        ]
      }
      const response = await axios.post(process.env.VUE_APP_BACKEND + "/admin/getIssuedCertificates", dto);
      if(response.data){
        alert(response.data);
      }
    }
  }

}
</script>

<style>

</style>