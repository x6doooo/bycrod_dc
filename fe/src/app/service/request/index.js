/**
 * Created by dx.yang on 2017/5/3.
 */


import axios from 'axios';
import _ from 'lodash';
import apis from '../apis';


let obj = {};

_.forEach(apis, (config, k) => {
    obj[k] = function(cfg) {
        cfg = cfg || {};
        let defaultConfig = _.cloneDeep(config);
        cfg = Object.assign(defaultConfig, cfg);
        return axios(cfg).then(resp => {
            if (resp.status != 200) {
                throw new Error(resp.statusText);
            } else if (resp.data.code != 0) {
                throw new Error(resp.data.data);
            } else {
                return resp.data.data;
            }
        });
    };
});

export default obj;
