/**
 * Created by dx.yang on 2017/5/5.
 */


import HiChart from '../../component/HiChart/index.vue';
import request from '../../service/request';
import moment from 'moment';
import axios from 'axios'
import model from './_model';


function pickerShortCuts(picker, offsetDays) {
    const end = new Date();
    const start = new Date();
    start.setTime(start.getTime() - 3600 * 1000 * 24 * offsetDays);
    picker.$emit('pick', [start, end]);
}



function formatSeries(data, cdl2Data, cdl3Data, cdl3insideData) {
    let candleStickSeries = [];
    let volumeSeries = [];
    let cdl2DataSeries = [];
    let cdl3DataSeries = [];
    let cdl3insideDataSeries = [];
    let size = data.Ts.length;
    for (let i = 0; i < size; i++) {
        let ts = data.Ts[i] * 1000;
        candleStickSeries.push([
            ts,
            data.Open[i],
            data.High[i],
            data.Low[i],
            data.Close[i],
        ]);
        volumeSeries.push([
            ts,
            data.Volume[i]
        ]);
        if (cdl2Data[i] != 0) {
            cdl2DataSeries.push({
                x: ts,
                title: '-',
                text: 'short'
            });
        }
        if (cdl3Data[i] != 0) {
            cdl3DataSeries.push({
                x: ts,
                title: '-',
                text: 'short'
            });
        }
        if (cdl3insideData[i] != 0) {
            cdl3insideDataSeries.push({
                x: ts,
                title: '+',
                text: 'long'
            })
        }
    }
    // console.log(candleStickSeries);
    let groupingUnits = [[
        'week',                         // unit name
        [1]                             // allowed multiples
    ], [
        'month',
        [1, 2, 3, 4, 6]
    ]];
    return [{
        type: 'candlestick',
        name: 'candl',
        data: candleStickSeries,
        dataGrouping: {
            units: groupingUnits
        }
    }, {
        type: 'column',
        name: 'Volume',
        data: volumeSeries,
        yAxis: 1,
        dataGrouping: {
            units: groupingUnits
        }
    }, {
        type: 'flags',
        data: cdl2DataSeries,
        onSeries: 'candl',
        shape: 'circlepin',
        width: 12
    }, {
        type: 'flags',
        data: cdl3DataSeries,
        onSeries: 'candl',
        shape: 'squarepin',
        width: 12
    }, {
        type: 'flags',
        data: cdl3insideDataSeries,
        onSeries: 'candl',
        shape: 'squarepin',
        width: 12
    }]
}


export default {
    name: 'ChartView',
    components: {
        HiChart
    },
    data() {
        return {

            code: '',

            loading: false,

            dateRangeType: 2,

            dateRange: '',
            pickerOptions: {
                shortcuts: [{
                    text: '最近一周',
                    onClick(picker) {
                        pickerShortCuts(picker, 7)
                    }
                }, {
                    text: '最近一个月',
                    onClick(picker) {
                        pickerShortCuts(picker, 30)
                    }
                }, {
                    text: '最近三个月',
                    onClick(picker) {
                        pickerShortCuts(picker, 90)
                    }
                }, {
                    text: '最近六个月',
                    onClick(picker) {
                        pickerShortCuts(picker, 180)
                    }
                }, {
                    text: '最近一年',
                    onClick(picker) {
                        pickerShortCuts(picker, 365)
                    }
                }]
            },
            chartSeries: [],
            chartOptions: model.chartOptions
        };
    },

    methods: {
        loadData() {
            this.loading = true;

            let params = {
                code: this.$route.params.code,
                dataType: 'daily',
            };

            if (this.dateRangeType === 2 && Array.isArray(this.dateRange)) {
                let startDate = this.dateRange[0];
                let endDate = this.dateRange[1];
                params.startDate = moment(startDate).format('YYYY-MM-DD');
                params.endDate = moment(endDate).format('YYYY-MM-DD');
            }

            let req1 = request.QueryTimeSeriesData({
                params
            });
            //     .then(data => {
            //     this.chartSeries = formatSeries(data);
            //     this.loading = false;
            // }).catch(e => {
            //     this.loading = false;
            // });

            let req2 = request.talibCdl2Crows({
                params
            });

            let req3 = request.talibCdl3BlackCrows({
                params
            });

            let req4 = request.talibCdl3Inside({
                params
            });

            this.loading = true;
            axios.all([req1, req2, req3, req4]).then(axios.spread((tsData, cdl2Data, cdl3Data, cdl3insideData) => {
                // console.log(tsData, markData);
                this.chartSeries = formatSeries(tsData, cdl2Data, cdl3Data, cdl3insideData);
                this.loading = false;
                // Both requests are now complete
            }));
            //     .then(data => {
            //     console.log(data);
            // }).catch(e => {});
        },

        back() {
            this.$router.push({
                name: 'MainView'
            })
        }
    },

    created() {},

}