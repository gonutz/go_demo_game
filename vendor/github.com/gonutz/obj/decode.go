package obj

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Vertices  [][4]float32
	TexCoords [][3]float32
	Normals   [][3]float32
	Faces     [][]FaceVertex
	Objects   []Object
}

type FaceVertex struct {
	VertexIndex   int
	TexCoordIndex int
	NormalIndex   int
}

type Object = struct {
	Name          string
	StartVertex   int
	StartTexCoord int
	StartNormal   int
	StartFace     int
	EndVertex     int
	EndTexCoord   int
	EndNormal     int
	EndFace       int
}

func (f *File) FindObject(name string) *Object {
	for i := range f.Objects {
		if f.Objects[i].Name == name {
			return &f.Objects[i]
		}
	}
	return nil
}

func Load(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Decode(f)
}

func Decode(r io.Reader) (*File, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s := string(data)
	s = strings.Replace(s, "\r\n", "\n", -1)
	lines := strings.Split(s, "\n")
	var f File
	for i, line := range lines {
		makeErr := func(msg string) error {
			return errors.New(fmt.Sprintf("%s in line %d: '%s'", msg, i+1, line))
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue // ignore empty lines
		} else if strings.HasPrefix(line, "#") {
			continue // ignore comments
		} else if strings.HasPrefix(line, "v ") {
			// vertex
			cols := strings.Split(strings.TrimSpace(line[2:]), " ")
			if len(cols) < 1 || len(cols) > 4 {
				return nil, makeErr("invalid vertex definition")
			}
			v := [4]float32{0, 0, 0, 1}
			for j, col := range cols {
				f, err := strconv.ParseFloat(col, 32)
				if err != nil {
					return nil, makeErr("invalid float in vertex definition")
				}
				v[j] = float32(f)
			}
			f.Vertices = append(f.Vertices, v)
		} else if strings.HasPrefix(line, "vt ") {
			// texture coordinate
			cols := strings.Split(strings.TrimSpace(line[3:]), " ")
			if len(cols) < 2 || len(cols) > 3 {
				return nil, makeErr("invalid texture coordinate definition")
			}
			uv := [3]float32{0, 0, 1}
			for j, col := range cols {
				f, err := strconv.ParseFloat(col, 32)
				if err != nil {
					return nil, makeErr("invalid float in texture coordinate definition")
				}
				uv[j] = float32(f)
			}
			f.TexCoords = append(f.TexCoords, uv)
		} else if strings.HasPrefix(line, "vn ") {
			// normal
			cols := strings.Split(strings.TrimSpace(line[3:]), " ")
			if len(cols) != 3 {
				return nil, makeErr("invalid normal definition")
			}
			var n [3]float32
			for j, col := range cols {
				f, err := strconv.ParseFloat(col, 32)
				if err != nil {
					return nil, makeErr("invalid float in normal definition")
				}
				n[j] = float32(f)
			}
			// normalize
			s := 1.0 / float32(math.Sqrt(float64(n[0]*n[0]+n[1]*n[1]+n[2]*n[2])))
			n[0] *= s
			n[1] *= s
			n[2] *= s
			f.Normals = append(f.Normals, n)
		} else if strings.HasPrefix(line, "f ") {
			// face
			cols := strings.Split(strings.TrimSpace(line[2:]), " ")
			if len(cols) < 3 {
				return nil, makeErr("invalid face definition, need at least 3 vertices")
			}
			var vertices []FaceVertex
			for _, col := range cols {
				parts := strings.Split(col, "/")
				if len(parts) == 0 || len(parts) > 3 {
					return nil, makeErr("invalid face vertex '" + col + "'")
				}
				v, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, makeErr("invalid vertex position index '" + parts[0] + "'")
				}
				var t int
				if len(parts) >= 2 && parts[1] != "" {
					t, err = strconv.Atoi(parts[1])
					if err != nil {
						return nil, makeErr("invalid texture coordinate index '" + parts[1] + "'")
					}
				}
				var n int
				if len(parts) >= 3 && parts[2] != "" {
					n, err = strconv.Atoi(parts[2])
					if err != nil {
						return nil, makeErr("invalid normal index '" + parts[1] + "'")
					}
				}
				vertices = append(vertices, FaceVertex{
					VertexIndex:   v - 1,
					TexCoordIndex: t - 1,
					NormalIndex:   n - 1,
				})
			}
			f.Faces = append(f.Faces, vertices)
		} else if strings.HasPrefix(line, "o ") {
			// object
			name := line[2:]

			if len(f.Objects) > 0 {
				// Remember the end of the last open object.
				o := &f.Objects[len(f.Objects)-1]
				o.EndVertex = len(f.Vertices)
				o.EndTexCoord = len(f.TexCoords)
				o.EndNormal = len(f.Normals)
				o.EndFace = len(f.Faces)
			}

			// Start a new object.
			f.Objects = append(f.Objects, Object{
				Name:          name,
				StartVertex:   len(f.Vertices),
				StartTexCoord: len(f.TexCoords),
				StartNormal:   len(f.Normals),
				StartFace:     len(f.Faces),
			})
		} else {
			continue // ignore unknown definition types
		}
	}

	if len(f.Objects) > 0 {
		// Remember the end of the last open object.
		o := &f.Objects[len(f.Objects)-1]
		o.EndVertex = len(f.Vertices)
		o.EndTexCoord = len(f.TexCoords)
		o.EndNormal = len(f.Normals)
		o.EndFace = len(f.Faces)
	}

	return &f, err
}
