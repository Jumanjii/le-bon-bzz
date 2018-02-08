from flask import Flask
import json
app = Flask(__name__)

@app.route("/fizzbuzz/<fizz>/<buzz>/<nbr1>/<nbr2>/<limit>")
def fizzbuzz(fizz, buzz, nbr1, nbr2, limit):
    nbr1  = int(nbr1)
    nbr2  = int(nbr2)
    limit = int(limit)
    result = []

    for i in range(1, limit):
		isFizz = i % nbr1 == 0
		isBuzz = i % nbr2 == 0

		if (isFizz and isBuzz):
			result.append(fizz + buzz)
		elif isFizz:
			result.append(fizz)
		elif isBuzz: 
			result.append(buzz)
		else:
			result.append(str(i))
    
    return json.dumps(result)