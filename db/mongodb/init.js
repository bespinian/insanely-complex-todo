disableTelemetry()
db.setProfilingLevel(2)

db = db.getSiblingDB("admin");
db.createUser({
    user: "monitoring",
    pwd: "monitoring",
    roles: [
        { role: "clusterMonitor", db: "admin" },
        { role: "read", db: "local" },
        { role: "read", db: "ict" }
    ]
});
