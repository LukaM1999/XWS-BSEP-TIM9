import Vue from 'vue'
import App from './App.vue'
import VueRouter from "vue-router";
import axios from 'axios'
import VueAxios from 'vue-axios'
import Toasted from 'vue-toasted';
import Vuesax from 'vuesax'
import 'vuesax/dist/vuesax.css' //Vuesax styles
import LandingPage from "@/components/LandingPage";
import HomePage from "@/components/HomePage";

Vue.config.productionTip = false
Vue.config.devtools

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


const routes = [
  {
    path: '/',
    name: 'landingPage',
    component: LandingPage,
    children: [

    ],
  },
  {
    path: '/home/:user',
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
}).$mount('#app')
