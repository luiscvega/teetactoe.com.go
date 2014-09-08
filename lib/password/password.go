package password

import (
	"crypto"

	"code.google.com/p/go.crypto/bcrypt"
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

func Crypt(password sring) {
    var salt = randomSalt();

    return digest(password, salt) + salt;
}

func randomSalt() string {
    return crypto.randomBytes(32).toString("hex");
}

func digest(password string, salt string) string {
    return crypto.pbkdf2Sync(
        password, salt,
        ITERATIONS, KEYLEN
    ).toString("hex");
};
