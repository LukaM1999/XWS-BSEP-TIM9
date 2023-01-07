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
              <div v-if="!tr.isApproved && tr.subjectId === userId" class="col d-flex justify-content-end  align-self-center">
                <vs-button primary style="font-size: medium" @click="approve(tr)">
                  <i class='bx bx-check'></i>
                </vs-button>
                <vs-button danger style="font-size: medium" @click="decline(tr)">
                  <i class='bx bx-x'></i>
                </vs-button>
              </div>
              <div class="col d-flex justify-content-center  align-self-center">
                <vs-button @click="viewProfile(tr)" dark>Profile</vs-button>
              </div>
            </div>
          </div>
          <div class="col"></div>
        </div>
      </div>
    </div>
    <div class="col"></div>
    <div v-if="recommendations.length > 0" class="col-4 mr-5" style="border-radius: 15px; background-color: lavenderblush" :key="refreshKey">
      <h1>Suggested: </h1>
      <div :key="i" v-for="(tr, i) in recommendations">
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
                <vs-button @click="connect(tr.id)" dark>Connect</vs-button>
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
import moment from "moment";

export default {
  name: "MyConnections",
  data() {
    return {
      connections: [],
      recommendations: [],
      userId: this.$store.getters.user?.id,
      fullname: '',
      refreshKey: 0,
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
      if(!response.data?.connections) return;
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
      if(!response.data?.recommendations) return;
      const recommendations = response.data.recommendations
       for (const recommendation of recommendations) {
         await this.getRecommendationFullName(recommendation)
       }
    },
    async getRecommendationFullName(id){
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${id}`)
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
      const fullName = profile.firstName + " " + profile.lastName
      this.recommendations.push({
        id: id,
        fullName: fullName
      })
      this.refreshKey += 1
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
    },
    async connect(id){
      const newConnection = {
        issuerId: this.$store.getters.user?.id,
        subjectId: id,
        date: moment().format()
      }
      const loading = this.$vs.loading();
      const response = await axios.post(process.env.VUE_APP_BACKEND + '/connection', newConnection).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error while creating connection',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      loading.close();
      this.recommendations = this.recommendations.filter(r => r.id !== id);
    },
    viewProfile(connection) {
      let profileId = connection.subjectId
      if(connection.subjectId === this.userId)
        profileId = connection.issuerId
      this.$router.push({name: 'profileInfo', params: { id: profileId }})
    }
  },
}
</script>

<style scoped>

</style>
