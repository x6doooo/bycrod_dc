/**
 * Created by dx.yang on 2017/5/3.
 */

import vue from 'vue';
import * as constants from './constants';

export default {
    [constants.POLL_TASK_STATUS](state, data) {
        _.forEach(data, (v, k) => {
            vue.set(state.taskStatus, k, v);
        });
    }
};