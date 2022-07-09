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
            <vs-navbar-item :active="active === 'events'" to="/admin" id="events">
              Events
            </vs-navbar-item>
          </div>
        </div>
        <template #right>
          <vs-button @click="logOut()">Log out</vs-button>
        </template>
      </vs-navbar>
    </div>
    <div class="row ms-5 me-5 justify-content-center" style="margin-top: 15%">
      <div class="col align-self-center">
        <div class="row">
          <div class="col">
            <vs-table>
              <template #header>
                <vs-input v-model="search" border placeholder="Search" />
              </template>
              <template #thead>
                <vs-tr>
                  <vs-th sort @click="logs = $vs.sortData($event ,logs, 'time')">
                    Time
                  </vs-th>
                  <vs-th sort @click="logs = $vs.sortData($event ,logs, 'level')">
                    Level
                  </vs-th>
                  <vs-th sort @click="logs = $vs.sortData($event ,logs, 'service')">
                    Service
                  </vs-th>
                  <vs-th sort @click="logs = $vs.sortData($event ,logs, 'msg')">
                    Message
                  </vs-th>
                </vs-tr>
              </template>
              <template #tbody>
                <vs-tr
                  :key="i"
                  v-for="(tr, i) in $vs.getPage($vs.getSearch(logs, search), page, max)"
                  :data="tr"
                >
                  <vs-td>
                    {{ tr.time }}
                  </vs-td>
                  <vs-td>
                    {{ tr.level }}
                  </vs-td>
                  <vs-td>
                    {{ tr.service }}
                  </vs-td>
                  <vs-td>
                    {{ tr.msg }}
                  </vs-td>
                  <template #expand>
                    {{ tr.fullContent}}
                  </template>
                </vs-tr>
              </template>
              <template #footer>
                <vs-pagination v-model="page" :length="$vs.getLength($vs.getSearch(logs, search), max)" />
              </template>
            </vs-table>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>


<script>
import axios from "axios";
import moment from "moment";

export default {
  name: "AdminHomePage",
  data() {
    return {
      active: 'events',
      logs: [],
      page: 1,
      max: 7,
      search: '',
    }
  },
  async mounted() {
    await this.getLogs();
  },
  methods: {
    logOut() {
      this.$store.commit('setToken', null);
      this.$store.commit('setUser', null);
      this.$router.push('/');
    },
    async getLogs(){
      let response = await axios.get(`${process.env.VUE_APP_BACKEND}/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/security/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/post/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/comment/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/reaction/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/job-offer/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/connection/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/post/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      response = await axios.get(`${process.env.VUE_APP_BACKEND}/interceptor/logs`);
      if(response.data){
        this.logs.push(...response.data.logs)
      }
      //this.logs = this.$vs.sortData(null, this.logs, 'time');
    }
  },
}
</script>

<style scoped>

</style>
