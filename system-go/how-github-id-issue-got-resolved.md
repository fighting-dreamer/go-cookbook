ssh directory earlier : 
```
▶ cd .ssh

~/.ssh
▶ ls
config                           known_hosts                      ssh_host_dsa_key_683dc920cf9.pub
id_ed25519                       mandapatiala
id_ed25519.pub                   mandapatiala.pub
```
config file earlier :
```
~/.ssh
▶ cat config
Host github.com
  Hostname github.com
  PreferredAuthentications publickey
  IdentityFile /Users/nipun.jindal/.ssh/mandapatiala
  User git
```

followed the following links : 
1. https://help.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh
2. https://help.github.com/en/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent
3. https://gist.github.com/BugBuster1701/80968cfa44b7c1730548d64fdaa30c17
4. https://gist.github.com/developius/c81f021eb5c5916013dc
5. specially this one from last link and responses : https://gist.github.com/developius/c81f021eb5c5916013dc#gistcomment-2957238

my config (now) : 
```
~/.ssh
▶ cat config
Host github.com
  Hostname github.com
  PreferredAuthentications publickey
  IdentityFile /Users/nipun.jindal/.ssh/fighting_dreamer
  User git
```

ssh directory (now) :
```
~/.ssh
▶ ls
config                           id_ed25519                       mandapatiala
fighting_dreamer                 id_ed25519.pub                   mandapatiala.pub
fighting_dreamer.pub             known_hosts                      ssh_host_dsa_key_683dc920cf9.pub
```
BTW, the fighting_dreamer and masdapatiala files are just copies with diff names

the solution was to use the ssh file with user-name rather random otehr filename

(Will be updating, if I am actuaally wrong here though!!!)