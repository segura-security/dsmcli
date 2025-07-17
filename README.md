# Segura® DSM CLI

**Segura® DSM CLI** is a unified command‑line tool to manage Segura® DevOps Secret Manager (DSM) services. With it, you can fetch and inject secrets from Segura® DSM into your local or CI/CD environments—centralizing sensitive variables and automating their consumption in build and deployment pipelines.

## 🚀 Quick Links  
- 🌐 [Segura® Official Site](https://segura.security)  
- 📖 [DSM Documentation](https://docs.senhasegura.io/docs/devops-secret-manager?utm_source=Github&utm_medium=Link&utm_campaign=dsm_cli)



## Key Concepts

- **Running Belt (`runb`)**  
  Reads secrets from Segura® DSM and injects them as environment variables for any script or process.

- **Mapping & Registration**  
  Use a mapping file to register or update secrets in DSM directly from your CI/CD pipeline.

- **CI/CD–ready**  
  Integrates natively with GitHub Actions, Azure DevOps, Jenkins, GitLab CI, Bamboo, Bitbucket, CircleCI, TeamCity—or any Unix shell.


## Installation

Download the latest binary for your platform from [BIN](https://github.com/segura-security/dsmcli/tree/main/bin) and place it in your `$PATH`.

As an executable binary, the installation is quite simple. Before deploying the plugin it's important to have an application configured using OAuth 2.0 and an authorization on Segura® DSM. For more information on how to register applications and authorizations, please check the [DSM manual in Help Center](https://docs.senhasegura.io/docs/how-to-manage-an-application-in-devops-secret-manager?utm_source=Github&utm_medium=Link&utm_campaign=dsm_cli).


## Using DSM CLI as Running Belt

As for today, Segura® DSM CLI can only be executed as Running Belt, which reads Secrets from Segura® DevOps Secret Manager module and inject them as environment variables.

The first thing needed is to place the executable into a directory of your environment or CI/CD tool together with a configuration file for authentication on Segura® DSM. After that, DSM CLI need information from the configured application such as its name, system and environment so it can retrieve the secrets.

For the configuration file, it should be a `.yaml` file containing the following information from DSM:

- **_SENHASEGURA_URL:_** The URL of your Segura® environment where DSM is enabled;
- **_SENHASEGURA_CLIENT_ID:_** An authorization Client ID for authentication.
- **_SENHASEGURA_CLIENT_SECRET:_** An authorization Client Secret for authentication.

DSM CLI accepts extra parameters. Here is an axample of a **full .config.yaml** file:

```yaml title=".config.yaml"
# Default properties needed for execution
SENHASEGURA_URL: "<Segura URL>"
SENHASEGURA_CLIENT_ID: "<Segura DSM Client ID>"
SENHASEGURA_CLIENT_SECRET: "<Segura DSM Client Secret>"
SENHASEGURA_MAPPING_FILE: "<Secrets variable name mapping file with path>"
SENHASEGURA_SECRETS_FILE: "<File name with path to inject Secret>"
SENHASEGURA_DISABLE_RUNB: 0

# Properties needed to delete GitLab variables
GITLAB_ACCESS_TOKEN: "<Your GitLab Access Token>"
CI_API_V4_URL: "<Your GitLab API URL as for V4>"
CI_PROJECT_ID: "<Your GitLab Project ID>"
```

> **Using Environment Variables**
> 
> Instead of using a configuration file, DSM CLI can use authentication information through CI/CD environment variables, making the configuration file optional.

To execute the binary you can run the following command line providing the needed information:

```bash
dsm runb \
    --application <application name> \
    --system <system name> \
    --environment <environment name> \
    --config <path to config file>
```
> **Using Environment Variables**
> 
> It is possible to use a **SENHASEGURA_CONFIG_FILE** environment variable to define the configuration file location.

Being agnostic means that it can run in any environment or CI/CD tool, but DSM CLI already comes with some additional configuration allowing you to integrate more seamlessly with your tool.

After executing the plugin with the necessary informations, it will collect all the environment variables running on that pipeline execution and send them to Segura® DSM.

Then, it will query for all the application secrets registered, injecting them in a file called **.runb.vars** by default or whatever is set on **SENHASEGURA_SECRETS_FILE** if provided, which can be sourcered on the system to update the environment variables with the new values through the command bellow:

```bash
source .runb.vars
```

This way, developers will not have to worry about injecting secrets during pipelines, for example. They can be managed directly via API or through Segura® DSM interface by any developer or security team member.

> **Security Best Practice**
> 
> Make sure to delete the variables file from the environment to prevent secret leakage.

> **CI/CD Solutions**
> 
> By default DSM CLI can parse the secrets and inject it on tools like GitHub, Azure DevOps, Bamboo, BitBucket, CircleCI, TeamCity and Linux (default option). You can change the default option with the --tool-name argument during its execution.


## Using DSM CLI to Register and Update Secrets

Using DSM CLI also allows developers to create or update secret values directly from the pipeline using a mapping file. This file makes it easy to identify secret variables through their names and automatically register them as secrets on Segura® DSM.

To do that, the only additional configuration needed is actually to provide the mapping file together with the executable and the configuration file. Here is an example of mapping file's content:

``` json
{
  "access_keys": [
    {
      "name": "ACCESS_KEY_VARIABLES",
      "type": "aws",
      "fields": {
        "access_key_id": "AWS_ACCESS_KEY_ID_VARIABLE",
        "secret_access_key": "AWS_SECRET_ACCESS_KEY_VARIABLE"
      }
    }
  ],
  "credentials": [
    {
      "name": "CREDENCIAL_VARIABLES",
      "fields": {
        "user": "USER_VARIABLE",
        "password": "PASSWORD_VARIABLE",
        "host": "HOST_VARIABLE"
      }
    }
  ],
  "key_value": [
    {
      "name": "GENERIC_VARIABLES",
      "fields": ["KEY_VALUE_VARIABLE"]
    }
  ]
}
```

This file can be broken down in 3 main blocks:

- **_access_keys:_** An array of objects composed by a `name` attribute, `type` and a sub-object `fields`, where this one is composed by an `access_key_id` and `secret_access_key`. These attribute values should be the name of the variable holding the values, so Segura® DSM will validate if the provided data exists on the Cloud IAM module and if it does it will register it as a secret for that provided authorization.
- **_credentials:_** An array of objects composed by a `name` and a sub-object `fields`, where this one is composed by `user`, `password` and `host`. The values of those attributes should be the name of the variables holding that information so Segura® DSM will validate if the provided data exists on the PAM Core module and if it does it will register it as a Secret for that provided authorization.
- **_key_value:_** An array of objects composed by `name` and a sub-array of `fields`, where the values of the array should be the name of the variables to be registered as secrets on Segura® DSM.

> **Mapping File Name**
> 
> To set the name of the file so DSM CLI can read it use the **SENHASEGURA_MAPPING_FILE** option in the configuration file or set its value as an environment variable pointing to the file full path.

> **Type Values**
> 
> Currently Segura® DSM only supports access keys through integration with **AWS**, **Azure** or **GCP**, so the **_type_** attribute informed should be one of the supported.
