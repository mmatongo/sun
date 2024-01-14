### USAGE
Obtain an API key from https://visualcrossing.com/

The args should be supplied in the order <location> -key=<api key>, anything other than that will break the app
```bash
go run main.go lusaka -key=YOUR_API_KEY
```

### OUTPUT
```bash
Lusaka, Zambia, Africa/Lusaka: 18°C, Rain, Overcast
00:00 - 18°C, Rain, Overcast
01:00 - 18°C, Rain, Overcast
02:00 - 18°C, Rain, Overcast
03:00 - 18°C, Rain, Overcast
04:00 - 18°C, Rain, Overcast
05:00 - 19°C, Rain, Overcast
06:00 - 19°C, Rain, Overcast
07:00 - 21°C, Rain, Partially cloudy
08:00 - 21°C, Rain, Partially cloudy
09:00 - 22°C, Rain, Partially cloudy
10:00 - 22°C, Rain, Partially cloudy
11:00 - 22°C, Rain, Overcast
12:00 - 22°C, Rain, Overcast
13:00 - 22°C, Rain, Overcast
14:00 - 22°C, Rain, Overcast
15:00 - 21°C, Rain, Overcast
16:00 - 20°C, Rain, Overcast
17:00 - 19°C, Rain, Overcast
18:00 - 18°C, Rain, Overcast
19:00 - 19°C, Rain, Overcast
20:00 - 19°C, Rain, Overcast
21:00 - 19°C, Rain, Overcast
```

The weather output will be for the current day up to 24 hours.
If a location is not supplied it will default to Bath for the city, the app will panic and crash if an API key is not supplied.

### Caveats
- The order of the args is important, this can be overcome by using flags
- Usage of Time.Now() means that if you try to look up the weather for another location it will base results ooff of your current time.
- Output is not the neatest.
- Code is a little messy
