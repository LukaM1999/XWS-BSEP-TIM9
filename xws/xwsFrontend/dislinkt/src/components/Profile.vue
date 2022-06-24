<style>
.vs-input{
    width: auto;
    font-size: 15pt;
}
.row{
    padding-bottom: 25px;
}
</style>
<template >
  <div class="row" style="margin-left: 10%; margin-right: 10%">
    <h2 style="margin-bottom: 2%">PROFILE</h2>
    <div class="col-3">
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Username" v-model="username"/>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Name" v-model="firstName"/>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Lastname" v-model="lastName"/>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Date of birth" v-model="dateOfBirth"/>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Email" v-model="email"/>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Gender" v-model="gender"/>
          </div>
          <vs-button block @click="updateProfile()" style="padding-top: 1em">
            Save
          </vs-button>
        </div>
      </div>
    </div>
    <div class="col-3">
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Biography" v-model="biography"/>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <table style="padding-left: 3em">
              <th>
                <h4 style="padding-top: 5pt">Skills</h4>
              </th>
              <tr v-for="item in skills" v-bind:key="item">
                <td>
                  <span>{{ item }}</span>
                </td>
                <td>
                  <vs-button icon @click="deleteSkill(item)" style="height: 2em; width: 2em">
                    <i class='bx bx-minus'></i>
                  </vs-button>
                </td>
              </tr>
            </table>
          </div>
        </div>
        <div class="col">
          <vs-button icon :active="activeSkill == 0"
                     @click="activeSkill=!activeSkill">
            <i class='bx bx-plus'></i>
          </vs-button>
          <vs-dialog v-model="activeSkill" prevent-close>
            <template #header>
              <h4 class="not-margin">
                Add skill
              </h4>
            </template>
            <div class="con-form">
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Skill" v-model="skill"/>
              </div>
            </div>
            <template #footer>
              <div class="footer-dialog">
                <vs-button block @click="addSkill()">
                  Save
                </vs-button>
              </div>
            </template>
          </vs-dialog>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <table style="padding-left: 3em">
              <th>
                <h4 style="padding-top: 5pt">Interests</h4>
              </th>
              <tr v-for="item in interests" v-bind:key="item">
                <td>
                  <span>{{ item }}</span>
                </td>
                <td>
                  <vs-button icon @click="deleteInterest(item)" style="height: 2em; width: 2em">
                    <i class='bx bx-minus'></i>
                  </vs-button>
                </td>
              </tr>
            </table>
          </div>
        </div>
        <div class="col">
          <vs-button icon :active="activeInterest == 0"
                     @click="activeInterest=!activeInterest">
            <i class='bx bx-plus'></i>
          </vs-button>
          <vs-dialog v-model="activeInterest" prevent-close>
            <template #header>
              <h4 class="not-margin">
                Add interest
              </h4>
            </template>
            <div class="con-form">
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Interest" v-model="interest"/>
              </div>
            </div>
            <template #footer>
              <div class="footer-dialog">
                <vs-button block @click="addInterest()">
                  Save
                </vs-button>
              </div>
            </template>
          </vs-dialog>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="center con-checkbox">
            <vs-checkbox v-model="isPrivate" :checked="isPrivate">
              Private
            </vs-checkbox>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <vs-input label-placeholder="Phone number" v-model="phoneNumber"/>
          </div>
        </div>
      </div>
    </div>
    <div class="col-3">
      <div class="row">
        <div class="col">
          <h3 style="padding-top: 3pt">Education</h3>
        </div>
        <div class="col">
          <vs-button
            icon
            :active="activeEducation == 0"
            @click="activeEducation=!activeEducation"
          >
            <i class='bx bx-plus'></i>
          </vs-button>
          <vs-dialog v-model="activeEducation" prevent-close>
            <template #header>
              <h4 class="not-margin">
                Add education
              </h4>
            </template>
            <div class="con-form">
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="School" v-model="school"/>
              </div>
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Degree" v-model="degree"/>
              </div>
              <div class="centerx">
                <vs-input label-placeholder="Field of study" v-model="fieldOfStudy"/>
              </div>
              <div style="padding-bottom: 1em; padding-top: 1em">
                <label for="example-datepicker1">When schooling began?</label>
                <b-form-datepicker id="example-datepicker1" :max="maxDate" v-model="startDate" class="mb-2"></b-form-datepicker>
              </div>
              <div style="padding-bottom: 1em">
                <label for="example-datepicker2">When schooling ended?</label>
                <b-form-datepicker id="example-datepicker2" :max="maxDate" v-model="endDate" class="mb-2"></b-form-datepicker>
              </div>
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Grade" v-model="grade"/>
              </div>
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Description" v-model="description"/>
              </div>
            </div>

            <template #footer>
              <div class="footer-dialog">
                <vs-button block @click="addEducation()">
                  Save
                </vs-button>
              </div>
            </template>
          </vs-dialog>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <div v-for="edu in education" v-bind:key="edu">
              <div class="row">
                <vs-card>
                  <template #title>
                    <h3>School: {{ edu.school }}</h3>
                  </template>
                  <template #text>
                    <p>
                      Degree: {{ edu.degree }}
                    </p>
                    <p>
                      Field of study: {{ edu.fieldOfStudy }}
                    </p>
                    <p>
                      Started: {{ edu.startDate }}
                    </p>
                    <p>
                      Ended: {{ edu.endDate }}
                    </p>
                    <p>
                      Grade: {{ edu.grade }}
                    </p>
                    <p>
                      Description: {{ edu.description }}
                    </p>
                    <vs-button danger icon @click="deleteEducation(edu)">
                      <i class='bx bx-minus'></i>
                    </vs-button>
                  </template>
                </vs-card>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="col-3">
      <div class="row">
        <div class="col">
          <h3 style="padding-top: 3pt">Work experience</h3>
        </div>
        <div class="col">
          <vs-button
            icon
            :active="activeWorkExperience == 0"
            @click="activeWorkExperience=!activeWorkExperience"
          >
            <i class='bx bx-plus'></i>
          </vs-button>
          <vs-dialog v-model="activeWorkExperience" prevent-close>
            <template #header>
              <h4 class="not-margin">
                Add work experience
              </h4>
            </template>
            <div class="con-form">
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Title" v-model="title"/>
              </div>
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Company" v-model="company"/>
              </div>
              <div class="centerx">
                <vs-input label-placeholder="Employment type" v-model="employmentType"/>
              </div>
              <div style="padding-bottom: 1em; padding-top: 1em">
                <label for="example-datepicker1">When schooling began?</label>
                <b-form-datepicker id="example-datepicker1" :max="maxDate" v-model="startDate" class="mb-2"></b-form-datepicker>
              </div>
              <div style="padding-bottom: 1em">
                <label for="example-datepicker2">When schooling ended?</label>
                <b-form-datepicker id="example-datepicker2" :max="maxDate" v-model="endDate" class="mb-2"></b-form-datepicker>
              </div>
              <div class="centerx" style="padding-bottom: 2em">
                <vs-input label-placeholder="Location" v-model="location"/>
              </div>
            </div>

            <template #footer>
              <div class="footer-dialog">
                <vs-button block @click="addWorkExperience()">
                  Save
                </vs-button>
              </div>
            </template>
          </vs-dialog>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <div class="centerx">
            <div v-for="work in workExperience" v-bind:key="work">
              <div class="row">
                <vs-card>
                  <template #title>
                    <h3>Company: {{ work.company }}</h3>
                  </template>
                  <template #text>
                    <p>
                      Title: {{ work.title }}
                    </p>
                    <p>
                      Employment type: {{ work.employmentType }}
                    </p>
                    <p>
                      Started: {{ work.startDate }}
                    </p>
                    <p>
                      Ended: {{ work.endDate }}
                    </p>
                    <p>
                      Location: {{ work.location }}
                    </p>
                    <vs-button danger icon @click="deleteWorkExperience(work)">
                      <i class='bx bx-minus'></i>
                    </vs-button>
                  </template>
                </vs-card>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment"

