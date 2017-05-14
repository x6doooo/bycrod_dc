/**
 * Created by dx.yang on 2017/5/3.
 */

export default {
    GetStockList: {
        url: '/api/stock/list'
    },
    PollTaskStatus: {
        url: '/api/common/task/status'
    },
    UpdateStockList: {
        url: '/api/stock/list/update'
    },
    ModifyWatchingState: {
        method: 'PUT',
        url: '/api/stock/watching'
    },
    LoadTimeSeriesData: {
        url: '/api/stock/timeSeriesData/load'
    },
    QueryTimeSeriesData: {
        url: '/api/stock/timeSeriesData/query'
    },

    //talib
    talibCdl2Crows: {
        url: '/api/stock/talib/Cdl2Crows'
    },
    talibCdl3BlackCrows: {
        url: '/api/stock/talib/Cdl3BlackCrows'
    },
    talibCdl3Inside: {
        url: '/api/stock/talib/Cdl3Inside'
    }
};


