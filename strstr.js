const fs = require("fs");

const lines = [];
const stream = fs.createReadStream('tmp/logs.log');

stream.on('data', data => {
    data
      .toString()
      .trim()
      .split('\n')
      .filter(line => (line.indexOf('OK') > -1))
      .map(line => line.split('OK ')[1])
      .forEach(line => {
        lines.push(line);
      });
});

stream.on('close', () => {
  console.log(lines.slice(0, 40));
});
