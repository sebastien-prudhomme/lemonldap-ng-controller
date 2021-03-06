/*
Copyright 2018 Mathieu Parent <math.parent@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

// VHost defines a LemonLDAP::NG virtual host
type VHost struct {
	ServerName      string
	LocationRules   map[string]string
	ExportedHeaders map[string]string
}

// DefaultLocationRules is the default location rules when not set
var DefaultLocationRules = map[string]string{
	"default": "accept",
}

// DefaultExportedHeaders is the default exported headers when not set
var DefaultExportedHeaders = map[string]string{
	"Auth-User": "$uid",
}

// NewVHost creates a new LemonLDAP::NG virtual host
func NewVHost(serverName string, locationRules, exportedHeaders map[string]string) *VHost {
	return &VHost{
		ServerName:      serverName,
		LocationRules:   locationRules,
		ExportedHeaders: exportedHeaders,
	}
}
