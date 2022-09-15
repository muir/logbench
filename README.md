# Logger benchmarks.

See the most recent [action](https://github.com/muir/logbench/actions) 
for the most recent results.


### Disable

When the log level is too low for the line to be output

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     29.45 ns/op |          0 B/op |     0 allocs/op |
|             Zap |     101.1 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |     11.53 ns/op |          0 B/op |     0 allocs/op |
|         ZeroLog |     8.466 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     2.070 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     29.58 ns/op |          0 B/op |     0 allocs/op |
|         PhusLog |     8.356 ns/op |          0 B/op |     0 allocs/op |

### Normal

One line with three attributes

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|        XopMilli |     439.8 ns/op |          0 B/op |     0 allocs/op |
|         Xop3339 |     592.8 ns/op |          0 B/op |     0 allocs/op |
|             Zap |     996.3 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |      1270 ns/op |        384 B/op |     1 allocs/op |
|         ZeroLog |     654.4 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     277.2 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     337.2 ns/op |          0 B/op |     0 allocs/op |
|  OneLogWithTime |     725.1 ns/op |         24 B/op |     1 allocs/op |
|         PhusLog |     338.7 ns/op |          0 B/op |     0 allocs/op |

### Interface

A line with a generic object that has three fields.  Note: Onelog can only serialize simple objects.

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     587.5 ns/op |          0 B/op |     0 allocs/op |
|             Zap |      1356 ns/op |        144 B/op |     2 allocs/op |
|        ZapSugar |      1439 ns/op |        208 B/op |     2 allocs/op |
|         ZeroLog |     939.2 ns/op |         48 B/op |     1 allocs/op |
|   OneLogNTChain |     369.8 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     610.7 ns/op |          0 B/op |     0 allocs/op |

### Printf

A line generated with printf

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     839.7 ns/op |         96 B/op |     2 allocs/op |
|        ZapSugar |      1324 ns/op |         96 B/op |     2 allocs/op |
|         ZeroLog |      1040 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     674.2 ns/op |         16 B/op |     1 allocs/op |

### Caller

A with three attributes and one stack frame

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      3473 ns/op |        208 B/op |     1 allocs/op |
|             Zap |      3100 ns/op |        432 B/op |     3 allocs/op |
|         ZeroLog |      3006 ns/op |        288 B/op |     4 allocs/op |
|         PhusLog |      1069 ns/op |        216 B/op |     2 allocs/op |

### Empty

The overhead of starting a trace/span for a new request

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      8852 ns/op |       8811 B/op |    21 allocs/op |
|      OTELStdout |     35475 ns/op |       3710 B/op |    67 allocs/op |

### Tenspan

The overhead of starting a trace/span for a new request AND ten sub-spans

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |    155591 ns/op |     104225 B/op |   243 allocs/op |
|      OTELStdout |    230185 ns/op |      36067 B/op |   637 allocs/op |
