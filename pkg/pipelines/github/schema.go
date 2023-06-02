package github

func get_schema() string {
	return `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://gitlab.com/.gitlab-ci.yml",
  "markdownDescription": "Gitlab has a built-in solution for doing CI called Gitlab CI. It is configured by supplying a file called ".gitlab-ci.yml", which will list all the jobs that are going to run for the project. A full list of all options can be found [here](https://docs.gitlab.com/ee/ci/yaml). [Learn More](https://docs.gitlab.com/ee/ci/index.html).",
  "type": "object",
  "properties": {
    "$schema": {
      "type": "string",
      "format": "uri"
    },
    "image": {
      "$ref": "#/definitions/image"
    },
    "services": {
      "$ref": "#/definitions/services"
    },
    "before_script": {
      "$ref": "#/definitions/before_script"
    },
    "after_script": {
      "$ref": "#/definitions/after_script"
    },
    "variables": {
      "$ref": "#/definitions/globalVariables"
    },
    "cache": {
      "$ref": "#/definitions/cache"
    },
    "!reference": {
      "$ref": "#/definitions/!reference"
    },
    "default": {
      "type": "object",
      "properties": {
        "after_script": {
          "$ref": "#/definitions/after_script"
        },
        "artifacts": {
          "$ref": "#/definitions/artifacts"
        },
        "before_script": {
          "$ref": "#/definitions/before_script"
        },
        "hooks": {
          "$ref": "#/definitions/hooks"
        },
        "cache": {
          "$ref": "#/definitions/cache"
        },
        "image": {
          "$ref": "#/definitions/image"
        },
        "interruptible": {
          "$ref": "#/definitions/interruptible"
        },
        "retry": {
          "$ref": "#/definitions/retry"
        },
        "services": {
          "$ref": "#/definitions/services"
        },
        "tags": {
          "$ref": "#/definitions/tags"
        },
        "timeout": {
          "$ref": "#/definitions/timeout"
        },
        "!reference": {
          "$ref": "#/definitions/!reference"
        }
      },
      "additionalProperties": false
    },
    "stages": {
      "type": "array",
      "markdownDescription": "Groups jobs into stages. All jobs in one stage must complete before next stage is executed. Defaults to ['build', 'test', 'deploy']. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#stages).",
      "default": [
        "build",
        "test",
        "deploy"
      ],
      "items": {
        "type": "string"
      },
      "uniqueItems": true,
      "minItems": 1
    },
    "include": {
      "markdownDescription": "Can be "IncludeItem" or "IncludeItem[]". Each "IncludeItem" will be a string, or an object with properties for the method if including external YAML file. The external content will be fetched, included and evaluated along the ".gitlab-ci.yml". [Learn More](https://docs.gitlab.com/ee/ci/yaml/#include).",
      "oneOf": [
        {
          "$ref": "#/definitions/include_item"
        },
        {
          "type": "array",
          "items": {
            "$ref": "#/definitions/include_item"
          }
        }
      ]
    },
    "pages": {
      "$ref": "#/definitions/job",
      "markdownDescription": "A special job used to upload static sites to Gitlab pages. Requires a "public/" directory with "artifacts.path" pointing to it. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#pages)."
    },
    "workflow": {
      "type": "object",
      "properties": {
        "name": {
          "$ref": "#/definitions/workflowName"
        },
        "rules": {
          "type": "array",
          "items": {
            "anyOf": [
              {
                "type": "object"
              },
              {
                "type": "array",
                "minLength": 1,
                "items": {
                  "type": "string"
                }
              }
            ],
            "properties": {
              "if": {
                "$ref": "#/definitions/if"
              },
              "changes": {
                "$ref": "#/definitions/changes"
              },
              "exists": {
                "$ref": "#/definitions/exists"
              },
              "variables": {
                "$ref": "#/definitions/rulesVariables"
              },
              "when": {
                "type": "string",
                "enum": [
                  "always",
                  "never"
                ]
              }
            },
            "additionalProperties": false
          }
        }
      }
    }
  },
  "patternProperties": {
    "^[.]": {
      "description": "Hidden keys.",
      "anyOf": [
        {
          "$ref": "#/definitions/job_template"
        },
        {
          "description": "Arbitrary YAML anchor."
        }
      ]
    }
  },
  "additionalProperties": {
    "$ref": "#/definitions/job"
  },
  "definitions": {
    "artifacts": {
      "type": "object",
      "markdownDescription": "Used to specify a list of files and directories that should be attached to the job if it succeeds. Artifacts are sent to Gitlab where they can be downloaded. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifacts).",
      "additionalProperties": false,
      "properties": {
        "paths": {
          "type": "array",
          "markdownDescription": "A list of paths to files/folders that should be included in the artifact. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactspaths).",
          "items": {
            "type": "string"
          },
          "minItems": 1
        },
        "exclude": {
          "type": "array",
          "markdownDescription": "A list of paths to files/folders that should be excluded in the artifact. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactsexclude).",
          "items": {
            "type": "string"
          },
          "minItems": 1
        },
        "expose_as": {
          "type": "string",
          "markdownDescription": "Can be used to expose job artifacts in the merge request UI. GitLab will add a link <expose_as> to the relevant merge request that points to the artifact. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactsexpose_as)."
        },
        "name": {
          "type": "string",
          "markdownDescription": "Name for the archive created on job success. Can use variables in the name, e.g. '$CI_JOB_NAME' [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactsname)."
        },
        "untracked": {
          "type": "boolean",
          "markdownDescription": "Whether to add all untracked files (along with 'artifacts.paths') to the artifact. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactsuntracked).",
          "default": false
        },
        "when": {
          "markdownDescription": "Configure when artifacts are uploaded depended on job status. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactswhen).",
          "default": "on_success",
          "type": "string",
          "enum": [
            "on_success",
            "on_failure",
            "always"
          ]
        },
        "expire_in": {
          "type": "string",
          "markdownDescription": "How long artifacts should be kept. They are saved 30 days by default. Artifacts that have expired are removed periodically via cron job. Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactsexpire_in).",
          "default": "30 days"
        },
        "reports": {
          "type": "object",
          "markdownDescription": "Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#artifactsreports).",
          "additionalProperties": false,
          "properties": {
            "junit": {
              "description": "Path for file(s) that should be parsed as JUnit XML result",
              "oneOf": [
                {
                  "type": "string",
                  "description": "Path to a single XML file"
                },
                {
                  "type": "array",
                  "description": "A list of paths to XML files that will automatically be concatenated into a single file",
                  "items": {
                    "type": "string"
                  },
                  "minItems": 1
                }
              ]
            },
            "browser_performance": {
              "type": "string",
              "description": "Path to a single file with browser performance metric report(s)."
            },
            "coverage_report": {
              "type": [
                "object",
                "null"
              ],
              "description": "Used to collect coverage reports from the job.",
              "properties": {
                "coverage_format": {
                  "description": "Code coverage format used by the test framework.",
                  "enum": [
                    "cobertura"
                  ]
                },
                "path": {
                  "description": "Path to the coverage report file that should be parsed.",
                  "type": "string",
                  "minLength": 1
                }
              }
            },
            "codequality": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with code quality report(s) (such as Code Climate)."
            },
            "dotenv": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files containing runtime-created variables for this job."
            },
            "lsif": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files containing code intelligence (Language Server Index Format)."
            },
            "sast": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with SAST vulnerabilities report(s)."
            },
            "dependency_scanning": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with Dependency scanning vulnerabilities report(s)."
            },
            "container_scanning": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with Container scanning vulnerabilities report(s)."
            },
            "dast": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with DAST vulnerabilities report(s)."
            },
            "license_management": {
              "$ref": "#/definitions/string_file_list",
              "description": "Deprecated in 12.8: Path to file or list of files with license report(s)."
            },
            "license_scanning": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with license report(s)."
            },
            "requirements": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with requirements report(s)."
            },
            "secret_detection": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with secret detection report(s)."
            },
            "metrics": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with custom metrics report(s)."
            },
            "terraform": {
              "$ref": "#/definitions/string_file_list",
              "description": "Path to file or list of files with terraform plan(s)."
            },
            "cyclonedx": {
              "$ref": "#/definitions/string_file_list",
              "markdownDescription": "Path to file or list of files with cyclonedx report(s). [Learn More](https://docs.gitlab.com/ee/ci/yaml/artifacts_reports.html#artifactsreportscyclonedx)."
            },
            "load_performance": {
              "$ref": "#/definitions/string_file_list",
              "markdownDescription": "Path to file or list of files with load performance testing report(s). [Learn More](https://docs.gitlab.com/ee/ci/yaml/artifacts_reports.html#artifactsreportsload_performance)."
            }
          }
        }
      }
    },
    "string_file_list": {
      "oneOf": [
        {
          "type": "string"
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      ]
    },
    "include_item": {
      "oneOf": [
        {
          "description": "Will infer the method based on the value. E.g. "https://..." strings will be of type "include:remote", and "/templates/..." or "templates/..." will be of type "include:local".",
          "type": "string",
          "format": "uri-reference",
          "pattern": "^(https?://|/?.?-?(?!\\w+://)\\w).+\\.ya?ml$"
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "local": {
              "description": "Relative path from local repository root ("/") to the "yaml"/"yml" file template. The file must be on the same branch, and does not work across git submodules.",
              "type": "string",
              "format": "uri-reference",
              "pattern": "\\.ya?ml$"
            },
            "rules": {
              "$ref": "#/definitions/rules"
            },
            "inputs": {
              "$ref": "#/definitions/inputs"
            }
          },
          "required": [
            "local"
          ]
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "project": {
              "description": "Path to the project, e.g. "group/project", or "group/sub-group/project" [Learn more](https://docs.gitlab.com/ee/ci/yaml/index.html#includefile).",
              "type": "string",
              "pattern": "(?:\\S/\\S|\\$\\S+)"
            },
            "ref": {
              "description": "Branch/Tag/Commit-hash for the target project.",
              "type": "string"
            },
            "file": {
              "oneOf": [
                {
                  "description": "Relative path from project root ("/") to the "yaml"/"yml" file template.",
                  "type": "string",
                  "pattern": "\\.ya?ml$"
                },
                {
                  "description": "List of files by relative path from project root ("/") to the "yaml"/"yml" file template.",
                  "type": "array",
                  "items": {
                    "type": "string",
                    "pattern": "\\.ya?ml$"
                  }
                }
              ]
            },
            "inputs": {
              "$ref": "#/definitions/inputs"
            }
          },
          "required": [
            "project",
            "file"
          ]
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "template": {
              "description": "Use a ".gitlab-ci.yml" template as a base, e.g. "Nodejs.gitlab-ci.yml".",
              "type": "string",
              "format": "uri-reference",
              "pattern": "\\.ya?ml$"
            },
            "inputs": {
              "$ref": "#/definitions/inputs"
            }
          },
          "required": [
            "template"
          ]
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "component": {
              "description": "Local path to component directory or full path to external component directory.",
              "type": "string",
              "format": "uri-reference"
            },
            "inputs": {
              "$ref": "#/definitions/inputs"
            }
          },
          "required": [
            "component"
          ]
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "remote": {
              "description": "URL to a "yaml"/"yml" template file using HTTP/HTTPS.",
              "type": "string",
              "format": "uri-reference",
              "pattern": "^https?://.+\\.ya?ml$"
            },
            "inputs": {
              "$ref": "#/definitions/inputs"
            }
          },
          "required": [
            "remote"
          ]
        }
      ]
    },
    "!reference": {
      "type": "array",
      "items": {
        "type": "string",
        "minLength": 1
      }
    },
    "image": {
      "oneOf": [
        {
          "type": "string",
          "minLength": 1,
          "description": "Full name of the image that should be used. It should contain the Registry part if needed."
        },
        {
          "type": "object",
          "description": "Specifies the docker image to use for the job or globally for all jobs. Job configuration takes precedence over global setting. Requires a certain kind of Gitlab runner executor.",
          "additionalProperties": false,
          "properties": {
            "name": {
              "type": "string",
              "minLength": 1,
              "description": "Full name of the image that should be used. It should contain the Registry part if needed."
            },
            "entrypoint": {
              "type": "array",
              "description": "Command or script that should be executed as the container's entrypoint. It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.",
              "minItems": 1
            },
            "pull_policy": {
              "markdownDescription": "Specifies how to pull the image in Runner. It can be one of "always", "never" or "if-not-present". The default value is "always". [Learn more](https://docs.gitlab.com/ee/ci/yaml/#imagepull_policy).",
              "default": "always",
              "oneOf": [
                {
                  "type": "string",
                  "enum": [
                    "always",
                    "never",
                    "if-not-present"
                  ]
                },
                {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "enum": [
                      "always",
                      "never",
                      "if-not-present"
                    ]
                  },
                  "minItems": 1,
                  "uniqueItems": true
                }
              ]
            }
          },
          "required": [
            "name"
          ]
        },
        {
          "type": "array",
          "minLength": 1,
          "items": {
            "type": "string"
          }
        }
      ],
      "markdownDescription": "Specifies the docker image to use for the job or globally for all jobs. Job configuration takes precedence over global setting. Requires a certain kind of Gitlab runner executor. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#image)."
    },
    "services": {
      "type": "array",
      "markdownDescription": "Similar to "image" property, but will link the specified services to the "image" container. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#services).",
      "items": {
        "oneOf": [
          {
            "type": "string",
            "minLength": 1,
            "description": "Full name of the image that should be used. It should contain the Registry part if needed."
          },
          {
            "type": "object",
            "description": "",
            "additionalProperties": false,
            "properties": {
              "name": {
                "type": "string",
                "description": "Full name of the image that should be used. It should contain the Registry part if needed.",
                "minLength": 1
              },
              "entrypoint": {
                "type": "array",
                "markdownDescription": "Command or script that should be executed as the container's entrypoint. It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array. [Learn More](https://docs.gitlab.com/ee/ci/services/index.html#available-settings-for-services)",
                "minItems": 1,
                "items": {
                  "type": "string"
                }
              },
              "pull_policy": {
                "markdownDescription": "Specifies how to pull the image in Runner. It can be one of "always", "never" or "if-not-present". The default value is "always". [Learn more](https://docs.gitlab.com/ee/ci/yaml/#servicepull_policy).",
                "default": "always",
                "oneOf": [
                  {
                    "type": "string",
                    "enum": [
                      "always",
                      "never",
                      "if-not-present"
                    ]
                  },
                  {
                    "type": "array",
                    "items": {
                      "type": "string",
                      "enum": [
                        "always",
                        "never",
                        "if-not-present"
                      ]
                    },
                    "minItems": 1,
                    "uniqueItems": true
                  }
                ]
              },
              "command": {
                "type": "array",
                "markdownDescription": "Command or script that should be used as the container's command. It will be translated to arguments passed to Docker after the image's name. The syntax is similar to Dockerfile's CMD directive, where each shell token is a separate string in the array. [Learn More](https://docs.gitlab.com/ee/ci/services/index.html#available-settings-for-services)",
                "minItems": 1,
                "items": {
                  "type": "string"
                }
              },
              "alias": {
                "type": "string",
                "markdownDescription": "Additional alias that can be used to access the service from the job's container. Read Accessing the services for more information. [Learn More](https://docs.gitlab.com/ee/ci/services/index.html#available-settings-for-services)",
                "minLength": 1
              },
              "variables": {
                "$ref": "#/definitions/jobVariables",
                "markdownDescription": "Additional environment variables that are passed exclusively to the service. Service variables cannot reference themselves. [Learn More](https://docs.gitlab.com/ee/ci/services/index.html#available-settings-for-services)"
              }
            },
            "required": [
              "name"
            ]
          }
        ]
      }
    },
    "id_tokens": {
      "type": "object",
      "markdownDescription": "Defines JWTs to be injected as environment variables.",
      "patternProperties": {
        ".*": {
          "type": "object",
          "properties": {
            "aud": {
              "oneOf": [
                {
                  "type": "string"
                },
                {
                  "type": "array",
                  "items": {
                    "type": "string"
                  },
                  "minItems": 1,
                  "uniqueItems": true
                }
              ]
            }
          },
          "required": [
            "aud"
          ],
          "additionalProperties": false
        }
      }
    },
    "secrets": {
      "type": "object",
      "markdownDescription": "Defines secrets to be injected as environment variables. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#secrets).",
      "patternProperties": {
        ".*": {
          "type": "object",
          "properties": {
            "vault": {
              "oneOf": [
                {
                  "type": "string",
                  "markdownDescription": "The secret to be fetched from Vault (e.g. 'production/db/password@ops' translates to secret 'ops/data/production/db', field "password"). [Learn More](https://docs.gitlab.com/ee/ci/yaml/#secretsvault)"
                },
                {
                  "type": "object",
                  "properties": {
                    "engine": {
                      "type": "object",
                      "properties": {
                        "name": {
                          "type": "string"
                        },
                        "path": {
                          "type": "string"
                        }
                      },
                      "required": [
                        "name",
                        "path"
                      ]
                    },
                    "path": {
                      "type": "string"
                    },
                    "field": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "engine",
                    "path",
                    "field"
                  ],
                  "additionalProperties": false
                }
              ]
            },
            "file": {
              "type": "boolean",
              "default": true,
              "markdownDescription": "Configures the secret to be stored as either a file or variable type CI/CD variable. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#secretsfile)"
            },
            "token": {
              "type": "string",
              "description": "Specifies the JWT variable that should be used to authenticate with Hashicorp Vault."
            }
          },
          "required": [
            "vault"
          ],
          "additionalProperties": false
        }
      }
    },
    "before_script": {
      "type": "array",
      "markdownDescription": "Defines scripts that should run *before* the job. Can be set globally or per job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#before_script).",
      "items": {
        "anyOf": [
          {
            "type": "string"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        ]
      }
    },
    "after_script": {
      "type": "array",
      "markdownDescription": "Defines scripts that should run *after* the job. Can be set globally or per job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#after_script).",
      "items": {
        "anyOf": [
          {
            "type": "string"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        ]
      }
    },
    "rules": {
      "type": [
        "array",
        "null"
      ],
      "markdownDescription": "Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#rules).",
      "items": {
        "anyOf": [
          {
            "type": "object",
            "additionalProperties": false,
            "properties": {
              "if": {
                "$ref": "#/definitions/if"
              },
              "changes": {
                "$ref": "#/definitions/changes"
              },
              "exists": {
                "$ref": "#/definitions/exists"
              },
              "variables": {
                "$ref": "#/definitions/rulesVariables"
              },
              "when": {
                "$ref": "#/definitions/when"
              },
              "start_in": {
                "$ref": "#/definitions/start_in"
              },
              "allow_failure": {
                "$ref": "#/definitions/allow_failure"
              },
              "needs": {
                "$ref": "#/definitions/rulesNeeds"
              }
            }
          },
          {
            "type": "string",
            "minLength": 1
          },
          {
            "type": "array",
            "minLength": 1,
            "items": {
              "type": "string"
            }
          }
        ]
      }
    },
    "workflowName": {
      "type": "string",
      "markdownDescription": "Defines the pipeline name. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#workflowname).",
      "minLength": 1,
      "maxLength": 255
    },
    "globalVariables": {
      "markdownDescription": "Defines default variables for all jobs. Job level property overrides global variables. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variables).",
      "type": "object",
      "patternProperties": {
        ".*": {
          "oneOf": [
            {
              "type": [
                "string",
                "number"
              ]
            },
            {
              "type": "object",
              "properties": {
                "value": {
                  "type": "string",
                  "markdownDescription": "Default value of the variable. If used with "options", "value" must be included in the array. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variablesvalue)"
                },
                "options": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  },
                  "minItems": 1,
                  "uniqueItems": true,
                  "markdownDescription": "A list of predefined values that users can select from in the **Run pipeline** page when running a pipeline manually. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variablesoptions)"
                },
                "description": {
                  "type": "string",
                  "markdownDescription": "Explains what the variable is used for, what the acceptable values are. Variables with "description" are prefilled when running a pipeline manually. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variablesdescription)."
                },
                "expand": {
                  "type": "boolean",
                  "markdownDescription": "If the variable is expandable or not. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variablesexpand)."
                }
              },
              "additionalProperties": false
            }
          ]
        },
        "additionalProperties": false
      }
    },
    "jobVariables": {
      "markdownDescription": "Defines variables for a job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variables).",
      "type": "object",
      "patternProperties": {
        ".*": {
          "oneOf": [
            {
              "type": [
                "string",
                "number"
              ]
            },
            {
              "type": "object",
              "properties": {
                "value": {
                  "type": "string"
                },
                "expand": {
                  "type": "boolean",
                  "markdownDescription": "Defines if the variable is expandable or not. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#variablesexpand)."
                }
              },
              "additionalProperties": false
            }
          ]
        },
        "additionalProperties": false
      }
    },
    "rulesVariables": {
      "markdownDescription": "Defines variables for a rule result. [Learn More](https://docs.gitlab.com/ee/ci/yaml/index.html#rulesvariables).",
      "type": "object",
      "patternProperties": {
        ".*": {
          "type": [
            "string",
            "number"
          ]
        },
        "additionalProperties": false
      }
    },
    "if": {
      "type": "string",
      "markdownDescription": "Expression to evaluate whether additional attributes should be provided to the job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#rulesif)."
    },
    "changes": {
      "markdownDescription": "Additional attributes will be provided to job if any of the provided paths matches a modified file. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#ruleschanges).",
      "anyOf": [
        {
          "type": "object",
          "additionalProperties": false,
          "required": [
            "paths"
          ],
          "properties": {
            "paths": {
              "type": "array",
              "description": "List of file paths.",
              "items": {
                "type": "string"
              }
            },
            "compare_to": {
              "type": "string",
              "description": "Ref for comparing changes."
            }
          }
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      ]
    },
    "exists": {
      "type": "array",
      "markdownDescription": "Additional attributes will be provided to job if any of the provided paths matches an existing file in the repository. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#rulesexists).",
      "items": {
        "type": "string"
      }
    },
    "timeout": {
      "type": "string",
      "markdownDescription": "Allows you to configure a timeout for a specific job (e.g. "1 minute", "1h 30m 12s"). [Learn More](https://docs.gitlab.com/ee/ci/yaml/index.html#timeout).",
      "minLength": 1
    },
    "start_in": {
      "type": "string",
      "markdownDescription": "Used in conjunction with 'when: delayed' to set how long to delay before starting a job. e.g. '5', 5 seconds, 30 minutes, 1 week, etc. [Learn More](https://docs.gitlab.com/ee/ci/jobs/job_control.html#run-a-job-after-a-delay).",
      "minLength": 1
    },
    "rulesNeeds": {
      "markdownDescription": "Use needs in rules to update job needs for specific conditions. When a condition matches a rule, the job's needs configuration is completely replaced with the needs in the rule. [Learn More](https://docs.gitlab.com/ee/ci/yaml/index.html#rulesneeds).",
      "type": "array",
      "items": {
        "oneOf": [
          {
            "type": "string"
          },
          {
            "type": "object",
            "additionalProperties": false,
            "properties": {
              "job": {
                "type": "string",
                "minLength": 1,
                "description": "Name of a job that is defined in the pipeline."
              },
              "artifacts": {
                "type": "boolean",
                "description": "Download artifacts of the job in needs."
              },
              "optional": {
                "type": "boolean",
                "description": "Whether the job needs to be present in the pipeline to run ahead of the current job."
              }
            },
            "required": [
              "job"
            ]
          }
        ]
      }
    },
    "allow_failure": {
      "markdownDescription": "Allow job to fail. A failed job does not cause the pipeline to fail. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#allow_failure).",
      "oneOf": [
        {
          "description": "Setting this option to true will allow the job to fail while still letting the pipeline pass.",
          "type": "boolean",
          "default": false
        },
        {
          "description": "Exit code that are not considered failure. The job fails for any other exit code.",
          "type": "object",
          "additionalProperties": false,
          "required": [
            "exit_codes"
          ],
          "properties": {
            "exit_codes": {
              "type": "integer"
            }
          }
        },
        {
          "description": "You can list which exit codes are not considered failures. The job fails for any other exit code.",
          "type": "object",
          "additionalProperties": false,
          "required": [
            "exit_codes"
          ],
          "properties": {
            "exit_codes": {
              "type": "array",
              "minItems": 1,
              "uniqueItems": true,
              "items": {
                "type": "integer"
              }
            }
          }
        }
      ]
    },
    "when": {
      "markdownDescription": "Describes the conditions for when to run the job. Defaults to 'on_success'. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#when).",
      "default": "on_success",
      "type": "string",
      "enum": [
        "on_success",
        "on_failure",
        "always",
        "never",
        "manual",
        "delayed"
      ]
    },
    "cache": {
      "markdownDescription": "Use "cache" to specify a list of files and directories to cache between jobs. You can only use paths that are in the local working copy. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cache)",
      "properties": {
        "key": {
          "markdownDescription": "Use the "cache:key" keyword to give each cache a unique identifying key. All jobs that use the same cache key use the same cache, including in different pipelines. Must be used with "cache:path", or nothing is cached. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cachekey).",
          "oneOf": [
            {
              "type": "string",
              "pattern": "^(?!.*\\/)^(.*[^.]+.*)$"
            },
            {
              "type": "object",
              "properties": {
                "files": {
                  "markdownDescription": "Use the "cache:key:files" keyword to generate a new key when one or two specific files change. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cachekeyfiles)",
                  "type": "array",
                  "items": {
                    "type": "string"
                  },
                  "minItems": 1,
                  "maxItems": 2
                },
                "prefix": {
                  "markdownDescription": "Use "cache:key:prefix" to combine a prefix with the SHA computed for "cache:key:files". [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cachekeyprefix)",
                  "type": "string"
                }
              }
            }
          ]
        },
        "paths": {
          "type": "array",
          "markdownDescription": "Use the "cache:paths" keyword to choose which files or directories to cache. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cachepaths)",
          "items": {
            "type": "string"
          }
        },
        "policy": {
          "type": "string",
          "markdownDescription": "Determines the strategy for downloading and updating the cache. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cachepolicy)",
          "default": "pull-push",
          "enum": [
            "pull",
            "push",
            "pull-push"
          ]
        },
        "unprotect": {
          "type": "boolean",
          "markdownDescription": "Use "unprotect: true" to set a cache to be shared between protected and unprotected branches.",
          "default": false
        },
        "untracked": {
          "type": "boolean",
          "markdownDescription": "Use "untracked: true" to cache all files that are untracked in your Git repository. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cacheuntracked)",
          "default": false
        },
        "when": {
          "type": "string",
          "markdownDescription": "Defines when to save the cache, based on the status of the job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#cachewhen).",
          "default": "on_success",
          "enum": [
            "on_success",
            "on_failure",
            "always"
          ]
        },
        "fallback_keys": {
          "type": "array",
          "markdownDescription": "List of keys to download cache from if no cache hit occurred for key",
          "items": {
            "type": "string"
          },
          "maxItems": 5
        }
      }
    },
    "filter_refs": {
      "type": "array",
      "description": "Filter job by different keywords that determine origin or state, or by supplying string/regex to check against branch/tag names.",
      "items": {
        "anyOf": [
          {
            "oneOf": [
              {
                "enum": [
                  "branches"
                ],
                "description": "When a branch is pushed."
              },
              {
                "enum": [
                  "tags"
                ],
                "description": "When a tag is pushed."
              },
              {
                "enum": [
                  "api"
                ],
                "description": "When a pipeline has been triggered by a second pipelines API (not triggers API)."
              },
              {
                "enum": [
                  "external"
                ],
                "description": "When using CI services other than Gitlab"
              },
              {
                "enum": [
                  "pipelines"
                ],
                "description": "For multi-project triggers, created using the API with 'CI_JOB_TOKEN'."
              },
              {
                "enum": [
                  "pushes"
                ],
                "description": "Pipeline is triggered by a "git push" by the user"
              },
              {
                "enum": [
                  "schedules"
                ],
                "description": "For scheduled pipelines."
              },
              {
                "enum": [
                  "triggers"
                ],
                "description": "For pipelines created using a trigger token."
              },
              {
                "enum": [
                  "web"
                ],
                "description": "For pipelines created using *Run pipeline* button in Gitlab UI (under your project's *Pipelines*)."
              }
            ]
          },
          {
            "type": "string",
            "description": "String or regular expression to match against tag or branch names."
          }
        ]
      }
    },
    "filter": {
      "oneOf": [
        {
          "type": "null"
        },
        {
          "$ref": "#/definitions/filter_refs"
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "refs": {
              "$ref": "#/definitions/filter_refs"
            },
            "kubernetes": {
              "enum": [
                "active"
              ],
              "description": "Filter job based on if Kubernetes integration is active."
            },
            "variables": {
              "type": "array",
              "markdownDescription": "Filter job by checking comparing values of CI/CD variables. [Learn More](https://docs.gitlab.com/ee/ci/jobs/job_control.html#cicd-variable-expressions).",
              "items": {
                "type": "string"
              }
            },
            "changes": {
              "type": "array",
              "description": "Filter job creation based on files that were modified in a git push.",
              "items": {
                "type": "string"
              }
            }
          }
        }
      ]
    },
    "retry": {
      "markdownDescription": "Retry a job if it fails. Can be a simple integer or object definition. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#retry).",
      "oneOf": [
        {
          "$ref": "#/definitions/retry_max"
        },
        {
          "type": "object",
          "additionalProperties": false,
          "properties": {
            "max": {
              "$ref": "#/definitions/retry_max"
            },
            "when": {
              "markdownDescription": "Either a single or array of error types to trigger job retry. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#retrywhen).",
              "oneOf": [
                {
                  "$ref": "#/definitions/retry_errors"
                },
                {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/retry_errors"
                  }
                }
              ]
            }
          }
        }
      ]
    },
    "retry_max": {
      "type": "integer",
      "description": "The number of times the job will be retried if it fails. Defaults to 0 and can max be retried 2 times (3 times total).",
      "default": 0,
      "minimum": 0,
      "maximum": 2
    },
    "retry_errors": {
      "oneOf": [
        {
          "const": "always",
          "description": "Retry on any failure (default)."
        },
        {
          "const": "unknown_failure",
          "description": "Retry when the failure reason is unknown."
        },
        {
          "const": "script_failure",
          "description": "Retry when the script failed."
        },
        {
          "const": "api_failure",
          "description": "Retry on API failure."
        },
        {
          "const": "stuck_or_timeout_failure",
          "description": "Retry when the job got stuck or timed out."
        },
        {
          "const": "runner_system_failure",
          "description": "Retry if there is a runner system failure (for example, job setup failed)."
        },
        {
          "const": "runner_unsupported",
          "description": "Retry if the runner is unsupported."
        },
        {
          "const": "stale_schedule",
          "description": "Retry if a delayed job could not be executed."
        },
        {
          "const": "job_execution_timeout",
          "description": "Retry if the script exceeded the maximum execution time set for the job."
        },
        {
          "const": "archived_failure",
          "description": "Retry if the job is archived and can’t be run."
        },
        {
          "const": "unmet_prerequisites",
          "description": "Retry if the job failed to complete prerequisite tasks."
        },
        {
          "const": "scheduler_failure",
          "description": "Retry if the scheduler failed to assign the job to a runner."
        },
        {
          "const": "data_integrity_failure",
          "description": "Retry if there is a structural integrity problem detected."
        }
      ]
    },
    "interruptible": {
      "type": "boolean",
      "markdownDescription": "Interruptible is used to indicate that a job should be canceled if made redundant by a newer pipeline run. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#interruptible).",
      "default": false
    },
    "inputs": {
      "markdownDescription": "Used to pass input values to included templates or components. [Learn More](https://docs.gitlab.com/ee/ci/yaml/includes.html#set-input-parameter-values-with-includeinputs).",
      "type": "object"
    },
    "job": {
      "allOf": [
        {
          "$ref": "#/definitions/job_template"
        }
      ]
    },
    "job_template": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "image": {
          "$ref": "#/definitions/image"
        },
        "services": {
          "$ref": "#/definitions/services"
        },
        "before_script": {
          "$ref": "#/definitions/before_script"
        },
        "after_script": {
          "$ref": "#/definitions/after_script"
        },
        "hooks": {
          "$ref": "#/definitions/hooks"
        },
        "rules": {
          "$ref": "#/definitions/rules"
        },
        "variables": {
          "$ref": "#/definitions/jobVariables"
        },
        "cache": {
          "$ref": "#/definitions/cache"
        },
        "id_tokens": {
          "$ref": "#/definitions/id_tokens"
        },
        "secrets": {
          "$ref": "#/definitions/secrets"
        },
        "script": {
          "markdownDescription": "Shell scripts executed by the Runner. The only required property of jobs. Be careful with special characters (e.g. ":", "{", "}", "&") and use single or double quotes to avoid issues. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#script)",
          "oneOf": [
            {
              "type": "string",
              "minLength": 1
            },
            {
              "type": "array",
              "items": {
                "anyOf": [
                  {
                    "type": "string"
                  },
                  {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  }
                ]
              },
              "minItems": 1
            }
          ]
        },
        "stage": {
          "description": "Define what stage the job will run in.",
          "anyOf": [
            {
              "type": "string",
              "minLength": 1
            },
            {
              "type": "array",
              "minLength": 1,
              "items": {
                "type": "string"
              }
            }
          ]
        },
        "only": {
          "$ref": "#/definitions/filter",
          "description": "Job will run *only* when these filtering options match."
        },
        "extends": {
          "description": "The name of one or more jobs to inherit configuration from.",
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "array",
              "items": {
                "type": "string"
              },
              "minItems": 1
            }
          ]
        },
        "needs": {
          "description": "The list of jobs in previous stages whose sole completion is needed to start the current job.",
          "type": "array",
          "items": {
            "oneOf": [
              {
                "type": "string"
              },
              {
                "type": "object",
                "additionalProperties": false,
                "properties": {
                  "job": {
                    "type": "string"
                  },
                  "artifacts": {
                    "type": "boolean"
                  },
                  "optional": {
                    "type": "boolean"
                  }
                },
                "required": [
                  "job"
                ]
              },
              {
                "type": "object",
                "additionalProperties": false,
                "properties": {
                  "pipeline": {
                    "type": "string"
                  },
                  "job": {
                    "type": "string"
                  },
                  "artifacts": {
                    "type": "boolean"
                  }
                },
                "required": [
                  "job",
                  "pipeline"
                ]
              },
              {
                "type": "object",
                "additionalProperties": false,
                "properties": {
                  "job": {
                    "type": "string"
                  },
                  "project": {
                    "type": "string"
                  },
                  "ref": {
                    "type": "string"
                  },
                  "artifacts": {
                    "type": "boolean"
                  }
                },
                "required": [
                  "job",
                  "project",
                  "ref"
                ]
              }
            ]
          }
        },
        "except": {
          "$ref": "#/definitions/filter",
          "description": "Job will run *except* for when these filtering options match."
        },
        "tags": {
          "$ref": "#/definitions/tags"
        },
        "allow_failure": {
          "$ref": "#/definitions/allow_failure"
        },
        "timeout": {
          "$ref": "#/definitions/timeout"
        },
        "when": {
          "$ref": "#/definitions/when"
        },
        "start_in": {
          "$ref": "#/definitions/start_in"
        },
        "dependencies": {
          "type": "array",
          "description": "Specify a list of job names from earlier stages from which artifacts should be loaded. By default, all previous artifacts are passed. Use an empty array to skip downloading artifacts.",
          "items": {
            "type": "string"
          }
        },
        "artifacts": {
          "$ref": "#/definitions/artifacts"
        },
        "environment": {
          "description": "Used to associate environment metadata with a deploy. Environment can have a name and URL attached to it, and will be displayed under /environments under the project.",
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "object",
              "additionalProperties": false,
              "properties": {
                "name": {
                  "type": "string",
                  "description": "The name of the environment, e.g. 'qa', 'staging', 'production'.",
                  "minLength": 1
                },
                "url": {
                  "type": "string",
                  "description": "When set, this will expose buttons in various places for the current environment in Gitlab, that will take you to the defined URL.",
                  "format": "uri",
                  "pattern": "^(https?://.+|\\$[A-Za-z]+)"
                },
                "on_stop": {
                  "type": "string",
                  "description": "The name of a job to execute when the environment is about to be stopped."
                },
                "action": {
                  "enum": [
                    "start",
                    "prepare",
                    "stop",
                    "verify",
                    "access"
                  ],
                  "description": "Specifies what this job will do. 'start' (default) indicates the job will start the deployment. 'prepare'/'verify'/'access' indicates this will not affect the deployment. 'stop' indicates this will stop the deployment.",
                  "default": "start"
                },
                "auto_stop_in": {
                  "type": "string",
                  "description": "The amount of time it should take before Gitlab will automatically stop the environment. Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'."
                },
                "kubernetes": {
                  "type": "object",
                  "description": "Used to configure the kubernetes deployment for this environment. This is currently not supported for kubernetes clusters that are managed by Gitlab.",
                  "properties": {
                    "namespace": {
                      "type": "string",
                      "description": "The kubernetes namespace where this environment should be deployed to.",
                      "minLength": 1
                    }
                  }
                },
                "deployment_tier": {
                  "type": "string",
                  "description": "Explicitly specifies the tier of the deployment environment if non-standard environment name is used.",
                  "enum": [
                    "production",
                    "staging",
                    "testing",
                    "development",
                    "other"
                  ]
                }
              },
              "required": [
                "name"
              ]
            }
          ]
        },
        "release": {
          "type": "object",
          "description": "Indicates that the job creates a Release.",
          "additionalProperties": false,
          "properties": {
            "tag_name": {
              "type": "string",
              "description": "The tag_name must be specified. It can refer to an existing Git tag or can be specified by the user.",
              "minLength": 1
            },
            "tag_message": {
              "type": "string",
              "description": "Message to use if creating a new annotated tag."
            },
            "description": {
              "type": "string",
              "description": "Specifies the longer description of the Release.",
              "minLength": 1
            },
            "name": {
              "type": "string",
              "description": "The Release name. If omitted, it is populated with the value of release: tag_name."
            },
            "ref": {
              "type": "string",
              "description": "If the release: tag_name doesn’t exist yet, the release is created from ref. ref can be a commit SHA, another tag name, or a branch name."
            },
            "milestones": {
              "type": "array",
              "description": "The title of each milestone the release is associated with.",
              "items": {
                "type": "string"
              }
            },
            "released_at": {
              "type": "string",
              "description": "The date and time when the release is ready. Defaults to the current date and time if not defined. Should be enclosed in quotes and expressed in ISO 8601 format.",
              "format": "date-time",
              "pattern": "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:Z|[+-][01]\\d:[0-5]\\d)$"
            },
            "assets": {
              "type": "object",
              "additionalProperties": false,
              "properties": {
                "links": {
                  "type": "array",
                  "description": "Include asset links in the release.",
                  "items": {
                    "type": "object",
                    "additionalProperties": false,
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "The name of the link.",
                        "minLength": 1
                      },
                      "url": {
                        "type": "string",
                        "description": "The URL to download a file.",
                        "minLength": 1
                      },
                      "filepath": {
                        "type": "string",
                        "description": "The redirect link to the url."
                      },
                      "link_type": {
                        "type": "string",
                        "description": "The content kind of what users can download via url.",
                        "enum": [
                          "runbook",
                          "package",
                          "image",
                          "other"
                        ]
                      }
                    },
                    "required": [
                      "name",
                      "url"
                    ]
                  },
                  "minItems": 1
                }
              },
              "required": [
                "links"
              ]
            }
          },
          "required": [
            "tag_name",
            "description"
          ]
        },
        "coverage": {
          "type": "string",
          "description": "Must be a regular expression, optionally but recommended to be quoted, and must be surrounded with '/'. Example: '/Code coverage: \\d+\\.\\d+/'",
          "format": "regex",
          "pattern": "^/.+/$"
        },
        "retry": {
          "$ref": "#/definitions/retry"
        },
        "parallel": {
          "description": "Parallel will split up a single job into several, and provide "CI_NODE_INDEX" and "CI_NODE_TOTAL" environment variables for the running jobs.",
          "oneOf": [
            {
              "type": "integer",
              "description": "Creates N instances of the same job that run in parallel.",
              "default": 0,
              "minimum": 2,
              "maximum": 200
            },
            {
              "type": "object",
              "properties": {
                "matrix": {
                  "type": "array",
                  "description": "Defines different variables for jobs that are running in parallel.",
                  "items": {
                    "type": "object",
                    "description": "Defines environment variables for specific job.",
                    "additionalProperties": {
                      "type": [
                        "string",
                        "number",
                        "array"
                      ]
                    }
                  },
                  "maxItems": 200
                }
              },
              "additionalProperties": false,
              "required": [
                "matrix"
              ]
            }
          ]
        },
        "interruptible": {
          "$ref": "#/definitions/interruptible"
        },
        "resource_group": {
          "type": "string",
          "description": "Limit job concurrency. Can be used to ensure that the Runner will not run certain jobs simultaneously."
        },
        "trigger": {
          "markdownDescription": "Trigger allows you to define downstream pipeline trigger. When a job created from trigger definition is started by GitLab, a downstream pipeline gets created. [Learn More](https://docs.gitlab.com/ee/ci/yaml/index.html#trigger).",
          "oneOf": [
            {
              "type": "object",
              "markdownDescription": "Trigger a multi-project pipeline. [Learn More](https://docs.gitlab.com/ee/ci/pipelines/multi_project_pipelines.html#specify-a-downstream-pipeline-branch).",
              "additionalProperties": false,
              "properties": {
                "project": {
                  "description": "Path to the project, e.g. "group/project", or "group/sub-group/project".",
                  "type": "string",
                  "pattern": "(?:\\S/\\S|\\$\\S+)"
                },
                "branch": {
                  "description": "The branch name that a downstream pipeline will use",
                  "type": "string"
                },
                "strategy": {
                  "description": "You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend",
                  "type": "string",
                  "enum": [
                    "depend"
                  ]
                },
                "forward": {
                  "description": "Specify what to forward to the downstream pipeline.",
                  "type": "object",
                  "additionalProperties": false,
                  "properties": {
                    "yaml_variables": {
                      "type": "boolean",
                      "description": "Variables defined in the trigger job are passed to downstream pipelines.",
                      "default": true
                    },
                    "pipeline_variables": {
                      "type": "boolean",
                      "description": "Variables added for manual pipeline runs and scheduled pipelines are passed to downstream pipelines.",
                      "default": false
                    }
                  }
                }
              },
              "required": [
                "project"
              ],
              "dependencies": {
                "branch": [
                  "project"
                ]
              }
            },
            {
              "type": "object",
              "description": "Trigger a child pipeline. [Learn More](https://docs.gitlab.com/ee/ci/pipelines/parent_child_pipelines.html).",
              "additionalProperties": false,
              "properties": {
                "include": {
                  "oneOf": [
                    {
                      "description": "Relative path from local repository root ("/") to the local YAML file to define the pipeline configuration.",
                      "type": "string",
                      "format": "uri-reference",
                      "pattern": "\\.ya?ml$"
                    },
                    {
                      "type": "array",
                      "description": "References a local file or an artifact from another job to define the pipeline configuration.",
                      "maxItems": 3,
                      "items": {
                        "oneOf": [
                          {
                            "type": "object",
                            "additionalProperties": false,
                            "properties": {
                              "local": {
                                "description": "Relative path from local repository root ("/") to the local YAML file to define the pipeline configuration.",
                                "type": "string",
                                "format": "uri-reference",
                                "pattern": "\\.ya?ml$"
                              }
                            }
                          },
                          {
                            "type": "object",
                            "additionalProperties": false,
                            "properties": {
                              "template": {
                                "description": "Name of the template YAML file to use in the pipeline configuration.",
                                "type": "string",
                                "format": "uri-reference",
                                "pattern": "\\.ya?ml$"
                              }
                            }
                          },
                          {
                            "type": "object",
                            "additionalProperties": false,
                            "properties": {
                              "artifact": {
                                "description": "Relative path to the generated YAML file which is extracted from the artifacts and used as the configuration for triggering the child pipeline.",
                                "type": "string",
                                "format": "uri-reference",
                                "pattern": "\\.ya?ml$"
                              },
                              "job": {
                                "description": "Job name which generates the artifact",
                                "type": "string"
                              }
                            },
                            "required": [
                              "artifact",
                              "job"
                            ]
                          },
                          {
                            "type": "object",
                            "additionalProperties": false,
                            "properties": {
                              "project": {
                                "description": "Path to another private project under the same GitLab instance, like "group/project" or "group/sub-group/project".",
                                "type": "string",
                                "pattern": "(?:\\S/\\S|\\$\\S+)"
                              },
                              "ref": {
                                "description": "Branch/Tag/Commit hash for the target project.",
                                "minLength": 1,
                                "type": "string"
                              },
                              "file": {
                                "description": "Relative path from repository root ("/") to the pipeline configuration YAML file.",
                                "type": "string",
                                "format": "uri-reference",
                                "pattern": "\\.ya?ml$"
                              }
                            },
                            "required": [
                              "project",
                              "file"
                            ]
                          }
                        ]
                      }
                    }
                  ]
                },
                "strategy": {
                  "description": "You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend",
                  "type": "string",
                  "enum": [
                    "depend"
                  ]
                },
                "forward": {
                  "description": "Specify what to forward to the downstream pipeline.",
                  "type": "object",
                  "additionalProperties": false,
                  "properties": {
                    "yaml_variables": {
                      "type": "boolean",
                      "description": "Variables defined in the trigger job are passed to downstream pipelines.",
                      "default": true
                    },
                    "pipeline_variables": {
                      "type": "boolean",
                      "description": "Variables added for manual pipeline runs and scheduled pipelines are passed to downstream pipelines.",
                      "default": false
                    }
                  }
                }
              }
            },
            {
              "markdownDescription": "Path to the project, e.g. "group/project", or "group/sub-group/project". [Learn More](https://docs.gitlab.com/ee/ci/yaml/index.html#trigger).",
              "type": "string",
              "pattern": "(?:\\S/\\S|\\$\\S+)"
            }
          ]
        },
        "inherit": {
          "type": "object",
          "markdownDescription": "Controls inheritance of globally-defined defaults and variables. Boolean values control inheritance of all default: or variables: keywords. To inherit only a subset of default: or variables: keywords, specify what you wish to inherit. Anything not listed is not inherited. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#inherit).",
          "properties": {
            "default": {
              "markdownDescription": "Whether to inherit all globally-defined defaults or not. Or subset of inherited defaults. [Learn more](https://docs.gitlab.com/ee/ci/yaml/#inheritdefault).",
              "oneOf": [
                {
                  "type": "boolean"
                },
                {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "enum": [
                      "after_script",
                      "artifacts",
                      "before_script",
                      "cache",
                      "image",
                      "interruptible",
                      "retry",
                      "services",
                      "tags",
                      "timeout"
                    ]
                  }
                }
              ]
            },
            "variables": {
              "markdownDescription": "Whether to inherit all globally-defined variables or not. Or subset of inherited variables. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#inheritvariables).",
              "oneOf": [
                {
                  "type": "boolean"
                },
                {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              ]
            }
          },
          "additionalProperties": false
        },
        "publish": {
          "description": "A path to a directory that contains the files to be published with Pages",
          "type": "string"
        }
      },
      "oneOf": [
        {
          "properties": {
            "when": {
              "enum": [
                "delayed"
              ]
            }
          },
          "required": [
            "when",
            "start_in"
          ]
        },
        {
          "properties": {
            "when": {
              "not": {
                "enum": [
                  "delayed"
                ]
              }
            }
          }
        }
      ]
    },
    "tags": {
      "type": "array",
      "minLength": 1,
      "markdownDescription": "Used to select runners from the list of available runners. A runner must have all tags listed here to run the job. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#tags).",
      "items": {
        "anyOf": [
          {
            "type": "string",
            "minLength": 1
          },
          {
            "type": "array",
            "minLength": 1,
            "items": {
              "type": "string"
            }
          }
        ]
      }
    },
    "hooks": {
      "type": "object",
      "markdownDescription": "Specifies lists of commands to execute on the runner at certain stages of job execution. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#hooks).",
      "properties": {
        "pre_get_sources_script": {
          "markdownDescription": "Specifies a list of commands to execute on the runner before updating the Git repository and any submodules. [Learn More](https://docs.gitlab.com/ee/ci/yaml/#hookspre_get_sources_script).",
          "oneOf": [
            {
              "type": "string",
              "minLength": 1
            },
            {
              "type": "array",
              "items": {
                "anyOf": [
                  {
                    "type": "string"
                  },
                  {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  }
                ]
              },
              "minItems": 1
            }
          ]
        }
      },
      "additionalProperties": false
    }
  }
}`
}
