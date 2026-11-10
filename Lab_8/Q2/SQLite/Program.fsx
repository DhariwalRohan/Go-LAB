// #r "nuget: Microsoft.Data.Sqlite, 8.0.0"


// open Microsoft.Data.Sqlite


// use connection =
//     new SqliteConnection("Data Source=students.db")
// connection.Open()
// // Create Table
// let createTable =
//     new SqliteCommand(
//         """CREATE TABLE IF NOT EXISTS Students
//            (Id INTEGER PRIMARY KEY, Name TEXT, Age INTEGER)""",
//         connection
//     )
// createTable.ExecuteNonQuery() |> ignore
// // Display Students
// let showStudents title =
//     printfn "\n%s" title
//     let cmd =
//         new SqliteCommand("SELECT * FROM Students", connection)
//     use reader = cmd.ExecuteReader()
//     let mutable found = false
//     while reader.Read() do
//         found <- true
//         printfn "ID: %d | Name: %s | Age: %d"
//             (reader.GetInt32(0))
//             (reader.GetString(1))
//             (reader.GetInt32(2))
//     if not found then
//         printfn "No students found."
// // Avoid Duplicate Record
// let cleanup =
//     new SqliteCommand(
//         "DELETE FROM Students WHERE Id = 1",
//         connection
//     )
// cleanup.ExecuteNonQuery() |> ignore
// // CREATE
// let insert =
//     new SqliteCommand(
//         "INSERT INTO Students VALUES (1, 'Rohan', 20)",
//         connection
//     )
// insert.ExecuteNonQuery() |> ignore
// showStudents "After CREATE"
// // READ
// showStudents "After READ"
// // UPDATE
// let update =
//     new SqliteCommand(
//         "UPDATE Students SET Age = 21 WHERE Id = 1",
//         connection
//     )
// update.ExecuteNonQuery() |> ignore
// showStudents "After UPDATE"
// // DELETE
// let delete =
//     new SqliteCommand(
//         "DELETE FROM Students WHERE Id = 1",
//         connection
//     )
// delete.ExecuteNonQuery() |> ignore
// showStudents "After DELETE"
