# Train Tickets Application

A Go application that calculates **tax-inclusive train ticket prices**.  
It can take ticket prices either from a **file** or directly from the **command line**, apply different tax rates, assign random departure and destination cities, and output the results as JSON or print them to the console.

---

## Features

- Read ticket prices from:
  - a text file (`prices.txt`)  
  - interactive command-line input  
- Apply multiple tax rates (0%, 7%, 10%, 15%)  
- Randomly assign **start** and **destination** cities  
- Save results to JSON files or display them in the terminal  
- Concurrent processing with Go routines and channels  

---

## About this project

This application was created as part of my early practice in learning the Go programming language.  
It is one of my **first Go projects**, built to explore and understand key concepts such as:

- Working with packages and project structure  
- Handling user input (from files and command line)  
- JSON encoding and decoding  
- Basic error handling  
- Using structs and interfaces  
- Concurrency with goroutines and channels