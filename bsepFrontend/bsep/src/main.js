import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import App from './App.vue'
import VueRouter from "vue-router";
import axios from 'axios'
import VueAxios from 'vue-axios'
import Toasted from 'vue-toasted';
import Vuesax from 'vuesax'
import 'vuesax/dist/vuesax.css' //Vuesax styles
import LandingPage from "@/components/LandingPage";
import HomePage from "@/components/HomePage";
import {jwtInterceptor} from "@/_helpers/jwt.interceptor";
import Vuelidate from 'vuelidate'


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
})

Vue.use(Vuelidate)



export const store = new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    user: null,
    token: null,
  },
  mutations: {
    setToken(state, token) {
      state.token = token
    },
    setUser(state, user) {
      state.user = user
    },
  },
  getters: {
    token(state) {
      return state.token
    },
    user(state) {
      return state.user
    },
  }
})

const routes = [
  {
    path: '/',
    name: 'landingPage',
    component: LandingPage,
    children: [

    ],
  },
  {
    path: '/home',
    name: 'homePage',
    component: HomePage,
    children: [

    ],
  },
]

export const router = new VueRouter({
  routes,
  mode: 'history'
})

export var vue = new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
