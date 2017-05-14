/**
 * Created by dx.yang on 2017/5/3.
 */


import request from '../../service/request';
import * as constants from './constants'


let poll_task_status_error_count = 0;
let poll_task_status_error_max = 5;
let poll_task_status_timer = 1;
let waitTime = 1000 * 5;

export default {
    [constants.POLL_TASK_STATUS](store, data) {
        let again = () => {
            poll_task_status_timer = setTimeout(() => {
                store.dispatch(constants.POLL_TASK_STATUS);
            }, waitTime)
        };
        request.PollTaskStatus().then(d => {
            if (!poll_task_status_timer) {
                return;
            }
            again();
            store.commit(constants.POLL_TASK_STATUS, d);
        }).catch(e => {
            poll_task_status_error_count += 1;
            if (poll_task_status_error_count < poll_task_status_error_max) {
                again();
            }
        })
    },
    [constants.STOP_POLL_TASK_STATUS](store, data) {
        clearTimeout(poll_task_status_timer);
        poll_task_status_error_max = null;
    },

    [constants.UpdateStockList](store, data) {
        request.UpdateStockList().then(data => {

        }).catch(e => {

        });
        store.commit(constants.POLL_TASK_STATUS, {
            UpdateStockList: true
        })
    }
}

