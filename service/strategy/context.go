package strategy

const (
    init_money float64 = 1, 000, 000.0
)

type deal struct {
    num    float64
    price  float64
    ts     int64
    method string
}

type store struct {
    num   float64
    price float64
}

type Context struct {
    money  float64
    stores map[string]store
    deals  []deal
}

func (me Context) Order(code string, price float64, storePosition float64) {

}

func NewContext() *Context {
    return &Context{
        money: init_money,
        stores: make(map[string]store),
        deals: []deal{},
    }
}
