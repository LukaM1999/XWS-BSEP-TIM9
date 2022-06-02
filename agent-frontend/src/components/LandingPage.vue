<template>
  <div style="overflow-x: hidden">
    <vs-navbar :color="'guide'" fixed shadow center-collapsed>
      <template #left>
        <div class="row">
          <div class="col">
            <img src="/logo.png" width="100" height="100" alt="">
          </div>
          <div class="col align-self-center">
            <p
              style="font-family: 'Bauhaus 93'; margin-bottom: 0rem; margin-left:-2rem ;  font-size: xxx-large; color: #048ce3">
              AGENTY</p>
          </div>
        </div>
      </template>
      <template #right>
        <vs-button class="vs-button--size-large" flat @click="openLoginDialog()">Login</vs-button>
        <vs-button class="vs-button--size-large" @click="openRegisterDialog()">Register</vs-button>
      </template>
    </vs-navbar>
    <div class="row justify-content-end">
      <div class="col-1">
        <vs-dialog :prevent-close="true" @close="resetRegister()" auto-width v-model="dialogRegister">
          <template #header>
            <h4 class="not-margin me-3 ms-3">
              Register
            </h4>
          </template>
          <div class="con-form" style="display: inline-block;">
            <div class="row">
              <div class="col">
                <vs-input required
                          :success="isSuccess($v.username)"
                          :danger="isInvalid($v.username)"
                          class="mt-2" primary v-model="$v.username.$model"
                          label-placeholder="Username">
                </vs-input>
                <label v-if="isError($v.username)" style="color: #FF9999; font-size: 10pt">Invalid input</label>
                <vs-input required
                          :success="isSuccess($v.password)"
                          :danger="isInvalid($v.password)"
                          class="mt-4" type="password" primary v-model="$v.password.$model" label-placeholder="Password"
                          :visiblePassword="hasVisiblePassword"
                          icon-after
                          @click-icon="hasVisiblePassword = !hasVisiblePassword">
                  <template #icon>
                    <i v-if="!hasVisiblePassword" class='bx bx-show-alt'></i>
                    <i v-else class='bx bx-hide'></i>
                  </template>
                </vs-input>
                <label v-if="isError($v.password)" style="color: #FF9999; font-size: 10pt">Minimum 8<br>and maximum 20
                  characters,<br>at least one uppercase letter,<br>one lowercase letter,<br>one number,<br> one special
                  character</label>
                <password v-model="password" :strength-meter-only="true" :secureLength="8"
                          :userInputs="[username]"/>
                <vs-input required
                          :success="isSuccess($v.confirmPassword)"
                          :danger="isInvalid($v.confirmPassword)" danger-text="asdasd"
                          class="mt-4" type="password" primary v-model="$v.confirmPassword.$model"
                          label-placeholder="Confirm password"/>
              </div>
            </div>
          </div>
          <template #footer>
            <div class="footer-dialog">
              <vs-button :disabled="!isRegisterValid()" @click="registerUser()" block>
                Register
              </vs-button>
            </div>
          </template>
        </vs-dialog>
      </div>
      <div class="col-1">
        <vs-dialog :prevent-close="true" @close="resetLogin()" auto-width v-model="dialogLogin">
          <template #header>
            <h4 class="not-margin me-3 ms-3">
              Login
            </h4>
          </template>
          <div class="con-form" style="display: inline-block;">
            <div class="row">
              <div class="col">
                <vs-input required class="mt-2" primary v-model="username" label-placeholder="Username"/>
                <vs-input required class="mt-4" type="password" primary v-model="password" label-placeholder="Password"
                          :visiblePassword="hasVisiblePassword"
                          icon-after
                          @click-icon="hasVisiblePassword = !hasVisiblePassword">
                  <template #icon>
                    <i v-if="!hasVisiblePassword" class='bx bx-show-alt'></i>
                    <i v-else class='bx bx-hide'></i>
                  </template>
                </vs-input>
              </div>
            </div>
          </div>
          <template #footer>
            <div class="footer-dialog">
              <vs-button :disabled="!isLoginValid()" @click="login()" block>
                Login
              </vs-button>
            </div>
          </template>
        </vs-dialog>
      </div>
    </div>
    <div class="row" style="padding: 5%">
      <h1 class="text-lg-center" style="">Companies</h1>
      <div class="row mt-2 justify-content-center">
        <div class="col-12 d-flex justify-content-center">
          <vs-input type="search" v-model="companySearch" warn label-placeholder="Search companies..." />
        </div>
      </div>
      <div class="col-4 mt-4" v-for="company in filteredCompanies" :key="company.id">
        <vs-card @click="$router.push({name: 'company-profile',  params: {companyName: company.name}})" style="padding-left: 5%;">
          <template #title>
            <h3>{{company.name}}</h3>
          </template>
          <template #text>
            <p class="text-lg-start">Address: <strong>{{company.country}} {{company.city}}, {{company.address}}</strong></p>
            <p class="text-lg-start">Employees: <strong>{{company.size}}</strong></p>
            <p class="text-lg-start">Industry: <strong>{{company.industry}}</strong></p>
            <br>
            <p class="text-lg-start">{{company.description}}</p>
          </template>
        </vs-card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Password from 'vue-password-strength-meter';
