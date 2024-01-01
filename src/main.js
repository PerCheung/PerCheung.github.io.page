import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.use(ElementUI);
Vue.config.productionTip = false
Vue.prototype.API_BASE_URL = 'http://124.221.87.224';
// Vue.prototype.API_BASE_URL = 'http://127.0.0.1:8080';

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')