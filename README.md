# Logger benchmarks.

See the most recent [action](https://github.com/muir/logbench/actions) 
for the most recent results.


### Disable

When the log level is too low for the line to be output

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     34.62 ns/op |          0 B/op |     0 allocs/op |
|             Zap |     128.3 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |     13.77 ns/op |          0 B/op |     0 allocs/op |
|         ZeroLog |     11.80 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     2.816 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     38.47 ns/op |          0 B/op |     0 allocs/op |
|         PhusLog |     12.90 ns/op |          0 B/op |     0 allocs/op |

### Normal

One line with three attributes

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|        XopMilli |     568.1 ns/op |          0 B/op |     0 allocs/op |
|         Xop3339 |     735.0 ns/op |          0 B/op |     0 allocs/op |
|             Zap |      1182 ns/op |        192 B/op |     1 allocs/op |
|        ZapSugar |      1476 ns/op |        384 B/op |     1 allocs/op |
|         ZeroLog |     741.5 ns/op |          0 B/op |     0 allocs/op |
|    OneLogNoTime |     370.2 ns/op |          0 B/op |     0 allocs/op |
|   OneLogNTChain |     380.8 ns/op |          0 B/op |     0 allocs/op |
|  OneLogWithTime |     857.0 ns/op |         24 B/op |     1 allocs/op |
|         PhusLog |     405.3 ns/op |          0 B/op |     0 allocs/op |

### Interface

A line with a generic object that has three fields.  Note: Onelog can only serialize simple objects.

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     760.3 ns/op |          0 B/op |     0 allocs/op |
|             Zap |      1566 ns/op |        144 B/op |     2 allocs/op |
|        ZapSugar |      1653 ns/op |        208 B/op |     2 allocs/op |
|         ZeroLog |      1043 ns/op |         48 B/op |     1 allocs/op |
|   OneLogNTChain |     439.6 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     798.2 ns/op |          0 B/op |     0 allocs/op |

### Printf

A line generated with printf

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     972.1 ns/op |         96 B/op |     2 allocs/op |
|        ZapSugar |      1563 ns/op |         96 B/op |     2 allocs/op |
|         ZeroLog |      1261 ns/op |         96 B/op |     2 allocs/op |
|         PhusLog |     850.2 ns/op |         16 B/op |     1 allocs/op |

### Caller

A with three attributes and one stack frame

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |      4801 ns/op |        208 B/op |     1 allocs/op |
|             Zap |      3573 ns/op |        432 B/op |     3 allocs/op |
|         ZeroLog |      3427 ns/op |        288 B/op |     4 allocs/op |
|         PhusLog |      1291 ns/op |        216 B/op |     2 allocs/op |

### Empty

The overhead of starting a trace/span for a new request

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |     10688 ns/op |       8859 B/op |    22 allocs/op |
|      OTELStdout |     44837 ns/op |       3710 B/op |    67 allocs/op |

### Tenspan

The overhead of starting a trace/span for a new request AND ten sub-spans

| Logger          | Rate            | Memory total    | Allocations     |
| --------------- | --------------- | --------------- | --------------- |
|             Xop |    129257 ns/op |     103455 B/op |   203 allocs/op |
|      OTELStdout |    277775 ns/op |      36070 B/op |   637 allocs/op |
