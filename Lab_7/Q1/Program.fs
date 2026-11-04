// module Program

// open System

// // Function to create asynchronous tasks
// let task (name: string) (delay: int) =
//     async {
//         printfn "%s started" name

//         do! Async.Sleep delay

//         printfn "%s finished" name

//         return sprintf "%s completed successfully" name
//     }

// // Main function
// let main argv =

//     printfn "Starting Parallel Workflows...\n"

//     // Three asynchronous tasks
//     let tasks =
//         [|
//             task "Task 1" 1000
//             task "Task 2" 2000
//             task "Task 3" 1500
//         |]

//     // Execute all tasks concurrently
//     let results =
//         tasks
//         |> Async.Parallel
//         |> Async.RunSynchronously

//     // Print all results
//     printfn "\nAll Tasks Completed!"
//     printfn "\nFinal Results:"

//     results |> Array.iter (printfn "%s")

//     0