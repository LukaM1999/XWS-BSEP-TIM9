<template>
 <div style="padding: 6rem">
   <div class="row justify-content-center">
     <div class="col-12 d-flex justify-content-center" style="margin-top: 6rem">
       <vs-input type="search"
       v-model="jobSearch" label-placeholder="Search job offers..." />
     </div>
   </div>
   <div class="row">
     <div class="col-4 mt-4" v-for="jobOffer in filteredJobs" :key="jobOffer.id">
       <vs-card style="padding-left: 5%;">
         <template #title>
           <div class="row">
             <div class="col">
               <p class="text-lg-end">{{jobOffer.createdAt | formatDate}} </p>
               <h3 class="text-start" style="font-weight: bold; font-size: x-large">{{jobOffer.position}}</h3>
             </div>
           </div>
         </template>
         <template #text>
           <div class="row mb-2">
             <div class="col">
               <h5 class="text-lg-start" v-if="jobOffer.profileId == undefined">Company</h5>
               <h5 class="text-lg-start" v-if="jobOffer.profileId != undefined">Employer</h5>
               <p class="text-lg-start mt-2"  style="font-size: large">{{jobOffer.company}}</p>

             </div>
           </div>
           <div class="row mt-3 mb-4">
             <div class="col">
               <h5 class="text-lg-start">Description</h5>
               <p class="text-lg-start mt-2" style="font-size: large">{{jobOffer.description}}</p>
             </div>
           </div>
           <div class="row mt-3 mb-4">
             <div class="col">
               <h5 class="text-lg-start">Criteria</h5>
               <p class="text-lg-start mt-2" style="font-size: large">{{jobOffer.criteria}}</p>
             </div>
           </div>
           <div class="row mt-3 mb-4">
             <div class="col">
               <h5 class="text-lg-start">Skills</h5>
               <div class="row justify-content-start">
                 <div class="col justify-content-start">
                   <table>
                     <tr v-for="item in jobOffer.skills" v-bind:key="item">
                       <td style="overflow: hidden; text-align: left">
                         <p style="font-size: large">{{ item }}</p>
                       </td>
                     </tr>
                   </table>
                 </div>
               </div>
             </div>
           </div>
         </template>
       </vs-card>
     </div>
   </div>
   <div class="row mt-6 ">
     <div class="col justify-content-center mt-4">
      <h3>Suggested </h3>
     </div>
   </div>

  <div class="row">
    <div class="col-4 mt-4" v-for="jobOffer in jobOffers" :key="jobOffer.id">
      <vs-card style="padding-left: 5%;">
        <template #title>
          <div class="row">
            <div class="col">
              <p class="text-lg-end">{{jobOffer.createdAt | formatDate}} </p>
              <h3 class="text-start" style="font-weight: bold; font-size: x-large">{{jobOffer.position}}</h3>
            </div>
          </div>
        </template>
        <template #text>
          <div class="row mb-2">
            <div class="col">
              <h5 class="text-lg-start" v-if="jobOffer.profileId == undefined">Company</h5>
              <h5 class="text-lg-start" v-if="jobOffer.profileId != undefined">Employer</h5>
              <p class="text-lg-start mt-2"  style="font-size: large">{{jobOffer.company}}</p>

            </div>
          </div>
          <div class="row mt-3 mb-4">
            <div class="col">
              <h5 class="text-lg-start">Description</h5>
              <p class="text-lg-start mt-2" style="font-size: large">{{jobOffer.description}}</p>
            </div>
          </div>
          <div class="row mt-3 mb-4">
            <div class="col">
              <h5 class="text-lg-start">Criteria</h5>
              <p class="text-lg-start mt-2" style="font-size: large">{{jobOffer.criteria}}</p>
            </div>
          </div>
          <div class="row mt-3 mb-4">
            <div class="col">
              <h5 class="text-lg-start">Skills</h5>
              <div class="row justify-content-start">
                <div class="col justify-content-start">
                  <table>
                    <tr v-for="item in jobOffer.skills" v-bind:key="item">
                      <td style="overflow: hidden; text-align: left">
                        <p style="font-size: large">{{ item }}</p>
                      </td>
                    </tr>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </template>
      </vs-card>
    </div>
  </div>

   </div>
</template>



<script>
import axios from 'axios';
import moment from "moment";

export default {
  name: "JobOffers",
  data() {
      return {
        currentUserId: null,
        profile: null,
        jobOffers: [],
        jobOffers1: [],
        recommendations: [],
        jobs: [],
        position: "",
        seniority: "",
        description: "",
        criteria: "",
        positions: [],
        skills: [],
        skill: "",
        jobSearch: "",
      }
    },
    async beforeMount() {
      await this.getProfile()
      await this.getJobs();
      await this.getRecommendations();
      await this.getJobOffer();
    },
    async mounted() {
      for(let j of this.jobs){
        if(j.profileId == undefined) continue
        await this.getEmployerData(j.profileId)
      }
    },
    computed: {
      filteredJobs() {
        let tempJobs = this.jobs;
        if (this.jobSearch == '') return [];
        tempJobs = tempJobs.filter((r) => {
          return r.company?.toLowerCase().includes(this.jobSearch.toLowerCase()) ||
          r.position?.toLowerCase().includes(this.jobSearch.toLowerCase())
        })
        return tempJobs
      }
    },
    methods : {
      async getProfile() {
        const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${this.$store.getters.user?.id}`)
        if (response.data) {
          this.profile = response.data.profile
        }
      },
      async getJobOffer() {
        for (let jobRecommendation of this.recommendations.jobRecommendations){
          const jobOfferId = jobRecommendation.jobId;
          await axios.get(`${process.env.VUE_APP_BACKEND}/job-offer/${jobOfferId}`).then(response => {
            const jobOffer = response.data.jobOffer;
            this.jobOffers1.push(jobOffer)
          });
        }
        for(let jobOffer of this.jobOffers1){
          if(jobOffer.profileId != undefined){
            const employerFullName = await this.getEmployerData(jobOffer.profileId)
            jobOffer.company = employerFullName
          }
          this.jobOffers.push(jobOffer)
        }
      },
      async getRecommendations() {
        const recommendationRequest = {
          profileId: this.$store.getters.user?.id,
          skills: this.profile.skills,
        };
        await axios.post(process.env.VUE_APP_BACKEND + '/job-offer/recommendation', recommendationRequest)
          .then(response => {
            this.recommendations = response.data;
          })
      },
      async getJobs() {
        await axios.get(process.env.VUE_APP_BACKEND + '/job-offer')
          .then(async response => {
            const jobs1 = response.data.jobOffers;
            for(let jobOffer of jobs1){
              if(jobOffer.profileId != undefined){
                const employerFullName = await this.getEmployerData(jobOffer.profileId)
                jobOffer.company = employerFullName
              }
              this.jobs.push(jobOffer)
            }
          })
      },
      async getEmployerData(id){
        const response = await axios.get(`${process.env.VUE_APP_BACKEND}/profile/${id}`)
        const fullName = response.data.profile.firstName + " " + response.data.profile.lastName
        return fullName
      }
    },
}

</script>

<style scoped>
/* Chrome, Safari, Edge, Opera */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

/* Firefox */
input[type=number] {
  -moz-appearance: textfield;
}

.dirty {
  border-color: #5A5;
  background: #EFE;
}

.dirty:focus {
  outline-color: #8E8;
}

.error {
  border-color: red;
  background: #FDD;
}

.error:focus {
  outline-color: #F99;
}


</style>
