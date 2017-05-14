// const gulpConf = require('./gulp.conf');

const path = require('path');
const gulp = require('gulp');
const fs = require('fs');
const url = require('url');
const util = require('util');
const querystring = require('querystring');
const httpProxy = require('http-proxy');


module.exports = function (conf) {

    let isMock = conf.isMock;
    let browser = 'default';
    // let baseDir = [
    //     conf.paths.tmp,
    //     conf.paths.src
    // ];
    let server = {
        ghostMode: false,
        baseDir: conf.baseDir,
        routes: null,
    };

    let proxy = httpProxy.createProxyServer({});

    // proxy.on('proxyReq', function(proxyReq, req, res, options) {
    //     var cookies = req.headers.cookie;
    //     cookies = cookies.split('; ');
    //     cookies.forEach(function(v, i) {
    //         v = v.split('=');
    //         if (v[0] === 'kscdigest') {
    //             v[1] = '0460163fc1223a47c3a62057f92a8be7-3501251849';
    //             cookies[i] = v.join('=');
    //         }
    //     });
    //     proxyReq.setHeader('Cookie', cookies);
    // });
    proxy.on('proxyReq', function(proxyReq, req, res, options) {
        proxyReq.setHeader('Host', conf.server.backendHost);
    });
    server.middleware = [
        function (req, res, next) {
            var theURL = url.parse(req.url);
            var query = theURL.query;
            query = querystring.parse(query);

            if (theURL.pathname.indexOf('/api-inner') === 0) {
                if (theURL.pathname.indexOf('/api-inner-mock') === 0) {
                    var p = theURL.pathname.replace('/api-inner-mock', '');
                    var filename = './mock' + p;
                    var thePath = filename + '.json';
                    var fileData = fs.readFileSync(thePath, 'utf8');
                    setTimeout(function () {
                        res.writeHead(200, {"Content-Type": "application/json"});
                        res.write(fileData, 'utf8');
                        res.end();
                    }, 500);
                    return;
                } else {
                    proxy.web(req, res, {
                        target: conf.server.backend
                    });
                    return;
                }
            }

            if (theURL.pathname === '/api-mock') {
                var filename = './mock/' + query.Service + '/' + query.Action;

                // var thePath = theURL.pathname.replace('/', filename) + '.json';
                var thePath = filename + '.json';
                var fileData = fs.readFileSync(thePath, 'utf8');
                setTimeout(function () {
                    res.writeHead(200, {"Content-Type": "application/json"});
                    res.write(fileData, 'utf8');
                    res.end();
                }, 500);
                return;
            }

            let condition1 = ['/api', '/api-sec', '/api-net'].indexOf(theURL.pathname) !== -1 && query.Action;
            let condition2 = theURL.pathname.indexOf('/api') === 0;


            // console.log(condition1, condition2);
            if (condition1 || condition2) {

                // to api
                if (isMock) {
                    var filename = './mock/' + query.Service + '/' + query.Action;

                    // var thePath = theURL.pathname.replace('/', filename) + '.json';
                    var thePath = filename + '.json';

                    var fileData = fs.readFileSync(thePath, 'utf8');
                    setTimeout(function () {
                        res.writeHead(200, {"Content-Type": "application/json"});
                        res.write(fileData, 'utf8');
                        res.end();
                    }, 500);
                } else {
                    // req.url = req.url.replace('/api', '/');
                    proxy.web(req, res, {
                        target: conf.server.backend
                    });
                }
            } else {
                next();
            }
        }
    ];
    return {
        server,
        open: false,
        notify: false,
        browser
    };
};
