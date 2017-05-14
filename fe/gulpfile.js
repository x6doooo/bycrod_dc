let conf = require('./config/config');
let mainGulpTask = require('./gulp_tasks/main');
let pkg = require('./package.json');
conf.pkg = pkg;

mainGulpTask(conf);