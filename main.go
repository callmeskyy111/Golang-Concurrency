package main

import (
	"fmt"

	"example.com/concurrency/filemanager"
	"example.com/concurrency/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool,len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		errorChans[idx] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[idx], errorChans[idx])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	// Err. Handling with multiple-channels
	for idx:= range taxRates {
    select {
	case err:= <-errorChans[idx]:
		if err != nil{
			fmt.Println("🔴 ERROR: ",err)
		}
	case <-doneChans[idx]:
		fmt.Println("Done ✅")	
	 }
	}

	// for _, errorChan := range errorChans{
	// 	<- errorChan
	// }

	// for _, doneChan := range doneChans{
	// 	<- doneChan
	// }
}
