<template>
  <div>
    <div>
      <div class="center examplex">
        <vs-navbar target-scroll="#padding-scroll-content" style="background-color: lavenderblush;" padding-scroll
                   center-collapsed v-model="active">
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
          <template #right>
            <vs-button flat :disabled="isLoginDisabled()" @click="openLoginDialog()" color="#be1d7b">Login</vs-button>
            <vs-button @click="openRegisterDialog()" color="#be1d7b" gradient>Get Started</vs-button>
          </template>
        </vs-navbar>
      </div>
      <vue-recaptcha style="position: absolute; bottom: 30px; left: 30px; z-index: 100000" v-if="showCaptcha"
                     ref="recaptcha" sitekey="6Len6gcgAAAAAK-QuPZGnklGAlC5aKsthR2aMKLx"
                     @verify="captchaVerified">

      </vue-recaptcha>
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
                            :success="isSuccess($v.email)"
                            :danger="isInvalid($v.email)"
                            class="mt-4" primary v-model="$v.email.$model" label-placeholder="Email"/>
                  <label v-if="isError($v.email)" style="color: #FF9999; font-size: 10pt">Invalid input</label>
                  <vs-input required
                            :success="isSuccess($v.firstName)"
                            :danger="isInvalid($v.firstName)"
                            class="mt-4" primary v-model="$v.firstName.$model" label-placeholder="First name"/>
                  <label v-if="isError($v.firstName)" style="color: #FF9999; font-size: 10pt">Invalid input</label>

                  <vs-input required
                            :success="isSuccess($v.lastName)"
                            :danger="isInvalid($v.lastName)"
                            class="mt-4" primary v-model="$v.lastName.$model" label-placeholder="Last name"/>
                  <label v-if="isError($v.lastName)" style="color: #FF9999; font-size: 10pt">Invalid input</label>

                  <vs-input required
                            :success="isSuccess($v.password)"
                            :danger="isInvalid($v.password)"
                            class="mt-4" type="password" primary v-model="$v.password.$model"
                            label-placeholder="Password"
                            :visiblePassword="hasVisiblePassword"
                            icon-after
                            @click-icon="hasVisiblePassword = !hasVisiblePassword">
                    <template #icon>
                      <i v-if="!hasVisiblePassword" class='bx bx-show-alt'></i>
                      <i v-else class='bx bx-hide'></i>
                    </template>
                  </vs-input>
                  <label v-if="isError($v.password)" style="color: #FF9999; font-size: 10pt">Minimum 8<br>and maximum 20
                    characters,<br>at least one uppercase letter,<br>one lowercase letter,<br>one number,<br> one
                    special
                    character</label>
                  <password v-model="password" :strength-meter-only="true" :secureLength="8"
                            :userInputs="[username, email, firstName, lastName]"/>
                  <vs-input required
                            :success="isSuccess($v.confirmPassword)"
                            :danger="isInvalid($v.confirmPassword)" danger-text="asdasd"
                            class="mt-4" type="password" primary v-model="$v.confirmPassword.$model"
                            label-placeholder="Confirm password"/>
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
          <vs-dialog :prevent-close="true" @close="resetLogin()" width="15%" v-model="dialogLogin">
            <template #header>
              <h4 class="not-margin me-3 ms-3">
                Login
              </h4>
            </template>
            <vs-navbar color="#be1d7b" text-white square v-model="activeTab">
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
                  <vs-input required class="mt-4" type="password" primary v-model="password"
                            label-placeholder="Password"
                            :visiblePassword="hasVisiblePassword"
                            icon-after
                            @click-icon="hasVisiblePassword = !hasVisiblePassword">
                    <template #icon>
                      <i v-if="!hasVisiblePassword" class='bx bx-show-alt'></i>
                      <i v-else class='bx bx-hide'></i>
                    </template>
                  </vs-input>
                  <vs-button flat size="small" @click="openEmailDialog">
                    Forgot password?
                  </vs-button>
                  <vs-dialog :prevent-close="true" @close="resetRegister()" auto-width v-model="dialogEmail">
                    <template #header>
                      <h4 class="not-margin me-3 ms-3">
                        Password recovery
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
                          <label v-if="isError($v.username)" style="color: #FF9999; font-size: 10pt">Invalid
                            input</label>

                          <vs-input required
                                    :success="isSuccess($v.email)"
                                    :danger="isInvalid($v.email)"
                                    class="mt-4" primary v-model="$v.email.$model"
                                    label-placeholder="Please enter email"/>
                          <label v-if="isError($v.email)" style="color: #FF9999; font-size: 10pt">Invalid input</label>
                        </div>
                      </div>
                    </div>
                    <template #footer>
                      <div class="footer-dialog">
                        <vs-button :disabled="!isEmailValid()" @click="passwordRecovery()" block>
                          Reset password
                        </vs-button>
                      </div>
                    </template>
                  </vs-dialog>
                </div>
              </div>
              <div v-if="activeTab === 'passwordless'" class="row">
                <div class="col">
                  <vs-input required class="mt-2" primary v-model="username" label-placeholder="Username"/>
                  <vs-input required class="mt-4" inputmode="numeric" pattern="[0-9]{6}" primary v-model="totp"
                            label-placeholder="TOTP"/>
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
                  <p>Secret: {{ this.secret }}</p>
                  <img :src="`data:image/png;base64,${this.qrCode}`"/>
                </div>
              </div>
            </div>
          </vs-dialog>
        </div>
      </div>
    </div>
    <div class="row" style="margin-top: 10rem; background-color: transparent;">
      <div class="col"></div>
      <div class="col d-flex justify-content-center" style="background-color: transparent">
        <Post :post="post"></Post>
      </div>
      <div class="col"></div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Password from 'vue-password-strength-meter';
