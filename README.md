# terraform-provider-gotemplate

* It is a fork of repo  alex-leonhardt/terraform-provider-gotemplate
* the code is modified to add IncMore and ResetCounter function in the go-template.
* IncMore function increase the counter till Counter < len
* And ResetCounter will reset the counter.
* the above function help to create join for maps.

## build and run tf
```
go build -o terraform-provider-gotemplate; tf init; tf plan && tf apply
```

## mixed json

when having a mix of json, like
```
{
  "m": "yolo",
  22
}
```

one can use the included `template funcs` to assert the type and change how one deals with the values/keys - example see:
https://gist.github.com/alex-leonhardt/8ed3f78545706d89d466434fb6870023

### template functions

to assert a type when dealing with mixed json, you have the following available:
- IncMore
- ResetCounter
and you can use them like this

```
[
    {
      "logConfiguration": {
        "logDriver": "json-file",
        "options": {
          "max-size": "100m",
          "max-file": "1",
          "labels": "source"
        }
      },
      "command": "{{ .Data.command }}"
      "cpu": {{ .Data.cpu }},
      "environment": [
      {{- $env_length := len .Data.env }}
      {{- range $k,$v := .Data.env }}
        {
          "name": "{{ $k }}",
          "value": "{{ $v }}"
        }{{if $.IncMore $env_length}},{{end}}
      {{- end }}
      ],
      "mountPoints": [
      {{- range $k,$v := .Data.volume_map }}
        {
          "readOnly": true,
          "containerPath": "{{ $k }}",
          "sourceVolume": "{{ $v }}"
          },
      {{- end }}
      ],
      "dockerSecurityOptions": [],
      "memory": ${mem},
      "image": "${image}",
      "dockerLabels": {
        "logger": "logspout",
        "source": "${var.source}"
      },
      "name": "${var.name}"
    }
]
```

