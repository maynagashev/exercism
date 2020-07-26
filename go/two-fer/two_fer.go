/*
`Two-fer` or `2-fer` is short for two for one. One for you and one for me.

Given a name, return a string with the message:

```text
One for X, one for me.
```

Where X is the given name.

However, if the name is missing, return the string:

```text
One for you, one for me.
```
 */

package twofer

// Build string with specified name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	//return fmt.Sprintf("One for %s, one for me.", name) // 386 ns/op 3.575s
	return "One for " + name + ", one for me." // actually faster 69.4 ns/op 1.701s
}
