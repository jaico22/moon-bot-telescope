# Moon Bot - Telescope
## Overview
_Telescope_ is a utility that actively observes and records the current price of Dogecoin using the Ticker endpoint provided by Kraken
## Installation and Setup
Everything you need to run telescope locally should be accessable in the Make file. 
`make build` builds the binary
`make setup` setups an alias for localhost and then runs docker-compose up to setup all dependent services
`make run` invokes the 'ScrapeData' lambda 
## Todos
- Implement data storage
## Stretch Goals
- Integrate data w/ google trends 
