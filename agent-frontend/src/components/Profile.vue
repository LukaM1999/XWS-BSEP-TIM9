<template>
  <div class="row justify-content-center"  style="padding: 5%">
    <h1 class="text-lg-center" style="">My profile</h1>
    <div class="col-5 justify-content-center">
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="firstName"
            label-placeholder="First name"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="lastName"
            label-placeholder="Last name"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="email"
            label-placeholder="Email"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="address"
            label-placeholder="Address"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="city"
            label-placeholder="City"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="country"
            label-placeholder="Country"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row mt-3">
        <div class="col">
          <vs-input
            warn
            v-model="phone"
            label-placeholder="Phone"
            class="mb-3">
          </vs-input>
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col d-flex justify-content-center">
          <vs-button @click="editProfile" class="vs-button--size-large" color="#7dcdec"><strong>Save changes</strong></vs-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Profile",
  data () {
    return {
      user: {
      },
      firstName: "",
      lastName: "",
      email: "",
      address: "",
      city: "",
      country: "",
      phone: "",
      role: "",
      positions: [],
    }
  },
  mounted() {
    this.$nextTick(() => {
      let elements = document.getElementsByTagName( "input");
      for (let e of elements) {
        e.style.width = "100%";
        e.style.height = "3rem";
      }

    });
    this.user = this.$store.getters.user;
    this.firstName = this.user.firstName;
    this.lastName = this.user.lastName;
    this.email = this.user.email;
    this.address = this.user.address;
    this.city = this.user.city;
    this.country = this.user.country;
    this.phone = this.user.phone;
    this.role = this.$store.getters.user?.role?.authority;
    this.$parent.active = 'profile';
  },
  methods: {
    async getPositions(){
      await axios.get(`${process.env.VUE_APP_BACKEND}/positions.json`,).then(response => {
        this.positions = response.data?.positions.flatMap(position => position.value)
      });
    },
    async editProfile(){
      let user = {
        username: this.user.username,
        firstName: this.firstName,
        lastName: this.lastName,
        email: this.email,
        address: this.address,
        city: this.city,
        country: this.country,
        phone: this.phone,
      }
      await axios.patch(`${process.env.VUE_APP_BACKEND}/user`, user).then(response => {
        if(response.status === 200){
          this.$store.commit('updateUser', user);
          this.$vs.notification({
            color: "success",
            title: "Success",
            position: "top-right",
            text: "Profile updated successfully!"
          });
        }
      }).catch(
        error => {
          this.$vs.notification({
            color: "danger",
            title: "Error",
            position: "top-right",
            text: "Profile update failed!"
          });
        }
      );
    },
  },
}
</script>

<style scoped>

</style>
