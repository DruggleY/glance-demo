package probe

type Function uint32

const FunctionPing Function = 0x01
const FunctionTraceroute Function = 0x02
const FunctionShowBGP Function = 0x04

type Continent uint32

const ContinentAsia Continent = 0x01
const ContinentEurope Continent = 0x02
const ContinentNorthAmerica Continent = 0x04
const ContinentSouthAmerica Continent = 0x08
const ContinentAfrica Continent = 0x10
const ContinentOceania Continent = 0x20
const ContinentAntarctica Continent = 0x40
