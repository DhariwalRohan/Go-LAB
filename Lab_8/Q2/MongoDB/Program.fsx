// #r "nuget: MongoDB.Driver, 2.30.0"

// open MongoDB.Bson
// open MongoDB.Driver

// // Connect to MongoDB
// let client = MongoClient("mongodb://localhost:27017")
// let database = client.GetDatabase("StudentDB")
// let collection = database.GetCollection<BsonDocument>("Students")

// // Create Student
// let student = BsonDocument()
// student.Add("Name", "Rohan")
// student.Add("Age", 20)

// collection.InsertOne(student)

// let filter = BsonDocument("_id", student.["_id"])

// // Display Students
// let showStudents title =
//     printfn "\n%s" title
//     let students = collection.Find(BsonDocument()).ToList()

//     if students.Count = 0 then
//         printfn "No students found."
//     else
//         students |> Seq.iter (fun s -> printfn "%s" (s.ToJson()))

// // CREATE
// showStudents "After CREATE"

// // READ
// showStudents "After READ"

// // UPDATE
// let update =
//     Builders<BsonDocument>.Update.Set("Age", 21)

// collection.UpdateOne(filter, update) |> ignore
// showStudents "After UPDATE"

// // DELETE
// collection.DeleteOne(filter) |> ignore
// showStudents "After DELETE"