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

	CurrencyDate, _ := time.Parse("02-01-2006", "28-04-2016")
	//CurrencyDate = time.Now().AddDate(0, 0, -2)
	c := currencyJournal.GetArchive(CurrencyDate)
	fmt.Println("==>",c)
}

