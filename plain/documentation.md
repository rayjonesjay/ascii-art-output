
### DOCUMENTATION: This file contains new features added to the codebase ###

### Objective(s):

	[*] If you add new functions to the code kindly add the function name, syntax what it takes and what it returns.
	[*] If the function is too long just provide a link to the line number of the file.

## How to provide a link to a function defined at line N:

Suppose you have a function called `myFunction` defined in `script.go`, in `main` branch in line number `42` (use `#L42` where L is line and 42 is line number where the function is defined). provide full path so that the link is correct.

If the function block is too long, provide the link to the file. 

Here is an example `don't click` on the link, i dont know where it will take you.

- [myFunction](https://github.com/your-username/my-repo/blob/main/src/script.go#L42):
Takes a string and returns a boolean value

```
	func myFunction(str string) bool {
		for _, char := range {
			....
		}
		...
		return bool
	}
```

- [isLower](https://github.com/your-username/my-repo/blob/main/src/script.go#L42):
takes a rune value and checks if it lowercase else false
```
	func isLower(r rune) bool {
		if r >= 'a' && r <= 'z'{
			return true
		}
		return false
	}
```

***NOTE** 
This documentation is for complex functions that one of our team members might have written and it is only clear to him.