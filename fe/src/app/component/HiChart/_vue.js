/**
 * Created by dx.yang on 2017/5/5.
 */



import Vue from 'vue'
import highcharts from 'highcharts/highstock'
require('highcharts/modules/no-data-to-display')(highcharts)



let defaultOptions = {
    chart: {
        type: 'areaspline',
        backgroundColor: '#fcfcfc',
        borderColor: '#d8d8d8',
        borderWidth: 1,
        height: 600
    },
    xAxis: {
        labels: {
            align: 'left'
        },
    },
    yAxis: {},
    tooltip: {
        borderColor: '#bfbfbf',
        shadow: false,
        valueSuffix: '',
        valueDecimals: 6
    },
    labels: {
        formatter: function () {
            return this.value.toFixed(8);//这里是两位小数，你要几位小数就改成几
        },
        style: {
            color: 'red'
        }
    },
    credits: {
        text: '',
        href: ''
    },

    plotOptions: {
        areaspline: {
            lineWidth: 1,
            fillOpacity: 0.2
        },
        series: {
            dataGrouping: {
                approximation: 'high',

                dateTimeLabelFormats: {
                    millisecond: ['%A, %b %e日, %H:%M:%S.%L', '%A, %b %e日, %H:%M:%S.%L', '-%H:%M:%S.%L'],
                    second: ['%A, %b %e日, %H:%M:%S', '%A, %b %e日, %H:%M:%S', '-%H:%M:%S'],
                    minute: ['%A, %b %e日, %H:%M', '%A, %b %e日, %H:%M', '-%H:%M'],
                    hour: ['%A, %b %e日, %H:%M', '%A, %b %e日, %H:%M', '-%H:%M'],
                    day: ['%A, %b %e日, %Y', '%A, %b %e日', '-%A, %b %e日, %Y'],
                    week: ['Settimana del %d/%m/%Y', '%A, %b %e日', '-%A, %b %e日, %Y'],
                    month: ['%B %Y', '%B', '-%B %Y'],
                    year: ['%Y', '%Y', '-%Y']
                }
            },
            tooltip: {
                dateTimeLabelFormats: {
                    millisecond: "%A, %b %e日, %H:%M:%S.%L",
                    second: "%A, %b %e日, %H:%M:%S",
                    minute: "%A, %b %e日, %H:%M",
                    hour: "%A, %b %e日, %H:%M",
                    day: "%A, %b %e日, %Y",
                    week: "Week from %A, %b %e日, %Y",
                    month: "%B %Y",
                    year: "%Y"
                }
            }
        }
    },
    exporting: {
        enabled: false
    },
    colors: ['#56aff0', '#ff8a00', '#50a157', /* 后面都是默认颜色 */ '#7cb5ec', '#434348', '#90ed7d', '#f7a35c',
        '#8085e9', '#f15c80', '#e4d354', '#2b908f', '#f45b5b', '#91e8e1'],
    legend: {
        enabled: true,
    },
    navigator: {
        margin: 10
    },
    rangeSelector: {
        enabled: false
    },
    title: {
        align: 'left',
        text: ''
    },
};



export default Vue.extend({
    name: 'HiChart',
    data() {
        return {
            chart: null
        };
    },
    props: {
        options: {
            type: Object,
            required: false,
            default: () => {
            }
        },
        series: {
            type: Array,
            required: false,
            default: () => []
        },
        loading: {
            type: Boolean,
            required: false,
            default: false
        },

    },
    watch: {
        series: {
            deep: true,
            handler() {
                this.setSeries(this.series);
            }
        },
        options: {
            deep: true,
            handler() {
                this.render(this.options);
            }
        },
        loading() {
            if (this.loading) {
                this.chart && this.chart.highcharts && this.chart.highcharts().showLoading();
            } else {
                this.chart && this.chart.highcharts && this.chart.highcharts().hideLoading();
            }
        }
    },
    methods: {
        setSeries(series) {
            this.render();
        },
        render() {
            let me = this;
            let options = this.options || {};
            options = _.defaultsDeep({}, options, defaultOptions);
            options.series = this.series;

            let containter = $(me.$el).find('.hi-chart-container-inner').get(0);
            me.chart = highcharts.stockChart(containter, options);
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.render();
        });
    }
});