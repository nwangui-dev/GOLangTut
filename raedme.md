| Array                               | Slice                                   |
| ----------------------------------- | --------------------------------------- |
| Owns the data                       | Refers to data                          |
| Fixed size                          | Can grow with `append()`                |
| Type includes its length (`[5]int`) | Type doesn't include a length (`[]int`) |
| Can't use `append()`                | Can use `append()`                      |
| Less common in everyday Go          | Used everywhere                         |

Modulus operator % -returns the division remainder example 10%3=3rem1 cz 3*3=9 10-9=1

    := creates a new variable       age := 20
    = change an existing variable   age = 20

    &    get an address
    *    follow an address

maps
structs
pointers
methods
interfaces
packages
error handling
file handling
JSON
modules
goroutines
channels
basic concurrency


 cleanest setup ;

Create an empty GitHub repository called GOLangTut
cd ~/PROJECTS/GOLangTut
git init
git branch -M main
git remote add origin git@github.com:nwangui-dev/GOLangTut.git
git add .
git commit -m "initial commit"
git push -u origin main

No git clone.

Make changes
     ↓
git status
     ↓
git add .
     ↓
git commit -m "what I changed"
     ↓
git push