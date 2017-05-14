/**
 * Created by dx.yang on 2017/5/3.
 */


import request from '../../service/request';
import {mapState} from 'vuex';
import * as commonConstants from '../../store/common/constants';

export default {
    name: 'MainView',

    data() {
        return {

            watchingStatusSelected: 'ALL',

            search: '',
            exchangeSelected: 'ALL',
            count: 0,
            currentPage: 1,
            currentRows: 50,
            data: []
        };
    },
    computed: {
        ...mapState(['common'])
    },
    methods: {

        UpdateStockList() {
            this.$store.dispatch(commonConstants.UpdateStockList)
        },

        handleSizeChange(size) {
            this.currentRows = size;
            this.GetStockList();
        },
        handleCurrentChange(page) {
            this.currentPage = page;
            this.GetStockList();
        },


        LoadTimeSeriesData() {
            request.LoadTimeSeriesData({
                params: {
                    type: 'daily'
                }
            }).then(data => {

            }).catch(e => {

            });
        },


        watchingChange(row) {
            request.ModifyWatchingState({
                params: {
                    code: row.Code,
                    state: row.Watching
                }
            }).catch(e => {
                row.Watching = !row.Watching;
            })
        },

        GetStockList() {
            let params = {
                page: this.currentPage,
                rows: this.currentRows,
            };

            let search = _.trim(this.search);
            if (search) {
                params.code = search;
            }

            let exchange = this.exchangeSelected;
            if (exchange && exchange !== 'ALL') {
                params.exchange = exchange;
            }

            let watching = this.watchingStatusSelected;
            if (watching && watching !== 'ALL') {
                params.watching = watching;
            }


            request.GetStockList({
                params
            }).then(data => {
                this.data = data.data;
                this.count = data.count;
                this.currentPage = data.page;
                this.currentRows = data.rows;
            });
        },
        exchangeChange() {
            this.GetStockList();
        },
        watchingStatusSelectedChange() {
            console.log(this.watchingStatus);
            this.GetStockList();
        },
        go(code) {
            this.$router.push({
                name: 'ChartView',
                params: {
                    code
                }
            })
        }
    },
    created() {
        this.GetStockList()
    }
};