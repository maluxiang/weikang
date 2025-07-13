import Vue from 'vue'
import App from './App'
import './uni.promisify.adaptor'
import axios from 'axios'
import store from './store'


const apps = new Vue({
  store,
  // ...
})
Vue.prototype.$axios = axios
Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
  ...App
})
app.$mount()
