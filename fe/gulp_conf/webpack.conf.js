const webpack = require('webpack');
// const gulpConf = require('./gulp.conf');
const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const FailPlugin = require('webpack-fail-plugin');
const autoprefixer = require('autoprefixer');

// const pkg = require('../package.json');

let baseWebpackConf = {
    module: {
        rules: [
            {
                test: /.json$/,
                loaders: [
                    'json-loader'
                ]
            },
            {
                test: /\.(css|scss|sass)$/,
                loaders: [
                    'style-loader',
                    'css-loader',
                    'sass-loader',
                    // 'postcss-loader',
                ]
            },
            {
                test: /\.js$/,
                // exclude: /node_modules/,

                // exclude: /node_modules(?!(\/|\\)ksc-vue-ui(\/|\\)src)/,
                exclude: /node_modules(?!(\/|\\)ksc-vue-ui2(\/|\\)src)/,
                loaders: [
                    'babel-loader'
                ]
            },
            {
                test: /.vue$/,
                loader: 'vue-loader',
                options: {
                    loaders: {
                        scss: 'style-loader!css-loader!sass-loader', // <style lang="scss">
                        sass: 'style-loader!css-loader!sass-loader?indentedSyntax' // <style lang="sass">
                    }
                }
            },
            {
                // Capture eot, ttf, svg, woff, and woff2 png
                test: /\.(woff2?|ttf|svg|eot)(\?v=\d+\.\d+\.\d+)?$/,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: './assets/fonts/[name].[ext]'
                    }
                },
            },
            {
                test: /\.(png|jpg|gif)(\?v=\d+\.\d+\.\d+)?$/,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: './assets/img/[name].[ext]'
                    }
                },
            }
        ]
    },
};


exports.dev = function(conf) {
    return Object.assign(baseWebpackConf, {
        plugins: [
            new webpack.optimize.OccurrenceOrderPlugin(),
            new webpack.NoEmitOnErrorsPlugin(),
            FailPlugin,
            new HtmlWebpackPlugin({
                template: conf.path.src('index.html')
            }),
            new webpack.optimize.CommonsChunkPlugin({
                name: 'commons',
                filename: 'scripts/commons.js',
                minChunks: 2
            }),
        ],
        resolve: {
            alias: {
                develop: conf.path_alias
            }
        },
        devtool: 'source-map',
        output: {
            path: path.join(process.cwd(), conf.paths.tmp),
            filename: 'scripts/[name]-[chunkhash].js',
            chunkFilename: 'scripts/chunks/[name]-[chunkhash].js'
        },
        entry: {
            index: `./${conf.path.src('index')}`,
            vendor: Object.keys(conf.pkg.dependencies)
        }
    });
};

exports.dist = function(conf, needUpload) {
    let c = Object.assign(baseWebpackConf, {
        plugins: [
            new webpack.optimize.OccurrenceOrderPlugin(),
            new webpack.NoEmitOnErrorsPlugin(),
            FailPlugin,
            new webpack.optimize.CommonsChunkPlugin({
                name: 'commons',
                filename: 'scripts/commons-[chunkhash].js',
                minChunks: 2
            }),
            new HtmlWebpackPlugin({
                template: conf.path.src('index.html')
            }),
            new webpack.DefinePlugin({
                'process.env.NODE_ENV': '"production"'
            }),
            new webpack.optimize.UglifyJsPlugin({
                compress: {unused: true, dead_code: true, warnings: false} // eslint-disable-line camelcase
            }),
            new ExtractTextPlugin('index-[chunkhash].css')
        ],
        resolve: {
            alias: {
                develop: conf.path_alias
            }
        },
        output: {
            path: path.join(process.cwd(), conf.paths.dist),
            filename: 'scripts/[name]-[chunkhash].js',
            chunkFilename: 'scripts/chunks/[name]-[chunkhash].js'
        },
        entry: {
            index: `./${conf.path.src('index')}`,
            vendor: Object.keys(conf.pkg.dependencies)
        }
    });

    if (needUpload) {
        c.output.publicPath = '//secnetweb.ksyun.com/' + conf.pkg.name + '/' + conf.pkg.version + '/';
    }

    return c;
};

