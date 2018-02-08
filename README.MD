## Le bon Fizzbuzz (jeu de mot, aucune pretention :-)

Je me suis un peu emporté et avait débuté une version en C mais au final cela met un certain temps donc bon.
Je voulais aussi faire un Dockerfile mais me suis dit autant prendre contact dans un premier temps.

## Dependencies
golang : No need, just Go \p/
nodejs:  npm install
python:  pip install -r requirements.txt

### Launch server (all on port 5000 so not at the same time)
golang : go run golang/main.go
nodejs:  npm start
python: FLASK_APP=python/api.py flask run

## Api Routes:

### GET /fizzbuzz/:str1/:str2/:nbr1/:nbr2/:limit 

Get the result of a fizzbuzz generation when **str1** is *fizz*, **str2** is *buzz*, **nbr1** and **nbr2** are the two dividends. **Limit** is the number to generate. Limit is cap with `const maxLimit = 10000` in golang because it's an array. Evolution would be to change it to the equivalent of a linked list in C.

