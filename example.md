Example data after subscribing to orderBookL2_25 filtered by XBTUSD

2020/09/06 21:36:00 connecting to wss://testnet.bitmex.com/realtime
2020/09/06 21:36:00 recv: {"info":"Welcome to the BitMEX Realtime API.","version":"2020-09-01T23:32:47.000Z","timestamp":"2020-09-06T19:36:00.078Z","docs":"https://testnet.bitmex.com/app/wsAPI","limit":{"remaining":35}}
2020/09/06 21:36:00 recv: {"success":true,"subscribe":"orderBookL2_25:XBTUSD","request":{"op":"subscribe","args":["orderBookL2_25:XBTUSD"]}}
2020/09/06 21:36:02 recv: {"table":"orderBookL2_25","action":"partial","keys":["symbol","id","side"],"types":{"symbol":"symbol","id":"long","side":"symbol","size":"long","price":"float"},"foreignKeys":{"symbol":"instrument","side":"side"},"attributes":{"symbol":"parted","id":"sorted"},"filter":{"symbol":"XBTUSD"},"data":[{"symbol":"XBTUSD","id":15598974300,"side":"Sell","size":130,"price":10257},{"symbol":"XBTUSD","id":15598974400,"side":"Sell","size":130,"price":10256},{"symbol":"XBTUSD","id":15598974500,"side":"Sell","size":130,"price":10255},{"symbol":"XBTUSD","id":15598974600,"side":"Sell","size":130,"price":10254},{"symbol":"XBTUSD","id":15598974650,"side":"Sell","size":2,"price":10253.5},{"symbol":"XBTUSD","id":15598974700,"side":"Sell","size":130,"price":10253},{"symbol":"XBTUSD","id":15598974800,"side":"Sell","size":132,"price":10252},{"symbol":"XBTUSD","id":15598974900,"side":"Sell","size":130,"price":10251},{"symbol":"XBTUSD","id":15598974950,"side":"Sell","size":20,"price":10250.5},{"symbol":"XBTUSD","id":15598975000,"side":"Sell","size":730,"price":10250},{"symbol":"XBTUSD","id":15598975050,"side":"Sell","size":150,"price":10249.5},{"symbol":"XBTUSD","id":15598975100,"side":"Sell","size":130,"price":10249},{"symbol":"XBTUSD","id":15598975200,"side":"Sell","size":130,"price":10248},{"symbol":"XBTUSD","id":15598975300,"side":"Sell","size":76,"price":10247},{"symbol":"XBTUSD","id":15598975400,"side":"Sell","size":234,"price":10246},{"symbol":"XBTUSD","id":15598975450,"side":"Sell","size":34447,"price":10245.5},{"symbol":"XBTUSD","id":15598975500,"side":"Sell","size":130,"price":10245},{"symbol":"XBTUSD","id":15598975600,"side":"Sell","size":130,"price":10244},{"symbol":"XBTUSD","id":15598975650,"side":"Sell","size":103,"price":10243.5},{"symbol":"XBTUSD","id":15598975700,"side":"Sell","size":130,"price":10243},{"symbol":"XBTUSD","id":15598975750,"side":"Sell","size":265,"price":10242.5},{"symbol":"XBTUSD","id":15598975800,"side":"Sell","size":130,"price":10242},{"symbol":"XBTUSD","id":15598975900,"side":"Sell","size":130,"price":10241},{"symbol":"XBTUSD","id":15598976000,"side":"Sell","size":130,"price":10240},{"symbol":"XBTUSD","id":15598976050,"side":"Sell","size":73120,"price":10239.5},{"symbol":"XBTUSD","id":15598976100,"side":"Buy","size":67592,"price":10239},{"symbol":"XBTUSD","id":15598976200,"side":"Buy","size":130,"price":10238},{"symbol":"XBTUSD","id":15598976300,"side":"Buy","size":130,"price":10237},{"symbol":"XBTUSD","id":15598976400,"side":"Buy","size":154,"price":10236},{"symbol":"XBTUSD","id":15598976450,"side":"Buy","size":23,"price":10235.5},{"symbol":"XBTUSD","id":15598976500,"side":"Buy","size":130,"price":10235},{"symbol":"XBTUSD","id":15598976600,"side":"Buy","size":130,"price":10234},{"symbol":"XBTUSD","id":15598976700,"side":"Buy","size":130,"price":10233},{"symbol":"XBTUSD","id":15598976800,"side":"Buy","size":130,"price":10232},{"symbol":"XBTUSD","id":15598976900,"side":"Buy","size":130,"price":10231},{"symbol":"XBTUSD","id":15598977000,"side":"Buy","size":130,"price":10230},{"symbol":"XBTUSD","id":15598977050,"side":"Buy","size":600,"price":10229.5},{"symbol":"XBTUSD","id":15598977100,"side":"Buy","size":280,"price":10229},{"symbol":"XBTUSD","id":15598977150,"side":"Buy","size":103,"price":10228.5},{"symbol":"XBTUSD","id":15598977200,"side":"Buy","size":130,"price":10228},{"symbol":"XBTUSD","id":15598977250,"side":"Buy","size":1141,"price":10227.5},{"symbol":"XBTUSD","id":15598977300,"side":"Buy","size":190,"price":10227},{"symbol":"XBTUSD","id":15598977400,"side":"Buy","size":130,"price":10226},{"symbol":"XBTUSD","id":15598977500,"side":"Buy","size":130,"price":10225},{"symbol":"XBTUSD","id":15598977600,"side":"Buy","size":130,"price":10224},{"symbol":"XBTUSD","id":15598977700,"side":"Buy","size":130,"price":10223},{"symbol":"XBTUSD","id":15598977800,"side":"Buy","size":130,"price":10222},{"symbol":"XBTUSD","id":15598977850,"side":"Buy","size":1650,"price":10221.5},{"symbol":"XBTUSD","id":15598977900,"side":"Buy","size":130,"price":10221},{"symbol":"XBTUSD","id":15598978000,"side":"Buy","size":130,"price":10220}]}
2020/09/06 21:36:05 recv: {"table":"orderBookL2_25","action":"update","data":[{"symbol":"XBTUSD","id":15598976050,"side":"Sell","size":73233}]}