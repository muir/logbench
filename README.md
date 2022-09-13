# Logger benchmarks.

See the most recent [action](https://github.com/muir/logbench/actions) 
for the most recent results.

### Disable

When the log level is too low for the line to be output

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     33.21 ns/op |          0 B/op |     0 allocs/op |
|             Zap |     106.8 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |     13.07 ns/op |          0 B/op |     0 allocs/op |
|         ZeroLog |     9.129 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     2.344 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     33.48 ns/op |          0 B/op |     0 allocs/op |
|         PhusLog |     9.481 ns/op |          0 B/op |     0 allocs/op |

### Normal

One line with three attributes

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|        XopMilli |     506.0 ns/op |          0 B/op |     0 allocs/op |
|         Xop3339 |     662.8 ns/op |          0 B/op |     0 allocs/op |
|             Zap |      1051 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |      1328 ns/op |        384 B/op |     1 allocs/op |
|         ZeroLog |     664.0 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     314.9 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     338.4 ns/op |          0 B/op |     0 allocs/op |
|  OneLogWithTime |     764.8 ns/op |         24 B/op |     1 allocs/op |
|         PhusLog |     353.6 ns/op |          0 B/op |     0 allocs/op |

### Interface

A line with a generic object that has three fields.  Note: Onelog can only serialize simple objects.

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     669.5 ns/op |          0 B/op |     0 allocs/op |
|             Zap |      1437 ns/op |        144 B/op |     2 allocs/op |
|        ZapSugar |      1532 ns/op |        208 B/op |     2 allocs/op |
|         ZeroLog |     940.2 ns/op |         48 B/op |     1 allocs/op |
|   OneLogNTChain |     386.1 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     691.7 ns/op |          0 B/op |     0 allocs/op |

### Printf

A line generated with printf

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     877.8 ns/op |         96 B/op |     2 allocs/op |
|        ZapSugar |      1408 ns/op |         96 B/op |     2 allocs/op |
|         ZeroLog |      1098 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     715.9 ns/op |         16 B/op |     1 allocs/op |

### Caller

A with three attributes and one stack frame

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      3550 ns/op |        208 B/op |     1 allocs/op |
|             Zap |      3228 ns/op |        432 B/op |     3 allocs/op |
|         ZeroLog |      3153 ns/op |        288 B/op |     4 allocs/op |
|         PhusLog |      1131 ns/op |        216 B/op |     2 allocs/op |

### Empty

The overhead of starting a trace/span for a new request

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      9034 ns/op |       8811 B/op |    21 allocs/op |
|            OTEL |      1692 ns/op |        967 B/op |     6 allocs/op |

### Tenspan

The overhead of starting a trace/span for a new request AND ten sub-spans

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |    139887 ns/op |     104209 B/op |   243 allocs/op |
|            OTEL |     19966 ns/op |      10730 B/op |    72 allocs/op |
