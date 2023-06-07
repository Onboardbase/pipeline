
# PipeMuta

A Pipeline configuration tool that generates pipeline configuration from a base configuration file.

This tends to help with Continous Integration and Delivery service migrations.

E.g moving from Gitlab to Github should be a breeze, and vice versa without worrying about how to make that possible.



## Development
$ git clone git@github.com:Onboardbase/pipeline.git
$ make dev

## Usage
Have a pipemuta yaml file that follows the Pipemuta configuration pattern

Then run:
---
```
pipemuta -i base.yaml -o gitlab-ci.yml -t gitlab
```
---

The -i option tells where the base configuration is, the -o specifies the name of the output file, and the -t specifies what type of configuration should be generated.


### Configuration



Defaults:

Jobs:

Workflows:
