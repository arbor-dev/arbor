# CONTRIBUTING
> Thanks for wanting to contribute!

We use milestones to organize our work. Take a look at unassigned tasks in each milestione

## What you should work on

If you want to help out check the issues and assign yourself to one. More accessible issues have the tag starter bug.

## Submitting Code Changes
If you are ready to put a pull request in, great!

## Things to verify:
- [ ] Make sure to tag the issue you are addressing in your comments if applicable. 
- [ ] Make sure any tests pass
- [ ] Make sure you stick to rough style guidelines
- [ ] Have you added the license header to any new files? 

After a quick discussion and maybe some changes requested, your code will get merged 


### Style Guidelines 
__Go__ (just use a linter basically):
```go
//Function Signatures:
func foo(arg1: int, arg2: bool) { 

//if/else 
if x == y {
  return x
}
return y 

//No non bracketed blocks
//wrong:
if x == 1
  return true 
//correct:
if x == 1 {
  return true
}
```

Basically use the go linter

### License Header 
```
Copyright © 2017, ACM@UIUC

This file is part of the Groot Project.  
 
The Arbor Project is open source software, released under the University of Illinois/NCSA Open Source License. 
You should have received a copy of this license in a file with the distribution.
```
