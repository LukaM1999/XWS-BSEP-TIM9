import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import OneSignalVue from "onesignal-vue";
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
import Post from "@/components/Post";
import ConnectionsPosts from "@/components/ConnectionsPosts";
import Messages from "@/components/Messages";
import Profile from "@/components/Profile";
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import './registerServiceWorker'
import MyPosts from "@/components/MyPosts";
import MyConnections from "@/components/MyConnections";
import ProfileInfo from "@/components/ProfileInfo";
import * as https from "https";
import AdminHomePage from "@/components/AdminHomePage";
import JobOffers from "@/components/JobOffers";
import moment from "moment";
import MyJobOffers from "@/components/MyJobOffers";

Vue.config.productionTip = false
Vue.config.devtools

jwtInterceptor()

Vue.use(Vuex)
Vue.use(VueRouter)
Vue.use(OneSignalVue)
Vue.use(VueAxios, axios)
Vue.axios.defaults.httpsAgent = new https.Agent({
  rejectUnauthorized: false,
  requestCert: false,
})
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

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

export const store = new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    user: null,
    token: null,
    failedLoginAttempts: 0,
    firebaseToken: null,
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
    },
    setFirebaseToken(state, token) {
      state.firebaseToken = token
    },
    setOneSignalToken(state, token) {
      state.oneSignalToken = token
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
    },
    firebaseToken(state) {
      return state.firebaseToken
    },
    oneSignalToken(state) {
      return state.oneSignalToken
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
    redirect: '/user/posts',
    children: [
      {
        path: 'posts',
        name: 'posts',
        component: ConnectionsPosts
      },
      {
        path: 'messages',
        name: 'messages',
        component: Messages
      },
      {
        path: 'profile',
        name: 'profile',
        component: Profile,
      },
      {
        path: 'my-posts',
        name: 'myPosts',
        component: MyPosts
      },
      {
        path: 'connections',
        name: 'connections',
        component: MyConnections
      },
      {
        path: 'profile-info',
        name: 'profileInfo',
        props: true,
        component: ProfileInfo
      },
      {
        path: 'job-offers',
        name: 'jobOffers',
        props: true,
        component: JobOffers
      },
      {
        path: 'my-job-offers',
        name: 'myJobOffers',
        props: true,
        component: MyJobOffers
      },
    ]
  },
  {
    path: '/admin',
    name: 'adminHomepage',
    component: AdminHomePage,
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

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('DD.MM.YYYY.')
  }
});

export var vue = new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
