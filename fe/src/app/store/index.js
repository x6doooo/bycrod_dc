/**
 * Created by dx.yang on 2017/5/2.
 */

import Vue from 'vue';
import Vuex from 'vuex';
// import mutations, {initialState as state} from './mutations';
// import * as actions from './actions';

Vue.use(Vuex);

// import nav from './nav';
import common from './common';
// import kec from './kec';
// import eip from './eip';
// import epc from './epc';

export default new Vuex.Store({
    modules: {
        // nav,
        common,
        // kec,
        // eip,
        // epc
    }
});




