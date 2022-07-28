package services

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    npmRegistryData, err := UnmarshalNpmRegistryData(bytes)
//    bytes, err = npmRegistryData.Marshal()

import "encoding/json"

func UnmarshalNpmRegistryData(data []byte) (NpmPackageData, error) {
	var r NpmPackageData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NpmPackageData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NpmPackageData struct {
	Name                   string                  `json:"name"`
	Version                string                  `json:"version"`
	Description            *string                 `json:"description,omitempty"`
	Main                   *string                 `json:"main,omitempty"`
	Repository             *Repository             `json:"repository,omitempty"`
	Keywords               []string                `json:"keywords,omitempty"`
	License                *string                 `json:"license,omitempty"`
	Bugs                   *Bugs                   `json:"bugs,omitempty"`
	Homepage               *string                 `json:"homepage,omitempty"`
	Dependencies           *map[string]string      `json:"dependencies,omitempty"`
	PeerDependencies       *map[string]string      `json:"peerDependencies,omitempty"`
	Exports                *map[string]interface{} `json:"exports,omitempty"`
	Browser                *map[string]interface{} `json:"browser,omitempty"`
	Browserify             *map[string]interface{} `json:"browserify,omitempty"`
	Scripts                *map[string]string      `json:"scripts,omitempty"`
	ID                     string                  `json:"_id"`
	NodeVersion            string                  `json:"_nodeVersion"`
	NpmVersion             string                  `json:"_npmVersion"`
	Dist                   Dist                    `json:"dist"`
	NpmUser                NpmUser                 `json:"_npmUser"`
	Directories            Directories             `json:"directories"`
	Maintainers            []NpmUser               `json:"maintainers"`
	NpmOperationalInternal NpmOperationalInternal  `json:"_npmOperationalInternal"`
	HasShrinkwrap          bool                    `json:"_hasShrinkwrap"`
}

type Browser struct {
	ServerJS *string `json:"./server.js,omitempty"`
}

type Browserify struct {
	Transform []string `json:"transform,omitempty"`
}

type Bugs struct {
	URL *string `json:"url,omitempty"`
}

type Dependencies struct {
	LooseEnvify *string `json:"loose-envify,omitempty"`
	Scheduler   *string `json:"scheduler,omitempty"`
}

type Directories struct {
}

type Dist struct {
	Integrity    *string     `json:"integrity,omitempty"`
	Shasum       *string     `json:"shasum,omitempty"`
	Tarball      *string     `json:"tarball,omitempty"`
	FileCount    *int64      `json:"fileCount,omitempty"`
	UnpackedSize *int64      `json:"unpackedSize,omitempty"`
	Signatures   []Signature `json:"signatures,omitempty"`
	NpmSignature *string     `json:"npm-signature,omitempty"`
}

type Signature struct {
	Keyid *string `json:"keyid,omitempty"`
	Sig   *string `json:"sig,omitempty"`
}

type Exports struct {
	Empty         *string `json:".,omitempty"`
	Client        *string `json:"./client,omitempty"`
	Server        *Server `json:"./server,omitempty"`
	ServerBrowser *string `json:"./server.browser,omitempty"`
	ServerNode    *string `json:"./server.node,omitempty"`
	Profiling     *string `json:"./profiling,omitempty"`
	TestUtils     *string `json:"./test-utils,omitempty"`
	PackageJSON   *string `json:"./package.json,omitempty"`
}

type Server struct {
	Deno    *string `json:"deno,omitempty"`
	Worker  *string `json:"worker,omitempty"`
	Browser *string `json:"browser,omitempty"`
	Default *string `json:"default,omitempty"`
}

type NpmUser struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type NpmOperationalInternal struct {
	Host *string `json:"host,omitempty"`
	Tmp  *string `json:"tmp,omitempty"`
}

type PeerDependencies struct {
	React *string `json:"react,omitempty"`
}

type Repository struct {
	Type      *string `json:"type,omitempty"`
	URL       *string `json:"url,omitempty"`
	Directory *string `json:"directory,omitempty"`
}

type Scripts struct {
	Start *string `json:"start,omitempty"`
}
