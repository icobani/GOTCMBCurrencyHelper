/*
 * B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
 * Created by Ibrahim.COBANI on 15-12-16.
 */

package GOTCMBCurrencyHelper

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	currencyHelper := new(Tarih_Date)

	CurrencyDate, _ := time.Parse("02-01-2006", "28-04-2016")
	//CurrencyDate = time.Now().AddDate(0, 0, -2)
	currencyHelper.GetToday(CurrencyDate)
}

