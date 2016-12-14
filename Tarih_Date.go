/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 12/12/2016    
* Time      : 22:18
* Developer : ibrahimcobani
*
*******/
package GOTCMBCurrencyHelper

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
)

type Tarih_Date struct {
	XMLName   xml.Name `xml:"Tarih_Date"`
	Tarih     string `xml:"Tarih,attr"`
	Date      string `xml:"Date,attr"`
	Bulten_No string `xml:"Bulten_No,attr"`
	Currency  []Currency `xml:"Currency"`
}

type Currency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}

func (c *Tarih_Date) GetToday(CurrencyDate time.Time) {

	for {
		url := "http://www.tcmb.gov.tr/kurlar/" + CurrencyDate.Format("200601") + "/" + CurrencyDate.Format("02012006") + ".xml"
		println(url)
		resp, err := http.Get(url)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {

			body, err2 := ioutil.ReadAll(resp.Body)
			println(err2 != nil)
			if err2 != nil {
				fmt.Println("=>", err2)
			}
			tarih := new(Tarih_Date)



			marshalErr := xml.Unmarshal(body, &tarih)
			if marshalErr != nil {
				fmt.Printf("error: %v", marshalErr)
				return
			}
			c = &Tarih_Date{}
			fmt.Println(tarih.Date, tarih.Bulten_No)
			fmt.Println("------------------")
			for _, curr := range tarih.Currency {
				fmt.Println(curr.Kod, curr.Isim, curr.ForexBuying)
			}
			break
		} else {
			println(err)
			CurrencyDate = CurrencyDate.AddDate(0, 0, -1)
		}

	}
}