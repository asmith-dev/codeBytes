# Encode or Decode Bytes

Two different programs are given here, 
* One takes user input, infers the intended type, and encodes into bytes
* The other allows the user to select a byte array from a list and to choose a decoding 
type, and then decodes that byte array accordingly.

Both of these have been created while having in mind the idea of creating a 
weakly-typed programming language, where variables are all stored in bytes without
storing their corresponding types. Then, variable calls would require denoting a 
decoding type.