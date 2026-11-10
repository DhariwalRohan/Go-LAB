// module Program

// open System

// // Base class
// [<AbstractClass>]
// type Shape() =
//     abstract member Area: unit -> float
//     default this.Area() = 0.0

// // Derived class 1
// type Circle(radius: float) =
//     inherit Shape()

//     override this.Area() =
//         Math.PI * radius * radius

// // Derived class 2
// type Rectangle(length: float, width: float) =
//     inherit Shape()

//     override this.Area() =
//         length * width

// [<EntryPoint>]
// let main argv =

//     let circle = Circle(5.0)
//     let rectangle = Rectangle(10.0, 4.0)

//     let shapes: Shape list = [circle; rectangle]

//     printfn "OOP in F#"
//     printfn "----------"

//     for shape in shapes do
//         printfn "Area = %.2f" (shape.Area())

//     0