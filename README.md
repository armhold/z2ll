# zip2latlong

Quick command line tool for looking up lat,long for a given zip code.

Zipcodes database courtesy of [federalgovernmentzipcodes.us](http://federalgovernmentzipcodes.us)

## Installation

Installing a stand-alone binary relies on [packr](https://github.com/gobuffalo/packr).

packr is used to generate the `a_zip2latlong-packr.go` source from the  
`assets/free-zipcode-database.csv` file, and thus bundle it into the binary.

```bash
go get -u https://github.com/gobuffalo/packr
```

```bash
git clone https://github.com/armhold/zip2latlong
cd zip2latlong
packr && go install ./cmd/z2ll
```

### Usage

```bash
z2ll 44120
41.470000,-81.670000
```
