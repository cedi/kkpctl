# Usefull commands

this page is dedicated to list some useful command lines with `kkpctl`

## Find all Clusters across all projects for a specific user-name or E-Mail address

### By Usernames

```bash
kkpctl get project --all -ojson | jq '.[]? | select( .owners[]?.name | contains("cedi") ) | .id' -r | xargs -n1 kkpctl get cluster --project
```

### By E-Mail

```bash
kkpctl get project --all -ojson | jq '.[]? | select( .owners[]?.email | contains("cedi@test.de") ) | .id' -r | xargs -n1 kkpctl get cluster --project
```

### Short explanation

* First I try to get all projects present in KKP as json objects, while adding the `--all` ensures that all projects are returned and not just the ones I am part of (only works if I'm an admin).
* In the jq filter term `.[]?`  the `?` protects the filter from nil-values. For example when `.owners[]?.email` is evaluated, but an element in the `.owners[]` array is empty, calling `.email` on it would cause a panic. By adding the `?` it simply skips the nil-value object.
* The `.id` at the end of the filter will only print the project id of a certain project, `-r` prints the raw output with no coloring
* `xargs -n1` then calls `kkpctl get cluster --project $projectID` for every found project id
