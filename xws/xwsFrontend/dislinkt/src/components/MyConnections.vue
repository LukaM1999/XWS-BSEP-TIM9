<template>
  <div style="margin-top: 10%" class="row">
    <div v-if="connections.length > 0" class="col-4 ml-5" style="border-radius: 15px; background-color: lavenderblush" :key="refreshKey">
      <h1>Connections: </h1>
      <div :key="i" v-for="(tr, i) in connections">
        <div class="row" style="margin-bottom: 5px; margin-top: 5px">
          <div class="col"></div>
          <div class="col-10">
            <div class="row" style="margin-bottom: 5px; margin-top: 5px">
              <div class="col d-flex justify-content-center align-self-center">
                <div class="center con-avatars">
                  <vs-avatar circle primary size="35">
                    <template #text>
                      {{tr.fullName}}
                    </template>
                  </vs-avatar>
                </div>
              </div>
              <div class="col d-flex justify-content-center align-self-end" style="color: black">
                <vs-navbar-item style="font-size: large; cursor: default">
                  {{tr.fullName}}
                </vs-navbar-item>
              </div>
              <div v-if="!tr.isApproved === true" class="col d-flex justify-content-end  align-self-center">
                <vs-button primary style="font-size: medium" @click="approve(tr)">
                  <i class='bx bx-check'></i>
                </vs-button>
                <vs-button danger style="font-size: medium" @click="decline(tr)">
                  <i class='bx bx-x'></i>
                </vs-button>
              </div>
              <div v-if="tr.isApproved === true" class="col d-flex justify-content-center  align-self-center">
                <vs-button dark>Profile</vs-button>
              </div>
            </div>
          </div>
          <div class="col"></div>
        </div>
      </div>
    </div>
    <div class="col"></div>
    <div v-if="connections.length > 0" class="col-4 mr-5" style="border-radius: 15px; background-color: lavenderblush" :key="refreshKey">
      <h1>Suggested: </h1>
      <div :key="i" v-for="(tr, i) in connections">
        <div class="row" style="margin-bottom: 5px; margin-top: 5px">
          <div class="col"></div>
          <div class="col-10">
            <div class="row" style="margin-bottom: 5px; margin-top: 5px">
              <div class="col d-flex justify-content-center align-self-center">
                <div class="center con-avatars">
                  <vs-avatar circle primary size="35">
                    <template #text>
                      {{tr.fullName}}
                    </template>
                  </vs-avatar>
                </div>
              </div>
              <div class="col d-flex justify-content-center align-self-end" style="color: black">
                <vs-navbar-item style="font-size: large; cursor: default">
                  {{tr.fullName}}
                </vs-navbar-item>
              </div>
              <div class="col d-flex justify-content-center  align-self-center">
                <vs-button dark>Add connection</vs-button>
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
import axios from "axios";

export default {
  name: "MyConnections",
  data() {
    return {
      connections: [],
      recommendations: [],
      userId: this.$store.getters.user?.id,
      fullname: '',
      refreshKey: 0
    }
  },
  async beforeMount() {
    await this.getMyConnections()
    await this.getRecommendations()
  },
  methods: {
    async getMyConnections() {
      const loading = this.$vs.loading();
      const response = await axios.get(process.env.VUE_APP_BACKEND + `/connection/${this.userId}`)
        .catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error getting connections',
            color: 'danger',
            position: 'top-right'
          });
          loading.close();
          throw error;
        });
      loading.close();
      this.connections = response.data.connections
      for (const connection of this.connections) {
        await this.getConnectionFullName(connection)
      }
    },
    async getConnectionFullName(connection) {
      let profileId = connection.subjectId
      if(connection.subjectId === this.userId)
        profileId = connection.issuerId
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${profileId}`)
        .catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error getting user',
            color: 'danger',
            position: 'top-right'
          });
          throw error;
        })
      const profile = response.data.profile
      this.fullname = profile.firstName + " " + profile.lastName
      connection.fullName = this.fullname
      this.refreshKey += 1
    },
    async getRecommendations(){
      const loading = this.$vs.loading();
      const response = await axios.get(process.env.VUE_APP_BACKEND + `/connection/user/${this.userId}/recommendation`)
        .catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error getting recommendations',
            color: 'danger',
            position: 'top-right'
          });
          loading.close();
          throw error;
        });
      loading.close();
      this.recommendations = response.data.recommendations
      console.log(this.recommendations)
      // for (const connection of this.connections) {
      //   await this.getConnectionFullName(connection)
      // }
    },
    async approve(c) {
      const loading = this.$vs.loading();
      const response = await axios.patch(process.env.VUE_APP_BACKEND + `/connection/${c.id}`)
        .catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error while approving',
            color: 'danger',
            position: 'top-right'
          });
          loading.close();
          throw error;
        });
      loading.close();
      c.isApproved = true
      this.refreshKey += 1
    },
    async decline(c) {
      const loading = this.$vs.loading();
      const response = await axios.delete(process.env.VUE_APP_BACKEND + `/connection/${c.id}`)
        .catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error while declining',
            color: 'danger',
            position: 'top-right'
          });
          loading.close();
          throw error;
        });
      loading.close();
      this.connections = this.connections.filter(con => con.id !== c.id)
    }
  },
}
</script>

<style scoped>

</style>
