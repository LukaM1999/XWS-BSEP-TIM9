<template>
  <div class="container">
    <div class="row justify-content-end">
      <div class="col-1">
        <vs-button @click="openRegisterDialog()">Register</vs-button>
        <vs-dialog :prevent-close="true" @close="resetRegister()" auto-width v-model="dialogRegister">
          <template #header>
            <h4 class="not-margin me-3 ms-3">
              Register
            </h4>
          </template>
          <div class="con-form" style="display: inline-block">
            <div class="row">
              <div class="col">
                <vs-input required class="mt-2" primary v-model="username" label-placeholder="Username"/>
                <vs-input required class="mt-4" primary v-model="email" label-placeholder="Email"/>
                <vs-input required class="mt-4" primary v-model="firstName" label-placeholder="First name"/>
                <vs-input required class="mt-4" primary v-model="lastName" label-placeholder="Last name"/>
                <vs-input required class="mt-4" type="password" primary v-model="password" label-placeholder="Password"
                          :visiblePassword="hasVisiblePassword"
                          icon-after
                          @click-icon="hasVisiblePassword = !hasVisiblePassword">
                  <template #icon>
                    <i v-if="!hasVisiblePassword" class='bx bx-show-alt'></i>
                    <i v-else class='bx bx-hide'></i>
                  </template>
                </vs-input>
                <password v-model="password" :strength-meter-only="true" :secureLength="8" :userInputs="[username, email, firstName, lastName]"/>
                <vs-input required class="mt-4" type="password" primary v-model="confirmPassword" label-placeholder="Confirm password"/>
                <vs-checkbox class="mt-4" primary v-model="setupOtp">Setup OTP authentication?</vs-checkbox>
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
        <vs-button @click="openLoginDialog()">Login</vs-button>
        <vs-dialog :prevent-close="true" @close="resetLogin()" width="15%" v-model="dialogLogin">
          <template #header>
            <h4 class="not-margin me-3 ms-3">
              Login
            </h4>
          </template>
          <vs-navbar color="#7d33ff" text-white square v-model="activeTab">
            <vs-navbar-item :active="activeTab === 'standard'" id="standard">
              Standard
            </vs-navbar-item>
            <vs-navbar-item :active="activeTab === 'passwordless'" id="passwordless">
              Passwordless
            </vs-navbar-item>
          </vs-navbar>
          <div class="con-form" style="display: inline-block; padding-top: 3.5rem">
            <div v-if="activeTab === 'standard'" class="row">
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
            <div v-if="activeTab === 'passwordless'" class="row">
              <div class="col">
                <vs-input required class="mt-2" primary v-model="username" label-placeholder="Username"/>
                <vs-input required class="mt-4" inputmode="numeric" pattern="[0-9]{6}" primary v-model="totp" label-placeholder="TOTP"/>
              </div>
            </div>
          </div>
          <template #footer>
            <div class="footer-dialog">
              <vs-button v-if="activeTab === 'standard'" :disabled="!isLoginValid()" @click="login()" block>
                Login
              </vs-button>
              <vs-button v-if="activeTab === 'passwordless'"
                         :disabled="!isPasswordlessLoginValid()"
                         @click="passwordlessLogin()" block>
                Login
              </vs-button>
            </div>
          </template>
        </vs-dialog>
        <vs-dialog :prevent-close="true" @close="resetOtp()" auto-width v-model="dialogOtp">
          <template #header>
            <h4 class="not-margin me-3 ms-3">
              OTP Setup
            </h4>
          </template>
          <div class="con-form" style="display: inline-block">
            <div class="row">
              <div class="col">
                <p>Secret: {{this.secret}}</p>
                <img :src="`data:image/png;base64,${this.qrCode}`"/>
              </div>
            </div>
          </div>
        </vs-dialog>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Password from 'vue-password-strength-meter';
import zxcvbn from "zxcvbn";

