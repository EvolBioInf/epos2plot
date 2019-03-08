set terminal epslatex color solid
set size 5/5., 4/3.
set format xy "\\Large$%g$"
set format y "\\Large $10^{%T}$"
set xlabel "\\Large Time ($\\times 10^3$ generations)"
set ylabel "\\Large$N (\\times 10^3)$"
#set format "\Large$%g$"
#set logscale xy
#set pointsize 2

set logscale xy
f(x)=10
set output "fig2a.tex"
plot [0.3:42][] "fig2a.dat" u ($1/1000):($2/1000) title "2.5\\% quantile" wi li lw 3,\
          "fig2a.dat" u ($1/1000):($3/1000) title "median" wi li lw 3,\
	  f(x) title "expected" wi li lw 3,\
          "fig2a.dat" u ($1/1000):($4/1000) title "97.5\\% quantile" wi li lw 3		
set output "chq2.tex"
plot [2:][10:10**5] "chq2.dat" u ($1/1000):($3/1000) title "" wi li lw 3
set output "chq3.tex"
plot [2:][10:10**5] "chq3.dat" u ($1/1000):($2/1000):($3/1000) title "" w yerrorlines lw 3
