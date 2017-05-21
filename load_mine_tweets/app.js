// npm install twitter
var Twitter = require('twitter');
// npm install mongodb
var MongoClient = require('mongodb').MongoClient;

var client = new Twitter({
  consumer_key: '1I7ttD6IeJfe762CNUtCRf03v',
  consumer_secret: '0jg36zGi1Sp0WukyLP0TCnQ7R6ggTJtP9tSyZkdtkqeMUn8wmY',
  access_token_key: '532988829-ijL3hxk8V2mDdxsRA2LWPJKaDz0QBrthZRp9vJSB',
  access_token_secret: 'jGpcyfoSrQhtadS6fBVUljFUcmrlBl9iGwAvQ5eusLcyX'
});



// the idea is to mine a couple of tweets into a container over night or so
// package that data into a container
// the container should grab that data and push some requests
// maybe that would be interesting to do from digital ocean using a swarm
// so manually increase the containers as we go

var url = 'mongodb://localhost:27017/test';
MongoClient.connect(url, function(err, db) {
  console.log("Connected correctly to server");
  
  var collection = db.collection('trump_tweets');
  
  var stream = client.stream('statuses/filter', {track: 'trump'});
  
  stream.on('data', function(event) {
    //console.log(event.text);
    
    collection.insertOne({text : event.text });  
  });
  
});









