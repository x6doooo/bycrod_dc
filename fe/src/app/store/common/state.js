/**
 * Created by dx.yang on 2017/5/3.
 */


export default {
    Exchanges: [{
        key: 'ALL',
        label: '全部交易所'
    }, {
        key: 'NYSE',
        label: '纽交所'
    }, {
        key: 'AMEX',
        label: '美交所'
    }, {
        key: 'NASDAQ',
        label: '纳斯达克'
    }],

    watchingStatus: [{
        key: 'ALL',
        label: '全部',
    }, {
        key: 'watching',
        label: '关注中',
    }, {
        key: 'notWatching',
        label: '未关注'
    }],

    taskStatus: {
        UpdateStockList: false,
        CleanCodeCollection: false,
        loadStockDataFromYahooApi: false
    },
};

