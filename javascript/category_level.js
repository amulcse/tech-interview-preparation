let links = [
    "s3://my-bucket/logs/2024/10/01/app.log",
    "s3://my-bucket/logs/2024/10/02/app.log",
    "s3://my-bucket/data/users/user1.json",
    "s3://my-bucket/data/users/user2.json",
    "s3://my-bucket/data/transactions/txn1.json",
    "s3://my-bucket/data/transactions/txn2.json",
    "s3://my-bucket/logs/2024/09/30/app.log"
]

let results = new Map()

for (let link of links) {
    let linkArray = link.split("s3://")[1].split("/")

    let c = results

    for (let l of linkArray) {
        if (!c.has(l)) {
            c.set(l, new Map())
        }
        c = c.get(l)
    }
}

function displayResults(results, depth) {

    if (results === undefined) {
        return
    }

    results.forEach((value, key) => {
        let indent = ("  ").repeat(depth) + "|--"
        console.log(`${indent} ${key}`)
        displayResults(value, depth + 1)
    });

}

displayResults(results, 1)