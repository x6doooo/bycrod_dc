/**
 * Created by dx.yang on 16/8/4.
 */


const path = require('path');

exports.paths = {
    src: 'src',
    dist: 'dist',
    tmp: '.tmp'
};

exports.path_alias = path.resolve(__dirname, '../' + exports.paths.src);

exports.server = {
    localServerDomain: 'http://dev.bycrod.com',
    backend: 'http://127.0.0.1:54321',
    backendHost: 'dev.bycrod.com'
};