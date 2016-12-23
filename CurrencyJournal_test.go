/*
 * B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
 * Created by Ibrahim.COBANI on 15-12-16.
 */

package GOTCMBCurrencyHelper

import (
	"testing"
	"time"
	"fmt"
	"runtime"
	"github.com/icobani/GOTCMBCurrencyHelper/config"
)

func Test(t *testing.T) {
	runtime.GOMAXPROCS(2)
	startTime := time.Now()
	config.Load()
	currencyJournal := new(CurrencyJournal)
	CurrencyDate, _ := time.Parse("02-01-2006", "16-04-1996")
	CurrencyDate, _ = time.Parse("02-01-2006", "01-01-2014")

	var durationDay int = 0;

	for CurrencyDate.Before(time.Now()) {
		CurrencyDate = CurrencyDate.AddDate(0, 0, 1)
		durationDay ++
	}

	CurrencyDate, _ = time.Parse("02-01-2006", "01-01-2014")

	for CurrencyDate.Before(time.Now()) {
		go func(CurrencyDate time.Time) {
			currencyJournal.GetArchive(CurrencyDate)
			durationDay--
		}(CurrencyDate)
		CurrencyDate = CurrencyDate.AddDate(0, 0, 1)
	}
	fmt.Println(durationDay)

	for durationDay > 0 {
		time.Sleep(10 * time.Millisecond)
		//println(durationDay)
	}

	elepsedTime := time.Since(startTime)
	fmt.Printf("Execution Time : %s", elepsedTime)
}

