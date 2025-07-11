const s3Permissions = [{
        path: "s3://company-bucket/hr/2024/salaries.xlsx",
        roles: ["hr", "admin"]
    },
    {
        path: "s3://company-bucket/hr/2023/salaries.xlsx",
        roles: ["hr"]
    },
    {
        path: "s3://company-bucket/hr/policies.pdf",
        roles: ["hr", "all"]
    },
    {
        path: "s3://company-bucket/engineering/designs/v1/ui.sketch",
        roles: ["eng"]
    },
    {
        path: "s3://company-bucket/engineering/designs/v2/ui.sketch",
        roles: ["eng", "admin"]
    },
    {
        path: "s3://company-bucket/engineering/internal/architecture.pdf",
        roles: ["eng"]
    },
    {
        path: "s3://company-bucket/legal/contracts/nda.docx",
        roles: ["legal", "admin"]
    },
]

let tree = new Map()

for (let permission of s3Permissions) {

    let {
        path,
        roles
    } = permission

    let pathSegments = path.split("s3://")[1].split("/")

    let currentNode = tree

    for (let segment of pathSegments) {

        if (!(currentNode.has(segment))) {
            currentNode.set(segment, {
                children: new Map(),
                roles: new Set()
            })
        }

        let c = currentNode.get(segment)
        roles.forEach(role => c.roles.add(role))

        currentNode = c.children

    }
}


function displayPathWithRoles(node, depth) {


    node.forEach((value, key) => {
        let indent = ("   ").repeat(depth) + "--"
        let role = [...value.roles].sort().join(", ")
        console.log(`${indent}${key} [${role}]`)
        displayPathWithRoles(value.children, depth + 1)
    });

}

displayPathWithRoles(tree, 1)