import zxcvbn from "zxcvbn";
import {email, helpers, minLength, required, sameAs} from "vuelidate/lib/validators";
import isPasswordCompromised from '@mathiscode/password-leak'

const isPasswordStrong = (value, vm) => zxcvbn(value, [vm.username, vm.email, vm.firstName, vm.lastName])?.score >= 3
const username = helpers.regex('username', /^[_a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[_a-zA-Z0-9]$/)
const password = helpers.regex('password', /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,20}$/)

export default {
  name: "LandingPage",
  components: {
    Password,
  },
  data() {
    return {
      dialogRegister: false,
      dialogLogin: false,
      hasVisiblePassword: false,
      username: "",
      password: "",
      confirmPassword: "",
      companies: [],
      companySearch: "",
    }
  },
  validations: {
    username: {
      required,
      username
    },
    password: {
      required,
      minLength: minLength(8),
      isPasswordStrong,
      password
    },
    confirmPassword: {
      sameAsPassword: sameAs('password')
    },
  },
  async mounted() {
    this.$store.commit('setUser', null);
    this.$store.commit('setToken', null);
    await this.getAllCompanies();
  },
  computed:{
    //Search by company name
    filteredCompanies() {
      return this.companies.filter(company => {
        return company.name.toLowerCase().includes(this.companySearch.toLowerCase())
      })
    },
  },
  methods: {
    async getAllCompanies(){
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/company`);
      if(response.status === 200){
        this.companies = response.data;
      }
    },
    isPasswordStrong() {
      return zxcvbn(this.password, [this.username, this.email, this.firstName, this.lastName])?.score >= 3
    },
    isSuccess(validation) {
      return !validation.$invalid
    },
    isInvalid(validation) {
      return validation.$invalid
    },
    isError(validation) {
      return validation.$error
    },
    openRegisterDialog() {
      this.dialogRegister = true;
    },
    resetRegister() {
      this.$v.$reset()
      this.username = "";
      this.password = "";
      this.confirmPassword = "";
    },
    isRegisterValid() {
      return !this.$v.$invalid;
    },
    async registerUser() {
      const loading = this.$vs.loading();
      if (!this.isRegisterValid()) {
        return;
      }
      if (this.password !== this.confirmPassword) {
        return;
      }
      const isCompromised = await isPasswordCompromised(this.password)
      if (isCompromised) {
        loading.close();
        this.$vs.notification({
          title: 'Error',
          text: 'Password is compromised',
          color: 'danger',
          position: 'top-right'
        })
        return;
      }
      const registeredUser = {
          username: this.username,
          password: this.password,
      };
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/auth/signup`, registeredUser)
        .catch(error => {
          this.$vs.notification({
            title: "Error",
            text: "Registration failed",
            color: "danger",
            position: "top-right"
          });
          loading.close();
          throw error;
        });
      loading.close()
      this.$vs.notification({
        title: "Success",
        text: "Registered successfully!",
        color: "success",
        position: "top-right"
      });
      this.dialogRegister = false;
      this.resetRegister();
    },
    openLoginDialog() {
      this.dialogLogin = true;
    },
    resetLogin() {
      this.username = "";
      this.password = "";
    },
    isLoginValid() {
      return this.username.length > 0 && this.password.length > 0;
    },

    async login() {
      if (!this.isLoginValid()) {
        return;
      }
      const loading = this.$vs.loading();
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/auth/login`, {
        username: this.username,
        password: this.password
      }).catch(error => {
        this.$vs.notification({
          title: "Error",
          text: "Invalid username/password",
          color: "danger",
          position: "top-right"
        });
        loading.close()
        throw error
      });
      loading.close()
      this.$store.commit("setToken", response.data?.accessToken);
      this.$store.commit("setUser", response.data?.user);
      await this.$router.push(`/home`);
    },
  }
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
