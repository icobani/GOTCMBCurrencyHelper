# GOTCMBCurrencyHelper
Currency Helper for TCMB Currency Service


  #####go-get github.com/icobani/GOTCMBCurrencyHelper.git



    import (
      "github.com/gin-gonic/gin"
      "net/http"
      "time"
      "github.com/icobani/GOTCMBCurrencyHelper"
      "github.com/icobani/b1ws.com/config"
      "github.com/icobani/b1ws.com/models"
      "strconv"
      "log"
    )
    func main() {
      currencyJournal := GOTCMBCurrencyHelper.GetArchive(CurrencyDate)
  
      for _, curr := range currencyJournal.Currencies {
  
        config.DB.Where("currency_year = ? AND currency_month = ? AND currency_day = ? AND code = ?",
          CurrencyYear, CurrencyMonth, CurrencyDay, curr.Code).First(&form)
        if config.DB.NewRecord(form) {
          form.CurrencyYear = CurrencyYear
          form.CurrencyMonth = CurrencyMonth
          form.CurrencyDay = CurrencyDay
          form.Code = curr.Code
          form.BanknoteBuying = curr.BanknoteBuying
          form.BanknoteSelling = curr.BanknoteSelling
          form.ForexBuying = curr.ForexBuying
          form.ForexSelling = curr.ForexSelling
          form.CrossRateUSD = curr.CrossRateUSD
          form.CrossRateOther = curr.CrossRateOther
          form.Unit = curr.Unit
          form.CallTimes += 1
          config.DB.Create(&form)
          form = models.ExchangeRate{}
      }
    }