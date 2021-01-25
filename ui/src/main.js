import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import store from "@/store";
import router from "@/router";
import Vuelidate from "vuelidate/src";

Vue.config.productionTip = false

Vue.use(Vuelidate)

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
