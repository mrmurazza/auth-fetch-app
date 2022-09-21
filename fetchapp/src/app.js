const express = require("express");
const morgan = require("morgan");
const helmet = require("helmet");
const bodyParser = require("body-parser");

require("dotenv").config();

const app = express();

app.use(helmet());
app.use(morgan("dev"));
app.use(bodyParser.json());

const { notFound, errorHandler } = require("./middlewares");

const resources = require("./controllers/resourcesController");
const { initCache } = require("./services/currencyService");
app.use("/api/v1/resources", resources);

app.use(notFound);
app.use(errorHandler);

initCache();

module.exports = app;
