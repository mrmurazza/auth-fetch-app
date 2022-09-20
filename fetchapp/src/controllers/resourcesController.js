const express = require('express');

const router = express.Router();

router.get('/', async (req, res, next) => {
  try {
    const data = [
        "test",
        "test2"
    ];
    res.json(data);
  } catch (error) {
    next(error);
  }
});

module.exports = router;