export default {
  name: "LandingPage",
  components: {
    Password
  },
  data() {
    return {
      dialogRegister: false,
      dialogLogin: false,
      activeTab: 'standard',
      hasVisiblePassword: false,
      username: "",
      email: "",
      firstName: "",
      lastName: "",
      password: "",
      confirmPassword: "",
      totp: "",
      setupOtp: false,
      dialogOtp: false,
      secret: "",
      qrCode: ""
    }
  },
  mounted() {
    this.$store.commit('setUser', null);
    this.$store.commit('setToken', null);
  },
  methods: {
    openRegisterDialog() {
      this.dialogRegister = true;
    },
    resetRegister() {
      this.username = "";
      this.email = "";
      this.firstName = "";
      this.lastName = "";
      this.password = "";
      this.confirmPassword = "";
      this.totp = "";
      this.setupOtp = false;
    },
    isRegisterValid() {
      return this.password.length >= 8 && this.password === this.confirmPassword &&
        zxcvbn(this.password, [this.username, this.email, this.firstName, this.lastName])?.score >= 3;
    },
    async registerUser() {
      if (!this.isRegisterValid()) {
        return;
      }
      if (this.password !== this.confirmPassword) {
        return;
      }
      const registeredUser = {
        user: {
          username: this.username,
          password: this.password,
          role: "user",
        },
        email: this.email,
        firstName: this.firstName,
        lastName: this.lastName,
      };
      const loading = this.$vs.loading();
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/security/user`, registeredUser)
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
        text: "Registered successfully! Verify your email to login",
        color: "success",
        position: "top-right"
      });
      this.dialogRegister = false;

      if(this.setupOtp){
        const response = await axios.get(`${process.env.VUE_APP_BACKEND}/security/setupOtp/${this.username}`)
          .catch(error => {
            this.$vs.notification({
              title: "Error",
              text: "OTP setup failed",
              color: "danger",
              position: "top-right"
            });
            loading.close()
            throw error;
          });
        loading.close()
        this.secret = response.data.secret;
        this.qrCode = response.data.qrCode;
        this.dialogOtp = true;
      }
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
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/security/login`, {
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
      this.$vs.notification({
        title: "Success",
        text: "Logged in successfully",
        color: "success",
        position: "top-right"
      });
      this.$store.commit("setToken", response.data?.accessToken);
      const user = await axios.get(`${process.env.VUE_APP_BACKEND}/security/user/${this.username}`)
        .catch(error => {
          this.$vs.notification({
            title: "Error",
            text: "Failed to get user details",
            color: "danger",
            position: "top-right"
          });
          throw error;
        });
      this.$store.commit("setUser", user.data?.user);
      await this.$router.push(`/${this.$store.getters.user?.role}`);
    },
    resetOtp() {

    },
    isPasswordlessLoginValid() {
      return this.username.length > 0 && this.totp.length > 0;
    },
    async passwordlessLogin() {
      if (!this.isPasswordlessLoginValid()) {
        return;
      }
      const loading = this.$vs.loading();
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/security/passwordlessLogin`, {
        username: this.username,
        otp: this.totp
      }).catch(error => {
        this.$vs.notification({
          title: "Error",
          text: "Invalid username/totp",
          color: "danger",
          position: "top-right"
        });
        loading.close()
        throw error
      });
      loading.close()
      this.$vs.notification({
        title: "Success",
        text: "Logged in successfully",
        color: "success",
        position: "top-right"
      });
      this.$store.commit("setToken", response.data?.accessToken);
      const user = await axios.get(`${process.env.VUE_APP_BACKEND}/security/user/${this.username}`)
        .catch(error => {
          this.$vs.notification({
            title: "Error",
            text: "Failed to get user details",
            color: "danger",
            position: "top-right"
          });
          throw error
        });
      this.$store.commit("setUser", user.data?.user);
      await this.$router.push(`/${this.$store.getters.user?.role}`);
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
</style>
