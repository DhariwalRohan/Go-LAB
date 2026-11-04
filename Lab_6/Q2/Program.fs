// module Program

// open System

// let isPrime n =
//     if n < 2 then
//         false
//     else
//         let rec check divisor =
//             if divisor * divisor > n then
//                 true
//             elif n % divisor = 0 then
//                 false
//             else
//                 check (divisor + 1)

//         check 2

// [<EntryPoint>]
// let main argv =
//     printfn "Enter a number:"
//     let n1 = int (Console.ReadLine())
//     printfn "Is %d prime? %b" n1 (isPrime n1)

//     printfn "\nEnter another number:"
//     let n2 = int (Console.ReadLine())
//     printfn "Is %d prime? %b" n2 (isPrime n2)

//     printfn "\nEnter one more number:"
//     let n3 = int (Console.ReadLine())
//     printfn "Is %d prime? %b" n3 (isPrime n3)

//     0