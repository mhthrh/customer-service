#!/usr/bin/bash
env_path="A3PATH"
env_value="/home/mohsen/Documents/Golang/x-bank"

shell_config_file="$(echo $SHELL | grep -o 'bash\|zsh')"

	if [ "$shell_config_file" = "bash" ]; then
      shell_config_file="$HOME/.bashrc"
  elif [ "$shell_config_file" = "zsh" ]; then
      shell_config_file="$HOME/.zshrc"
  else
      echo "Unsupported shell: $shell_config_file"
      exit 1
  fi

  echo "export $env_path=\"$env_value\"" >> "$shell_config_file"

  source "$shell_config_file"