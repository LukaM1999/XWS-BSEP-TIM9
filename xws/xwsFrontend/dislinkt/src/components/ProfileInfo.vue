<template>
  <div>
    <div class="row" style="margin-left: 5%; margin-top: 13%; margin-right: 5%">
      <div class="col-md-10">
        <div class="row">
          <div class="col-md-4">
            <vs-card style="overflow-wrap: anywhere">
              <template #title>
                <h1>{{ firstName }} {{ lastName }}</h1>
              </template>
              <template #text>
                <div class="row justify-content-center">
                  <div class="col justify-content-center d-grid">
                    <vs-card class="text-center" v-if="isPrivate == false">
                      <template #text>
                        <h6>
                          Date of birth: {{dateOfBirth}}
                        </h6>
                      </template>
                    </vs-card>
                    <vs-card v-if="isPrivate == false">
                      <template #text>
                        <h6>
                          Email: {{email}}
                        </h6>
                      </template>
                    </vs-card>
                    <vs-card v-if="isPrivate == false">
                      <template #text>
                        <h6>
                          Phone number: {{phoneNumber}}
                        </h6>
                      </template>
                    </vs-card>
                    <vs-card v-if="isPrivate == false">
                      <template #text>
                        <h6>
                          Gender: {{gender}}
                        </h6>
                      </template>
                    </vs-card>
                    <vs-card v-if="isPrivate == false">
                      <template #text>
                        <h6>
                          Biography: {{biography}}
                        </h6>
                      </template>
                    </vs-card>
                  </div>
                </div>
              </template>
            </vs-card>
            <vs-card style="padding-top: 1em" v-if="isPrivate == false" >
              <template #title>
                <h6>Skills:</h6>
              </template>
              <template #text>
                <div v-for="item in skills" v-bind:key="item">
                  <p>{{item}}</p>
                </div>
              </template>
            </vs-card>
            <vs-card v-if="isPrivate == false" style="padding-top: 1em; width: 30rem">
              <template #title>
                <h6>Interests:</h6>
              </template>
              <template #text>
                <div v-for="item in interests" v-bind:key="item">
                  <p>{{item}}</p>
                </div>
              </template>
            </vs-card>
            <vs-card v-if="isPrivate == false" style="padding-top: 1em; width: 30rem">
              <template #title>
                <h6>Education:</h6>
              </template>
              <template #text>
                <div v-for="(item, i) in education" v-bind:key="i">
                  <vs-card>
                    <template #text>
                      <div class="row">
                        <div class="col-md-12">
                          <h6>
                            School: {{item.school}}
                          </h6>
                          <br>
                          <h6>
                            Field of study: {{item.fieldOfStudy}}
                          </h6>
                        </div>
                      </div>
                      <div class="row">
                        <div class="col-md-6">
                          <h6>
                            Degree: {{item.degree}}
                          </h6>
                          <br>
                          <h6>
                            Grade: {{item.grade}}
                          </h6>
                        </div>
                        <div class="col-md-6">
                          <h6>
                            Started: {{formatDate(item.startDate)}}
                          </h6>
                          <h6>
                            Ended: {{formatDate(item.endDate)}}
                          </h6>
                        </div>
                      </div>
                      <h6>
                        Description: {{item.description}}
                      </h6>
                    </template>
                  </vs-card>
                </div>
              </template>
            </vs-card>
            <vs-card v-if="isPrivate == false" style="padding-top: 1em; width: 30rem">
              <template #title>
                <h6>Work experience:</h6>
              </template>
              <template #text>
                <div v-for="(item, i) in education" v-bind:key="i">
                  <vs-card>
                    <template #text>
                      <h6>
                        Title: {{item.title}}
                      </h6>
                      <br>
                      <h6>
                        Company: {{item.company}} ({{getEmploymentType(item.employmentType)}})
                      </h6>
                      <div class="row">
                        <div class="col-md-6">
                          <h6>
                            Started: {{formatDate(item.startDate)}}
                          </h6>
                        </div>
                        <div class="col-md-6">
                          <h6>
                            Ended: {{formatDate(item.endDate)}}
                          </h6>
                        </div>
                      </div>
                      <div class="row">
                        <h6>
                          Location: {{item.location}}
                        </h6>
                      </div>
                    </template>
                  </vs-card>
                </div>
              </template>
            </vs-card>
          </div>
          <div class="col-md-8">
            <div v-for="post in posts" v-bind:key="post.id">
              <div class="row">
                <div class="col-12 d-flex justify-content-center">
                  <Post :post="post" class="mb-4" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-md-2">
        <table>
          <tr>
            <td>
              <vs-button block >
                Follow
              </vs-button>
            </td>
            <td>
              <vs-button block>
                Block
              </vs-button>
            </td>
          </tr>
        </table>
        <vs-card>
          <template #title>
            <h3>People you may know</h3>
          </template>
          <template #text>
            <div v-for="(item, i) in suggestions" v-bind:key="i">
              <vs-card>
                <template #text>
                  <h6>
                    {{item.firstName}} {{item.lastName}}
                  </h6>
                  <vs-button block >
                    Follow
                  </vs-button>
                </template>
              </vs-card>
              <vs-card>
                <template #text>
                  <h6>
                    Company: {{item.company}}
                  </h6>
                </template>
              </vs-card>
              <vs-card>
                <template #text>
                  <h6>
                    Employment type: {{item.employmentType}}
                  </h6>
                </template>
              </vs-card>
              <vs-card>
                <template #text>
                  <h6>
                    Location: {{item.location}}
                  </h6>
                </template>
              </vs-card>
              <vs-card>
                <template #text>
                  <h6>
                    Started: {{item.startDate}}
                  </h6>
                </template>
              </vs-card>
              <vs-card>
                <template #text>
                  <h6>
                    Ended: {{item.endDate}}
                  </h6>
                </template>
              </vs-card>
            </div>
          </template>
        </vs-card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
