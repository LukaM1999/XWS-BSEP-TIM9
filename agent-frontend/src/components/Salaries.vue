<template>
<div style="padding: 6rem">
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
    }
  },
  mounted() {
    this.$parent.active = 'salaries';
    this.getSalaries();
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
  },
}
</script>

<style scoped>

p {
  color: white
}

</style>
