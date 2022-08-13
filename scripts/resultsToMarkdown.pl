#!/usr/bin/perl 

my $lastgroup;

while (<>) {
	next unless /allocs/;
	next unless s/^Benchmark//;
	die unless m/^([A-Z][a-z]+)/;
	my $group = $1;
	s/^([A-Z][a-z]+)//;
	if ($group ne $lastgroup) {
		print "\n";
		print "### $group\n";
		print "\n";
		print "| Logger          | Rate            | Memory total    | Allocations |\n";
		print "| --------------- | --------------- | --------------- | --------------- |\n";
		$lastgroup = $group
	}
	chomp;
	my @cols = split(/\t/, $_);
	$cols[0] =~ s/-\d+\s*$//;
	foreach (@cols) {
		s/\s\s+/ /g;
		chomp;
		s/^\s//;
		$_ = ( " "x (15 - length($_) )) . $_
	}
	print "| $cols[0] | $cols[2] | $cols[3] | $cols[4] |\n";
}

__END__
BenchmarkDisableXop
BenchmarkDisableXop-4             	372288250	        32.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisableZap
BenchmarkDisableZap-4             	100000000	       100.5 ns/op	     192 B/op	       1 allocs/op
BenchmarkDisableZeroLog
BenchmarkDisableZeroLog-4         	1000000000	        10.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisableOneLog
BenchmarkDisableOneLog-4          	1000000000	         2.009 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisableOneLogChain
BenchmarkDisableOneLogChain-4     	357913698	        33.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisablePhusLog
BenchmarkDisablePhusLog-4         	1000000000	        10.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalXop
BenchmarkNormalXop-4              	11590264	       999.2 ns/op	    1280 B/op	       4 allocs/op
BenchmarkNormalZap
BenchmarkNormalZap-4              	11354103	      1052 ns/op	     192 B/op	       1 allocs/op
BenchmarkNormalZeroLog
BenchmarkNormalZeroLog-4          	17905940	       670.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOneLog
BenchmarkNormalOneLog-4           	38576966	       311.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOneLogChain
BenchmarkNormalOneLogChain-4      	35965369	       333.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalPhusLog
BenchmarkNormalPhusLog-4          	33730822	       355.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceXop
BenchmarkInterfaceXop-4           	10070061	      1204 ns/op	    1280 B/op	       4 allocs/op
BenchmarkInterfaceZap
BenchmarkInterfaceZap-4           	 7828550	      1509 ns/op	     208 B/op	       2 allocs/op
BenchmarkInterfaceZeroLog
BenchmarkInterfaceZeroLog-4       	12636904	       954.4 ns/op	      48 B/op	       1 allocs/op
BenchmarkInterfaceOneLogChain
BenchmarkInterfaceOneLogChain-4   	30336339	       390.6 ns/op	      96 B/op	       2 allocs/op
BenchmarkInterfacePhusLog
BenchmarkInterfacePhusLog-4       	17456988	       684.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrintfXop
BenchmarkPrintfXop-4              	 8703260	      1382 ns/op	    1376 B/op	       6 allocs/op
BenchmarkPrintfZap
BenchmarkPrintfZap-4              	 8569634	      1404 ns/op	      96 B/op	       2 allocs/op
BenchmarkPrintfZeroLog
BenchmarkPrintfZeroLog-4          	10997137	      1103 ns/op	      96 B/op	       2 allocs/op
BenchmarkPrintfPhusLog
BenchmarkPrintfPhusLog-4          	16999431	       709.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkCallerXop
BenchmarkCallerXop-4              	 3164554	      3742 ns/op	    1488 B/op	       5 allocs/op
BenchmarkCallerZap
BenchmarkCallerZap-4              	 3686038	      3245 ns/op	     432 B/op	       3 allocs/op
BenchmarkCallerZeroLog
BenchmarkCallerZeroLog-4          	 3793507	      3158 ns/op	     288 B/op	       4 allocs/op
BenchmarkCallerPhusLog
BenchmarkCallerPhusLog-4          	10533097	      1141 ns/op	     216 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	317.455s

