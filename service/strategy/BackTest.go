package strategy

func BackTest(strategyName string) {

    strategyInstance := makeStrategyInstance(strategyName)
    strategyInstance.Init()



}
