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
}

