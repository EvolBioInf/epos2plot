# `Epos2plot` Converts `epos` Output to Plots
  
## Dependencies
`Epos2plot` is written in Go and requires a working
[Go environment](https://golang.org/).

## Installation
Download `epos2plot` and install it
```
go get -u  github.com/evolbioinf/epos2plot
go install github.com/evolbioinf/epos2plot
```

## Examples
The example data is based on Figure 2a of @liu15:exp: The population size
is 10,000 and 10 sequences of length 10Mbp were simulated with 1000
replications using the coalescent simulator
[`mspms`](https://pypi.org/project/msprime/) [@kel16:eff]. Its output was
converted to [site frequency spectra](https://en.wikipedia.org/wiki/Allele_frequency_spectrum) using
[`sfs`](http://github.com/evolbioinf/sfs), and the site frequency
spectra in turn were converted to population sizes using
[`epos`](http://github.com/evolbioinf/epos): 
```
for i in $(seq 1000); do
	mspms 10 1 -t 12310 -r 9750 10000000 -eN 0.066 0.3 | 
	sfs -f |
	epos -u 1.2e-1
done > fig2a_10.epos
```
The resulting output is located at
```
$GOPATH/src/github.com/evolbioinf/epos2plot/data/fig2a_10.epos.bz2
```
and can be displayed in "raw" format using the program [`pipePlot`](http://github.com/evolbioinf/pipeplot):
```
bzcat $GOPATH/src/github.com/evolbioinf/epos2plot/data/fig2a_10.epos.bz2 |
epos2plot -r | 
pipePlot
```
Instead of plotting raw `epos` results, 
`epos2plot` by default computes 2.5% and 97.5% quantiles around the median: 
```
bzcat $GOPATH/src/github.com/evolbioinf/epos2plot/data/fig2a_10.epos.bz2 |
epos2plot | 
head
#Time	LowerQ	Median	UpperQ	SampleSize
0	9550	10200	18700	986
463	9550	10200	18700	986
482	9550	10200	18700	986
487	9550	10200	18700	986
491	9550	10200	18700	986
502	9550	10200	18700	986
505	9550	10200	18700	986
506	9550	10200	18700	986
510	9540	10200	18700	986
```
where `Time` contains the time in generations, 
`LowerQ` the lower quantile of the population size, `Median` its
median, and `UpperQ` its upper
quantile. Notice how these values are close to the expected
10,000. The column `SampleSize`, finally, lists the number of data points from which the
quantiles were computed. As one last example, plot the median:
```
bzcat $GOPATH/src/github.com/evolbioinf/epos2plot/data/fig2a_10.epos.bz2 |
epos2plot | 
cut -f 1,3 |
pipePlot
```
## References
