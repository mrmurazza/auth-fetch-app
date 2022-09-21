function notFound(req, res, next) {
    res.status(404);
    const error = new Error("Not Found", req.originalUrl);
    next(error);
}

function errorHandler(err, req, res, next) {
    res.status(res.statusCode || 500);
    var result = {
        message: err.message,
    };

    const debug = process.env.DEBUG_MODE === "true";
    console.log(`debug mode ${debug}`);
    if (debug) {
        result.stack = err.stack;
    }

    res.json(result);
}

module.exports = {
    notFound,
    errorHandler,
};
