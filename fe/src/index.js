import Vue from 'vue';

require('font-awesome/css/font-awesome.css');
require('./app/style/common.sass');

window.$ = require('jquery');
window._ = require('lodash');

require('./app/service/socket');

import 'element-ui/lib/theme-default/index.css'
import ElementUI from 'element-ui'

Vue.use(ElementUI);

import store from './app/store';
import VueRouter from 'vue-router';
Vue.use(VueRouter);

import router from './app/router';

export default new Vue({
    el: '#root',
    store,
    router,
    render: h => h('router-view')
});
