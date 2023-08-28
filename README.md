### description

mltctl is a CLI for generating starter code for common machine learning utilities. Currently, the CLI provides the ability to generate basic Argo and Airflow workflows on the fly.

Currently, the templating only supports basic naming of the Argo image names to use. In the future, this functionality will be expanded such that different models can be specified to generate starter code for different types of ML models.

### installation

```go build -o mlctl```

### use

Fill in the contents of `template_config` with values for the image names used in your Argo workflow.

``````./mlctl template argo -t ./<destination-dir>``````