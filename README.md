# golangexamples
Concurrent and Distributed Systems – Assignment 3
Deadline (no extension possible!): 6th April, 2019 – 11:55 PM.

Overview:
You need to code a golang package, named exactly golangexamples. It should have
following exported functions (use exact same name and arguments):
• golangexamples.ConcatSlice(sliceToConcat []byte) returns string
It should have an exported function to iterate over a byte slice. You provide the
slice as an argument and it returns the contents of the slice concatenated
together and separated by a dash (-).
• golangexamples.Encrypt(sliceToEncrypt []byte, ceaserCount int)
It should have a function to encrypt the provided byte slice. You provide the slice
name address and it changes the original slice by encrypting its text. The
encryption should be done by the Caesar cypher, with the number provided as an
argument with the call.

• golangexamples.EZGreetings(name string) returns string
It invokes the greetings package provided by your instructor and returns the
resulting string. The greetings package is available at
github.com/ehteshamz/greetings
