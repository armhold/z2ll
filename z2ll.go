package zip2latlong

import (
	"encoding/csv"
	"fmt"
	"github.com/gobuffalo/packr"
	"io"
	"log"
	"strconv"
)

type ZipCode struct {
	Code string // "90210"
	City string // "Beverly Hills"
	Lat  float64
	Long float64
}

type Locator struct {
	m map[string]*ZipCode
}

func NewLocator() (*Locator, error) {
	l := &Locator{m: make(map[string]*ZipCode)}

	return l, l.buildMapFromCSV()
}

func (l *Locator) buildMapFromCSV() error {
	box := packr.NewBox("./assets")

	f, err := box.Open("free-zipcode-database.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)

	for i := 0; true; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// skip header
		if i == 0 {
			continue
		}
		// RecordNumber","Zipcode","ZipCodeType","City","State","LocationType","Lat","Long","Xaxis","Yaxis","Zaxis","WorldRegion","Country","LocationText","Location","Decommisioned","TaxReturnsFiled","EstimatedPopulation","TotalWages","Notes"
		recordNumber, zipCode, _, city, state, _, lat, long, _, _, _, worldRegion, country, locationText, location, _, _, _, _, _ := record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[17], record[18], record[19]

		// for debugging
		if false {
			fmt.Println(i, recordNumber, zipCode, city, state, lat, long, worldRegion, country, locationText, location)
			fmt.Printf("i: %d\n", i)
		}

		// TODO: make these nullable?
		if lat == "" {
			lat = "0.0"
		}

		if long == "" {
			long = "0.0"
		}

		//fmt.Printf("parsing: %s, %s, %s\n", zipCode, lat, long)
		l.m[zipCode] = &ZipCode{Code: zipCode, City: city, Lat: float64FromString(lat), Long: float64FromString(long)}
	}

	return nil
}

func float64FromString(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}

	return float64(value)
}

func (l *Locator) LookupZip(zip string) (ZipCode, error) {
	if zc, ok := l.m[zip]; ok {
		return *zc, nil
	}

	return ZipCode{}, fmt.Errorf("zip code not found: %s", zip)
}
