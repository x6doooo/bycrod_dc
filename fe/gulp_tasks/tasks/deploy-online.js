/**
 * Created by dx.yang on 16/6/17.
 */

var parseString = require('xml2js').parseString;
var KS3 = require('ks3');
var path = require('path');
var Promise = require('promise');
var client;
var config;
try {
    config = require('../../../../config/deploy-online-config');
    client = new KS3(config.accessKey, config.secretKey);
} catch(e) {
    //....
}

function apiHandler(method, params) {
    return new Promise(function(resolve, reject) {
        var args = [];
        if (params) {
            args = [params];
        }
        args.push(function(e, d) {
            if (e) {
                reject(e);
            }
            parseString(d, function(e, parsedData) {
                if (e) {
                    reject(e);
                }
                parsedData = JSON.stringify(parsedData, null, 4);
                resolve(parsedData);
            });
        });
        method.apply(null, args);
    });
}

function putBucketACL(acl) {
    return apiHandler(client.bucket.putACL, {
        Bucket: config.bucket,
        ACL: acl
    });
}

function getBucketACL() {
    return apiHandler(client.bucket.getACL, {
        Bucket: config.bucket
    });
}

function getService() {
    return apiHandler(client.service.get);
}

function listObjects() {
    return apiHandler(client.bucket.get, {
        Bucket: config.bucket
    });
}

function putObject(params) {
    return apiHandler(client.object.put, params);
}

var APIs = {
    putBucketACL: putBucketACL,
    getBucketACL: getBucketACL,
    getService: getService,
    listObjects: listObjects,
    putObject: putObject
};

function go(which, params) {
    APIs[which](params).then(function(d) {
        console.log(d);
        console.log('done');
    }).catch(function(e) {
        throw e;
    });
}



function uploadDist(cb) {
    var fse = require('fs-extra');
    var packageJSON = require('../../../../package.json');
    //secnetweb.fe.ksyun.com/project/
    var prefix = [
        packageJSON.name,
        packageJSON.version
    ].join('/');


    var failedCount = 0;
    var failedCountMax = 10;

    
    var files = [];
    
    function upload() {
        setTimeout(function() {


            if (files.length) {

                if (failedCount > failedCountMax) {
                    console.log('-----------------------');
                    console.log('upload failed!');
                    console.log(files.join('\n'));
                    cb();
                    return;
                }

                var f = files.pop();
                APIs.putObject({
                    Bucket: config.bucket,
                    Key: f.key,
                    filePath: f.path,
                    ACL: 'public-read'
                }).then(function(d) {
                    console.log('[done]');
                    upload();
                }).catch(function(e) {
                    console.log('[error]', e, f);

                    // 上传错误计数
                    failedCount++;

                    // 重新上传
                    files.unshift(f);
                    console.log('re-upload');
                    upload();
                });
            } else {
                cb();
            }
        }, 500);
    }
    
    var rootPath = path.resolve(__dirname, '../../../../dist');
    fse.walk(rootPath)
        .on('data', function(item) {
            //var path = item.path.replace(__dirname + '/dist', '');
            var thePath = item.path.replace(rootPath, '');
            var isDic = thePath.indexOf('.') === -1;
            var isIndex = thePath.indexOf('index.html') !== -1;
            var isStaticFile = !isDic && !isIndex;
            if (isStaticFile) {
                files.push({
                    key: prefix + thePath,
                    path: item.path
                });
            }
        }).on('end', function() {
            //console.log(files.join('\n'));
            // files.forEach(function(f) {

            //     console.log(JSON.stringify(f));

            // });
            upload();
        });
}

exports.uploadDist = uploadDist;
// go('listObjects');


