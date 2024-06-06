

go build -o git-tool
./git-tool clone --file projects.yaml --branch feature-branch
./git-tool pull --file projects.yaml --branch feature-branch
./git-tool replace --file projects.yaml --old "oldString" --new "newString"
./git-tool push --file projects.yaml --branch feature-branch