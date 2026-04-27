# My Go Learning Journey

Welcome to my Go Language learning repository! This repository tracks my progress, projects, and the core concepts I've learned throughout the course.

## What I Learned By Section

### Section 2: Go Fundamentals & Basic Projects
In this section, I laid the foundation for programming in Go by building practical, small-scale applications.
**Key Learnings:**
- **Variables & Data Types:** How to declare and use basic data types (strings, integers, floats, booleans).
- **Control Flow:** Using `if/else` statements for logic and `for` loops for iteration.
- **Functions:** Writing reusable blocks of code and understanding return values.
- **I/O Operations:** Taking user input via the console and formatting output using the `fmt` package.
- **File Handling & Errors:** Reading from and writing to local text files, and handling basic errors gracefully.
*Projects built: Investment Calculator, Profit Calculator, and a rudimentary Banking App.*

### Section 3: Packages & Modules
This section taught me how to transition from writing single-file scripts to building structured applications.
**Key Learnings:**
- **Go Modules:** Initializing projects using `go mod init` to manage dependencies.
- **Code Organization:** Splitting logic across multiple files and custom packages to keep code clean and maintainable.
- **Exporting/Importing:** Understanding Go's visibility rules (capitalizing variables and function names to make them public/exported across packages).
*Projects built: A modularized System application demonstrating package imports.*

### Section 4: Pointers
This section was a deep dive into how Go handles memory, which is crucial for writing efficient applications.
**Key Learnings:**
- **Memory Addresses:** Understanding how variables are stored in memory.
- **Pointer Operators:** Using the `&` operator to find a variable's memory address, and the `*` operator to read or modify the value at that address (dereferencing).
- **Pass by Value vs. Pass by Reference:** Learning why and how to pass pointers to functions so that the original variable can be modified without creating unnecessary copies in memory.

### Section 5: Structs and Custom Types
In this section, I learned how to group related data together to model real-world concepts.
**Key Learnings:**
- **Defining Structs:** Creating custom, complex data types with multiple fields.
- **Initialization:** Creating instances of structs using struct literals and field names.
- **Structs and Pointers:** Passing structs to functions via pointers to prevent copying large amounts of data and to allow functions to mutate the struct's fields directly.
- **User Input to Structs:** Dynamically populating struct fields based on user input.
*Projects built: User Information Manager.*