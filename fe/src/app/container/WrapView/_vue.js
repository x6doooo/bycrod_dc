/**
 * Created by dx.yang on 2017/5/5.
 */


import * as commonConstants from '../../store/common/constants'

export default {
    name: 'WrapView',
    beforeRouteEnter(to, from, next) {
        next(vm => {
            vm.$store.dispatch(commonConstants.POLL_TASK_STATUS)
        });
    },
    beforeRouteLeave(to, from, next) {
        this.$store.dispatch(commonConstants.STOP_POLL_TASK_STATUS);
        next();
    },
}