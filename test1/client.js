var fs = require('fs');
var https = require('https');

var get = https.request({
  path: '/', hostname: 'localhost', port: 7252,
  ca: fs.readFileSync('cert/root.pem'),
  agent: false,
  rejectUnauthorized: true,
}, function(response) {
  response.on('data', (d) => {
    process.stdout.write(d);
  });
});

get.on('error', function(e) {
  console.error(e)
  console.error("error", e)
  console.error("error", JSON.stringify(e))
});

get.end();
