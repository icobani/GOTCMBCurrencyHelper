/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 15/12/2016    
* Time      : 18:22
* Developer : ibrahimcobani
*
*******/
package GOTCMBCurrencyHelper

import (
	"net/http"
	"encoding/xml"
	"time"
	"io"
	"bytes"
	"strings"
	"strconv"
	"log"
)

type CurrencyJournal struct {
	Id         string `json:"_id" bson:"_id"`
	Date       time.Time
	JournalNo  string
	Currencies []Currency
}

type Currency struct {
	Code            string
	CrossOrder      int
	Unit            int
	CurrencyNameTR  string
	CurrencyName    string
	ForexBuying     float64
	ForexSelling    float64
	BanknoteBuying  float64
	BanknoteSelling float64
	CrossRateUSD    float64
	CrossRateOther  float64
}

func GetArchive(CurrencyDate time.Time) CurrencyJournal {
	numberOfTrial := 0
	MaxnumberOfTrial := 30
	ghostDate := CurrencyDate
	t := new(tarih_Date)
	cj := t.getArchive(CurrencyDate, ghostDate)
	for {
		numberOfTrial +=1
		if (cj.Id == "") {
			CurrencyDate = CurrencyDate.AddDate(0, 0, -1)
			cj = t.getArchive(CurrencyDate, ghostDate)
			if (cj.Id != "") {
				break
			}
		} else {
			break
		}
		current_time := time.Now().Local()
		if (numberOfTrial > MaxnumberOfTrial) || (CurrencyDate.After(current_time) ){
			break
		}
	}
	return cj
}

type tarih_Date struct {
	XMLName   xml.Name   `xml:"Tarih_Date"`
	Tarih     string     `xml:"Tarih,attr"`
	Date      string     `xml:"Date,attr"`
	Bulten_No string     `xml:"Bulten_No,attr"`
	Currency  []currency `xml:"Currency"`
}

type currency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
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

func (c *tarih_Date) getArchive(CurrencyDate time.Time, GhostDate time.Time) CurrencyJournal {
	cj := CurrencyJournal{}
	var resp *http.Response
	var err error
	var url string
	url = "http://www.tcmb.gov.tr/kurlar/" + CurrencyDate.Format("200601") + "/" + CurrencyDate.Format("02012006") + ".xml"
	log.Println(url)
	resp, err = http.Get(url)

	if err != nil {
		log.Println(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			tarih := new(tarih_Date)
			d := xml.NewDecoder(resp.Body)
			d.CharsetReader = CharsetReader
			marshalErr := d.Decode(&tarih)
			if marshalErr != nil {
				log.Printf("error: %v", marshalErr)
			}

			c = &tarih_Date{}
			cj.Id = GhostDate.Format("20060102")
			cj.Date = GhostDate
			cj.JournalNo = tarih.Bulten_No
			cj.Currencies = make([]Currency, len(tarih.Currency))
			for i, curr := range tarih.Currency {
				cj.Currencies[i].Code = curr.CurrencyCode
				cj.Currencies[i].CurrencyName = curr.CurrencyName
				cj.Currencies[i].CurrencyNameTR = curr.Isim
				cj.Currencies[i].BanknoteBuying, err = strconv.ParseFloat(curr.BanknoteBuying, 64)
				cj.Currencies[i].BanknoteSelling, err = strconv.ParseFloat(curr.BanknoteSelling, 64)
				cj.Currencies[i].ForexBuying, err = strconv.ParseFloat(curr.ForexBuying, 64)
				cj.Currencies[i].ForexSelling, err = strconv.ParseFloat(curr.ForexSelling, 64)
				cj.Currencies[i].CrossOrder, err = strconv.Atoi(curr.CrossOrder)
				cj.Currencies[i].CrossRateOther, err = strconv.ParseFloat(curr.CrossRateOther, 64)
				cj.Currencies[i].CrossRateUSD, err = strconv.ParseFloat(curr.CrossRateUSD, 64)
				cj.Currencies[i].Unit, err = strconv.Atoi(curr.Unit)
			}

		} else {
			cj = CurrencyJournal{}
		}

	}
	return cj
}
