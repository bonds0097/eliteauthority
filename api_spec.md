/api/commodities | GET | Returns all commodities with basic data (galactic avg., rare?)
/api/commodities | POST | Creates new commodity.
/api/commodities/{:slug}/search | POST | Searches for commodity with certain parameters (max buy price, min sell price, distance from station, number of results)
/api/snapshot | POST | Creates new commodity snapshot.
/api/stations | POST | Creates new station.
/api/stations | GET | Returns all stations with basic data.
/api/stations/search | POST | Searches for stations with certain parameters (distance from system, number of results, type)
/api/stations/{:slug} | PUT | Updates station.
/api/stations/{:slug} | GET | Returns data for specific station.