const express = require('express');
const os = require('os');
const app = express();
const port = 3000;

app.get('/', (req, res) => {
  const hostname = os.hostname();
  const timestamp = new Date().toISOString();
  const networkInterfaces = os.networkInterfaces();
  const ips = Object.values(networkInterfaces)
    .flat()
    .map((iface) => iface.address);

  res.json({
    'Data': {
      'Hostname': hostname,
      'Timestamp': timestamp,
      'IP Addresses': ips
    },
  });

});

app.listen(port, () => console.log(`App running on port ${port}`));
