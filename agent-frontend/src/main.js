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
import moment from 'moment'
import LandingPage from "@/components/LandingPage";
import {isTokenExpired, jwtInterceptor} from "@/_helpers/jwt.interceptor";
import CompanyProfile from "@/components/CompanyProfile";
import Comments from "@/components/Comments";
import Salaries from "@/components/Salaries";
import Interviews from "@/components/Interviews";
import Overview from "@/components/Overview";
import HomePage from "@/components/HomePage";
import Companies from "@/components/Companies";
import Profile from "@/components/Profile";
import OwnerCompanies from "@/components/OwnerCompanies";
import Requests from "@/components/Requests";
import JobOffers from "@/components/JobOffers";

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
    updateUser(state, user) {
      state.user.firstName = user.firstName
      state.user.lastName = user.lastName
      state.user.email = user.email
      state.user.address = user.address
      state.user.city = user.city
      state.user.country = user.country
      state.user.phone = user.phone
      state.user.dislinktUsername = user.dislinktUsername
      state.user.dislinktToken = user.dislinktToken
    }

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
    name: 'landing-page',
    component: LandingPage
  },
  {
    path: '/company-profile/:companyName',
    name: 'company-profile',
    component: CompanyProfile,
    redirect: {name: 'company-overview'},
    children: [
      {
        path: 'overview',
        name: 'company-overview',
        component: Overview
      },
      {
        path: 'comments',
        name: 'company-comments',
        component: Comments
      },
      {
        path: 'salaries',
        name: 'company-salaries',
        component: Salaries
      },
      {
        path: 'interviews',
        name: 'company-interviews',
        component: Interviews
      },
      {
        path: 'job-offers',
        name: 'company-job-offers',
        component: JobOffers
      }
    ]
  },
  {
    path: '/home',
    name: 'home-page',
    component: HomePage,
    children: [
      {
        path: 'requests',
        name: 'admin-requests',
        component: Requests,
      },
      {
        path: 'my-companies',
        name: 'owner-companies',
        component: OwnerCompanies
      },
      {
        path: 'companies',
        name: 'companies',
        component: Companies
      },
      {
        path: 'profile',
        name: 'profile',
        component: Profile,
      },
  ]
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
