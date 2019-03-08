./cmd/plotSum data/chq.dat > tmp.out
DIFF=$(diff tmp.out data/chq.out)
if [ "$DIFF" == "" ] 
then
    printf "Test(plotSum)\t\t\tpass\n"
else
    printf "Test(plotSum)\t\t\tfail\n"
    echo ${DIFF}
fi

./cmd/plotSum -e data/chq.dat > tmp.out
DIFF=$(diff tmp.out data/chqE.out)
if [ "$DIFF" == "" ] 
then
    printf "Test(plotSum -e)\t\tpass\n"
else
    printf "Test(plotSum -e)\t\tfail\n"
    echo ${DIFF}
fi

rm tmp.out
