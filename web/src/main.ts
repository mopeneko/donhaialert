import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import NavBar from "buefy";

Vue.config.productionTip = false;

Vue.use(NavBar);

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
