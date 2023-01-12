'use strict';
import { Express } from "express";
const express = require('express');
var cors = require('cors');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app:Express = express();
app.use(cors());
app.get('/', (req, res): void => {
  res.send('Hello World');
});

app.listen(PORT, HOST, () => {
  console.log(`Running on http://${HOST}:${PORT}`);
}); 