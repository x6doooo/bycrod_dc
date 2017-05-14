const gulp = require('gulp');
const gutil = require('gulp-util');

const webpack = require('webpack');
const gulpConf = require('../../gulp_conf/gulp.conf');
const browsersync = require('browser-sync');


exports.wrapper = function (watch, conf, done) {


    const webpackBundler = webpack(conf);

    const webpackChangeHandler = (err, stats) => {
        if (err) {
            gulpConf.errorHandler('Webpack')(err);
        }
        gutil.log(stats.toString({
            colors: true,
            chunks: false,
            hash: false,
            version: false
        }));
        if (done) {
            done();
            done = null;
        } else {
            browsersync.reload();
        }
    };

    if (watch) {
        webpackBundler.watch(200, webpackChangeHandler);
    } else {
        webpackBundler.run(webpackChangeHandler);
    }
}
