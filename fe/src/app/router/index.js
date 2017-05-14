/**
 * Created by dx.yang on 2017/2/13.
 */


import VueRouter from 'vue-router'

import WrapView from '../container/WrapView/index.vue'
import MainView from '../container/MainView/index.vue'
import ChartView from '../container/ChartView/index.vue'


export default new VueRouter({
    mode: 'hash',
    routes: [
        {
            name: 'WrapView',
            path: '/',
            component: WrapView,
            children: [{
                name: 'MainView',
                path: '/main',
                component: MainView
            }, {
                name: 'ChartView',
                path: '/chart/:code',
                component: ChartView
            }]
        },
        // {
        //     name: 'MainView',
        //     path: '/',
        //     component: MainView,
        //     // children: [
        //     //     {
        //     //         name: 'KecView',
        //     //         path: 'kec',
        //     //         component: KecView
        //     //     },
        //     // ]
        // },
        // {
        //     path: '/',
        //     redirect: '/prod/kec'
        // }
    ],
});
