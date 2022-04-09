<template>
  <div>
    <div class="row"><h1>My certificates</h1></div>
    <div class="row mt-3">
      <div v-for="c in myCertificates" :key="c.serialNumberSubject" class="col d-flex justify-content-center">
        <vs-card>
          <template #title>
            <h3>{{c.commonNameSubject}}</h3>
          </template>
          <template #img>
            <img :src="'/' + c.authoritySubject + '.svg'" alt="" width="200" height="200">
          </template>
          <template #text>
            <p style="text-align: left">
              <strong>Subject username:</strong> {{c.usernameSubject}}<br>
              <strong>Subject name:</strong> {{c.nameSubject}}<br>
              <strong>Subject surname:</strong> {{c.surnameSubject}}<br>
              <strong>Subject country:</strong> {{c.countrySubject}}<br>
              <strong>Subject serial number:</strong> {{c.serialNumberSubject}}<br>
              <strong>Subject authority:</strong> {{c.authoritySubject}}<br>
              <strong>Extensions:</strong> {{c.keyUsages}}<br>


              <strong>Valid from:</strong> {{c.validFromSubject}}<br>
              <strong>Valid until:</strong> {{c.validUntilSubject}}<br>

              <strong>Issuer username:</strong> {{c.usernameIssuer}}<br>
              <strong>Issuer common name:</strong> {{c.commonNameIssuer}}<br>
            </p>
          </template>
          <template #interactions>
            <vs-button danger icon>
              Revoke
            </vs-button>
            <vs-button class="btn-primary" primary>
              Download
            </vs-button>
          </template>
        </vs-card>
      </div>
    </div>
    <div class="row mt-5"><h1>Issued certificates</h1></div>
    <div class="row mb-5 mt-5">
      <div v-for="c in issued" :key="c.serialNumberSubject" class="col d-flex justify-content-center">
      <vs-card>
        <template #title>
          <h3>{{c.commonNameSubject}}</h3>
        </template>
        <template #img>
          <img :src="'/' + c.authoritySubject + '.svg'" alt="" width="200" height="200">
        </template>
        <template #text>
          <p style="text-align: left">
            <strong>Subject username:</strong> {{c.usernameSubject}}<br>
            <strong>Subject name:</strong> {{c.nameSubject}}<br>
            <strong>Subject surname:</strong> {{c.surnameSubject}}<br>
            <strong>Subject country:</strong> {{c.countrySubject}}<br>
            <strong>Subject serial number:</strong> {{c.serialNumberSubject}}<br>
            <strong>Subject authority:</strong> {{c.authoritySubject}}<br>
            <strong>Extensions:</strong> {{c.keyUsages}}<br>


            <strong>Valid from:</strong> {{c.validFromSubject}}<br>
            <strong>Valid until:</strong> {{c.validUntilSubject}}<br>

            <strong>Issuer username:</strong> {{c.usernameIssuer}}<br>
            <strong>Issuer common name:</strong> {{c.commonNameIssuer}}<br>
          </p>
        </template>
        <template #interactions>
          <vs-button danger icon>
            Revoke
          </vs-button>
          <vs-button class="btn-primary" primary>
            Download
          </vs-button>
        </template>
      </vs-card>
    </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "HomePage",
  data() {
    return {
      myCertificates: [],
      issued: [],
      path: "",
    };
  },
  async mounted() {
    const response = await axios.get(`${process.env.VUE_APP_BACKEND}/user/login/${this.$route.params.user}`);
    if(response.data) {
      console.table(response.data);
      for(let i = 0; i < response.data.length; i++) {
        if(response.data[i].usernameSubject !== this.$route.params.user) {
          this.issued.push(response.data[i]);
        } else {
          this.myCertificates.push(response.data[i]);
        }
      }
    }
  },
  methods: {

  },
  props: {

  }
}
</script>

<style scoped>

</style>