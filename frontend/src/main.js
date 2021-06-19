import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import vuetify from './plugins/vuetify'
import utils from './utils'
import VueQr from 'vue-qr'

Vue.use(VueQr)

Vue.config.productionTip = false
Vue.prototype.utils = utils

new Vue({
    router,
    vuetify,

    render: function (h) {
        return h(App)
    }
}).$mount('#app')
