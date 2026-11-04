// module Program

// open System
// open System.Text.RegularExpressions

// // Task 1: Reverse string and check palindrome
// let reverseString (str: string) =
//     new string (str.ToCharArray() |> Array.rev)

// let isPalindrome (str: string) =
//     str.Equals(reverseString str, StringComparison.OrdinalIgnoreCase)

// // Task 2: Transform string
// let transformString (str: string) =
//     if str.Length > 0 && str |> Seq.forall Char.IsLetter then
//         str.ToUpper()
//     elif str.Length > 0 && str |> Seq.forall Char.IsDigit then
//         let number = int str
//         string (number * 2)
//     else
//         str

// // Task 3: Validate email
// let isValidEmail (email: string) =
//     Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

// [<EntryPoint>]
// let main argv =

//     printfn "TASK 1: Reverse String and Palindrome"

//     let strings = ["madam"; "hello"; "level"]

//     strings |> List.iter (fun str ->
//         printfn "String: %s | Reverse: %s | Palindrome: %b"
//             str (reverseString str) (isPalindrome str))

//     printfn "\nTASK 2: Transform String"

//     let values = ["hello"; "123"; "abc123"]

//     values |> List.iter (fun str ->
//         printfn "Input: %s | Output: %s" str (transformString str))

//     printfn "\nTASK 3: Email Validation"

//     let emails = [
//         "abc@gmail.com"
//         "rohan123@yahoo.com"
//         "invalid-email"
//     ]

//     emails |> List.iter (fun email ->
//         printfn "Email: %s | Valid: %b" email (isValidEmail email))

//     0