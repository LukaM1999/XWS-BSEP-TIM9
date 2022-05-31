import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import App from './App.vue'
import VueRouter from "vue-router";
import axios from 'axios'
import VueAxios from 'vue-axios'
import Toasted from 'vue-toasted';
import Vuesax from 'vuesax'
import Vuelidate from 'vuelidate'
import 'vuesax/dist/vuesax.css' //Vuesax styles
import LandingPage from "@/components/LandingPage";
import {isTokenExpired, jwtInterceptor} from "@/_helpers/jwt.interceptor";
import HomePage from "@/components/HomePage";
import PasswordReset from "@/components/PasswordReset";

Vue.config.productionTip = false
Vue.config.devtools

jwtInterceptor()

Vue.use(Vuex)
Vue.use(VueRouter)
Vue.use(VueAxios, axios)
Vue.use(Toasted, {
  position: 'top-right',
  duration: 3000,
  keepOnHover: true,
})
Vue.use(Vuesax, {
  // options here
  colors: {
    primary: '#be1d7b',
  },
  vsCard: {
    minWidth: '100%',
    background: 'transparent',
  }
})

Vue.use(Vuelidate)

export const store = new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    user: null,
    token: null,
    failedLoginAttempts: 0,
  },
  mutations: {
    setToken(state, token) {
      state.token = token
    },
    setUser(state, user) {
      state.user = user
    },
    incrementFailedLoginAttempts(state) {
      state.failedLoginAttempts++
    },
    resetFailedLoginAttempts(state) {
      state.failedLoginAttempts = 0
    }
  },
  getters: {
    token(state) {
      return state.token
    },
    user(state) {
      return state.user
    },
    failedLoginAttempts(state) {
      return state.failedLoginAttempts
    }
  }
})


const routes = [
  {
    path: '/',
    name: 'landingPage',
    component: LandingPage
  },
  {
    path: '/user',
    name: 'userHomepage',
    component: HomePage,
    children: [

      ]
  },
  {
    path: '/admin',
    name: 'adminHomepage',
    component: HomePage
  },
  {
    path: '/passwordRecovery',
    name: 'passwordRecovery',
    component: PasswordReset
  },
]

export const router = new VueRouter({
  routes,
  mode: 'history'
})

function isAuthorized(role) {
  const token = store.getters.token
  const storedRole = store.getters.user?.role
  if(isTokenExpired(token) || storedRole !== role) return false
  return true

}

router.beforeEach((to, from, next) => {
  if (to.path.indexOf('user') !== -1) {
    if (!isAuthorized('user')) {
      alert('Unauthorized!')
      next('/')
    } else next()
  }
  else if (to.path.indexOf('admin') !== -1) {
    if (!isAuthorized('admin')) {
      alert('Unauthorized!')
      next('/')
    } else next()
  }
  else next()
})

export var vue = new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
