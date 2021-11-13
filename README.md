# Settings

A settings/configuration manager written in Golang and running on etcd. The motivation to write this comes from the concept of how etcs is used by Kubernetes to maintain state and configuration.

## What I'm try to accomplish

- Write golang :P
- Try to solve a problem that I have at work.
- Get to learn about etcd.
- Write a open source tool.

## What this is provides

- A service which allows to:
  - create/add new setting
  - update setting
  - delete setting
  - find a setting
- A service which sends out events so that services can update itself. So when any of the above functions are called appropriate events are triggered.