import Post from "@/components/Post";

export default {
  name: "ProfileInfo",
  props: ["id"],
  components:{Post},
  data() {
    return {
      user: {},
      username: "",
      firstName: "",
      lastName: "",
      dateOfBirth: Date(),
      email: "",
      gender: "",
      biography: "",
      skills: [],
      skill: "",
      interest: "",
      interests: [],
      isPrivate: true,
      phoneNumber: "",
      workExperience: [],
      education: [],
      recomendations: [],
      posts: [],
    }
  },
  mounted() {
    this.getProfile();
    this.getRecomendations();
    this.getMyPosts();
  },
  methods: {
    async getProfile() {
      const loading = this.$vs.loading();
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${this.id}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting user',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      this.user = response.data?.profile;
      this.username = this.user.username;
      this.firstName = this.user.firstName;
      this.lastName = this.user.lastName;
      this.dateOfBirth = moment(this.user.dateOfBirth).format('l')
      this.email = this.user.email;
      this.gender = this.user.gender;
      this.biography = this.user.biography;
      if (this.user?.skills?.length > 0)
        this.skills = this.user.skills;
      if (this.user?.interests?.length > 0)
        this.interests = this.user.interests;
      if (this.user?.isPrivate != undefined)
        this.isPrivate = this.user.isPrivate;
      this.phoneNumber = this.user.phoneNumber;
      if (this.user?.workExperience?.length > 0) {
        this.workExperience = this.user.workExperience;
      }
      if (this.user?.education?.length > 0) {
        this.education = this.user.education;
      }
      loading.close();
    },
    async getRecomendations() {
      const loading = this.$vs.loading();
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/connection/user/${this.$store.getters.user?.id}/recommendation`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting recomendations',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      for (const item in response.data) {
        const profile = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${item.id}`).catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Error getting user',
            color: 'danger',
            position: 'top-right'
          });
          throw error;
        });
        this.recomendations.push(profile);
      }
      loading.close();
    },
    async getMyPosts() {
      const loading = this.$vs.loading();
      const response = await axios.get(`https://localhost:8000/post/profile/${this.id}`).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error getting posts',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      loading.close();
      this.posts = response.data.posts;
      this.sortPosts()
    },
    sortPosts(){
      this.posts = this.posts.sort((a, b) => moment(b.createdAt) - moment(a.createdAt))
    },
    formatDate(date){
      return moment(date).format('l')
    },
    getEmploymentType(employmentType){
      if (employmentType == 0)
        return "FULL_TIME"
      if (employmentType == 1)
        return "PART_TIME"
      if (employmentType == 2)
        return "SELF_EMPLOYED"
      if (employmentType == 3)
        return "FREELANCE"
      if (employmentType == 4)
        return "CONTRACT"
      if (employmentType == 5)
        return "INTERNSHIP"
      if (employmentType == 6)
        return "APPRENTICESHIP"
      if (employmentType == 7)
        return "SEASONAL"

      return "FULL_TIME"
    }
  }
}
</script>

<style>
</style>
