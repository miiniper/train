# train
get 12306 train list info
## usage
```
Usage:
  train [flags]
  train [command]

Available Commands:
  help        Help about any command
  list        get train list
  station     Print Chinese to shorthand
  version     Print the version number of Demo

Flags:
  -h, --help   help for train

eg:

train list -t BJP -f SHH -d 2020-10-10
┌─────────┬───────────┬──────────┬───────────────┬────────────┬─────────────┬───────────┬───────────┬──────────┐
│ TrainNo │ StartTime │ Date     │ BusinessClass │ FirstClass │ SecondClass │ SoftBerth │ HardBerth │ HardSeat │
├─────────┼───────────┼──────────┼───────────────┼────────────┼─────────────┼───────────┼───────────┼──────────┤
│ D706    │ 21:20     │ 20201010 │               │            │ 有          │ 有        │ 有        │          │
│ D710    │ 21:24     │ 20201010 │               │            │ 有          │ 有        │ 有        │          │
└─────────┴───────────┴──────────┴───────────────┴────────────┴─────────────┴───────────┴───────────┴──────────┘

train station -s 上海
SHH 

```


## todo
- train info
- other