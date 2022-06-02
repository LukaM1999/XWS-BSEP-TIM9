<template>
<div style="padding: 6rem">
  <vs-dialog v-model="showModal" ref="salary" width="600px">
    <template #header>
      <h4 class="not-margin">
        <b>Add salary</b>
      </h4>
    </template>
    <div class="con-form">
      <div class="row mt-2 justify-content-center">
        <div class="col-5">
          <vs-select  v-model="position" label-placeholder="Position">
            <vs-option v-for="position in positions" :label="position" :value="position" :key="position">
              {{position}}
            </vs-option>
          </vs-select>
        </div>
        <div class="col-5">
          <vs-select  v-model="seniority" label-placeholder="Seniority">
            <vs-option label="Junior" value="Junior">Junior</vs-option>
            <vs-option label="Medior" value="Medior">Medior</vs-option>
            <vs-option label="Senior" value="Senior">Senior</vs-option>
          </vs-select>
        </div>
      </div>
      <div class="row mt-4 justify-content-center">
        <div class="col-5">
          <vs-select  v-model="engagement" label-placeholder="Engagement">
            <vs-option label="Full-time" value="Full-time">Full-time</vs-option>
            <vs-option label="Part-time" value="Part-time">Part-time</vs-option>
            <vs-option label="Freelance" value="Freelance">Freelance</vs-option>
            <vs-option label="Internship" value="Internship">Internship</vs-option>
          </vs-select>
        </div>
        <div class="col-5 d-inline-flex">
          <div class="row">
            <div class="col">
              <vs-radio v-model="currentlyEmployed" label="Current employee" :val="1">Current employee</vs-radio>
            </div>
          </div>
          <div class="row">
            <div class="col">
              <vs-radio v-model="currentlyEmployed" label="Former employee" :val="0">Former employee</vs-radio>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row mt-4 justify-content-center">
      <div class="col-5">
        <vs-input v-model="salary" min="0" max="20000" label-placeholder="Net salary" type="number">
          <template #icon>
            <i class='bx bx-euro'></i>
          </template>
        </vs-input>
      </div>
      <div class="col-5">
      </div>
    </div>
    <template #footer>
      <div class="footer-dialog">
        <div class="row justify-content-end">
          <div class="col d-flex justify-content-end">
            <vs-button @click="addSalary" class="btn-primary">Add salary</vs-button>
          </div>
        </div>
      </div>
    </template>
  </vs-dialog>
  <div v-if="role === 'USER'" class="row justify-content-end">
    <div class="col d-flex justify-content-end">
      <vs-button @click="openModal" class="vs-button--size-large" color="#7dcdec"><strong>Add salary</strong></vs-button>
    </div>
  </div>
  <div class="row justify-content-center">
    <div class="col-3 border-bottom border-1">
      <p>Position (Seniority)</p>
    </div>
    <div class="col-3 border-bottom border-1">
      <p>Salary</p>
    </div>
  </div>
  <div class="row mt-4 justify-content-center" v-for="salary in salaries" :key="salary.position">
    <div class="col-3 border-bottom border-1">
      <p><strong>{{salary.position}}</strong></p>
    </div>
    <div class="col-3 border-bottom border-1">
      <p><strong>{{salary.salary}} &euro;</strong></p>
    </div>
  </div>
</div>
</template>

<script>
import axios from "axios";

export default {
  name: "Salaries",
  data() {
    return {
      salaries: [],
      showModal: false,
      role: "",
      positions: [],
      seniority: "",
      position: "",
      engagement: "",
      currentlyEmployed: 1,
      salary: 0,
    }
  },
  mounted() {
    this.$parent.active = 'salaries';
    this.getSalaries();
    this.getPositions();
    this.role = this.$store.getters.user?.role?.authority;
  },
  methods: {
    async getSalaries() {
      await axios.get(`${process.env.VUE_APP_BACKEND}/company/${this.$route.params.companyName}/salary/average`).then(response => {
        const positions = Object.keys(response.data);
        this.salaries = positions.flatMap(position => {
          return {
            position,
            salary: response.data[position],
          }
        });
      });
    },
    async getPositions(){
      await axios.get(`${process.env.VUE_APP_BACKEND}/positions.json`,).then(response => {
        this.positions = response.data?.positions.flatMap(position => position.value)
      });
    },
    resetSalaryForm() {
      this.seniority = "";
      this.position = "";
      this.engagement = "";
      this.currentlyEmployed = 1;
      this.salary = 0;
    },
    openModal(){
      this.showModal = true;
    },
    async addSalary() {
      const salary = {
        companyName: this.$route.params.companyName,
        position: `${this.position} (${this.seniority})`,
        engagement: this.engagement,
        currentlyEmployed: Boolean(this.currentlyEmployed),
        monthlyNetSalary: this.salary,
      };
      const loading = this.$vs.loading({
        container: this.$refs.salary,
        color: 'primary',
        scale: 0.6,
        center: true
      });
      await axios.post(`${process.env.VUE_APP_BACKEND}/company/salary`, salary).then(response => {
        loading.close()
        this.$vs.notification({
          color: 'success',
          title: 'Success',
          text: 'Salary added successfully',
          position: 'top-right',
        });
        this.getSalaries();
        this.showModal = false;
        this.resetSalaryForm();
      }).catch(error => {
        loading.close()
        this.$vs.notification({
          color: 'danger',
          title: 'Error',
          text: 'Error adding salary',
          position: 'top-right',
        });
        throw error;
      });
    },
  },
}
</script>

<style scoped>

p {
  color: white
}

</style>
