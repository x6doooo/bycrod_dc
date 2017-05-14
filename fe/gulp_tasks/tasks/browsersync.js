// const gulp = require('gulp');
const browserSync = require('browser-sync');
const spa = require('browser-sync-spa');

const browserSyncConf = require('../../gulp_conf/browsersync.conf');
const open = require('opn');
// const conf = require('../conf/gulp.conf');

browserSync.use(spa());


exports.devServer = function (conf, isMock) {
    let devConf = {
        baseDir: [conf.paths.tmp],
        server: conf.server,
        isMock
    };
    return function (done) {
        browserSync.init(browserSyncConf(devConf));
        open(conf.server.localServerDomain);
        done();
    };
};

exports.distServer = function (conf) {
    let distConf = {
        baseDir: [conf.paths.dist],
        server: conf.server,
        isMock: false
    };
    return function (done) {
        browserSync.init(browserSyncConf(distConf));
        open(conf.server.localServerDomain);
        done();
    };
};

exports.reload = function (cb) {
    browserSync.reload();
    cb();
};
