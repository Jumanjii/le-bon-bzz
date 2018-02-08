const restify = require('restify');

// Server
const server = restify.createServer();
server.listen(5000, function() {
  console.log('%s listening at %s', server.name, server.url);
});

// Lib utils

/**
 * FizzBuzz Generator
 * 
 * @param {String} fizz  String to use as Fizz in FizzBuzz 
 * @param {String} buzz  String to use as Buzz in FizzBuzz
 * @param {Int}    nbr1  Dividend link to Fizz  
 * @param {Int}    nbr2  Dividend link to Buzz
 * @param {Int}    limit Limit of the list
 * 
 * @returns {Array[String]} List generated
 */
const generateFBList = (fizz, buzz, nbr1, nbr2, limit) => {
    const result = [];
    
    for (let i = 1, msg = "";i <= limit; i++, msg = "") {

        msg += i % nbr1 == 0 ? fizz : ""
        msg += i % nbr2 == 0 ? buzz : ""
        
        result.push(msg || `${i}`)
    }

    return result;
}

//Controller

/**
 * Controller of the endpoint (can be move in a seperate controller file if evolution)
 * 
 * @param {*} req  Request
 * @param {*} res  Response
 * @param {*} next Next function
 */
function getFizzbuzz(req, res, next) {
    const result = generateFBList(req.params.str1, req.params.str2, req.params.nbr1, req.params.nbr2, req.params.limit);
    res.send(result);
    next();
}

// Routes
server.get('/fizzbuzz/:str1/:str2/:nbr1/:nbr2/:limit', 
            checkMaxLimitMiddleware,
            getFizzbuzz);