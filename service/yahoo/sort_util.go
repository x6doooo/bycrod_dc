package yahoo


type ByHigh []QuoteItem
func (me ByHigh) Len() int {
    return len(me)
}
func (me ByHigh) Swap(i, j int) {
    me[i], me[j] = me[j], me[i]
}
func (me ByHigh) Less(i, j int) bool {
    return me[i].High < me[i].High
}

type ByLow []QuoteItem
func (me ByLow) Len() int {
    return len(me)
}
func (me ByLow) Swap(i, j int) {
    me[i], me[j] = me[j], me[i]
}
func (me ByLow) Less(i, j int) bool {
    return me[i].Low < me[i].Low
}