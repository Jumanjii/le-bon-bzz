package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const maxLimit = 10000

func generateFBList(fizz string, buzz string, nbr1 int, nbr2 int, limit int) (jsonList []byte) {
	var result [maxLimit]string
	var i = 1

	for ; i <= limit; i++ {
		isFizz := i%nbr1 == 0
		isBuzz := i%nbr2 == 0

		if isFizz && isBuzz {
			result[i-1] = fizz + buzz // I read concat like that is not optimal but it's my current level in Go for this test. Will be happy to learn another way ! :-)
		} else if isFizz {
			result[i-1] = fizz
		} else if isBuzz {
			result[i-1] = buzz
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	jsonList, _ = json.Marshal(result[0 : i-1])
	return
}

func parseIntParameters(nbr1 string, nbr2 string, limit string) (int1 int, int2 int, intLimit int, err error) {
	int1, err1 := strconv.Atoi(nbr1)
	int2, err2 := strconv.Atoi(nbr2)
	intLimit, err3 := strconv.Atoi(limit)

	// Really uggly :-s
	// TODO : An error handler that can take a list of error
	err = nil
	if err1 != nil {
		err = err1
	} else if err2 != nil {
		err = err2
	} else if err3 != nil {
		err = err3
	}

	return
}

func getFizzbuzz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fizz := vars["str1"]
	buzz := vars["str2"]
	nbr1, nbr2, limit, err := parseIntParameters(vars["nbr1"], vars["nbr2"], vars["limit"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "A parameter should be a number and isn't: %v , %v, %v \n", vars["nbr1"], vars["nbr2"], vars["limit"])
		return
	} else if limit > maxLimit {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Limit excedding the maximum authorized. Max limit : %d\n", maxLimit)
		return
	}

	message := generateFBList(fizz, buzz, nbr1, nbr2, limit)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(message))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{str1}/{str2}/{nbr1}/{nbr2}/{limit}", getFizzbuzz)

	fmt.Printf("Launching server on port :5000\n")
	log.Fatal(http.ListenAndServe(":5000", r))
}
