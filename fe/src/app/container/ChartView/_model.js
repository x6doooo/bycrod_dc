/**
 * Created by dx.yang on 2017/5/5.
 */


export default {
    chartOptions: {
        chart: {
            height: 700
        },
        rangeSelector: {
            selected: 1
        },

        title: {
            text: ''
        },

        yAxis: [{
            labels: {
                align: 'right',
                x: -3
            },
            title: {
                text: ''
            },
            height: '80%',
            lineWidth: 2
        }, {
            labels: {
                align: 'right',
                x: -3
            },
            title: {
                text: 'Volume'
            },
            top: '85%',
            height: '15%',
            offset: 0,
            lineWidth: 2
        }],

        tooltip: {
            split: true
        },
    }
};
