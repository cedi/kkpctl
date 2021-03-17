# Command Tree

`kkpctl` uses a cli syntax, similar to `kubectl`, consiting of a _verb_ and _nouns_.

In the following tree, you can see how the commands are structured. The first command is always the verb, while the second is always a noun.

The verbs are `add`, `delete`, `describe`, `edit`, and `get`.
The nouns are `cluster`, `nodedeployment`, and `project`. You can use them together with all verbs.
Some verbs do more advanced operations - like `edit`. This makes it necessary to specify a _adjective_ - like `upgrade`. This may seem confusing, but think of it as "What do I need to do? -> Edit. What do I want to edit? -> The Cluster. What do I want to do with the cluster? -> upgrading it."


There are special commands the top level which are not verbs. These are `completion`, `config`, and `ctx`.
They are not related to any objects inside KKP, this is why they are not accessible via the regular verbs.

```
│
│   # this are the verbs with the nouns which can be added, deleted, described, or retrieved from KKP
│
├── add # <------------------ verb
│   ├── cluster # <---------- noun
│   ├── nodedeployment
│   └── project
├── delete
│   ├── cluster
│   ├── nodedeployment
│   └── project
├── describe
│   ├── cluster
│   ├── nodedeployment
│   └── project
├── edit
│   └── cluster
│       └── upgrade
├── get
│   ├── cluster
│   ├── datacenter
│   ├── kubeconfig
│   ├── nodedeployment
│   ├── project
│   └── version
│
│   # from here on are special commands to work with kkpctl
│
├── completion
│   ├── bash
│   ├── fish
│   ├── powershell
│   └── zsh
├── config
│   ├── add
│   │   ├── cloud
│   │   ├── node
│   │   ├── operatingsystem
│   │   └── provider
│   ├── get
│   │   └── cloud
│   └── set
│       └── cloud
│           └── bearer
├── ctx
│   ├── get
│   │   └── cloud
│   └── set
│       └── cloud
└── help
```