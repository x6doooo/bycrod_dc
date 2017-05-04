package stock

import (
    "gopkg.in/mgo.v2/bson"
)

type Record struct {
    Id            bson.ObjectId `bson:"_id,omitempty"`
    Watching      bool
    NoUpdate      bool `bson:"no_update"`
    Symbol        string
    Exchange      string
    Code          string
    Name          string
    Current       float64                         // 当前价格
    Percentage    float64                         // 涨跌幅度
    Change        float64                         // 涨跌额
    Open          float64
    High          float64
    Low           float64
    Close         float64
    LastClose     float64 `bson:"last_close"`
    High52week    float64                         // 52周 最高
    Low52week     float64                         // 52周 最低
    Volume        float64                         // 成交量
    VolumeAverage float64 `bson:"volume_average"` // 平均成交量
    MarketCapital float64 `bson:"market_capital"` // 市值
    Eps           float64                         // 每股收益
    PeTtm         float64 `bson:"pe_ttm"`         // 市盈率 TTM
    PeLyr         float64 `bson:"pe_lyr"`         // 市盈率 LYR
    Beta          float64                         // beta值   风险指标？
    TotalShares   float64 `bson:"total_shares"`   // 总股本
    AfterHours    float64 `bson:"after_hours"`
    AfterHoursPct float64 `bson:"after_hours_pct"`
    AfterHoursChg float64 `bson:"after_hours_chg"`
    Dividend      float64                         // 股息/红利
    Yield         float64                         // 收益
    TurnoverRate  float64 `bson:"turnover_rate"`  // 换手率
    InstOwn       float64 `bson:"inst_own"`       // 机构持股
    RiseStop      float64 `bson:"rise_stop"`
    FallStop      float64 `bson:"fall_stop"`
    Amount        float64
    NetAssets     string `bson:"net_assets"`      // 美股净资产
    Pb            float64                         // 市净率MRQ
}

