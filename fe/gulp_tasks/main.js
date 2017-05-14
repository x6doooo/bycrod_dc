/**
 * Created by dx.yang on 2017/2/10.
 */


const gulp = require('gulp');
// const HubRegistry = require('gulp-hub');
// const browserSync = require('browser-sync');
const path = require('path');
const child_process = require('child_process');

// conf
const webpackConf = require('../gulp_conf/webpack.conf');

// tasks
const browserSyncTasks = require('./tasks/browsersync');
const webpackTasks = require('./tasks/webpack');
const misc = require('./tasks/misc');

const ontlineTasks = require('./tasks/online');

module.exports = function(conf) {

    conf.paths.tasks = './tasks';

    let pathMap = {};
    for (const pathName in conf.paths) {
        if (conf.paths.hasOwnProperty(pathName)) {
            pathMap[pathName] = function pathJoin() {
                const pathValue = conf.paths[pathName];
                const funcArgs = Array.prototype.slice.call(arguments);
                const joinArgs = [pathValue].concat(funcArgs);
                return path.join.apply(this, joinArgs);
            };
        }
    }
    conf.path = pathMap;

    gulp.task('clean', misc.clean(conf));
    gulp.task('other', misc.other(conf));


    gulp.task('browserSync', browserSyncTasks.devServer(conf));
    gulp.task('browserSync:mock', browserSyncTasks.devServer(conf, true));
    gulp.task('browserSync:dist', browserSyncTasks.distServer(conf));

    gulp.task('webpack:watch', done => {
        webpackTasks.wrapper(true, webpackConf.dev(conf), done);
    });

    gulp.task('webpack:dist', done => {
        process.env.NODE_ENV = 'production';
        webpackTasks.wrapper(false, webpackConf.dist(conf), done);
    });


const tagReg = /v\d+\.\d+\.\d+/;
function compareTag(tag_a, tag_b) {
    tag_a = tag_a.replace('v', '').split('.');
    tag_b = tag_b.replace('v', '').split('.');
    var i = 0;
    var c = 0;
    while(i < 3) {
        c = tag_a[i] - tag_b[i];
        if (c) {
            break;
        }
        i++;
    }
    return c;
}
    gulp.task('webpack:dist:online', done => {

            var list = child_process.execSync('git tag -l');

    var tags = list.toString().split('\n');

    var maxTag = 'v0.0.0';
    tags.forEach(function(tag) {
        if (!tag || !tagReg.test(tag)) {
            return;
        }
        var compareResult = compareTag(tag, maxTag);
        if (compareResult > 0) {
            maxTag = tag;
        }
    });

    //console.log(maxTag);
    //return;

    //var ver = packageJSON.version;
    var ver = maxTag.replace('v', '').split('.')//.replace('v', '');

    // ver[2] = ver[2] * 1 + 1;
    if (ver[2] * 1 < 99) {
        ver[2] = ver[2] * 1 + 1;
    } else {
        ver[2] = 1;
        if (ver[1] * 1 < 99) {
            ver[1] = ver[1] * 1 + 1;
        } else {
            ver[1] = 1;
            ver[0] = ver[0] * 1 + 1;
        }
    }

    conf.pkg.version = ver.join('.');

        process.env.NODE_ENV = 'production';
        webpackTasks.wrapper(false, webpackConf.dist(conf, true), done);
    })

    function watch(done) {
        console.log(pathMap.tmp('index.html'));
        gulp.watch(pathMap.tmp('index.html'), browserSyncTasks.reload);
        done();
    }

    // online
    gulp.task('online-update-version', ontlineTasks.updateVersion(conf));
    gulp.task('online-rewrite-info', ontlineTasks.rewriteInfo(conf));
    gulp.task('online-ks3', ontlineTasks.ks3);

    gulp.task('watch', watch);
    gulp.task('build', gulp.series(gulp.parallel('other', 'webpack:dist')));
    gulp.task('build:online:beCareful', gulp.series(gulp.parallel(
        'other', 
        'webpack:dist:online'
    ),
        'online-rewrite-info',
        'online-update-version',
        'online-ks3'
    ));

    gulp.task('build:test', gulp.series(gulp.parallel(
        'other',
        'webpack:dist'
        ),
        'online-rewrite-info'
    ));

    gulp.task('default', gulp.series('clean', 'build'));
    gulp.task('serve', gulp.series('other', 'webpack:watch', 'watch', 'browserSync'));
    gulp.task('serve:mock', gulp.series('other', 'webpack:watch', 'watch', 'browserSync:mock'));
    gulp.task('serve:dist', gulp.series('browserSync:dist'));
};