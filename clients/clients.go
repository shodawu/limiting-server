package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MultiRun(iRuns int, iDuration, iRetry time.Duration) {
	for iCli := 1; iCli <= iRuns; iCli++ {
		go func(i int) {
			for {
				res, err := http.Get("http://localhost:1234/hit")
				if err != nil {
					fmt.Printf("Cli:%v, err: %s\n", i, err)
					time.Sleep(iRetry * time.Second)
					continue
				}
				data, _ := ioutil.ReadAll(res.Body)
				fmt.Printf("Cli:%v, data: %s\n", i, data)
				res.Body.Close()

				time.Sleep(iDuration * time.Millisecond)
			}
		}(iCli)
	}
	select {}
}
