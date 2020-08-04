# Moon Bot - Telescope
## Overview
_Telescope_ is a utility that actively observes and records the current price of Dogecoin using the Ticker endpoint provided by Kraken
## Installation and Setup
Everything you need to run telescope locally should be accessable in the Make file. 

`make build` builds the binaries

`make setup` setups an alias for localhost and then runs docker-compose up to setup all dependent services

`make run` will setup a local lambda enviorment that can be invoked with either the aws cli or via postman

To test data scraping, send a POST to `http://127.0.0.1:3001/2015-03-31/functions/ScrapeData/invocations`

### Viewing Data
An API is available to view data collected from the scraping tool.

To test viewing data, send a POST to `http://127.0.0.1:3001/2015-03-31/functions/ViewData/invocations`
with the following body:
```
{
	"StartDate": "{start time stamp : string}",
	"EndDate": "{end time stamp : string}",
	"Version": 2
}
```

The response body is as follows
```
{
    "RequestCopy": {
	"StartDate": "{start time stamp : string}",
	"EndDate": "{end time stamp : string}",
	"Version": 2
    },
    "Records": [
        {
            "DateTime": "{time the record was recorded : string}",
            "AskingPrice": {asking price : decimal},
            "BiddingPrice": {bidding price : decimal}
        }
    ],
    "NumberOfRecords": {Count of items in "Record" : int}
}
```

## Todos
## Stretch Goals
- Integrate data w/ google trends 
