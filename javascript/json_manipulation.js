let input = `{
  "incidents": [
    {
      "id": "P200",
      "title": "Critical memory leak detected",
      "severity": "critical",
      "created_at": "2025-06-09T08:00:00Z",
      "escalations": [
        {
          "id": "P200-1",
          "title": "Memory consumption over 90%",
          "severity": "high",
          "created_at": "2025-06-09T08:10:00Z",
          "escalations": [
            {
              "id": "P200-1-1",
              "title": "Container restart loop",
              "severity": "critical",
              "created_at": "2025-06-09T08:20:00Z"
            }
          ]
        }
      ]
    },
    {
      "id": "P201",
      "title": "Service degraded - slow response",
      "severity": "medium",
      "created_at": "2025-06-09T09:30:00Z"
    },
    {
      "id": "P202",
      "title": "Disk space warning",
      "severity": "low",
      "created_at": "2025-06-09T07:45:00Z",
      "escalations": []
    },
    {
      "id": "P203",
      "title": "Failed login attempts spike",
      "severity": "high",
      "created_at": "2025-06-09T06:00:00Z",
      "escalations": [
        {
          "id": "P203-1",
          "title": "IP block list updated",
          "severity": "medium",
          "created_at": "2025-06-09T06:15:00Z",
          "escalations": [
            {
              "id": "P203-1-1",
              "title": "Firewall rule applied",
              "severity": "high",
              "created_at": "2025-06-09T06:25:00Z"
            },
            {
              "id": "P203-1-2",
              "title": "Alert sent to security team",
              "severity": "medium",
              "created_at": "2025-06-09T06:30:00Z"
            }
          ]
        }
      ]
    },
    {
      "id": "P204",
      "title": "API endpoint failure",
      "severity": "critical",
      "created_at": "2025-06-09T10:00:00Z",
      "escalations": [
        {
          "id": "P204-1",
          "title": "Timeouts increasing",
          "severity": "high",
          "created_at": "2025-06-09T10:05:00Z",
          "escalations": [
            {
              "id": "P204-1-1",
              "title": "Backend service lagging",
              "severity": "medium",
              "created_at": "2025-06-09T10:10:00Z"
            }
          ]
        }
      ]
    },
    {
      "id": "P205",
      "title": "Configuration drift detected",
      "severity": "medium",
      "created_at": "2025-06-09T11:30:00Z",
      "escalations": null
    }
  ]
}
`

let result = []

function getIncidentData(input, depth = 0) {

    let requireData = {
        "id": input.id,
        "title": input.title,
        "severity": input.severity,
        "created_at": input.created_at,
        "depth": depth,
    }

    result.push(requireData)

    depth = depth + 1

    if (input.escalations && input.escalations.length > 0) {
        for (let i of input.escalations) {
            getIncidentData(i, depth)
        }
    }

}

function getHours(timestamp) {
    let difference = new Date() - new Date(timestamp)
    let diffHrs = difference / (60 * 60 * 1000)
    return `${diffHrs.toFixed(1)} hours ago`

}

let inputData = JSON.parse(input)

for (let incident of inputData.incidents) {
    getIncidentData(incident, 0)
}

let severityOrder = {
    "critical": 4,
    'high': 3,
    'medium': 2,
    'low': 1
}

result.sort((a, b) => {
    if (severityOrder[b.severity] == severityOrder[a.severity]) {
        return new Date(severityOrder[a.created_at]) - new Date(severityOrder[b.created_at])
    }
    return severityOrder[b.severity] - severityOrder[a.severity]
})

for (let r of result) {
    let indent = "  ".repeat(r.depth)
    console.log(`${indent}[${r.severity.toUpperCase()}] -- ${r.id} - ${r.title} (${getHours(r.created_at)})`)
}