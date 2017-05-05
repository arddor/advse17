var Twitter = require('twitter');


var client = new Twitter({
  consumer_key: '1I7ttD6IeJfe762CNUtCRf03v',
  consumer_secret: '0jg36zGi1Sp0WukyLP0TCnQ7R6ggTJtP9tSyZkdtkqeMUn8wmY',
  access_token_key: '532988829-ijL3hxk8V2mDdxsRA2LWPJKaDz0QBrthZRp9vJSB',
  access_token_secret: 'jGpcyfoSrQhtadS6fBVUljFUcmrlBl9iGwAvQ5eusLcyX'
});


// var client = new Twitter({
//   consumer_key: process.env.TWITTER_CONSUMER_KEY,
//   consumer_secret: process.env.TWITTER_CONSUMER_SECRET,
//   access_token_key: process.env.TWITTER_ACCESS_TOKEN_KEY,
//   access_token_secret: process.env.TWITTER_ACCESS_TOKEN_SECRET
// });

// the idea is to mine a couple of tweets into a container over night or so
// package that data into a container
// the container should grab that data and push some requests
// maybe that would be interesting to do from digital ocean using a swarm
// so manually increase the containers as we go

var stream = client.stream('statuses/filter', {track: 'javascript'});
stream.on('data', function(event) {
  console.log(event && event.text);
});

stream.on('error', function(error) {
  throw error;
});

client.stream('statuses/filter', {track: 'javascript'}, function(stream) {
  stream.on('data', function(event) {
    console.log(event && event.text);
  });

  stream.on('error', function(error) {
    throw error;
  });
});


// _ = require('lodash')
// const isTweet = _.conforms({
//   contributors: _.isObject,
//   id_str: _.isString,
//   text: _.isString,
// })
