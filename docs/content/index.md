---
title: Get Started
---

# Get Started
- This will help you get started with UTDocs

# Installation
- Install UTDocs CLI from Github [Releases](https://github.com/gaurishhs/utdocs/releases)

# Creating a new Project
- Create a new project using the following command:
```bash
utdocs init <project-name>
```
- If no project name is provided, the project will be created in the current directory

# Config
- UTDocs doesn't generate sidebar on it's own, You need to write the sidebar config in `config.json` file
- The config file is located in the root of the project
- The config file is in JSON format

### Sidebar
- The sidebar is an array of objects
- Each object represents a section in the sidebar

```json 
{
    "sidebaritems": [
        {
            "name": "Section 1",
            "link": "/section1"
        }
    ]
}
```

# Building the Project
- To build the project, run the following command:
```bash
utdocs build <project-dir>
```
- If no project directory is provided, the project in the current directory will be built   
- The built project will be located in the `build` directory
- The built output is in plain HTML5 format and can be hosted on any web server or static hosting service