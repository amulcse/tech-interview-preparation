// Advanced Password Strength Validator
// Write a function isStrongPassword(password string) bool that returns true if the given password is strong according to all the criteria, and false otherwise.
// A password is considered strong if it meets the following criteria:
// It has at least 6 characters and at most 20 characters.
// It does not contain any sequence of characters that can be found in a common list of passwords (e.g., “password”, “123456”, “qwerty”).
// It is not a palindrome (i.e., it does not read the same forwards and backwards).
// It does not contain any sequence of characters that forms a recognizable pattern (e.g., “123456”, “abcdef”).
// It contains at least one lowercase letter, one uppercase letter, one digit, and one special character (punctuation or symbol).
// It does not contain three or more repeating characters in a row.
// It does not have more than two consecutive lowercase letters, two consecutive uppercase letters, or two consecutive digits.

function getCharType(c) {
    if (c >= 'a' && c <= 'z') {
        return 'lowercase'
    } else if (c >= 'A' && c <= 'Z') {
        return 'uppercase'
    } else if (c >= '0' && c <= '9') {
        return 'digit'
    } else {
        return 'special_char'
    }
}

function isStrongPassword(password) {

    let response = {
        "isValid": false,
        "error": null
    }

    let characterCount = password.length
    if (characterCount < 6 || characterCount > 20) {
        response.error = "character count does not match"
        return response
    }

    let commonPasswordPatterns = ["password", "123456", "qwerty"]
    if (commonPasswordPatterns.includes(password.toLowerCase())) {
        response.error = "it has known common pattern"
        return response
    }

    let reverseString = ((password.split("")).reverse()).join("")
    if (password == reverseString) {
        response.error = "palindrome present"
        return response
    }

    if (/[a-z]{3,}/.test(password)) {
        console.log("more than three consecutive lower case")
    }

    if (/[A-Z]{3,}/.test(password)) {
        console.log("more than three uppercase character")
    }

    let lastSceneCharType = ''
    let characterTypeWiseCount = {}
    for (let c of password) {

        let charType = getCharType(c)
        if (charType == lastSceneCharType) {
            response.error = "Character repeat " + charType + " " + lastSceneCharType
            return response
        }

        if (!characterTypeWiseCount[charType]) {
            characterTypeWiseCount[charType] = 0
        }
        characterTypeWiseCount[charType]++


        lastSceneCharType = charType
    }

    if (Object.keys(characterTypeWiseCount).length < 4) {
        response.error = "It does not contain all kind of characters"
        return response
    }

    response.isValid = true
    return response;

}


let password = "zzzaadddaaa";

console.log("sort password", password.split("").sort().join(""))

console.log(`${password} is valid? - ${JSON.stringify(isStrongPassword(password))}`)