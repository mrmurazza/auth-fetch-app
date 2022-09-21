const express = require("express");
const getCommodities = require("../clients/commodities");
const getCurrencyConversion = require("../clients/currency");
const getResources = require("../services/resourcesService");

const router = express.Router();

router.get("/", async (req, res, next) => {
    try {
        const data2 = await getResources();
        // const data3 = await getCurrencyConversion(1, "IDR", "USD");
        // const data = ["test", "test2"];
        res.json(data2);
    } catch (error) {
        next(error);
    }
});

router.get("/aggregate", async (req, res, next) => {
    try {
        const data2 = await getCommodities();
        const data = ["test", "test2"];
        res.json(data2);
    } catch (error) {
        next(error);
    }
});

module.exports = router;
