# Goxygen ("JHipster" like) [[{01_PM.low_code,io.http]]

* https://github.com/Shpota/goxygen

* Goxygen generates a Web-project-skeleton with:
  - BACKEND : Go server
  - FRONTEND: Angular|React*|Vue
  - Configuration: Dockerfile+docker-compose files (development|production environments).
  - DDBB:  MongoDB*|MySQL|PostgreSQL

* USSAGE: # go 1.17+
  ```
  | $ go run github.com/shpota/goxygen@latest \
  |   init \
  |   --frontend react \                               angular|react*|vue
  |   --db postgres \                                  mongo* |mysql | postgres
  |   my-app
  |
  | $ cd my-app
  | $ docker-compose up # http://localhost:8080.
  ```

* Layout of skeleton project (React/MongoDB example)

  ```
  | my-app
  | ├─ server               # Go project files
  | │  ├ db                 # MongoDB communications
  | │  ├ model              # domain objects
  | │  ├ web                # REST APIs, web server
  | │  ├ server.go          # the starting point of the server
  | │  └ go.mod             # server dependencies
  | ├─ webapp
  | │  ├ public             # icons, static files, and index.html
  | │  ├ src
  | │  │ ├ App.js           # the main React component
  | │  │ ├ App.css          # App component-specific styles
  | │  │ ├ index.js         # the entry point of the application
  | │  │ └ index.css        # global styles
  | │  ├ package.json       # front end dependencies
  | │  ├ .env.development   # holds API endpoint for dev environment
  | │  └ .env.production    # API endpoint for prod environment
  | │
  | ├─ Dockerfile           # builds back end and front end together
  | ├─ docker-compose.yml   # prod environment deployment descriptor
  | ├─ docker-compose-dev.yml # runs local MongoDB for development needs
  | ├─ init-db.js           # creates a MongoDB collection with test data
  | ├─ .dockerignore        # specifies files ignored in Docker builds
  | ├─ .gitignore
  | └─ README.md            # guide on how to use the generated repo
  ```
  [[}]]
