<template>
  <div class="row justify-content-center d-flex">
    <vs-dialog :prevent-close="true" not-close @close="resetForm()" auto-width v-model="dialogPassword">
      <template #header>
        <h4 class="not-margin me-3 ms-3">
          Password recovery
        </h4>
      </template>
      <div class="con-form" style="display: inline-block;">
        <div class="row">
          <div class="col">
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
            <label v-if="isError($v.password)" style="color: #FF9999; font-size: 10pt">Minimum 8<br>and maximum 20 characters,<br>at least one uppercase letter,<br>one lowercase letter,<br>one number,<br> one special character</label>
            <password v-model="password" :strength-meter-only="true" :secureLength="8"
                      />
            <vs-input required
                      :success="isSuccess($v.confirmPassword)"
                      :danger="isInvalid($v.confirmPassword)" danger-text="asdasd"
                      class="mt-4" type="password" primary v-model="$v.confirmPassword.$model"
                      label-placeholder="Confirm password"/></div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <vs-button :disabled="isFormValid()" @click="passwordUpdate()" block>
            Reset password
          </vs-button>
        </div>
      </template>
    </vs-dialog>
  </div>
</template>

<script>
import axios from "axios";
import Password from 'vue-password-strength-meter';
import zxcvbn from "zxcvbn";
import {helpers, minLength, required, sameAs} from "vuelidate/lib/validators";

const isPasswordStrong = (value, vm) => zxcvbn(value, [vm.username, vm.email, vm.firstName, vm.lastName])?.score >= 3
const password = helpers.regex('password', /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,20}$/)

export default {
  name: "PasswordReset",
  data() {
    return {
      dialogPassword: true,
      password: '',
      confirmPassword: '',
      hasVisiblePassword: false
    }
  },
  components: {
    Password
  },
  validations: {
    password: {
      required,
      minLength: minLength(8),
      isPasswordStrong,
      password
    },
    confirmPassword: {
      required,
      minLength: minLength(8),
      sameAs: sameAs('password')
    }
  },
  methods: {
    resetForm() {
      this.$v.$reset()
      this.password = ''
      this.confirmPassword = ''
    },
    isFormValid() {
      return this.$v.$invalid
    },
    isSuccess(field) {
      return !field.$invalid
    },
    isInvalid(field) {
      return field.$invalid
    },
    isError(field) {
      return field.$error
    },
    async passwordUpdate() {
      if (this.isFormValid()) {
        return
      }
      const loading = this.$vs.loading();
      await axios.patch(`${process.env.VUE_APP_BACKEND}/security/updatePassword`, {
        password: this.password,
        token: this.$route.query.token
      }).then(response => {
        this.dialogPassword = false
        this.$vs.notification({
          title: 'Success',
          text: 'Password has been changed',
          color: 'success',
          position: 'top-right'
        })
        loading.close();
        setTimeout(() => {this.$router.push("/")}, 5000)
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
  }
}
</script>

<style scoped>

</style>
