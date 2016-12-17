/*
 * B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
 * Created by Ibrahim.COBANI on 15-12-16.
 */

package GOTCMBCurrencyHelper

import (
	"testing"
	"time"
	"fmt"
)

func Test(t *testing.T) {
	currencyJournal := new(CurrencyJournal)

	CurrencyDate, _ := time.Parse("02-01-2006", "16-04-1996")
	for CurrencyDate.Before(time.Now()) {
		c := currencyJournal.GetArchive(CurrencyDate)
		fmt.Println(CurrencyDate, c.Currencies[1].ForexBuying)
		CurrencyDate = CurrencyDate.AddDate(0, 0, 1)
	}

	//fmt.Println(CurrencyDate)
	//for d := CurrencyDate; d < time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC); d = d.AddDate(0, 0, 1) {
	//	fmt.Println(d)
	//c := currencyJournal.GetArchive(CurrencyDate)
	//fmt.Println(d, c.Currencies[1].ForexBuying)
	//}

}

