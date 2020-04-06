### Bitbucket Server

#### List Projects
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects
```

#### List Repos in Project 
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos
```

#### Show Repository Detail
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos/git-reader
```

#### Show Pull Requests
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos/git-reader/pull-requests


#### List Tags
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos/git-reader/tags
```

#### List Branches
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos/git-reader/branches
```

#### List Files
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos/git-reader/files/
```
#### Retrieve Raw File
```
curl -s -H"Authorization: Bearer $GIT_READER" \
http://localhost:7990/rest/api/1.0/projects/TOOL/repos/git-reader/raw/.gitignore
```