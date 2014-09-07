package password

import (
	"crypto"
)

const (
	ITERATIONS = 5000
	KEYLEN     = 64
)

func Check(password string, hash string) {
    var sha512 = hash.substring(0, 128);
    var salt = hash.substring(128);

    return digest(password, salt) === sha512;
};

exports.crypt = function crypt(password) {
    var salt = randomSalt();

    return digest(password, salt) + salt;
}

function randomSalt() {
    return crypto.randomBytes(32).toString("hex");
}

function digest(password, salt) {
    return crypto.pbkdf2Sync(
        password, salt,
        ITERATIONS, KEYLEN
    ).toString("hex");
};

