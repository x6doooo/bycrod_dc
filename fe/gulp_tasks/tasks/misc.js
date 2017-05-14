const path = require('path');

const gulp = require('gulp');
const del = require('del');
const filter = require('gulp-filter');

const conf = require('../../gulp_conf/gulp.conf');



exports.clean = function(conf) {
    return function() {
        return del([conf.paths.dist, conf.paths.tmp]);
    }
};


exports.other = function(conf) {
    return function() {
        const fileFilter = filter(file => {
            //console.log(file.path);
            return file.stat.isFile()
        });
        return gulp.src([
            // path.join('./others', '/**/*'),
            path.join(conf.paths.src, '/**/*'),
            path.join(`!${conf.paths.src}`, '/**/*.{sass,scss,js,html,vue,pug,jade,svg,woff,tff,eot}'),
            path.join(conf.paths.src, '/assets/**/*.js')
        ])
            .pipe(fileFilter)
            .pipe(gulp.dest(conf.paths.tmp));
    };
};

