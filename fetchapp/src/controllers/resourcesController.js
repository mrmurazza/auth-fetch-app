const express = require("express");
const { authenticateAdmin } = require("../middlewares/jwt-auth");
const {
    getResources,
    aggregateResources,
} = require("../services/resourcesService");

const router = express.Router();

router.get("/", async (req, res, next) => {
    try {
        const data = await getResources();
        res.json(data);
    } catch (error) {
        next(error);
    }
});

router.use("/aggregate", authenticateAdmin);
router.get("/aggregate", async (req, res, next) => {
    try {
        const data = await aggregateResources();
        res.json({
            data: data,
        });
    } catch (error) {
        next(error);
    }
});

module.exports = router;
