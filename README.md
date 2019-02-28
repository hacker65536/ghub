# ghub

set options of repo


## explanation

if you have multiple github accounts, and using SSH.

Set or change user.name user.email and remote url using a yamlfile


`~/.ssh/config`
```sshconfig
host one_github.com
 hostname github.com
 user git
 identityfile ~/.ssh/one_github

host two_github.com
 hostname github.com
 user git
 identityfile ~/.ssh/two_github

host three_github.com
 hostname github.com
 user git
 identityfile ~/.ssh/three_github
```


`.ghub.yml`
```yml
username: hacker65536
useremail: s.hacker65536@gmail.com
host: one_github.com
repo: ghub2
```


## usage
`.ghub.yml` 
```yml
username: username
useremail: useremail 
host: <alternate github.com default=github.com> 
repo: <local dir different from remote>     
secrets: <true / false> (git secretes)
```

[secretes](https://github.com/awslabs/git-secrets)
