rs.initiate(
    {
        _id : 'rs1',
        members: [
            { _id : 0, host : "mongo_one:27017"   },
            { _id : 1, host : "mongo_two:27017"   },
            { _id : 2, host : "mongo_three:27017" }
        ]
    }
)