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

my-go-app/
├── cmd/                      # 🚀 Entry points for executables
│   └── myapp/                # One subdirectory per binary
│       └── main.go           # The thin shell that boots the app
├── internal/                 # 🔒 Core business logic (private to this module)
│   ├── api/                  # HTTP/gRPC handlers, routers, and transport layers
│   ├── db/                   # Database logic and migrations
│   └── user/                 # Domain-specific logic (e.g., UserService, User struct)
├── pkg/                      # 🌐 Code safe for other external projects to import
│   └── util/                 # General helpers (strings, dates, etc.)
├── config/                   # ⚙️ Configuration loaders (YAML, env vars)
├── go.mod                    # 📦 Module definition and dependencies
├── go.sum                    # Checksums for dependencies
└── Makefile                  # 🛠️ Automation shortcuts (build, test, run)


cmd/ contains the main entry points e'g the main.go file , it shld only wire up dependencies, load configurations and invoke the internal logic

internal/ code here cannot be inported by outside modules, perfect place for all your private application-specific bst logic
        ├── /api or handlers     # Your HTTP/gRPC handlers and routers ...handles incoming requests, decodes requests, calls the appropriate methods on the service layer then encodes the response 
        |- /configs: Configuration files (e.g., config.yaml).
        ├── /service  # Business logic layer
        ├── /repository # Data access layer
        └── /domain     # Core data structures and domain logic ;plain go structs that represent bst entities 
pkg/ code safe to be imported and used by external applications ,

pkg/ explicitly used for code you want to share with other external ...Before you place a package here, ask yourself: “Will another project ever need to import this?” If the answer is no, it belongs in /internal.

GO.MOD sits at the root level and defines the module path and dependency tree


URL-Shortener/
├── app/
│   └── app.go              # App initialization (DB setup, router setup, server start)
├── configs/
│   └── configs.go          # Environment variables & DB connection strings
├── controllers/
│   └── controller.go       # Route declarations (mapping endpoints to handlers)
├── db/
│   └── db.go               # GORM database connection & auto-migrations
├── handlers/
│   └── handlers.go         # Request handlers (HTTP input binding, responses, logic)
├── models/
│   └── url.go              # Database structs/models
├── server/
│   ├── routes.go           # Echo wrapper helper functions (GET, POST, etc.)
│   └── server.go           # Echo instance setup, CORS, and server execution
├── store/
│   └── store.go            # Global state access (holds GORM DB reference)
├── go.mod
├── go.sum
└── main.go                 # Entry point (calls app.Start())