export default {
  name: "Profile",
  data() {
    return {
      activeEducation: false,
      activeWorkExperience: false,
      activeSkill: false,
      activeInterest: false,
      maxDate: new Date(),
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
      school: "",
      degree: "",
      fieldOfStudy: "",
      startDate: Date(),
      endDate: Date(),
      grade: "",
      description: "",
      title: "",
      company: "",
      employmentType: "",
      location: "",
    }
  },
  methods: {
    deleteWorkExperience(w){
      for(let i=0; i<this.workExperience.length; i++){
        if(this.workExperience[i] == w){
          this.workExperience.splice(i, 1);
        }
      }
    },
    deleteEducation(e){
      for(let i=0; i<this.education.length; i++){
        if(this.education[i] == e){
          this.education.splice(i, 1);
        }
      }
    },
    deleteSkill(s){
      for(let i=0; i<this.skills.length; i++){
        if(this.skills[i] == s){
          this.skills.splice(i, 1);
        }
      }
    },
    deleteInterest(s){
      for(let i=0; i<this.interests.length; i++){
        if(this.interests[i] == s){
          this.interests.splice(i, 1);
        }
      }
    },
    addSkill(){
      this.skills.push(this.skill);
      this.skill = "";
      this.activeSkill = false;
    },
    addInterest(){
      this.interests.push(this.interest);
      this.interest = "";
      this.activeInterest = false;
    },
    addEducation(){
      this.education.push({
        school: this.school,
        degree: this.degree,
        fieldOfStudy: this.fieldOfStudy,
        startDate: moment(this.startDate).format(),
        endDate: moment(this.endDate).format(),
        grade: this.grade,
        description: this.description
      })
      this.school = "";
      this.degree = "";
      this.fieldOfStudy = "";
      this.startDate = Date();
      this.endDate = Date();
      this.grade = "";
      this.description = "";
      this.activeEducation = false;
    },
    addWorkExperience(){
      this.workExperience.push({
        title: this.title,
        company: this.company,
        employmentType: this.employmentType,
        location: this.location,
        endDate: moment(this.endDate).format(),
        startDate: moment(this.startDate).format()
      })
      this.title = "";
      this.company = "";
      this.employmentType = "";
      this.startDate = Date();
      this.endDate = Date();
      this.location = "";
      this.activeWorkExperience = false;
    },
    async updateProfile(){
      const loading = this.$vs.loading();
      const response = await axios.put(`${process.env.VUE_APP_BACKEND}/profile/${this.$store.getters.user?.id}`,
        {
          id: this.$store.getters.user?.id,
          username: this.username,
          firstName: this.firstName,
          lastName: this.lastName,
          dateOfBirth: moment(this.dateOfBirth).format(),
          phoneNumber: this.phoneNumber,
          email: this.email,
          gender: this.gender,
          isPrivate: this.isPrivate,
          biography: this.biography,
          education: this.education,
          workExperience: this.workExperience,
          skills: this.skills,
          interests: this.interests
        }
      ).catch(error => {
        this.$vs.notification({
          title: 'Error',
          text: 'Error updating profile',
          color: 'danger',
          position: 'top-right'
        });
        loading.close();
        throw error;
      });
      this.$vs.notification({
        title: 'Success',
        text: 'Successfully updated profile',
        color: 'success',
        position: 'top-right'
      });
      loading.close();
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
      console.log(response);
      this.user = response.data?.profile;
      this.username = this.user.username;
      this.firstName = this.user.firstName;
      this.lastName = this.user.lastName;
      this.dateOfBirth = this.user.dateOfBirth;
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
  },
  mounted() {
    this.getProfile();
  }
}
</script>

<style scoped>

</style>
