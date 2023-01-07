'use strict';
import { Express } from "express";
const express = require('express');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app:Express = express();
app.get('/', (req, res): void => {
  res.send('Hello World');
});

app.listen(PORT, HOST, () => {
  console.log(`Running on http://${HOST}:${PORT}`);
}); 