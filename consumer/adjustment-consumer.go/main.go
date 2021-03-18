package main

func main() {
	mq.NewConsumer("stock_adjust", handleAdjustment)
}

func handleAdjustment() {

}
