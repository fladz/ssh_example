# ssh_example
Super simple example of SSH using Golang using password-less authentication with private key file.

## Set up password-less SSH authentication ##
Follow normal ssh-keygen process to generate id_rsa file:

  ssh-keygen -t rsa
  (which creates id_rsa file under ~/.ssh/)

Append the line to server's ~/.ssh/authorized_keys file.

  scp -p ~/.ssh/id_rsa {SERVER}:~/id_rsa.pub
  
  ssh {SERVER}
  
  cd
  * Create /.ssh/ folder if not exists yet.
  mkdir .ssh
 
  cd .ssh
  cat ~/id_rsa.pub >> authorized_keys

Once this is set up, you can use PublicKeys authentication method.
