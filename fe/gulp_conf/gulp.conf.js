'use strict';

/**
 *  This file contains the variables used in other gulp files
 *  which defines tasks
 *  By design, we only put there very generic config values
 *  which are used in several places to keep good readability
 *  of the tasks
 */

// const path = require('path');
const gutil = require('gulp-util');

/**
 *  The main paths of your project handle these with care
 */

/**
 * used on gulp dist
 */
exports.htmlmin = {
    ignoreCustomFragments: [/{{.*?}}/]
};


/**
 *  Common implementation for an error handler of a Gulp plugin
 */
exports.errorHandler = function (title) {
    return function (err) {
        gutil.log(gutil.colors.red(`[${title}]`), err.toString());
        this.emit('end');
    };
};

// exports.server = {
//     localServerDomain: 'http://abc.console.ksyun.com:3000',
//     // backend: 'http://10.111.17.86',
//     backend: 'http://network.console.ksyun.com',
//     // back
// };
