TAG     := $(shell git describe | sed 's/-.*//')
DATE    := $(shell git log -1 --format=%ai $(TAG) | sed 's/-/ /g' | awk '{printf "\\\\DTMdisplaydate{%s}{%s}{%s}{-1}\n", $$1, $$2, $$3}')
VERSION := $(shell git describe | sed 's/^[vV]//; s/-.*//')
all:
	echo $(VERSION) > version.tex
	echo $(DATE)    > date.tex
	gnuplot fig.gp
	latex epos2plot
	bibtex epos2plot
	latex epos2plot
	latex epos2plot
	dvipdf epos2plot
clean:	
	rm -f epos2plot.pdf *.log *.bbl *.blg *.dvi
