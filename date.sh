git log |
    grep Date |
    head -n 1 |
    awk '{printf "\"%s_%s_%s_%s,_%s\"", $2, $3, $4, $6, $5}'