import zxcvbn from "zxcvbn";
import {email, helpers, minLength, required, sameAs} from "vuelidate/lib/validators";
import {VueRecaptcha} from 'vue-recaptcha';
import isPasswordCompromised from '@mathiscode/password-leak'
import Post from "@/components/Post";

const isPasswordStrong = (value, vm) => zxcvbn(value, [vm.username, vm.email, vm.firstName, vm.lastName])?.score >= 3
const name = helpers.regex('name', /^[A-Z][a-z]+$/)
const username = helpers.regex('username', /^[_a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[_a-zA-Z0-9]$/)
const password = helpers.regex('password', /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,20}$/)

export default {
  name: "LandingPage",
  components: {
    Post,
    Password,
    VueRecaptcha
  },
  data() {
    return {
      dialogRegister: false,
      dialogLogin: false,
      dialogEmail: false,
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
      qrCode: "",
      showCaptcha: false,
      active: 'guide',
      post: {
        id: "6210611b624b2da721f63fe3",
        createdAt: "0001-01-01T00:00:00Z",
        profile: {
          id: "62706d1b624b3da748f63fe3",
          firstName: "Luka",
          lastName: "Miletic"
        },
        content: {
          text: "Check out my github page"
        }
      }
    }
  },
  validations: {
    username: {
      required,
      username
    },
    email: {
      required,
      email
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
    firstName: {
      required,
      name
    },
    lastName: {
      required,
      name
    }
  },
  mounted() {
    this.$store.commit('setUser', null);
    this.$store.commit('setToken', null);
    if (this.isLoginDisabled()) {
      this.showCaptcha = true;
    }
  },
  methods: {
    isPasswordStrong() {
      return zxcvbn(this.password, [this.username, this.email, this.firstName, this.lastName])?.score >= 3
    },
    isLoginDisabled() {
      return this.$store.getters.failedLoginAttempts > 5
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
      this.email = "";
      this.firstName = "";
      this.lastName = "";
      this.password = "";
      this.confirmPassword = "";
      this.totp = "";
      this.setupOtp = false;
    },
    isRegisterValid() {
      return !this.$v.$invalid;
    },
    isEmailValid() {
      return !this.$v.email.$invalid && !this.$v.username.$invalid;
    },
    async passwordRecovery() {
      const loading = this.$vs.loading();
      await axios.post(`${process.env.VUE_APP_BACKEND}/security/recoverPassword`, {
        username: this.username,
        email: this.email,
      })
        .then(response => {
          this.dialogEmail = false;
          this.$vs.notification({
            title: 'Success',
            text: 'Email has been sent',
            color: 'success',
            position: 'top-right'
          })
          loading.close();
        }).catch(error => {
          this.$vs.notification({
            title: 'Error',
            text: 'Something went wrong',
            color: 'danger',
            position: 'top-right'
          })
          loading.close();
          throw error
        })
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
        user: {
          username: this.username,
          password: this.password,
          role: "user",
        },
        email: this.email,
        firstName: this.firstName,
        lastName: this.lastName,
      };
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

      if (this.setupOtp) {
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
    captchaVerified() {
      setTimeout(() => {
        this.showCaptcha = false;
        this.$store.commit('resetFailedLoginAttempts')
      }, 2000);
    },
    openLoginDialog() {
      this.dialogLogin = true;
    },
    openEmailDialog() {
      this.dialogEmail = true;
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
        this.$store.commit('incrementFailedLoginAttempts');
        if (this.isLoginDisabled()) {
          this.dialogLogin = false;
          this.resetLogin();
          this.showCaptcha = true;
          this.$refs.recaptcha.execute();
        }
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
        loading.close()
        this.$vs.notification({
          title: "Error",
          text: "Invalid username/totp",
          color: "danger",
          position: "top-right"
        });
        this.$store.commit('incrementFailedLoginAttempts');
        if (this.isLoginDisabled()) {
          this.dialogLogin = false;
          this.resetLogin();
          this.showCaptcha = true;
          this.$refs.recaptcha.execute();
        }
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
    }
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

.vs-card{
  background: transparent;
  min-width: 100%;
}

</style>
