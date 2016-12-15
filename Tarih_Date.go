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
	"encoding/xml"
	"fmt"
	"time"
	"io"
	"bytes"
	"strings"
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


//*********************

type CharsetISO88591er struct {
	r   io.ByteReader
	buf *bytes.Buffer
}

func NewCharsetISO88591(r io.Reader) *CharsetISO88591er {
	buf := bytes.Buffer{}
	return &CharsetISO88591er{r.(io.ByteReader), &buf}
}

func (cs *CharsetISO88591er) Read(p []byte) (n int, err error) {
	for _ = range p {
		if r, err := cs.r.ReadByte(); err != nil {
			break
		} else {
			cs.buf.WriteRune(rune(r))
		}
	}
	return cs.buf.Read(p)
}

func isCharset(charset string, names []string) bool {
	charset = strings.ToLower(charset)
	for _, n := range names {
		if charset == strings.ToLower(n) {
			return true
		}
	}
	return false
}

func IsCharsetISO88591(charset string) bool {
	// http://www.iana.org/assignments/character-sets
	// (last updated 2010-11-04)
	names := []string{
		// Name
		"ISO_8859-1:1987",
		// Alias (preferred MIME name)
		"ISO-8859-1",
		// Aliases
		"iso-ir-100",
		"ISO_8859-1",
		"ISO-8859-9",
		"latin1",
		"l1",
		"IBM819",
		"CP819",
		"csISOLatin1",
	}
	return isCharset(charset, names)
}

func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if IsCharsetISO88591(charset) {
		return NewCharsetISO88591(input), nil
	}
	return input, nil
}

//********************


func (c *Tarih_Date) GetToday(CurrencyDate time.Time) {

	for {
		url := "http://www.tcmb.gov.tr/kurlar/" + CurrencyDate.Format("200601") + "/" + CurrencyDate.Format("02012006") + ".xml"
		println(url)
		resp, err := http.Get(url)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {

			//body, err2 := ioutil.ReadAll(resp.Body)
			//println(err2 != nil)
			//if err2 != nil {
			//	fmt.Println("=>", err2)
			//}
			tarih := new(Tarih_Date)
			//fmt.Println(&body)

			d := xml.NewDecoder(resp.Body)
			d.CharsetReader = CharsetReader
			//marshalErr := xml.Unmarshal(body, &tarih)
			marshalErr := d.Decode(&tarih)
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