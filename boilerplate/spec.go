// Package boilerplate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package boilerplate

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xcW2/bOBb+K4Z2H+067VyA8Vs6wQ4ymc4YbfZli0A4tk5s2hKlklRST+H/viCpCyWR",
	"lOw6aZLOW2KSh+R3vnPhIe0vwTJNspQiFTyYfQn4co0JqD+RCiII8lfndLVFcZFypPLzjKUZMtki/4Mo",
	"CiMQKP8WuwyDWcAFI3QV7MeqMefIjEZCBa6QydZbEtuHyYZwBckCmLWdRHaBJAqjcpXWVrluTHJw9NhA",
	"Em7zmMDaOu0mT2JYhxGoD+wSdJc1RIR5e8R5rxBBItj6RG1RIAO6amynXm6SRuSWoEc9VQ+3jirEwi3a",
	"e0hhW+BAwwQEMmKdiQsQObc07cflJ+lig0shO1fEu8JYU/FklHMzZ5tSjlQw4MTZJWdkm8d54urAMUEu",
	"nHNTwgnE1aY6S+/RZ2uF/2Z4G8yCf01rA54W1jutETTGSBHmDoYJqEacilEUEnBjYEI4aIEfygHG4CuX",
	"DfvZ1sT3ZJxbYXyoH9NMCd0jNznLuYsnfTSKMMxYGjl4fjIlh3aTquWoTu2V1M0ZbmABIhTC4XkEA8pj",
	"EOgEyq9v0xha2l4Kcoeh02uNH8IHKcXYfMyJXXqSx0B6JuJbHmYkJmtXqJQd7mFDFrZmL+wfDAt/AqiT",
	"KBSwzukDGoMlLNTj1eQhbIBpqAc5vWs56FyPUTLoagVxuEAG2zVxTaQ7KfUfaivmhE9Dbyc0hQWyDcQO",
	"d9kmhw+uNdAoRsZfXadbpO+RZ9IDdgETsrmTqhpzqvZkPTRhStIIY/7qLXAsJ50A3XUnjkCAIdNQNGMp",
	"U12IwIRb+xQfAGOwU7Mi57BS28PPkGSxWma+XCLnwdiiD9ST+xhe7GQOK0JBkJS+k4OGb1quLSxZGzry",
	"xxKFaqvDMiElzILDi8JOt4S/oTiPY8eR7yD8iqlMgd8Vhr9DdA+xzkjdHuF4SK3yvzOEkzrhPx2spdDv",
	"Css5snwBa6CaVqdEtC36O8H1HQg4JT0NeS8XwSqCe0uPJXyDwrcp6eUAVWacYU/GOQQqR/b6csBqeLrz",
	"Bf9awBrurSnvWYM2EKBdSPNk4TparcFRCvYcxLetsmCzNdkeW9cnUTioNNtZgFFwg8R5ctatHGPkrvah",
	"1wI9GxlQIlJFNffRUhfmnNscXreLka62kLk7NNQ1uOjXXys57EzsPcg4K3EYQ0KOqq72Y0Mix0jiAIw4",
	"BOUZPeyuZRzkFBI8BsBhp5geE+43Mm0ELkN3sr9TeD4Bvx+bop5TjIumX+eT+oichjnLxdDyrjfV71YM",
	"Y3VhmJX9rCuQESTcVBLs7WkcHXW14vf0vkvgqtk595ALmJ46v6pLbiDxlS3lSppLaNvBI4SBI708JJlv",
	"g3W7e4deM9Q+MMx5fmhdty6Uu9RbtttidbeXXkK4AirIwX7hHayBE34P7+H+YHfbkxZ5Zja9/a8MQeB7",
	"/JQjF901eNK8nrcYNu/V7dH1YEOSwv04YPgpJwyjYPZRL9JYUmv+zmSl5JthEDnB8StIuzbD1xxSBnCr",
	"qHOsaIFRLaq9As9u/4AsZUDnyHRfoBfXf9n2e+FT+ZWPrO+c2f28J7nnc4Z3BO+N5kWaxgjUsveLigBX",
	"BQhq6uY8HiR85RxnzOEhrCABZ6uAMNvag9wGKeEuLzsk0PgctKOt9KzdMNi8AHbfDMeQtZfV6pEx2Arc",
	"ujtwIo3T9RjnoPdE9hNwR31ZcfK2zEb+dj1BSEUjO2k3zR0yPcv8AHeozt7XDCh3+hZYFJZ2iNeQspVY",
	"NYGtAqE8YZlyVTWI1+OeA3tvz0Yo8vVuWazeprmusenBGpJv/KAaG7d56qSMtQP2w+ujfIR8yUgmaRXM",
	"gh3yUcpGNA3GRgVnh9bqTdN8m4L0I0Mpq3gpaMorXiCOLeRvefp6U0PAMYOK48h8TFnnMfKAo0tHvsKQ",
	"x0r/m0UgUANWBQ63sQ4571QUN/yun4edLNegCOEo8k1ObMRz5ab18DdnZz9Pzl5Pzt5cv/5pdvbj7Oyn",
	"/9kkFRIa5xX/mosRzRNAj0NoKUGFqGXOiNh9kP5NY/wWgSE7z4XOpaQRrREiZIGKdlKAbEwZ+VsFgXo7",
	"kJErlImSOh/epnL8MqUCllKV+3HLNP9a3Oacj87nl6OLdJknSEUlkAgFX90lGAd3yLge+VoCkGZIISPB",
	"LPjh1dmr18E4yEDopx7TylJWKLo+4TcUozJ/kexSs15GuqVsYEV1Rgl8c3ZWbgapkghZFpOlGjnd8JTW",
	"b74HhpBhbwYUli3YrhqaC2Yfmzr7eLO/GQc8TxJgu9ZuBay49Gj6/xspZ6ptZVLbihMz3XNUpG0W7Ey/",
	"p/TBIJG+matFKip9ypHtaiapXGFsINdlrX2gSiW8A2++vQ7tbwpOo9LfW8ooVdv83KLiiTKOqVfT53Hc",
	"nMBtL91Y11F8U/5cq/xoNjSl/UESIhzi4qLtKZGE7k5FALeODqVCFfO9fCjPef0MKHt+A/NrPVk7IdTl",
	"pg4Ft5EOZyn3IFxVqTooz1Pehdns/swsTqV3b9Nod2pKtGt9FgKoG2kzwRcsx/3z9gMmFRwEHQefJ8s0",
	"whXSSaGAySKNdpNCdQoWJ4s53OFEfU/ATWN5/Bmp808nOA1gc3V4Ch6UIJ4zmkUfEqABXPnxW3Dlz3T0",
	"azHlQZzxqOlI6iiUnNT5QqK9JkuM+oV5yx7V54OIc6G6djR3GfV5wGacvLwo3Zc8MdTei0QdXR+YPDwr",
	"IniBHxzmHtY1PJJXsN1CWJA2tlEs5qUEktM7hY15FrTQ5kHdwtPxCM9C+1/hCZIhhYP6TYb1+FA3fvsD",
	"e/Va/VSn9MQCZNIAMdaXgpOsvhXsP50XN4kj4yrRBm73wvGf4/njpeV2JZU8sLW6GTE0zJbRshtXu1w4",
	"PLp+ntzf309uU5ZMchYjlREgamLerNjfkhj/tD+WGwe3qUi7u1A/ZyEbWQJCppeEguKN921E1Z1Q8cOb",
	"YMj12YAhPY+Y+n45wv1DDu3fg+hdS/dOyniDUW+rseTGAqvltCbvXmg153rmWU3J7+OsTqSTLLp1W9x1",
	"Oppf/GeguV2n8+j2YRNZ+wOTx0hlC5jq1fba79frt0Lfq9zhlQ8bA76QaFLZy37acCCu3FUnpCP9VRlH",
	"6Vz36errsn4/Mzes2huzq0V5EtmJ6SOe/yG34oAda7+1jw/OrUZvdyNTHQMSraevyIf14dbvnj1sctXW",
	"Ur/Pr6rz0+odkJsasoenSG/Wns8X1puwR74C7aTj5W+GGPRqC03E5YGH3s4s6qrEPYW6brp84idrz3fc",
	"TsXhDp1KumryWAk6vNpWMrCbmDRJ+oh19+5rPAuUGpaXWF7TO2vno4Wy+5IU7Z6slPhiPLndD0lLFu58",
	"pMkNI4T1lg0GRK75i8pAhuq+hbrdyAUMqaO9AwFbZyHNeFL9LStp5hfrT+MpG9su8TM+1CBWjwIn9cNb",
	"J5LVN6uKcqcNz/a3r/550NT97YnTKNiijVLNnSa7sqUTbD0i3eu36GK57lJAvzrtZ8FcDm9t+rI9UZ9r",
	"vLwoa+rmEKuH7Eo+1E8+WAD3PtS18KCAuL3z0fvBkf1ZeXk3o9xM7ov5GeziFCId9dVvmXl9mvq5CZsj",
	"Kxse1Xs4flHD4zEaLqFcc4me/v9mr5XC7kpTy1kczILpNE6XEK9TLmY///LLz1PISLC/2f8/AAD//9Wk",
	"5p+WWQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
