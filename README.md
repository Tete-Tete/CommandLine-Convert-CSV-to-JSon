# CSV to JSON Converter - Go Command-line Application

This project is writing a Command-line which uses golang language that converts a CSV file to a JSON lines file. The purpose of this exe file is to provide an easy way to convert CSV file into JSON. 

##  Overview
This exe file reads the CSV data and creates a JSON lines file, where each line contains a JSON object representing one row of data. 

## Features
- **CSVTOJL**: This is the main exe file that we use to conver CSV file to JSON
- **Input data**: The CSV file that you want to convert to JSON
- **Output data**: The JSON file that converts from your CSV input file 

## Requirements
- Go programming language installed (version 1.16 or higher recommended).

## Installation From Git and Set up for your own computer
### Step 1: Clone the Repository
Clone this repository to your local machine:
```sh
git clone <https://github.com/Tete-Tete/CommandLine-Convert-CSV-to-JSon>.git
```

### Step 2: Run the Application
To run the Go application, run the following command in your terminal:
```sh
go build -o csvtojl
```
This will create a file named `csvtojl` in your current directory.

After running the application, type the following command:
```sh
./csvtojl input.csv output.jl   
```
Replace `input.csv` with the path to your CSV file, and `output.jl` with the output JSON file.

## Command-line Arguments
- **input.csv**: The path where your CSV File located. 
- **output.jl**: The path where your output JSON file will be located.

## Example
Given an input CSV file (`housesInput.csv`) which is provided in my github like this:
```csv
name,age,salary
Alice,30,55000
Bob,25,48000
```
The output JSON file (`housesOutput.jl`) will look like this:
```json
{"name":"Alice","age":30,"salary":55000}
{"name":"Bob","age":25,"salary":48000}
```

## Testing the Application
1. **Unit Tests**: You can add unit tests for the critical components of the application to ensure its reliability. A `main_test.go` file can be used for testing different scenarios.
2. **Validation Tests**: Use online JSON validators like [JSONLint](https://jsonlint.com/) to ensure the output JSON is valid.

## Error Handling
The application handles various types of errors:
- **Missing arguments**: Ensures the user provides both input and output file paths.
- **File errors**: Catches errors when opening, reading, or writing files.
- **Conversion errors**: Handles errors during data conversion from CSV to JSON.

If any error occurs, the program will send a message with first world is "Sry"

## Conclusion

This project is for use Go command-Line application to do the convert from CSV To JSON. By the above steps and instructions, you should can convery easyly. If you met some problems, you are feel free to reach out via Github issues. 

