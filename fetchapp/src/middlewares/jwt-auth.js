const jwt = require("jsonwebtoken");

const ROLE_USER = "USER";
const ROLE_ADMIN = "ADMIN";

function authenticate(req, res, next) {
    const authorizationHeader = req.get("Authorization");

    if (!authorizationHeader.includes("Bearer")) {
        res.status(401);
        next(new Error("Invalid Token"));
    }

    tokenString = authorizationHeader.replace("Bearer ", "");

    try {
        const jwtSecret = process.env.JWT_SECRET_KEY || "";
        var decoded = jwt.verify(tokenString, jwtSecret);
        req.userInfo = decoded.data;
        next();
    } catch (err) {
        res.status(401);
        next(err);
    }
}

function authenticateAdmin(req, res, next) {
    const userInfo = req.userInfo;

    if (userInfo.role != ROLE_ADMIN) {
        res.status(403);
        next(new Error("You are forbidden to access this"));
    }

    next();
}

module.exports = {
    authenticate,
    authenticateAdmin,
};
