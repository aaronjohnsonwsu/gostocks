package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    "time"
    )



func main() {
    data := getSingleCompanyData(searchObj{"orcl",time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),"d"})
    fmt.Print(data)
}

// Name: searchObj
// Summary: Search object contains the parameters of a Yahoo Finance API GET call

type searchObj struct {
    symbol string
    fromDate time.Time
    toDate time.Time
    dataType string
}


// Name: buildQueryString
// Parameters: Search Object (Type searchObj)
// Returns: Query String (Type string)
// Summary: Takes a searchObj and builds a query string to make a REST call to the Yahoo Finance API
// --------------------------------------------------------------------------------------------------
//     Here is the API query string information:
// --------------------------------------------------------------------------------------------------
//     sn	Ticker symbol (YHOO in the example)
//     a	The "from month" - 1
//     b	The "from day" (two digits)
//     c	The "from year"
//     d	The "to month" - 1
//     e	The "to day" (two digits)
//     f	The "to year"
//     g	d for day, m for month, y for yearly     

func buildYahooFinanceDataQueryString(search searchObj )(string){
    var query string = "http://table.finance.yahoo.com/table.csv?"
    query += "s=" + search.symbol
    query += "&a=" + strconv.Itoa(int(search.fromDate.Month()) - 1)
    query += "&b=" + strconv.Itoa(int(search.fromDate.Day()))
    query += "&c=" + strconv.Itoa(int(search.fromDate.Year()))
    query += "&d=" + strconv.Itoa(int(search.toDate.Month()) - 1)
    query += "&e=" + strconv.Itoa(int(search.toDate.Day()))
    query += "&f=" + strconv.Itoa(int(search.toDate.Year()))
    query += "&g=" + search.dataType
    
    return query
}


// Name: getYahooFinanceData
// Parameters: query string (Type string)
// Returns: response data from the REST call (Type string)
// Summary: Takes a query and makes a REST call and returns the data from the response.

func getYahooFinanceData(query string) (string){
    resp, err := http.Get(query)
    if err != nil {
        //Error has occured
        return string(err.Error())
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    
    return string(data)
}


// Name: getSingleCompanyData
// Parameters: search critera for a single company (Type searchObj)
// Returns: response data from the REST call (Type string)
// Summary: Takes a search criteria then builds a query string then makes a REST call the yahoo finanace API

func getSingleCompanyData(search searchObj) (string){
    searchQuery := buildYahooFinanceDataQueryString(search)
    return getYahooFinanceData(searchQuery)
}


// Name: getAllCompanySymbols
// Parameters: none
// Returns: array of symbols (Type []string)
// Summary: Retrieves a list of company symbols and returns the list in an array of strings

func getAllCompanySymbols() ([]string){
    return []string{}
}