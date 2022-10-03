# Logger benchmarks.

See the most recent [action](https://github.com/muir/logbench/actions) 
for the most recent results.

### Disable

When the log level is too low for the line to be output

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     39.76 ns/op |          0 B/op |     0 allocs/op |
|             Zap |     125.7 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |     15.74 ns/op |          0 B/op |     0 allocs/op |
|         ZeroLog |     12.59 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     2.816 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     40.56 ns/op |          0 B/op |     0 allocs/op |
|         PhusLog |     12.07 ns/op |          0 B/op |     0 allocs/op |

### Normal

One line with three attributes

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      1070 ns/op |         32 B/op |     1 allocs/op |
|             Zap |      1299 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |      1693 ns/op |        384 B/op |     1 allocs/op |
|         ZeroLog |     794.4 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     394.0 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     405.1 ns/op |          0 B/op |     0 allocs/op |
|  OneLogWithTime |     951.2 ns/op |         24 B/op |     1 allocs/op |
|         PhusLog |     424.7 ns/op |          0 B/op |     0 allocs/op |

### Interface

A line with a generic object that has three fields.  Note: Onelog can only serialize simple objects.

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      1312 ns/op |         32 B/op |     1 allocs/op |
|             Zap |      1703 ns/op |        144 B/op |     2 allocs/op |
|        ZapSugar |      1844 ns/op |        208 B/op |     2 allocs/op |
|         ZeroLog |      1154 ns/op |         48 B/op |     1 allocs/op |
|   OneLogNTChain |     480.9 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     828.3 ns/op |          0 B/op |     0 allocs/op |

### Printf

A line generated with printf

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      1492 ns/op |        128 B/op |     3 allocs/op |
|        ZapSugar |      1735 ns/op |         96 B/op |     2 allocs/op |
|         ZeroLog |      1350 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     884.9 ns/op |         16 B/op |     1 allocs/op |

### Caller

A with three attributes and one stack frame

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      5773 ns/op |        240 B/op |     2 allocs/op |
|             Zap |      3981 ns/op |        432 B/op |     3 allocs/op |
|         ZeroLog |      3900 ns/op |        288 B/op |     4 allocs/op |
|         PhusLog |      1409 ns/op |        216 B/op |     2 allocs/op |

### Empty

The overhead of starting a trace/span for a new request

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     16446 ns/op |       7735 B/op |    22 allocs/op |
|      OTELStdout |     43915 ns/op |       3704 B/op |    67 allocs/op |
