package algorithms;

import "strings"

// Encrypts data using the RLE compression method.
// RLE scans the data and replaces repeating consecutive characters with two characters, them being a binary character that has the hexadecimal value of how many of those characters have repeated, followed by a single one of that character that has been repeated.
// Example: "aaaaaaaaaabbbbbbbbb" -> "(NEWLINE)a(TAB)b", where "a" repeats ten times and "b" repeats nine times, represented by a newline and a tab.
func rle(data string) (string, error) {
  if len(data) == 0 { return "", nil } // Handle empty data

  
}
