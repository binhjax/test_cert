const https = require('https');
const fs = require('fs');

const options = {
  key: fs.readFileSync('cert/leaf.key'),
  cert: fs.readFileSync('cert/leaf.pem'),
};

https.createServer(options, (req, res) => {
  res.writeHead(200);
  res.end('hello world\n');
}).listen(7252);
