{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "eventing.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default "eventing" .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "eventing.labels" -}}
component: {{ .Release.Name }}
helm.sh/chart: {{ include "eventing.chart" . }}
{{ include "eventing.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "eventing.postUpgrade.labels" -}}
component: {{ .Release.Name }}
job: post-upgrade-hook
{{- end }}

{{- define "eventing.postInstall.labels" -}}
component: {{ .Release.Name }}
job: post-install-hook
{{- end }}

{{/*
Selector labels
*/}}
{{- define "eventing.selectorLabels" -}}
app.kubernetes.io/name: {{ include "eventing.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Expand the name of the chart.
*/}}
{{- define "eventing.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "eventing.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Create the name of the secret to use
*/}}
{{- define "eventing.secretName" -}}
{{ default (include "eventing.fullname" .) .Values.global.secretName }}
{{- end }}

{{/*
Create a URL for container images
*/}}
{{- define "imageurl" -}}
{{- $registry := default $.reg.path $.img.containerRegistryPath -}}
{{- if hasKey $.img "directory" -}}
{{- printf "%s/%s/%s:%s" $registry $.img.directory $.img.name $.img.version -}}
{{- else -}}
{{- printf "%s/%s:%s" $registry $.img.name $.img.version -}}
{{- end -}}
{{- end -}}

{{/*
Create a URL for container images, without version number
*/}}
{{- define "shortimageurl" -}}
{{- $registry := default $.reg.path $.img.containerRegistryPath -}}
{{- if hasKey $.img "directory" -}}
{{- printf "%s/%s/%s" $registry $.img.directory $.img.name -}}
{{- else -}}
{{- printf "%s/%s" $registry $.img.name -}}
{{- end -}}
{{- end -}}
